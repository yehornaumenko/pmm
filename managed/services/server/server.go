// Copyright (C) 2023 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package server implements pmm-managed Server API.
package server

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm/api/serverpb"
	"github.com/percona/pmm/managed/models"
	"github.com/percona/pmm/managed/utils/envvars"
	"github.com/percona/pmm/version"
)

// Server represents service for checking PMM Server status and changing settings.
type Server struct {
	db                   *reform.DB
	vmdb                 prometheusService
	agentsState          agentsStateUpdater
	vmalert              vmAlertService
	vmalertExternalRules vmAlertExternalRules
	checksService        checksService
	templatesService     templatesService
	supervisord          supervisordService
	telemetryService     telemetryService
	awsInstanceChecker   *AWSInstanceChecker
	grafanaClient        grafanaClient
	haService            haService
	updater              *Updater

	l *logrus.Entry

	pmmUpdateAuthFileM sync.Mutex
	pmmUpdateAuthFile  string

	envRW       sync.RWMutex
	envSettings *models.ChangeSettingsParams

	sshKeyM sync.Mutex

	serverpb.UnimplementedServerServer
}

type pmmUpdateAuth struct {
	AuthToken string `json:"auth_token"`
}

// Params holds the parameters needed to create a new service.
type Params struct {
	DB                   *reform.DB
	AgentsStateUpdater   agentsStateUpdater
	VMDB                 prometheusService
	VMAlert              prometheusService
	ChecksService        checksService
	TemplatesService     templatesService
	VMAlertExternalRules vmAlertExternalRules
	Supervisord          supervisordService
	TelemetryService     telemetryService
	AwsInstanceChecker   *AWSInstanceChecker
	GrafanaClient        grafanaClient
	Updater              *Updater
}

// NewServer returns new server for Server service.
func NewServer(params *Params) (*Server, error) {
	path := os.TempDir()
	if _, err := os.Stat(path); err != nil {
		return nil, errors.WithStack(err)
	}
	path = filepath.Join(path, "pmm-update.json")

	s := &Server{
		db:                   params.DB,
		vmdb:                 params.VMDB,
		agentsState:          params.AgentsStateUpdater,
		vmalert:              params.VMAlert,
		checksService:        params.ChecksService,
		templatesService:     params.TemplatesService,
		vmalertExternalRules: params.VMAlertExternalRules,
		supervisord:          params.Supervisord,
		telemetryService:     params.TelemetryService,
		awsInstanceChecker:   params.AwsInstanceChecker,
		grafanaClient:        params.GrafanaClient,
		updater:              params.Updater,
		l:                    logrus.WithField("component", "server"),
		pmmUpdateAuthFile:    path,
		envSettings:          &models.ChangeSettingsParams{},
	}
	return s, nil
}

// UpdateSettingsFromEnv updates settings in the database with environment variables values.
// It returns only validation or database errors; invalid environment variables are logged and skipped.
func (s *Server) UpdateSettingsFromEnv(env []string) []error {
	s.envRW.Lock()
	defer s.envRW.Unlock()

	envSettings, errs, warns := envvars.ParseEnvVars(env)
	for _, w := range warns {
		s.l.Warnln(w)
	}
	if len(errs) != 0 {
		return errs
	}

	err := s.db.InTransaction(func(tx *reform.TX) error {
		_, err := models.UpdateSettings(tx, envSettings)
		return err
	})
	if err != nil {
		return []error{err}
	}
	s.envSettings = envSettings
	return nil
}

// Version returns PMM Server version.
func (s *Server) Version(ctx context.Context, req *serverpb.VersionRequest) (*serverpb.VersionResponse, error) {
	// for API testing of authentication, panic handling, etc.
	if req.Dummy != "" {
		switch {
		case strings.HasPrefix(req.Dummy, "panic-"):
			switch req.Dummy {
			case "panic-error":
				panic(errors.New("panic-error"))
			case "panic-fmterror":
				panic(fmt.Errorf("panic-fmterror"))
			default:
				panic(req.Dummy)
			}

		case strings.HasPrefix(req.Dummy, "grpccode-"):
			code, err := strconv.Atoi(strings.TrimPrefix(req.Dummy, "grpccode-"))
			if err != nil {
				return nil, err
			}
			grpcCode := codes.Code(code)
			return nil, status.Errorf(grpcCode, "gRPC code %d (%s)", grpcCode, grpcCode)
		}
	}

	res := &serverpb.VersionResponse{
		// always return something in this field:
		// it is used by PMM 1.x's pmm-client for compatibility checking
		Version: version.Version,

		Managed: &serverpb.VersionInfo{
			Version:     version.Version,
			FullVersion: version.FullCommit,
		},

		DistributionMethod: s.telemetryService.DistributionMethod(),
	}
	if t, err := version.Time(); err == nil {
		res.Managed.Timestamp = timestamppb.New(t)
	}

	v := s.updater.InstalledPMMVersion()
	res.Version = v.Version
	res.Server = &serverpb.VersionInfo{
		Version:     v.Version,
		FullVersion: v.FullVersion,
	}
	if v.BuildTime != nil {
		res.Server.Timestamp = timestamppb.New(*v.BuildTime)
	}

	return res, nil
}

