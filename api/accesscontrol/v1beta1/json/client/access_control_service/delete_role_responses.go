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

// DeleteRoleReader is a Reader for the DeleteRole structure.
type DeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRoleOK creates a DeleteRoleOK with default headers values
func NewDeleteRoleOK() *DeleteRoleOK {
	return &DeleteRoleOK{}
}

/*
DeleteRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteRoleOK struct {
	Payload interface{}
}

// IsSuccess returns true when this delete role Ok response has a 2xx status code
func (o *DeleteRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete role Ok response has a 3xx status code
func (o *DeleteRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role Ok response has a 4xx status code
func (o *DeleteRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete role Ok response has a 5xx status code
func (o *DeleteRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role Ok response a status code equal to that given
func (o *DeleteRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete role Ok response
func (o *DeleteRoleOK) Code() int {
	return 200
}

func (o *DeleteRoleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/accesscontrol/roles/{role_id}][%d] deleteRoleOk %s", 200, payload)
}

func (o *DeleteRoleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/accesscontrol/roles/{role_id}][%d] deleteRoleOk %s", 200, payload)
}

func (o *DeleteRoleOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleDefault creates a DeleteRoleDefault with default headers values
func NewDeleteRoleDefault(code int) *DeleteRoleDefault {
	return &DeleteRoleDefault{
		_statusCode: code,
	}
}

/*
DeleteRoleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteRoleDefault struct {
	_statusCode int

	Payload *DeleteRoleDefaultBody
}

// IsSuccess returns true when this delete role default response has a 2xx status code
func (o *DeleteRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete role default response has a 3xx status code
func (o *DeleteRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete role default response has a 4xx status code
func (o *DeleteRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete role default response has a 5xx status code
func (o *DeleteRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete role default response a status code equal to that given
func (o *DeleteRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete role default response
func (o *DeleteRoleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRoleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/accesscontrol/roles/{role_id}][%d] DeleteRole default %s", o._statusCode, payload)
}

func (o *DeleteRoleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/accesscontrol/roles/{role_id}][%d] DeleteRole default %s", o._statusCode, payload)
}

func (o *DeleteRoleDefault) GetPayload() *DeleteRoleDefaultBody {
	return o.Payload
}

func (o *DeleteRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(DeleteRoleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DeleteRoleDefaultBody delete role default body
swagger:model DeleteRoleDefaultBody
*/
type DeleteRoleDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DeleteRoleDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this delete role default body
func (o *DeleteRoleDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteRoleDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("DeleteRole default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeleteRole default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete role default body based on the context it is used
func (o *DeleteRoleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteRoleDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeleteRole default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeleteRole default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteRoleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRoleDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteRoleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeleteRoleDefaultBodyDetailsItems0 delete role default body details items0
swagger:model DeleteRoleDefaultBodyDetailsItems0
*/
type DeleteRoleDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// delete role default body details items0
	DeleteRoleDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *DeleteRoleDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv DeleteRoleDefaultBodyDetailsItems0

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
		o.DeleteRoleDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o DeleteRoleDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.DeleteRoleDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.DeleteRoleDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this delete role default body details items0
func (o *DeleteRoleDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete role default body details items0 based on context it is used
func (o *DeleteRoleDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteRoleDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRoleDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteRoleDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}