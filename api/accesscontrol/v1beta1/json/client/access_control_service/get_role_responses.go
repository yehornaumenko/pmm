// Code generated by go-swagger; DO NOT EDIT.

package access_control_service

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

// GetRoleReader is a Reader for the GetRole structure.
type GetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoleOK creates a GetRoleOK with default headers values
func NewGetRoleOK() *GetRoleOK {
	return &GetRoleOK{}
}

/*
GetRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetRoleOK struct {
	Payload *GetRoleOKBody
}

// IsSuccess returns true when this get role Ok response has a 2xx status code
func (o *GetRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get role Ok response has a 3xx status code
func (o *GetRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role Ok response has a 4xx status code
func (o *GetRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role Ok response has a 5xx status code
func (o *GetRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get role Ok response a status code equal to that given
func (o *GetRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get role Ok response
func (o *GetRoleOK) Code() int {
	return 200
}

func (o *GetRoleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/accesscontrol/roles/{role_id}][%d] getRoleOk %s", 200, payload)
}

func (o *GetRoleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/accesscontrol/roles/{role_id}][%d] getRoleOk %s", 200, payload)
}

func (o *GetRoleOK) GetPayload() *GetRoleOKBody {
	return o.Payload
}

func (o *GetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetRoleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleDefault creates a GetRoleDefault with default headers values
func NewGetRoleDefault(code int) *GetRoleDefault {
	return &GetRoleDefault{
		_statusCode: code,
	}
}

/*
GetRoleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetRoleDefault struct {
	_statusCode int

	Payload *GetRoleDefaultBody
}

// IsSuccess returns true when this get role default response has a 2xx status code
func (o *GetRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get role default response has a 3xx status code
func (o *GetRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get role default response has a 4xx status code
func (o *GetRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get role default response has a 5xx status code
func (o *GetRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get role default response a status code equal to that given
func (o *GetRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get role default response
func (o *GetRoleDefault) Code() int {
	return o._statusCode
}

func (o *GetRoleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/accesscontrol/roles/{role_id}][%d] GetRole default %s", o._statusCode, payload)
}

func (o *GetRoleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/accesscontrol/roles/{role_id}][%d] GetRole default %s", o._statusCode, payload)
}

func (o *GetRoleDefault) GetPayload() *GetRoleDefaultBody {
	return o.Payload
}

func (o *GetRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetRoleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetRoleDefaultBody get role default body
swagger:model GetRoleDefaultBody
*/
type GetRoleDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetRoleDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get role default body
func (o *GetRoleDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoleDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetRole default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetRole default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get role default body based on the context it is used
func (o *GetRoleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoleDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetRole default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetRole default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRoleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoleDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetRoleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetRoleDefaultBodyDetailsItems0 get role default body details items0
swagger:model GetRoleDefaultBodyDetailsItems0
*/
type GetRoleDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// get role default body details items0
	GetRoleDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *GetRoleDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv GetRoleDefaultBodyDetailsItems0

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
		o.GetRoleDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o GetRoleDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.GetRoleDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.GetRoleDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this get role default body details items0
func (o *GetRoleDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get role default body details items0 based on context it is used
func (o *GetRoleDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoleDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoleDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetRoleDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetRoleOKBody get role OK body
swagger:model GetRoleOKBody
*/
type GetRoleOKBody struct {
	// role id
	RoleID int64 `json:"role_id,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// filter
	Filter string `json:"filter,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this get role OK body
func (o *GetRoleOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get role OK body based on context it is used
func (o *GetRoleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoleOKBody) UnmarshalBinary(b []byte) error {
	var res GetRoleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}