// Readiness returns an error when some PMM Server component is not ready yet or is being restarted.
// It can be used as for Docker health check or Kubernetes readiness probe.
func (s *Server) Readiness(ctx context.Context, req *serverpb.ReadinessRequest) (*serverpb.ReadinessResponse, error) { //nolint:revive
	var notReady bool
	for n, svc := range map[string]healthChecker{
		"grafana":         s.grafanaClient,
		"victoriametrics": s.vmdb,
		"vmalert":         s.vmalert,
	} {
		if err := svc.IsReady(ctx); err != nil {
			s.l.Errorf("%s readiness check failed: %+v", n, err)
			notReady = true
		}
	}

	if notReady {
		return nil, status.Error(codes.Internal, "PMM Server is not ready yet.")
	}

	return &serverpb.ReadinessResponse{}, nil
}

// LeaderHealthCheck checks if the instance is the leader in a cluster.
// Returns an error if the instance isn't the leader.
// It's used for HA purpose.
func (s *Server) LeaderHealthCheck(ctx context.Context, req *serverpb.LeaderHealthCheckRequest) (*serverpb.LeaderHealthCheckResponse, error) { //nolint:revive
	if s.haService.IsLeader() {
		return &serverpb.LeaderHealthCheckResponse{}, nil
	}
	return nil, status.Error(codes.FailedPrecondition, "this PMM Server isn't the leader")
}

// CheckUpdates checks PMM Server updates availability.
func (s *Server) CheckUpdates(ctx context.Context, req *serverpb.CheckUpdatesRequest) (*serverpb.CheckUpdatesResponse, error) {
	s.envRW.RLock()
	updatesDisabled := s.envSettings.DisableUpdates
	s.envRW.RUnlock()

	if req.OnlyInstalledVersion {
		installedPMMVersion := s.updater.InstalledPMMVersion()
		return &serverpb.CheckUpdatesResponse{
			Installed: &serverpb.VersionInfo{
				Version:     installedPMMVersion.Version,
				FullVersion: installedPMMVersion.FullVersion,
				Timestamp:   timestamppb.New(*installedPMMVersion.BuildTime),
			},
		}, nil
	}

	if req.Force {
		if err := s.updater.ForceCheckUpdates(ctx); err != nil {
			return nil, err
		}
	}

	v, lastCheck := s.updater.LastCheckUpdatesResult(ctx)
	if v == nil {
		return nil, status.Error(codes.Unavailable, "failed to check for updates")
	}

	res := convertResponseToProto(v)

	s.l.Infof("%+v", res.AvailableVersions)
	if updatesDisabled {
		res.UpdateAvailable = false
	}

	res.LastCheck = timestamppb.New(lastCheck)

	if v.Installed.BuildTime != nil {
		t := v.Installed.BuildTime.UTC().Truncate(24 * time.Hour) // return only date
		res.Installed.Timestamp = timestamppb.New(t)
	}

	if v.Latest.DockerImage != "" {
		t := v.Latest.BuildTime.UTC().Truncate(24 * time.Hour) // return only date
		res.Latest.Timestamp = timestamppb.New(t)
	}

	return res, nil
}

