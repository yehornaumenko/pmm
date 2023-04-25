// Copyright 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package apiserver holds logic for operating API server.
package apiserver

import (
	"context"
	"net"
	"os"
	"sync"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"

	"github.com/percona/pmm/admin/services/status"
	"github.com/percona/pmm/admin/services/update"
	"github.com/percona/pmm/api/updatepb"
)

const (
	SocketPath      = "/srv/pmm-server-upgrade.sock"
	shutdownTimeout = 1 * time.Second
)

// Server allows for running API server.
type Server struct {
	EnableDebug bool

	l                  *logrus.Entry
	gRPCMessageMaxSize int

	// gRPCServer stores reference to the current server which is running. It is nil if the server is not running.
	gRPCServer   *grpc.Server
	gRPCServerMu sync.Mutex

	updateServer *update.Server

	// stopped channel receives a message once the server is stopped.
	stopped chan struct{}
}

// New returns new Server.
func New(ctx context.Context, upgrader upgrader, gRPCMessageMaxSize int) *Server {
	l := logrus.WithField("component", "api-server/gRPC")

	updateServer := update.New(ctx, upgrader)

	return &Server{
		l:                  l,
		gRPCMessageMaxSize: gRPCMessageMaxSize,
		stopped:            make(chan struct{}),
		updateServer:       updateServer,
	}
}

// Start starts API server. If the server is already running, it's a no-op.
func (s *Server) Start(ctx context.Context) {
	s.gRPCServerMu.Lock()
	defer s.gRPCServerMu.Unlock()

	if s.gRPCServer != nil {
		s.l.Info("API server already running. Skipping start")
		return
	}

	s.gRPCServer = grpc.NewServer(
		grpc.MaxRecvMsgSize(s.gRPCMessageMaxSize),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)

	updatepb.RegisterStatusServer(s.gRPCServer, status.New())
	updatepb.RegisterUpdateServer(s.gRPCServer, s.updateServer)

	if s.EnableDebug {
		s.l.Debug("Reflection and channelz are enabled")
		reflection.Register(s.gRPCServer)
		channelz.RegisterChannelzServiceToServer(s.gRPCServer)
	}

	// run server until it is stopped gracefully
	go func() {
		var err error
		for {
			s.l.Infof("Starting gRPC server on unix://%s", SocketPath)
			err = os.Remove(SocketPath)
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				s.l.Panic(err)
			}

			var listener net.Listener
			listener, err = net.Listen("unix", SocketPath)
			if err != nil {
				s.l.Panic(err)
			}

			err = s.gRPCServer.Serve(listener) // listener will be closed when this method returns
			if err == nil || errors.Is(err, grpc.ErrServerStopped) {
				break
			}
		}
		if err != nil {
			s.l.Errorf("Failed to serve: %s", err)
			return
		}
		s.l.Debug("Server stopped")
	}()

	go func() {
		select {
		case <-s.stopped:
		case <-ctx.Done():
			s.Stop()
		}
	}()
}

// Stop tries to stop server gracefully or force the stop after a timeout.
func (s *Server) Stop() {
	s.gRPCServerMu.Lock()
	defer s.gRPCServerMu.Unlock()

	if s.gRPCServer == nil {
		return
	}

	stopped := make(chan struct{})
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	go func() {
		<-shutdownCtx.Done()
		s.gRPCServer.Stop()
		close(stopped)
	}()
	s.gRPCServer.GracefulStop()
	shutdownCancel()
	<-stopped

	s.gRPCServer = nil

	// Notify the server has been stopped
	select {
	case s.stopped <- struct{}{}:
	default:
	}
}