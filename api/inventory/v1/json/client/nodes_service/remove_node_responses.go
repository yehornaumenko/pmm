// Code generated by go-swagger; DO NOT EDIT.

package nodes_service

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

// RemoveNodeReader is a Reader for the RemoveNode structure.
type RemoveNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveNodeOK creates a RemoveNodeOK with default headers values
func NewRemoveNodeOK() *RemoveNodeOK {
	return &RemoveNodeOK{}
}

/*
RemoveNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveNodeOK struct {
	Payload interface{}
}

// IsSuccess returns true when this remove node Ok response has a 2xx status code
func (o *RemoveNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove node Ok response has a 3xx status code
func (o *RemoveNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove node Ok response has a 4xx status code
func (o *RemoveNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove node Ok response has a 5xx status code
func (o *RemoveNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove node Ok response a status code equal to that given
func (o *RemoveNodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove node Ok response
func (o *RemoveNodeOK) Code() int {
	return 200
}

func (o *RemoveNodeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/inventory/nodes/{node_id}][%d] removeNodeOk %s", 200, payload)
}

func (o *RemoveNodeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/inventory/nodes/{node_id}][%d] removeNodeOk %s", 200, payload)
}

func (o *RemoveNodeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveNodeDefault creates a RemoveNodeDefault with default headers values
func NewRemoveNodeDefault(code int) *RemoveNodeDefault {
	return &RemoveNodeDefault{
		_statusCode: code,
	}
}

/*
RemoveNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveNodeDefault struct {
	_statusCode int

	Payload *RemoveNodeDefaultBody
}

// IsSuccess returns true when this remove node default response has a 2xx status code
func (o *RemoveNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this remove node default response has a 3xx status code
func (o *RemoveNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this remove node default response has a 4xx status code
func (o *RemoveNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this remove node default response has a 5xx status code
func (o *RemoveNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this remove node default response a status code equal to that given
func (o *RemoveNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the remove node default response
func (o *RemoveNodeDefault) Code() int {
	return o._statusCode
}

func (o *RemoveNodeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/inventory/nodes/{node_id}][%d] RemoveNode default %s", o._statusCode, payload)
}

func (o *RemoveNodeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/inventory/nodes/{node_id}][%d] RemoveNode default %s", o._statusCode, payload)
}

func (o *RemoveNodeDefault) GetPayload() *RemoveNodeDefaultBody {
	return o.Payload
}

func (o *RemoveNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RemoveNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
RemoveNodeDefaultBody remove node default body
swagger:model RemoveNodeDefaultBody
*/
type RemoveNodeDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RemoveNodeDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this remove node default body
func (o *RemoveNodeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveNodeDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("RemoveNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this remove node default body based on the context it is used
func (o *RemoveNodeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveNodeDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoveNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RemoveNodeDefaultBodyDetailsItems0 remove node default body details items0
swagger:model RemoveNodeDefaultBodyDetailsItems0
*/
type RemoveNodeDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// remove node default body details items0
	RemoveNodeDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *RemoveNodeDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv RemoveNodeDefaultBodyDetailsItems0

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
		o.RemoveNodeDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o RemoveNodeDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.RemoveNodeDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.RemoveNodeDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this remove node default body details items0
func (o *RemoveNodeDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove node default body details items0 based on context it is used
func (o *RemoveNodeDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveNodeDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveNodeDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RemoveNodeDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}