func convertResponseToProto(v *version.UpdateCheckResult) *serverpb.CheckUpdatesResponse {
	newerVersions := make([]*serverpb.DockerVersionInfo, 0, len(v.NewerVersions))
	for _, nv := range v.NewerVersions {
		newerVersions = append(newerVersions, &serverpb.DockerVersionInfo{
			Version:          nv.Version.String(),
			ReleaseNotesUrl:  nv.ReleaseNotesURL,
			ReleaseNotesText: nv.ReleaseNotesText,
			Tag:              nv.DockerImage,
			Timestamp:        timestamppb.New(nv.BuildTime),
		})
	}

	return &serverpb.CheckUpdatesResponse{
		Installed: &serverpb.VersionInfo{
			Version:     v.Installed.Version,
			FullVersion: v.Installed.FullVersion,
			Timestamp:   timestamppb.New(*v.Installed.BuildTime),
		},
		Latest: &serverpb.DockerVersionInfo{
			Version:          v.Latest.Version.String(),
			ReleaseNotesUrl:  v.Latest.ReleaseNotesURL,
			ReleaseNotesText: v.Latest.ReleaseNotesText,
			Tag:              v.Latest.DockerImage,
			Timestamp:        timestamppb.New(v.Latest.BuildTime),
		},
		AvailableVersions: newerVersions,
		UpdateAvailable:   v.Latest.DockerImage != "",
	}
}

// StartUpdate starts PMM Server update.
func (s *Server) StartUpdate(ctx context.Context, req *serverpb.StartUpdateRequest) (*serverpb.StartUpdateResponse, error) { //nolint:revive
	s.envRW.RLock()
	updatesDisabled := s.envSettings.DisableUpdates
	s.envRW.RUnlock()

	if updatesDisabled {
		return nil, status.Error(codes.FailedPrecondition, "Updates are disabled via PMM_DISABLE_UPDATES environment variable.")
	}

	newImage := req.GetNewImage()
	if newImage == "" {
		_, latest, err := s.updater.latest(ctx)
		if err != nil {
			s.l.WithError(err).Error("Failed to get latest version")
			newImage = defaultLatestPMMImage
		} else {
			newImage = latest.DockerImage
		}
	}
	err := s.updater.StartUpdate(ctx, newImage)
	if err != nil {
		return nil, err
	}

	authToken := uuid.New().String()
	if err := s.writeUpdateAuthToken(authToken); err != nil {
		return nil, err
	}

	return &serverpb.StartUpdateResponse{
		AuthToken: authToken,
		LogOffset: 0,
	}, nil
}

// UpdateStatus returns PMM Server update status.
func (s *Server) UpdateStatus(ctx context.Context, req *serverpb.UpdateStatusRequest) (*serverpb.UpdateStatusResponse, error) {
	token, err := s.readUpdateAuthToken()
	if err != nil {
		return nil, err
	}
	if subtle.ConstantTimeCompare([]byte(req.AuthToken), []byte(token)) == 0 {
		return nil, status.Error(codes.PermissionDenied, "Invalid authentication token.")
	}

	// wait up to 30 seconds for new log lines
	var lines []string
	var newOffset uint32
	var done bool
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for ctx.Err() == nil {
		done = !s.updater.IsRunning()
		if done {
			// give supervisord a second to flush logs to file
			time.Sleep(time.Second)
		}

		lines, newOffset, err = s.updater.UpdateLog(req.GetLogOffset())
		if err != nil {
			s.l.Warn(err)
		}

		if len(lines) != 0 || done {
			break
		}

		time.Sleep(200 * time.Millisecond)
	}

	return &serverpb.UpdateStatusResponse{
		LogLines:  lines,
		LogOffset: newOffset,
		Done:      done,
	}, nil
}

// writeUpdateAuthToken writes authentication token for getting update status and logs to the file.
//
// We can't rely on Grafana for authentication or on PostgreSQL for storage as their configuration
// is being changed during update.
func (s *Server) writeUpdateAuthToken(token string) error {
	s.pmmUpdateAuthFileM.Lock()
	defer s.pmmUpdateAuthFileM.Unlock()

	a := &pmmUpdateAuth{
		AuthToken: token,
	}
	f, err := os.OpenFile(s.pmmUpdateAuthFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o600|os.ModeExclusive)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			s.l.Error(err)
		}
	}()

	return errors.WithStack(json.NewEncoder(f).Encode(a))
}

// readUpdateAuthToken reads authentication token for getting update status and logs from the file.
func (s *Server) readUpdateAuthToken() (string, error) {
	s.pmmUpdateAuthFileM.Lock()
	defer s.pmmUpdateAuthFileM.Unlock()

	f, err := os.OpenFile(s.pmmUpdateAuthFile, os.O_RDONLY, os.ModeExclusive)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			s.l.Error(err)
		}
	}()

	var a pmmUpdateAuth
	err = json.NewDecoder(f).Decode(&a)
	return a.AuthToken, errors.WithStack(err)
}

