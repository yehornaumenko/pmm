// Code generated by go-swagger; DO NOT EDIT.

package actions_service

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
)

// StartPTSummaryActionReader is a Reader for the StartPTSummaryAction structure.
type StartPTSummaryActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPTSummaryActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartPTSummaryActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartPTSummaryActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartPTSummaryActionOK creates a StartPTSummaryActionOK with default headers values
func NewStartPTSummaryActionOK() *StartPTSummaryActionOK {
	return &StartPTSummaryActionOK{}
}

/*
StartPTSummaryActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartPTSummaryActionOK struct {
	Payload *StartPTSummaryActionOKBody
}

// IsSuccess returns true when this start Pt summary action Ok response has a 2xx status code
func (o *StartPTSummaryActionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start Pt summary action Ok response has a 3xx status code
func (o *StartPTSummaryActionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start Pt summary action Ok response has a 4xx status code
func (o *StartPTSummaryActionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start Pt summary action Ok response has a 5xx status code
func (o *StartPTSummaryActionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start Pt summary action Ok response a status code equal to that given
func (o *StartPTSummaryActionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start Pt summary action Ok response
func (o *StartPTSummaryActionOK) Code() int {
	return 200
}

func (o *StartPTSummaryActionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/actions:startNodeAction][%d] startPtSummaryActionOk %s", 200, payload)
}

func (o *StartPTSummaryActionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/actions:startNodeAction][%d] startPtSummaryActionOk %s", 200, payload)
}

func (o *StartPTSummaryActionOK) GetPayload() *StartPTSummaryActionOKBody {
	return o.Payload
}

func (o *StartPTSummaryActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartPTSummaryActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPTSummaryActionDefault creates a StartPTSummaryActionDefault with default headers values
func NewStartPTSummaryActionDefault(code int) *StartPTSummaryActionDefault {
	return &StartPTSummaryActionDefault{
		_statusCode: code,
	}
}

/*
StartPTSummaryActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartPTSummaryActionDefault struct {
	_statusCode int

	Payload *StartPTSummaryActionDefaultBody
}

// IsSuccess returns true when this start PT summary action default response has a 2xx status code
func (o *StartPTSummaryActionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this start PT summary action default response has a 3xx status code
func (o *StartPTSummaryActionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this start PT summary action default response has a 4xx status code
func (o *StartPTSummaryActionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this start PT summary action default response has a 5xx status code
func (o *StartPTSummaryActionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this start PT summary action default response a status code equal to that given
func (o *StartPTSummaryActionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the start PT summary action default response
func (o *StartPTSummaryActionDefault) Code() int {
	return o._statusCode
}

func (o *StartPTSummaryActionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/actions:startNodeAction][%d] StartPTSummaryAction default %s", o._statusCode, payload)
}

func (o *StartPTSummaryActionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/actions:startNodeAction][%d] StartPTSummaryAction default %s", o._statusCode, payload)
}

func (o *StartPTSummaryActionDefault) GetPayload() *StartPTSummaryActionDefaultBody {
	return o.Payload
}

func (o *StartPTSummaryActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartPTSummaryActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
StartPTSummaryActionBody start PT summary action body
swagger:model StartPTSummaryActionBody
*/
type StartPTSummaryActionBody struct {
	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Node ID for this Action.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this start PT summary action body
func (o *StartPTSummaryActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start PT summary action body based on context it is used
func (o *StartPTSummaryActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTSummaryActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTSummaryActionBody) UnmarshalBinary(b []byte) error {
	var res StartPTSummaryActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartPTSummaryActionDefaultBody start PT summary action default body
swagger:model StartPTSummaryActionDefaultBody
*/
type StartPTSummaryActionDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartPTSummaryActionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start PT summary action default body
func (o *StartPTSummaryActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPTSummaryActionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("StartPTSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPTSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start PT summary action default body based on the context it is used
func (o *StartPTSummaryActionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPTSummaryActionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartPTSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPTSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartPTSummaryActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTSummaryActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartPTSummaryActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartPTSummaryActionDefaultBodyDetailsItems0 start PT summary action default body details items0
swagger:model StartPTSummaryActionDefaultBodyDetailsItems0
*/
type StartPTSummaryActionDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// start PT summary action default body details items0
	StartPTSummaryActionDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *StartPTSummaryActionDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv StartPTSummaryActionDefaultBodyDetailsItems0

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
		o.StartPTSummaryActionDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o StartPTSummaryActionDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.StartPTSummaryActionDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.StartPTSummaryActionDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this start PT summary action default body details items0
func (o *StartPTSummaryActionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start PT summary action default body details items0 based on context it is used
func (o *StartPTSummaryActionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTSummaryActionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTSummaryActionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartPTSummaryActionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
StartPTSummaryActionOKBody start PT summary action OK body
swagger:model StartPTSummaryActionOKBody
*/
type StartPTSummaryActionOKBody struct {
	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start PT summary action OK body
func (o *StartPTSummaryActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start PT summary action OK body based on context it is used
func (o *StartPTSummaryActionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTSummaryActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTSummaryActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartPTSummaryActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}