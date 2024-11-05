// Code generated by go-swagger; DO NOT EDIT.

package dump_service

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

// StartDumpReader is a Reader for the StartDump structure.
type StartDumpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDumpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartDumpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartDumpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartDumpOK creates a StartDumpOK with default headers values
func NewStartDumpOK() *StartDumpOK {
	return &StartDumpOK{}
}

/*
StartDumpOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartDumpOK struct {
	Payload *StartDumpOKBody
}

// IsSuccess returns true when this start dump Ok response has a 2xx status code
func (o *StartDumpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start dump Ok response has a 3xx status code
func (o *StartDumpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start dump Ok response has a 4xx status code
func (o *StartDumpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start dump Ok response has a 5xx status code
func (o *StartDumpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start dump Ok response a status code equal to that given
func (o *StartDumpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start dump Ok response
func (o *StartDumpOK) Code() int {
	return 200
}

func (o *StartDumpOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/dumps:start][%d] startDumpOk %s", 200, payload)
}

func (o *StartDumpOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/dumps:start][%d] startDumpOk %s", 200, payload)
}

func (o *StartDumpOK) GetPayload() *StartDumpOKBody {
	return o.Payload
}

func (o *StartDumpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartDumpOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDumpDefault creates a StartDumpDefault with default headers values
func NewStartDumpDefault(code int) *StartDumpDefault {
	return &StartDumpDefault{
		_statusCode: code,
	}
}

/*
StartDumpDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartDumpDefault struct {
	_statusCode int

	Payload *StartDumpDefaultBody
}

// IsSuccess returns true when this start dump default response has a 2xx status code
func (o *StartDumpDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this start dump default response has a 3xx status code
func (o *StartDumpDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this start dump default response has a 4xx status code
func (o *StartDumpDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this start dump default response has a 5xx status code
func (o *StartDumpDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this start dump default response a status code equal to that given
func (o *StartDumpDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the start dump default response
func (o *StartDumpDefault) Code() int {
	return o._statusCode
}

func (o *StartDumpDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/dumps:start][%d] StartDump default %s", o._statusCode, payload)
}

func (o *StartDumpDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/dumps:start][%d] StartDump default %s", o._statusCode, payload)
}

func (o *StartDumpDefault) GetPayload() *StartDumpDefaultBody {
	return o.Payload
}

func (o *StartDumpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartDumpDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
StartDumpBody start dump body
swagger:model StartDumpBody
*/
type StartDumpBody struct {
	// service names
	ServiceNames []string `json:"service_names"`

	// start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// end time
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// export qan
	ExportQAN bool `json:"export_qan,omitempty"`

	// ignore load
	IgnoreLoad bool `json:"ignore_load,omitempty"`
}

// Validate validates this start dump body
func (o *StartDumpBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartDumpBody) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"start_time", "body", "date-time", o.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *StartDumpBody) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(o.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"end_time", "body", "date-time", o.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start dump body based on context it is used
func (o *StartDumpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartDumpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartDumpBody) UnmarshalBinary(b []byte) error {
	var res StartDumpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartDumpDefaultBody start dump default body
swagger:model StartDumpDefaultBody
*/
type StartDumpDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartDumpDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start dump default body
func (o *StartDumpDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartDumpDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("StartDump default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartDump default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start dump default body based on the context it is used
func (o *StartDumpDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartDumpDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartDump default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartDump default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartDumpDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartDumpDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartDumpDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartDumpDefaultBodyDetailsItems0 start dump default body details items0
swagger:model StartDumpDefaultBodyDetailsItems0
*/
type StartDumpDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// start dump default body details items0
	StartDumpDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *StartDumpDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv StartDumpDefaultBodyDetailsItems0

	rcv.AtType = stage1.AtType
	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "@type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		o.StartDumpDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o StartDumpDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}

	stage1.AtType = o.AtType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.StartDumpDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.StartDumpDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this start dump default body details items0
func (o *StartDumpDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start dump default body details items0 based on context it is used
func (o *StartDumpDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartDumpDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartDumpDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartDumpDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartDumpOKBody start dump OK body
swagger:model StartDumpOKBody
*/
type StartDumpOKBody struct {
	// dump id
	DumpID string `json:"dump_id,omitempty"`
}

// Validate validates this start dump OK body
func (o *StartDumpOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start dump OK body based on context it is used
func (o *StartDumpOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartDumpOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartDumpOKBody) UnmarshalBinary(b []byte) error {
	var res StartDumpOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}