// convertSettings merges database settings and settings from environment variables into API response.
// Checking if PMM is connected to Platform is separated from settings for security and concurrency reasons.
func (s *Server) convertSettings(settings *models.Settings, connectedToPlatform bool) *serverpb.Settings {
	res := &serverpb.Settings{
		UpdatesDisabled:  settings.Updates.Disabled,
		TelemetryEnabled: !settings.Telemetry.Disabled,
		MetricsResolutions: &serverpb.MetricsResolutions{
			Hr: durationpb.New(settings.MetricsResolutions.HR),
			Mr: durationpb.New(settings.MetricsResolutions.MR),
			Lr: durationpb.New(settings.MetricsResolutions.LR),
		},
		SttCheckIntervals: &serverpb.STTCheckIntervals{
			RareInterval:     durationpb.New(settings.SaaS.STTCheckIntervals.RareInterval),
			StandardInterval: durationpb.New(settings.SaaS.STTCheckIntervals.StandardInterval),
			FrequentInterval: durationpb.New(settings.SaaS.STTCheckIntervals.FrequentInterval),
		},
		DataRetention:        durationpb.New(settings.DataRetention),
		SshKey:               settings.SSHKey,
		AwsPartitions:        settings.AWSPartitions,
		SttEnabled:           !settings.SaaS.STTDisabled,
		AzurediscoverEnabled: settings.Azurediscover.Enabled,
		PmmPublicAddress:     settings.PMMPublicAddress,

		AlertingEnabled:         !settings.Alerting.Disabled,
		BackupManagementEnabled: !settings.BackupManagement.Disabled,
		ConnectedToPlatform:     connectedToPlatform,

		TelemetrySummaries: s.telemetryService.GetSummaries(),

		EnableAccessControl: settings.AccessControl.Enabled,
		DefaultRoleId:       uint32(settings.DefaultRoleID),
	}

	return res
}

// GetSettings returns current PMM Server settings.
func (s *Server) GetSettings(ctx context.Context, req *serverpb.GetSettingsRequest) (*serverpb.GetSettingsResponse, error) { //nolint:revive
	s.envRW.RLock()
	defer s.envRW.RUnlock()

	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	_, err = models.GetPerconaSSODetails(ctx, s.db.Querier)

	return &serverpb.GetSettingsResponse{
		Settings: s.convertSettings(settings, err == nil),
	}, nil
}

func (s *Server) validateChangeSettingsRequest(ctx context.Context, req *serverpb.ChangeSettingsRequest) error {
	metricsRes := req.MetricsResolutions

	if req.PmmPublicAddress != "" && req.RemovePmmPublicAddress {
		return status.Error(codes.InvalidArgument, "Both pmm_public_address and remove_pmm_public_address are present.")
	}

	if req.SshKey != "" {
		if err := s.validateSSHKey(ctx, req.SshKey); err != nil {
			return err
		}
	}

	// check request parameters compatibility with environment variables

	if req.EnableUpdates && s.envSettings.DisableUpdates {
		return status.Error(codes.FailedPrecondition, "Updates are disabled via PMM_DISABLE_UPDATES environment variable.")
	}

	// ignore req.DisableTelemetry and req.DisableStt even if they are present since that will not change anything
	if req.EnableTelemetry && s.envSettings.DisableTelemetry {
		return status.Error(codes.FailedPrecondition, "Telemetry is disabled via PMM_DISABLE_TELEMETRY environment variable.")
	}

	// ignore req.EnableAlerting even if they are present since that will not change anything
	if req.DisableAlerting && s.envSettings.EnableAlerting {
		return status.Error(codes.FailedPrecondition, "Alerting is enabled via ENABLE_ALERTING environment variable.")
	}

	// ignore req.DisableAzurediscover even if they are present since that will not change anything
	if req.DisableAzurediscover && s.envSettings.EnableAzurediscover {
		return status.Error(codes.FailedPrecondition, "Azure Discover is enabled via PMM_ENABLE_AZURE_DISCOVER environment variable.")
	}

	if !canUpdateDurationSetting(metricsRes.GetHr().AsDuration(), s.envSettings.MetricsResolutions.HR) {
		return status.Error(codes.FailedPrecondition, "High resolution for metrics is set via PMM_METRICS_RESOLUTION_HR (or PMM_METRICS_RESOLUTION) environment variable.") //nolint:lll
	}

	if !canUpdateDurationSetting(metricsRes.GetMr().AsDuration(), s.envSettings.MetricsResolutions.MR) {
		return status.Error(codes.FailedPrecondition, "Medium resolution for metrics is set via PMM_METRICS_RESOLUTION_MR environment variable.")
	}

	if !canUpdateDurationSetting(metricsRes.GetLr().AsDuration(), s.envSettings.MetricsResolutions.LR) {
		return status.Error(codes.FailedPrecondition, "Low resolution for metrics is set via PMM_METRICS_RESOLUTION_LR environment variable.")
	}

	if !canUpdateDurationSetting(req.DataRetention.AsDuration(), s.envSettings.DataRetention) {
		return status.Error(codes.FailedPrecondition, "Data retention for queries is set via PMM_DATA_RETENTION environment variable.")
	}

	return nil
}

