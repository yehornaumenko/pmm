// Code generated by go-swagger; DO NOT EDIT.

package artifacts

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

// ListArtifactsReader is a Reader for the ListArtifacts structure.
type ListArtifactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArtifactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArtifactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListArtifactsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListArtifactsOK creates a ListArtifactsOK with default headers values
func NewListArtifactsOK() *ListArtifactsOK {
	return &ListArtifactsOK{}
}

/*
ListArtifactsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListArtifactsOK struct {
	Payload *ListArtifactsOKBody
}

func (o *ListArtifactsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Artifacts/List][%d] listArtifactsOk  %+v", 200, o.Payload)
}

func (o *ListArtifactsOK) GetPayload() *ListArtifactsOKBody {
	return o.Payload
}

func (o *ListArtifactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListArtifactsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArtifactsDefault creates a ListArtifactsDefault with default headers values
func NewListArtifactsDefault(code int) *ListArtifactsDefault {
	return &ListArtifactsDefault{
		_statusCode: code,
	}
}

/*
ListArtifactsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListArtifactsDefault struct {
	_statusCode int

	Payload *ListArtifactsDefaultBody
}

// Code gets the status code for the list artifacts default response
func (o *ListArtifactsDefault) Code() int {
	return o._statusCode
}

func (o *ListArtifactsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Artifacts/List][%d] ListArtifacts default  %+v", o._statusCode, o.Payload)
}

func (o *ListArtifactsDefault) GetPayload() *ListArtifactsDefaultBody {
	return o.Payload
}

func (o *ListArtifactsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListArtifactsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListArtifactsDefaultBody list artifacts default body
swagger:model ListArtifactsDefaultBody
*/
type ListArtifactsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListArtifactsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list artifacts default body
func (o *ListArtifactsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListArtifacts default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListArtifacts default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list artifacts default body based on the context it is used
func (o *ListArtifactsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListArtifacts default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListArtifacts default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListArtifactsDefaultBodyDetailsItems0 list artifacts default body details items0
swagger:model ListArtifactsDefaultBodyDetailsItems0
*/
type ListArtifactsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list artifacts default body details items0
func (o *ListArtifactsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list artifacts default body details items0 based on context it is used
func (o *ListArtifactsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListArtifactsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListArtifactsOKBody list artifacts OK body
swagger:model ListArtifactsOKBody
*/
type ListArtifactsOKBody struct {
	// artifacts
	Artifacts []*ListArtifactsOKBodyArtifactsItems0 `json:"artifacts"`
}

// Validate validates this list artifacts OK body
func (o *ListArtifactsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateArtifacts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactsOKBody) validateArtifacts(formats strfmt.Registry) error {
	if swag.IsZero(o.Artifacts) { // not required
		return nil
	}

	for i := 0; i < len(o.Artifacts); i++ {
		if swag.IsZero(o.Artifacts[i]) { // not required
			continue
		}

		if o.Artifacts[i] != nil {
			if err := o.Artifacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listArtifactsOk" + "." + "artifacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listArtifactsOk" + "." + "artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list artifacts OK body based on the context it is used
func (o *ListArtifactsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateArtifacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactsOKBody) contextValidateArtifacts(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Artifacts); i++ {
		if o.Artifacts[i] != nil {
			if err := o.Artifacts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listArtifactsOk" + "." + "artifacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listArtifactsOk" + "." + "artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactsOKBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListArtifactsOKBodyArtifactsItems0 Artifact represents single backup artifact.
swagger:model ListArtifactsOKBodyArtifactsItems0
*/
type ListArtifactsOKBodyArtifactsItems0 struct {
	// Machine-readable artifact ID.
	ArtifactID string `json:"artifact_id,omitempty"`

	// Artifact name
	Name string `json:"name,omitempty"`

	// Database vendor e.g. PostgreSQL, MongoDB, MySQL.
	Vendor string `json:"vendor,omitempty"`

	// Machine-readable location ID.
	LocationID string `json:"location_id,omitempty"`

	// Location name.
	LocationName string `json:"location_name,omitempty"`

	// Machine-readable service ID.
	ServiceID string `json:"service_id,omitempty"`

	// Service name.
	ServiceName string `json:"service_name,omitempty"`

	// DataModel is a model used for performing a backup.
	// Enum: [DATA_MODEL_INVALID PHYSICAL LOGICAL]
	DataModel *string `json:"data_model,omitempty"`

	// BackupStatus shows the current status of execution of backup.
	// Enum: [BACKUP_STATUS_INVALID BACKUP_STATUS_PENDING BACKUP_STATUS_IN_PROGRESS BACKUP_STATUS_PAUSED BACKUP_STATUS_SUCCESS BACKUP_STATUS_ERROR BACKUP_STATUS_DELETING BACKUP_STATUS_FAILED_TO_DELETE]
	Status *string `json:"status,omitempty"`

	// Artifact creation time.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// BackupMode specifies backup mode.
	// Enum: [BACKUP_MODE_INVALID SNAPSHOT INCREMENTAL PITR]
	Mode *string `json:"mode,omitempty"`
}

// Validate validates this list artifacts OK body artifacts items0
func (o *ListArtifactsOKBodyArtifactsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDataModel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
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

var listArtifactsOkBodyArtifactsItems0TypeDataModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DATA_MODEL_INVALID","PHYSICAL","LOGICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listArtifactsOkBodyArtifactsItems0TypeDataModelPropEnum = append(listArtifactsOkBodyArtifactsItems0TypeDataModelPropEnum, v)
	}
}

const (

	// ListArtifactsOKBodyArtifactsItems0DataModelDATAMODELINVALID captures enum value "DATA_MODEL_INVALID"
	ListArtifactsOKBodyArtifactsItems0DataModelDATAMODELINVALID string = "DATA_MODEL_INVALID"

	// ListArtifactsOKBodyArtifactsItems0DataModelPHYSICAL captures enum value "PHYSICAL"
	ListArtifactsOKBodyArtifactsItems0DataModelPHYSICAL string = "PHYSICAL"

	// ListArtifactsOKBodyArtifactsItems0DataModelLOGICAL captures enum value "LOGICAL"
	ListArtifactsOKBodyArtifactsItems0DataModelLOGICAL string = "LOGICAL"
)

// prop value enum
func (o *ListArtifactsOKBodyArtifactsItems0) validateDataModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listArtifactsOkBodyArtifactsItems0TypeDataModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListArtifactsOKBodyArtifactsItems0) validateDataModel(formats strfmt.Registry) error {
	if swag.IsZero(o.DataModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateDataModelEnum("data_model", "body", *o.DataModel); err != nil {
		return err
	}

	return nil
}

var listArtifactsOkBodyArtifactsItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BACKUP_STATUS_INVALID","BACKUP_STATUS_PENDING","BACKUP_STATUS_IN_PROGRESS","BACKUP_STATUS_PAUSED","BACKUP_STATUS_SUCCESS","BACKUP_STATUS_ERROR","BACKUP_STATUS_DELETING","BACKUP_STATUS_FAILED_TO_DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listArtifactsOkBodyArtifactsItems0TypeStatusPropEnum = append(listArtifactsOkBodyArtifactsItems0TypeStatusPropEnum, v)
	}
}

const (

	// ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSINVALID captures enum value "BACKUP_STATUS_INVALID"
	ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSINVALID string = "BACKUP_STATUS_INVALID"

	// ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSPENDING captures enum value "BACKUP_STATUS_PENDING"
	ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSPENDING string = "BACKUP_STATUS_PENDING"

	// ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSINPROGRESS captures enum value "BACKUP_STATUS_IN_PROGRESS"
	ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSINPROGRESS string = "BACKUP_STATUS_IN_PROGRESS"

	// ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSPAUSED captures enum value "BACKUP_STATUS_PAUSED"
	ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSPAUSED string = "BACKUP_STATUS_PAUSED"

	// ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSSUCCESS captures enum value "BACKUP_STATUS_SUCCESS"
	ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSSUCCESS string = "BACKUP_STATUS_SUCCESS"

	// ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSERROR captures enum value "BACKUP_STATUS_ERROR"
	ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSERROR string = "BACKUP_STATUS_ERROR"

	// ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSDELETING captures enum value "BACKUP_STATUS_DELETING"
	ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSDELETING string = "BACKUP_STATUS_DELETING"

	// ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSFAILEDTODELETE captures enum value "BACKUP_STATUS_FAILED_TO_DELETE"
	ListArtifactsOKBodyArtifactsItems0StatusBACKUPSTATUSFAILEDTODELETE string = "BACKUP_STATUS_FAILED_TO_DELETE"
)

// prop value enum
func (o *ListArtifactsOKBodyArtifactsItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listArtifactsOkBodyArtifactsItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListArtifactsOKBodyArtifactsItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

func (o *ListArtifactsOKBodyArtifactsItems0) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var listArtifactsOkBodyArtifactsItems0TypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BACKUP_MODE_INVALID","SNAPSHOT","INCREMENTAL","PITR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listArtifactsOkBodyArtifactsItems0TypeModePropEnum = append(listArtifactsOkBodyArtifactsItems0TypeModePropEnum, v)
	}
}

const (

	// ListArtifactsOKBodyArtifactsItems0ModeBACKUPMODEINVALID captures enum value "BACKUP_MODE_INVALID"
	ListArtifactsOKBodyArtifactsItems0ModeBACKUPMODEINVALID string = "BACKUP_MODE_INVALID"

	// ListArtifactsOKBodyArtifactsItems0ModeSNAPSHOT captures enum value "SNAPSHOT"
	ListArtifactsOKBodyArtifactsItems0ModeSNAPSHOT string = "SNAPSHOT"

	// ListArtifactsOKBodyArtifactsItems0ModeINCREMENTAL captures enum value "INCREMENTAL"
	ListArtifactsOKBodyArtifactsItems0ModeINCREMENTAL string = "INCREMENTAL"

	// ListArtifactsOKBodyArtifactsItems0ModePITR captures enum value "PITR"
	ListArtifactsOKBodyArtifactsItems0ModePITR string = "PITR"
)

// prop value enum
func (o *ListArtifactsOKBodyArtifactsItems0) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listArtifactsOkBodyArtifactsItems0TypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListArtifactsOKBodyArtifactsItems0) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(o.Mode) { // not required
		return nil
	}

	// value enum
	if err := o.validateModeEnum("mode", "body", *o.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list artifacts OK body artifacts items0 based on context it is used
func (o *ListArtifactsOKBodyArtifactsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactsOKBodyArtifactsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactsOKBodyArtifactsItems0) UnmarshalBinary(b []byte) error {
	var res ListArtifactsOKBodyArtifactsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
