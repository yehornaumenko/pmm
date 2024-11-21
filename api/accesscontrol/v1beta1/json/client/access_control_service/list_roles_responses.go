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

// ListRolesReader is a Reader for the ListRoles structure.
type ListRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRolesOK creates a ListRolesOK with default headers values
func NewListRolesOK() *ListRolesOK {
	return &ListRolesOK{}
}

/*
ListRolesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListRolesOK struct {
	Payload *ListRolesOKBody
}

// IsSuccess returns true when this list roles Ok response has a 2xx status code
func (o *ListRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list roles Ok response has a 3xx status code
func (o *ListRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list roles Ok response has a 4xx status code
func (o *ListRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list roles Ok response has a 5xx status code
func (o *ListRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list roles Ok response a status code equal to that given
func (o *ListRolesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list roles Ok response
func (o *ListRolesOK) Code() int {
	return 200
}

func (o *ListRolesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/accesscontrol/roles][%d] listRolesOk %s", 200, payload)
}

func (o *ListRolesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/accesscontrol/roles][%d] listRolesOk %s", 200, payload)
}

func (o *ListRolesOK) GetPayload() *ListRolesOKBody {
	return o.Payload
}

func (o *ListRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListRolesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRolesDefault creates a ListRolesDefault with default headers values
func NewListRolesDefault(code int) *ListRolesDefault {
	return &ListRolesDefault{
		_statusCode: code,
	}
}

/*
ListRolesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListRolesDefault struct {
	_statusCode int

	Payload *ListRolesDefaultBody
}

// IsSuccess returns true when this list roles default response has a 2xx status code
func (o *ListRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list roles default response has a 3xx status code
func (o *ListRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list roles default response has a 4xx status code
func (o *ListRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list roles default response has a 5xx status code
func (o *ListRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list roles default response a status code equal to that given
func (o *ListRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list roles default response
func (o *ListRolesDefault) Code() int {
	return o._statusCode
}

func (o *ListRolesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/accesscontrol/roles][%d] ListRoles default %s", o._statusCode, payload)
}

func (o *ListRolesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/accesscontrol/roles][%d] ListRoles default %s", o._statusCode, payload)
}

func (o *ListRolesDefault) GetPayload() *ListRolesDefaultBody {
	return o.Payload
}

func (o *ListRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListRolesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListRolesDefaultBody list roles default body
swagger:model ListRolesDefaultBody
*/
type ListRolesDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListRolesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list roles default body
func (o *ListRolesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRolesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListRoles default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListRoles default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list roles default body based on the context it is used
func (o *ListRolesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRolesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListRoles default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListRoles default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListRolesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRolesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListRolesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListRolesDefaultBodyDetailsItems0 list roles default body details items0
swagger:model ListRolesDefaultBodyDetailsItems0
*/
type ListRolesDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`

	// list roles default body details items0
	ListRolesDefaultBodyDetailsItems0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *ListRolesDefaultBodyDetailsItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {
		// at type
		AtType string `json:"@type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ListRolesDefaultBodyDetailsItems0

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
		o.ListRolesDefaultBodyDetailsItems0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o ListRolesDefaultBodyDetailsItems0) MarshalJSON() ([]byte, error) {
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

	if len(o.ListRolesDefaultBodyDetailsItems0) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.ListRolesDefaultBodyDetailsItems0)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this list roles default body details items0
func (o *ListRolesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list roles default body details items0 based on context it is used
func (o *ListRolesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListRolesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRolesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListRolesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListRolesOKBody list roles OK body
swagger:model ListRolesOKBody
*/
type ListRolesOKBody struct {
	// roles
	Roles []*ListRolesOKBodyRolesItems0 `json:"roles"`
}

// Validate validates this list roles OK body
func (o *ListRolesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRolesOKBody) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(o.Roles) { // not required
		return nil
	}

	for i := 0; i < len(o.Roles); i++ {
		if swag.IsZero(o.Roles[i]) { // not required
			continue
		}

		if o.Roles[i] != nil {
			if err := o.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listRolesOk" + "." + "roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listRolesOk" + "." + "roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list roles OK body based on the context it is used
func (o *ListRolesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRolesOKBody) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Roles); i++ {
		if o.Roles[i] != nil {

			if swag.IsZero(o.Roles[i]) { // not required
				return nil
			}

			if err := o.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listRolesOk" + "." + "roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listRolesOk" + "." + "roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListRolesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRolesOKBody) UnmarshalBinary(b []byte) error {
	var res ListRolesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListRolesOKBodyRolesItems0 list roles OK body roles items0
swagger:model ListRolesOKBodyRolesItems0
*/
type ListRolesOKBodyRolesItems0 struct {
	// role id
	RoleID int64 `json:"role_id,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// filter
	Filter string `json:"filter,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this list roles OK body roles items0
func (o *ListRolesOKBodyRolesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list roles OK body roles items0 based on context it is used
func (o *ListRolesOKBodyRolesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListRolesOKBodyRolesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRolesOKBodyRolesItems0) UnmarshalBinary(b []byte) error {
	var res ListRolesOKBodyRolesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}