// ChangeSettings changes PMM Server settings.
func (s *Server) ChangeSettings(ctx context.Context, req *serverpb.ChangeSettingsRequest) (*serverpb.ChangeSettingsResponse, error) { //nolint:cyclop,maintidx
	s.envRW.RLock()
	defer s.envRW.RUnlock()

	if err := s.validateChangeSettingsRequest(ctx, req); err != nil {
		return nil, err
	}

	var newSettings, oldSettings *models.Settings
	errTX := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		if oldSettings, err = models.GetSettings(tx); err != nil {
			return errors.WithStack(err)
		}

		metricsRes := req.MetricsResolutions
		sttCheckIntervals := req.SttCheckIntervals
		settingsParams := &models.ChangeSettingsParams{
			DisableUpdates:   req.DisableUpdates,
			EnableUpdates:    req.EnableUpdates,
			DisableTelemetry: req.DisableTelemetry,
			EnableTelemetry:  req.EnableTelemetry,
			STTCheckIntervals: models.STTCheckIntervals{
				RareInterval:     sttCheckIntervals.GetRareInterval().AsDuration(),
				StandardInterval: sttCheckIntervals.GetStandardInterval().AsDuration(),
				FrequentInterval: sttCheckIntervals.GetFrequentInterval().AsDuration(),
			},
			MetricsResolutions: models.MetricsResolutions{
				HR: metricsRes.GetHr().AsDuration(),
				MR: metricsRes.GetMr().AsDuration(),
				LR: metricsRes.GetLr().AsDuration(),
			},
			DataRetention:          req.DataRetention.AsDuration(),
			AWSPartitions:          req.AwsPartitions,
			SSHKey:                 req.SshKey,
			EnableSTT:              req.EnableStt,
			DisableSTT:             req.DisableStt,
			EnableAzurediscover:    req.EnableAzurediscover,
			DisableAzurediscover:   req.DisableAzurediscover,
			PMMPublicAddress:       req.PmmPublicAddress,
			RemovePMMPublicAddress: req.RemovePmmPublicAddress,

			EnableAlerting:          req.EnableAlerting,
			DisableAlerting:         req.DisableAlerting,
			EnableBackupManagement:  req.EnableBackupManagement,
			DisableBackupManagement: req.DisableBackupManagement,

			EnableAccessControl:  req.EnableAccessControl,
			DisableAccessControl: req.DisableAccessControl,
		}

		var errInvalidArgument *models.InvalidArgumentError
		newSettings, err = models.UpdateSettings(tx, settingsParams)
		switch {
		case err == nil:
		case errors.As(err, &errInvalidArgument):
			return status.Errorf(codes.InvalidArgument, "Invalid argument: %s.", errInvalidArgument.Details)
		default:
			return errors.WithStack(err)
		}

		// absent value means "do not change"
		if req.SshKey != "" {
			if err = s.writeSSHKey(req.SshKey); err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})
	if errTX != nil {
		return nil, errTX
	}

	if err := s.UpdateConfigurations(ctx); err != nil {
		return nil, err
	}

	// If STT intervals are changed reset timers.
	if oldSettings.SaaS.STTCheckIntervals != newSettings.SaaS.STTCheckIntervals {
		s.checksService.UpdateIntervals(
			newSettings.SaaS.STTCheckIntervals.RareInterval,
			newSettings.SaaS.STTCheckIntervals.StandardInterval,
			newSettings.SaaS.STTCheckIntervals.FrequentInterval)
	}

	// When STT moved from disabled state to enabled force checks download and execution.
	var sttStarted bool
	if oldSettings.SaaS.STTDisabled && !newSettings.SaaS.STTDisabled {
		sttStarted = true
		if err := s.checksService.StartChecks(nil); err != nil {
			s.l.Error(err)
		}
	}

	// When STT moved from enabled state to disabled drop all existing STT alerts.
	if !oldSettings.SaaS.STTDisabled && newSettings.SaaS.STTDisabled {
		s.checksService.CleanupAlerts()
	}

	// When telemetry state is switched force alert templates and STT checks files collection.
	// If telemetry switched off that will drop previously downloaded files.
	if oldSettings.Telemetry.Disabled != newSettings.Telemetry.Disabled {
		s.templatesService.CollectTemplates(ctx)
		if !sttStarted {
			s.checksService.CollectAdvisors(ctx)
		}
	}

	if isAgentsStateUpdateNeeded(req.MetricsResolutions) {
		if err := s.agentsState.UpdateAgentsState(ctx); err != nil {
			return nil, err
		}
	}

	_, err := models.GetPerconaSSODetails(ctx, s.db.Querier)

	return &serverpb.ChangeSettingsResponse{
		Settings: s.convertSettings(newSettings, err == nil),
	}, nil
}

