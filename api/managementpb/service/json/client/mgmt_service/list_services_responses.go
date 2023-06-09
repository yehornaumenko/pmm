// Code generated by go-swagger; DO NOT EDIT.

package mgmt_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListServicesReader is a Reader for the ListServices structure.
type ListServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListServicesOK creates a ListServicesOK with default headers values
func NewListServicesOK() *ListServicesOK {
	return &ListServicesOK{}
}

/*
ListServicesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListServicesOK struct {
	Payload *ListServicesOKBody
}

func (o *ListServicesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Service/List][%d] listServicesOk  %+v", 200, o.Payload)
}

func (o *ListServicesOK) GetPayload() *ListServicesOKBody {
	return o.Payload
}

func (o *ListServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesDefault creates a ListServicesDefault with default headers values
func NewListServicesDefault(code int) *ListServicesDefault {
	return &ListServicesDefault{
		_statusCode: code,
	}
}

/*
ListServicesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListServicesDefault struct {
	_statusCode int

	Payload *ListServicesDefaultBody
}

// Code gets the status code for the list services default response
func (o *ListServicesDefault) Code() int {
	return o._statusCode
}

func (o *ListServicesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Service/List][%d] ListServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListServicesDefault) GetPayload() *ListServicesDefaultBody {
	return o.Payload
}

func (o *ListServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListServicesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListServicesBody list services body
swagger:model ListServicesBody
*/
type ListServicesBody struct {
	// Return only Services running on that Node.
	NodeID string `json:"node_id,omitempty"`

	// ServiceType describes supported Service types.
	// Enum: [SERVICE_TYPE_INVALID MYSQL_SERVICE MONGODB_SERVICE POSTGRESQL_SERVICE PROXYSQL_SERVICE HAPROXY_SERVICE EXTERNAL_SERVICE]
	ServiceType *string `json:"service_type,omitempty"`

	// Return only services in this external group.
	ExternalGroup string `json:"external_group,omitempty"`
}

// Validate validates this list services body
func (o *ListServicesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listServicesBodyTypeServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SERVICE_TYPE_INVALID","MYSQL_SERVICE","MONGODB_SERVICE","POSTGRESQL_SERVICE","PROXYSQL_SERVICE","HAPROXY_SERVICE","EXTERNAL_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listServicesBodyTypeServiceTypePropEnum = append(listServicesBodyTypeServiceTypePropEnum, v)
	}
}

const (

	// ListServicesBodyServiceTypeSERVICETYPEINVALID captures enum value "SERVICE_TYPE_INVALID"
	ListServicesBodyServiceTypeSERVICETYPEINVALID string = "SERVICE_TYPE_INVALID"

	// ListServicesBodyServiceTypeMYSQLSERVICE captures enum value "MYSQL_SERVICE"
	ListServicesBodyServiceTypeMYSQLSERVICE string = "MYSQL_SERVICE"

	// ListServicesBodyServiceTypeMONGODBSERVICE captures enum value "MONGODB_SERVICE"
	ListServicesBodyServiceTypeMONGODBSERVICE string = "MONGODB_SERVICE"

	// ListServicesBodyServiceTypePOSTGRESQLSERVICE captures enum value "POSTGRESQL_SERVICE"
	ListServicesBodyServiceTypePOSTGRESQLSERVICE string = "POSTGRESQL_SERVICE"

	// ListServicesBodyServiceTypePROXYSQLSERVICE captures enum value "PROXYSQL_SERVICE"
	ListServicesBodyServiceTypePROXYSQLSERVICE string = "PROXYSQL_SERVICE"

	// ListServicesBodyServiceTypeHAPROXYSERVICE captures enum value "HAPROXY_SERVICE"
	ListServicesBodyServiceTypeHAPROXYSERVICE string = "HAPROXY_SERVICE"

	// ListServicesBodyServiceTypeEXTERNALSERVICE captures enum value "EXTERNAL_SERVICE"
	ListServicesBodyServiceTypeEXTERNALSERVICE string = "EXTERNAL_SERVICE"
)

