// Code generated by go-swagger; DO NOT EDIT.

package backups

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

// ListScheduledBackupsReader is a Reader for the ListScheduledBackups structure.
type ListScheduledBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListScheduledBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListScheduledBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListScheduledBackupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListScheduledBackupsOK creates a ListScheduledBackupsOK with default headers values
func NewListScheduledBackupsOK() *ListScheduledBackupsOK {
	return &ListScheduledBackupsOK{}
}

/*
ListScheduledBackupsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListScheduledBackupsOK struct {
	Payload *ListScheduledBackupsOKBody
}

func (o *ListScheduledBackupsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/ListScheduled][%d] listScheduledBackupsOk  %+v", 200, o.Payload)
}

func (o *ListScheduledBackupsOK) GetPayload() *ListScheduledBackupsOKBody {
	return o.Payload
}

func (o *ListScheduledBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListScheduledBackupsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScheduledBackupsDefault creates a ListScheduledBackupsDefault with default headers values
func NewListScheduledBackupsDefault(code int) *ListScheduledBackupsDefault {
	return &ListScheduledBackupsDefault{
		_statusCode: code,
	}
}

/*
ListScheduledBackupsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListScheduledBackupsDefault struct {
	_statusCode int

	Payload *ListScheduledBackupsDefaultBody
}

// Code gets the status code for the list scheduled backups default response
func (o *ListScheduledBackupsDefault) Code() int {
	return o._statusCode
}

func (o *ListScheduledBackupsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/ListScheduled][%d] ListScheduledBackups default  %+v", o._statusCode, o.Payload)
}

func (o *ListScheduledBackupsDefault) GetPayload() *ListScheduledBackupsDefaultBody {
	return o.Payload
}

func (o *ListScheduledBackupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListScheduledBackupsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListScheduledBackupsDefaultBody list scheduled backups default body
swagger:model ListScheduledBackupsDefaultBody
*/
type ListScheduledBackupsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListScheduledBackupsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list scheduled backups default body
func (o *ListScheduledBackupsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListScheduledBackupsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListScheduledBackups default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListScheduledBackups default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list scheduled backups default body based on the context it is used
func (o *ListScheduledBackupsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListScheduledBackupsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListScheduledBackups default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListScheduledBackups default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListScheduledBackupsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListScheduledBackupsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListScheduledBackupsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListScheduledBackupsDefaultBodyDetailsItems0 list scheduled backups default body details items0
swagger:model ListScheduledBackupsDefaultBodyDetailsItems0
*/
type ListScheduledBackupsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list scheduled backups default body details items0
func (o *ListScheduledBackupsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list scheduled backups default body details items0 based on context it is used
func (o *ListScheduledBackupsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListScheduledBackupsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListScheduledBackupsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListScheduledBackupsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListScheduledBackupsOKBody list scheduled backups OK body
swagger:model ListScheduledBackupsOKBody
*/
type ListScheduledBackupsOKBody struct {
	// scheduled backups
	ScheduledBackups []*ListScheduledBackupsOKBodyScheduledBackupsItems0 `json:"scheduled_backups"`
}

// Validate validates this list scheduled backups OK body
func (o *ListScheduledBackupsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateScheduledBackups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListScheduledBackupsOKBody) validateScheduledBackups(formats strfmt.Registry) error {
	if swag.IsZero(o.ScheduledBackups) { // not required
		return nil
	}

	for i := 0; i < len(o.ScheduledBackups); i++ {
		if swag.IsZero(o.ScheduledBackups[i]) { // not required
			continue
		}

		if o.ScheduledBackups[i] != nil {
			if err := o.ScheduledBackups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listScheduledBackupsOk" + "." + "scheduled_backups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listScheduledBackupsOk" + "." + "scheduled_backups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list scheduled backups OK body based on the context it is used
func (o *ListScheduledBackupsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateScheduledBackups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListScheduledBackupsOKBody) contextValidateScheduledBackups(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.ScheduledBackups); i++ {
		if o.ScheduledBackups[i] != nil {
			if err := o.ScheduledBackups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listScheduledBackupsOk" + "." + "scheduled_backups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listScheduledBackupsOk" + "." + "scheduled_backups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListScheduledBackupsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListScheduledBackupsOKBody) UnmarshalBinary(b []byte) error {
	var res ListScheduledBackupsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListScheduledBackupsOKBodyScheduledBackupsItems0 ScheduledBackup represents scheduled task for backup.
swagger:model ListScheduledBackupsOKBodyScheduledBackupsItems0
*/
type ListScheduledBackupsOKBodyScheduledBackupsItems0 struct {
	// Machine-readable ID.
	ScheduledBackupID string `json:"scheduled_backup_id,omitempty"`

	// Machine-readable service ID.
	ServiceID string `json:"service_id,omitempty"`

	// Service name.
	ServiceName string `json:"service_name,omitempty"`

	// Machine-readable location ID.
	LocationID string `json:"location_id,omitempty"`

	// Location name.
	LocationName string `json:"location_name,omitempty"`

	// How often backup will be run in cron format.
	CronExpression string `json:"cron_expression,omitempty"`

	// First backup wouldn't happen before this time.
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// Artifact name.
	Name string `json:"name,omitempty"`

	// Description.
	Description string `json:"description,omitempty"`

	// Delay between each retry. Should have a suffix in JSON: 1s, 1m, 1h.
	RetryInterval string `json:"retry_interval,omitempty"`

	// How many times to retry a failed backup before giving up.
	Retries int64 `json:"retries,omitempty"`

	// If scheduling is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// DataModel is a model used for performing a backup.
	// Enum: [DATA_MODEL_INVALID PHYSICAL LOGICAL]
	DataModel *string `json:"data_model,omitempty"`

	// Database vendor e.g. PostgreSQL, MongoDB, MySQL.
	Vendor string `json:"vendor,omitempty"`

	// Last run.
	// Format: date-time
	LastRun strfmt.DateTime `json:"last_run,omitempty"`

	// Next run.
	// Format: date-time
	NextRun strfmt.DateTime `json:"next_run,omitempty"`

	// How many artifacts keep. 0 - unlimited.
	Retention int64 `json:"retention,omitempty"`

	// BackupMode specifies backup mode.
	// Enum: [BACKUP_MODE_INVALID SNAPSHOT INCREMENTAL PITR]
	Mode *string `json:"mode,omitempty"`
}

// Validate validates this list scheduled backups OK body scheduled backups items0
func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDataModel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastRun(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNextRun(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", o.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var listScheduledBackupsOkBodyScheduledBackupsItems0TypeDataModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DATA_MODEL_INVALID","PHYSICAL","LOGICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listScheduledBackupsOkBodyScheduledBackupsItems0TypeDataModelPropEnum = append(listScheduledBackupsOkBodyScheduledBackupsItems0TypeDataModelPropEnum, v)
	}
}

const (

	// ListScheduledBackupsOKBodyScheduledBackupsItems0DataModelDATAMODELINVALID captures enum value "DATA_MODEL_INVALID"
	ListScheduledBackupsOKBodyScheduledBackupsItems0DataModelDATAMODELINVALID string = "DATA_MODEL_INVALID"

	// ListScheduledBackupsOKBodyScheduledBackupsItems0DataModelPHYSICAL captures enum value "PHYSICAL"
	ListScheduledBackupsOKBodyScheduledBackupsItems0DataModelPHYSICAL string = "PHYSICAL"

	// ListScheduledBackupsOKBodyScheduledBackupsItems0DataModelLOGICAL captures enum value "LOGICAL"
	ListScheduledBackupsOKBodyScheduledBackupsItems0DataModelLOGICAL string = "LOGICAL"
)

// prop value enum
func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) validateDataModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listScheduledBackupsOkBodyScheduledBackupsItems0TypeDataModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) validateDataModel(formats strfmt.Registry) error {
	if swag.IsZero(o.DataModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateDataModelEnum("data_model", "body", *o.DataModel); err != nil {
		return err
	}

	return nil
}

func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) validateLastRun(formats strfmt.Registry) error {
	if swag.IsZero(o.LastRun) { // not required
		return nil
	}

	if err := validate.FormatOf("last_run", "body", "date-time", o.LastRun.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) validateNextRun(formats strfmt.Registry) error {
	if swag.IsZero(o.NextRun) { // not required
		return nil
	}

	if err := validate.FormatOf("next_run", "body", "date-time", o.NextRun.String(), formats); err != nil {
		return err
	}

	return nil
}

var listScheduledBackupsOkBodyScheduledBackupsItems0TypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BACKUP_MODE_INVALID","SNAPSHOT","INCREMENTAL","PITR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listScheduledBackupsOkBodyScheduledBackupsItems0TypeModePropEnum = append(listScheduledBackupsOkBodyScheduledBackupsItems0TypeModePropEnum, v)
	}
}

const (

	// ListScheduledBackupsOKBodyScheduledBackupsItems0ModeBACKUPMODEINVALID captures enum value "BACKUP_MODE_INVALID"
	ListScheduledBackupsOKBodyScheduledBackupsItems0ModeBACKUPMODEINVALID string = "BACKUP_MODE_INVALID"

	// ListScheduledBackupsOKBodyScheduledBackupsItems0ModeSNAPSHOT captures enum value "SNAPSHOT"
	ListScheduledBackupsOKBodyScheduledBackupsItems0ModeSNAPSHOT string = "SNAPSHOT"

	// ListScheduledBackupsOKBodyScheduledBackupsItems0ModeINCREMENTAL captures enum value "INCREMENTAL"
	ListScheduledBackupsOKBodyScheduledBackupsItems0ModeINCREMENTAL string = "INCREMENTAL"

	// ListScheduledBackupsOKBodyScheduledBackupsItems0ModePITR captures enum value "PITR"
	ListScheduledBackupsOKBodyScheduledBackupsItems0ModePITR string = "PITR"
)

// prop value enum
func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listScheduledBackupsOkBodyScheduledBackupsItems0TypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(o.Mode) { // not required
		return nil
	}

	// value enum
	if err := o.validateModeEnum("mode", "body", *o.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list scheduled backups OK body scheduled backups items0 based on context it is used
func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListScheduledBackupsOKBodyScheduledBackupsItems0) UnmarshalBinary(b []byte) error {
	var res ListScheduledBackupsOKBodyScheduledBackupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