// UpdateConfigurations updates supervisor config and requests configuration update for VictoriaMetrics components.
func (s *Server) UpdateConfigurations(ctx context.Context) error {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return errors.Wrap(err, "failed to get settings")
	}
	ssoDetails, err := models.GetPerconaSSODetails(ctx, s.db.Querier)
	if err != nil {
		if !errors.Is(err, models.ErrNotConnectedToPortal) {
			return errors.Wrap(err, "failed to get SSO details")
		}
	}
	if err := s.supervisord.UpdateConfiguration(settings, ssoDetails); err != nil {
		return errors.Wrap(err, "failed to update supervisord configuration")
	}
	s.vmdb.RequestConfigurationUpdate()
	s.vmalert.RequestConfigurationUpdate()
	return nil
}

func (s *Server) validateSSHKey(_ context.Context, sshKey string) error {
	_, _, _, _, err := ssh.ParseAuthorizedKey([]byte(sshKey)) //nolint:dogsled
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid SSH key.")
	}

	return nil
}

func (s *Server) writeSSHKey(sshKey string) error {
	s.sshKeyM.Lock()
	defer s.sshKeyM.Unlock()

	username := "pmm"
	usr, err := user.Lookup(username)
	if err != nil {
		return errors.WithStack(err)
	}
	sshDirPath := path.Join(usr.HomeDir, ".ssh")
	if err = os.MkdirAll(sshDirPath, 0o700); err != nil {
		return errors.WithStack(err)
	}

	uid, err := strconv.Atoi(usr.Uid)
	if err != nil {
		return errors.WithStack(err)
	}
	gid, err := strconv.Atoi(usr.Gid)
	if err != nil {
		return errors.WithStack(err)
	}
	if err = os.Chown(sshDirPath, uid, gid); err != nil {
		return errors.WithStack(err)
	}
	keysPath := path.Join(sshDirPath, "authorized_keys")
	if err = os.WriteFile(keysPath, []byte(sshKey), 0o600); err != nil {
		return errors.WithStack(err)
	}
	if err = os.Chown(keysPath, uid, gid); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// AWSInstanceCheck checks AWS EC2 instance ID.
func (s *Server) AWSInstanceCheck(ctx context.Context, req *serverpb.AWSInstanceCheckRequest) (*serverpb.AWSInstanceCheckResponse, error) { //nolint:revive
	if err := s.awsInstanceChecker.check(req.InstanceId); err != nil {
		return nil, err
	}
	return &serverpb.AWSInstanceCheckResponse{}, nil
}

// isAgentsStateUpdateNeeded - checks metrics resolution changes,
// if it was changed, agents state must be updated.
func isAgentsStateUpdateNeeded(mr *serverpb.MetricsResolutions) bool {
	if mr == nil {
		return false
	}
	if mr.Lr == nil && mr.Hr == nil && mr.Mr == nil {
		return false
	}
	return true
}

func canUpdateDurationSetting(newValue, envValue time.Duration) bool {
	if newValue == 0 || envValue == 0 || newValue == envValue {
		return true
	}

	return false
}

// check interfaces.
var (
	_ serverpb.ServerServer = (*Server)(nil)
)