// prop value enum
func (o *ListServicesBody) validateServiceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listServicesBodyTypeServiceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListServicesBody) validateServiceType(formats strfmt.Registry) error {
	if swag.IsZero(o.ServiceType) { // not required
		return nil
	}

	// value enum
	if err := o.validateServiceTypeEnum("body"+"."+"service_type", "body", *o.ServiceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list services body based on context it is used
func (o *ListServicesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesBody) UnmarshalBinary(b []byte) error {
	var res ListServicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListServicesDefaultBody list services default body
swagger:model ListServicesDefaultBody
*/
type ListServicesDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListServicesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list services default body
func (o *ListServicesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListServices default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list services default body based on the context it is used
func (o *ListServicesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListServices default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListServicesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListServicesDefaultBodyDetailsItems0 list services default body details items0
swagger:model ListServicesDefaultBodyDetailsItems0
*/
type ListServicesDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list services default body details items0
func (o *ListServicesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services default body details items0 based on context it is used
func (o *ListServicesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListServicesOKBody list services OK body
swagger:model ListServicesOKBody
*/
type ListServicesOKBody struct {
	// List of Services.
	Services []*ListServicesOKBodyServicesItems0 `json:"services"`
}

// Validate validates this list services OK body
func (o *ListServicesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesOKBody) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(o.Services) { // not required
		return nil
	}

	for i := 0; i < len(o.Services); i++ {
		if swag.IsZero(o.Services[i]) { // not required
			continue
		}

		if o.Services[i] != nil {
			if err := o.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list services OK body based on the context it is used
func (o *ListServicesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesOKBody) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Services); i++ {
		if o.Services[i] != nil {
			if err := o.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBody) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListServicesOKBodyServicesItems0 list services OK body services items0
swagger:model ListServicesOKBodyServicesItems0
*/
type ListServicesOKBodyServicesItems0 struct {
	// Unique service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Service type.
	ServiceType string `json:"service_type,omitempty"`

	// User-defined name unique across all Services.
	ServiceName string `json:"service_name,omitempty"`

	// Database name.
	DatabaseName string `json:"database_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Node name where this instance runs.
	NodeName string `json:"node_name,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// External group name.
	ExternalGroup string `json:"external_group,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Creation timestamp.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Last update timestamp.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// List of agents related to this service.
	Agents []*ListServicesOKBodyServicesItems0AgentsItems0 `json:"agents"`

	// Service status.
	//
	//  - STATUS_INVALID: In case we don't support the db vendor yet.
	//  - UP: The service is up.
	//  - DOWN: The service is down.
	//  - UNKNOWN: The service's status cannot be known (e.g. there are no metrics yet).
	// Enum: [STATUS_INVALID UP DOWN UNKNOWN]
	Status *string `json:"status,omitempty"`
}

// Validate validates this list services OK body services items0
func (o *ListServicesOKBodyServicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesOKBodyServicesItems0) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0) validateAgents(formats strfmt.Registry) error {
	if swag.IsZero(o.Agents) { // not required
		return nil
	}

	for i := 0; i < len(o.Agents); i++ {
		if swag.IsZero(o.Agents[i]) { // not required
			continue
		}

		if o.Agents[i] != nil {
			if err := o.Agents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var listServicesOkBodyServicesItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATUS_INVALID","UP","DOWN","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listServicesOkBodyServicesItems0TypeStatusPropEnum = append(listServicesOkBodyServicesItems0TypeStatusPropEnum, v)
	}
}

const (

	// ListServicesOKBodyServicesItems0StatusSTATUSINVALID captures enum value "STATUS_INVALID"
	ListServicesOKBodyServicesItems0StatusSTATUSINVALID string = "STATUS_INVALID"

	// ListServicesOKBodyServicesItems0StatusUP captures enum value "UP"
	ListServicesOKBodyServicesItems0StatusUP string = "UP"

	// ListServicesOKBodyServicesItems0StatusDOWN captures enum value "DOWN"
	ListServicesOKBodyServicesItems0StatusDOWN string = "DOWN"

	// ListServicesOKBodyServicesItems0StatusUNKNOWN captures enum value "UNKNOWN"
	ListServicesOKBodyServicesItems0StatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *ListServicesOKBodyServicesItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listServicesOkBodyServicesItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListServicesOKBodyServicesItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this list services OK body services items0 based on the context it is used
func (o *ListServicesOKBodyServicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAgents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesOKBodyServicesItems0) contextValidateAgents(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Agents); i++ {
		if o.Agents[i] != nil {
			if err := o.Agents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyServicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListServicesOKBodyServicesItems0AgentsItems0 list services OK body services items0 agents items0
swagger:model ListServicesOKBodyServicesItems0AgentsItems0
*/
type ListServicesOKBodyServicesItems0AgentsItems0 struct {
	// Unique agent identifier.
	AgentID string `json:"agent_id,omitempty"`

	// True if the agent password is set.
	IsAgentPasswordSet bool `json:"is_agent_password_set,omitempty"`

	// Agent type.
	AgentType string `json:"agent_type,omitempty"`

	// AWS Access Key.
	AWSAccessKey string `json:"aws_access_key,omitempty"`

	// True if AWS Secret Key is set.
	IsAWSSecretKeySet bool `json:"is_aws_secret_key_set,omitempty"`

	// Creation timestamp.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// List of disabled collector names.
	DisabledCollectors []string `json:"disabled_collectors"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// Log level for exporter.
	LogLevel string `json:"log_level,omitempty"`

	// Limit query length in QAN.
	MaxQueryLength int32 `json:"max_query_length,omitempty"`

	// Limit query log size in QAN.
	MaxQueryLogSize string `json:"max_query_log_size,omitempty"`

	// Path under which metrics are exposed, used to generate URI.
	MetricsPath string `json:"metrics_path,omitempty"`

	// Scheme to generate URI to exporter metrics endpoints.
	MetricsScheme string `json:"metrics_scheme,omitempty"`

	// A unique node identifier.
	NodeID string `json:"node_id,omitempty"`

	// True if password for connecting the agent to the database is set.
	IsPasswordSet bool `json:"is_password_set,omitempty"`

	// The pmm-agent identifier.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetrics bool `json:"push_metrics,omitempty"`

	// True if query examples are disabled.
	QueryExamplesDisabled bool `json:"query_examples_disabled,omitempty"`

	// True if RDS basic metrics are disdabled.
	RDSBasicMetricsDisabled bool `json:"rds_basic_metrics_disabled,omitempty"`

	// True if RDS enhanced metrics are disdabled.
	RDSEnhancedMetricsDisabled bool `json:"rds_enhanced_metrics_disabled,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Actual Agent status.
	Status string `json:"status,omitempty"`

	// Last known table count.
	TableCount int32 `json:"table_count,omitempty"`

	// Tablestats group collectors are disabled if there are more than that number of tables.
	// 0 means tablestats group collectors are always enabled (no limit).
	// Negative value means tablestats group collectors are always disabled.
	TableCountTablestatsGroupLimit int32 `json:"table_count_tablestats_group_limit,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// HTTP basic auth username for collecting metrics.
	Username string `json:"username,omitempty"`

	// Last update timestamp.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// Agent version.
	Version string `json:"version,omitempty"`

	// True if Agent is running and connected to pmm-managed.
	IsConnected bool `json:"is_connected,omitempty"`

	// azure options
	AzureOptions *ListServicesOKBodyServicesItems0AgentsItems0AzureOptions `json:"azure_options,omitempty"`

	// mongo db options
	MongoDBOptions *ListServicesOKBodyServicesItems0AgentsItems0MongoDBOptions `json:"mongo_db_options,omitempty"`

	// mysql options
	MysqlOptions *ListServicesOKBodyServicesItems0AgentsItems0MysqlOptions `json:"mysql_options,omitempty"`

	// postgresql options
	PostgresqlOptions *ListServicesOKBodyServicesItems0AgentsItems0PostgresqlOptions `json:"postgresql_options,omitempty"`
}

// Validate validates this list services OK body services items0 agents items0
func (o *ListServicesOKBodyServicesItems0AgentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAzureOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongoDBOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMysqlOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePostgresqlOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) validateAzureOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.AzureOptions) { // not required
		return nil
	}

	if o.AzureOptions != nil {
		if err := o.AzureOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) validateMongoDBOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.MongoDBOptions) { // not required
		return nil
	}

	if o.MongoDBOptions != nil {
		if err := o.MongoDBOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongo_db_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongo_db_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) validateMysqlOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.MysqlOptions) { // not required
		return nil
	}

	if o.MysqlOptions != nil {
		if err := o.MysqlOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mysql_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mysql_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) validatePostgresqlOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.PostgresqlOptions) { // not required
		return nil
	}

	if o.PostgresqlOptions != nil {
		if err := o.PostgresqlOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postgresql_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postgresql_options")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list services OK body services items0 agents items0 based on the context it is used
func (o *ListServicesOKBodyServicesItems0AgentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAzureOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMongoDBOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMysqlOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePostgresqlOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) contextValidateAzureOptions(ctx context.Context, formats strfmt.Registry) error {
	if o.AzureOptions != nil {
		if err := o.AzureOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) contextValidateMongoDBOptions(ctx context.Context, formats strfmt.Registry) error {
	if o.MongoDBOptions != nil {
		if err := o.MongoDBOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongo_db_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongo_db_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) contextValidateMysqlOptions(ctx context.Context, formats strfmt.Registry) error {
	if o.MysqlOptions != nil {
		if err := o.MysqlOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mysql_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mysql_options")
			}
			return err
		}
	}

	return nil
}

func (o *ListServicesOKBodyServicesItems0AgentsItems0) contextValidatePostgresqlOptions(ctx context.Context, formats strfmt.Registry) error {
	if o.PostgresqlOptions != nil {
		if err := o.PostgresqlOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postgresql_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postgresql_options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyServicesItems0AgentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListServicesOKBodyServicesItems0AgentsItems0AzureOptions list services OK body services items0 agents items0 azure options
swagger:model ListServicesOKBodyServicesItems0AgentsItems0AzureOptions
*/
type ListServicesOKBodyServicesItems0AgentsItems0AzureOptions struct {
	// Azure client ID.
	ClientID string `json:"client_id,omitempty"`

	// True if Azure client secret is set.
	IsClientSecretSet bool `json:"is_client_secret_set,omitempty"`

	// Azure resource group.
	ResourceGroup string `json:"resource_group,omitempty"`

	// Azure subscription ID.
	SubscriptionID string `json:"subscription_id,omitempty"`

	// Azure tenant ID.
	TenantID string `json:"tenant_id,omitempty"`
}

// Validate validates this list services OK body services items0 agents items0 azure options
func (o *ListServicesOKBodyServicesItems0AgentsItems0AzureOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body services items0 agents items0 azure options based on context it is used
func (o *ListServicesOKBodyServicesItems0AgentsItems0AzureOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0AzureOptions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0AzureOptions) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyServicesItems0AgentsItems0AzureOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListServicesOKBodyServicesItems0AgentsItems0MongoDBOptions list services OK body services items0 agents items0 mongo DB options
swagger:model ListServicesOKBodyServicesItems0AgentsItems0MongoDBOptions
*/
type ListServicesOKBodyServicesItems0AgentsItems0MongoDBOptions struct {
	// True if TLS certificate is set.
	IsTLSCertificateKeySet bool `json:"is_tls_certificate_key_set,omitempty"`

	// True if TLS certificate file password is set.
	IsTLSCertificateKeyFilePasswordSet bool `json:"is_tls_certificate_key_file_password_set,omitempty"`

	// TLS CA certificate.
	TLSCa string `json:"tls_ca,omitempty"`

	// MongoDB auth mechanism.
	AuthenticationMechanism string `json:"authentication_mechanism,omitempty"`

	// MongoDB auth database.
	AuthenticationDatabase string `json:"authentication_database,omitempty"`

	// MongoDB stats collections.
	StatsCollections []string `json:"stats_collections"`

	// MongoDB collections limit.
	CollectionsLimit int32 `json:"collections_limit,omitempty"`

	// True if all collectors are enabled.
	EnableAllCollectors bool `json:"enable_all_collectors,omitempty"`
}

// Validate validates this list services OK body services items0 agents items0 mongo DB options
func (o *ListServicesOKBodyServicesItems0AgentsItems0MongoDBOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body services items0 agents items0 mongo DB options based on context it is used
func (o *ListServicesOKBodyServicesItems0AgentsItems0MongoDBOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0MongoDBOptions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0MongoDBOptions) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyServicesItems0AgentsItems0MongoDBOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListServicesOKBodyServicesItems0AgentsItems0MysqlOptions list services OK body services items0 agents items0 mysql options
swagger:model ListServicesOKBodyServicesItems0AgentsItems0MysqlOptions
*/
type ListServicesOKBodyServicesItems0AgentsItems0MysqlOptions struct {
	// TLS CA certificate.
	TLSCa string `json:"tls_ca,omitempty"`

	// TLS certificate.
	TLSCert string `json:"tls_cert,omitempty"`

	// True if TLS key is set.
	IsTLSKeySet bool `json:"is_tls_key_set,omitempty"`
}

// Validate validates this list services OK body services items0 agents items0 mysql options
func (o *ListServicesOKBodyServicesItems0AgentsItems0MysqlOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body services items0 agents items0 mysql options based on context it is used
func (o *ListServicesOKBodyServicesItems0AgentsItems0MysqlOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0MysqlOptions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0MysqlOptions) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyServicesItems0AgentsItems0MysqlOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListServicesOKBodyServicesItems0AgentsItems0PostgresqlOptions list services OK body services items0 agents items0 postgresql options
swagger:model ListServicesOKBodyServicesItems0AgentsItems0PostgresqlOptions
*/
type ListServicesOKBodyServicesItems0AgentsItems0PostgresqlOptions struct {
	// TLS CA certificate.
	SslCa string `json:"ssl_ca,omitempty"`

	// TLS certificate.
	SslCert string `json:"ssl_cert,omitempty"`

	// True if TLS key is set.
	IsSslKeySet bool `json:"is_ssl_key_set,omitempty"`
}

// Validate validates this list services OK body services items0 agents items0 postgresql options
func (o *ListServicesOKBodyServicesItems0AgentsItems0PostgresqlOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body services items0 agents items0 postgresql options based on context it is used
func (o *ListServicesOKBodyServicesItems0AgentsItems0PostgresqlOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0PostgresqlOptions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyServicesItems0AgentsItems0PostgresqlOptions) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyServicesItems0AgentsItems0PostgresqlOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
