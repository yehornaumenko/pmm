// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetKubernetesClusterReader is a Reader for the GetKubernetesCluster structure.
type GetKubernetesClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetKubernetesClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKubernetesClusterOK creates a GetKubernetesClusterOK with default headers values
func NewGetKubernetesClusterOK() *GetKubernetesClusterOK {
	return &GetKubernetesClusterOK{}
}

/*
GetKubernetesClusterOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetKubernetesClusterOK struct {
	Payload *GetKubernetesClusterOKBody
}

func (o *GetKubernetesClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/Get][%d] getKubernetesClusterOk  %+v", 200, o.Payload)
}

func (o *GetKubernetesClusterOK) GetPayload() *GetKubernetesClusterOKBody {
	return o.Payload
}

func (o *GetKubernetesClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetKubernetesClusterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesClusterDefault creates a GetKubernetesClusterDefault with default headers values
func NewGetKubernetesClusterDefault(code int) *GetKubernetesClusterDefault {
	return &GetKubernetesClusterDefault{
		_statusCode: code,
	}
}

/*
GetKubernetesClusterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetKubernetesClusterDefault struct {
	_statusCode int

	Payload *GetKubernetesClusterDefaultBody
}

// Code gets the status code for the get kubernetes cluster default response
func (o *GetKubernetesClusterDefault) Code() int {
	return o._statusCode
}

func (o *GetKubernetesClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/Get][%d] GetKubernetesCluster default  %+v", o._statusCode, o.Payload)
}

func (o *GetKubernetesClusterDefault) GetPayload() *GetKubernetesClusterDefaultBody {
	return o.Payload
}

func (o *GetKubernetesClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetKubernetesClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetKubernetesClusterBody get kubernetes cluster body
swagger:model GetKubernetesClusterBody
*/
type GetKubernetesClusterBody struct {
	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`
}

// Validate validates this get kubernetes cluster body
func (o *GetKubernetesClusterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get kubernetes cluster body based on context it is used
func (o *GetKubernetesClusterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetKubernetesClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetKubernetesClusterBody) UnmarshalBinary(b []byte) error {
	var res GetKubernetesClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetKubernetesClusterDefaultBody get kubernetes cluster default body
swagger:model GetKubernetesClusterDefaultBody
*/
type GetKubernetesClusterDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetKubernetesClusterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get kubernetes cluster default body
func (o *GetKubernetesClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetKubernetesClusterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetKubernetesCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetKubernetesCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get kubernetes cluster default body based on the context it is used
func (o *GetKubernetesClusterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetKubernetesClusterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetKubernetesCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetKubernetesCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetKubernetesClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetKubernetesClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetKubernetesClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetKubernetesClusterDefaultBodyDetailsItems0 get kubernetes cluster default body details items0
swagger:model GetKubernetesClusterDefaultBodyDetailsItems0
*/
type GetKubernetesClusterDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get kubernetes cluster default body details items0
func (o *GetKubernetesClusterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get kubernetes cluster default body details items0 based on context it is used
func (o *GetKubernetesClusterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetKubernetesClusterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetKubernetesClusterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetKubernetesClusterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetKubernetesClusterOKBody get kubernetes cluster OK body
swagger:model GetKubernetesClusterOKBody
*/
type GetKubernetesClusterOKBody struct {
	// kube auth
	KubeAuth *GetKubernetesClusterOKBodyKubeAuth `json:"kube_auth,omitempty"`
}

// Validate validates this get kubernetes cluster OK body
func (o *GetKubernetesClusterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKubeAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetKubernetesClusterOKBody) validateKubeAuth(formats strfmt.Registry) error {
	if swag.IsZero(o.KubeAuth) { // not required
		return nil
	}

	if o.KubeAuth != nil {
		if err := o.KubeAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getKubernetesClusterOk" + "." + "kube_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getKubernetesClusterOk" + "." + "kube_auth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get kubernetes cluster OK body based on the context it is used
func (o *GetKubernetesClusterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateKubeAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetKubernetesClusterOKBody) contextValidateKubeAuth(ctx context.Context, formats strfmt.Registry) error {
	if o.KubeAuth != nil {
		if err := o.KubeAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getKubernetesClusterOk" + "." + "kube_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getKubernetesClusterOk" + "." + "kube_auth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetKubernetesClusterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetKubernetesClusterOKBody) UnmarshalBinary(b []byte) error {
	var res GetKubernetesClusterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetKubernetesClusterOKBodyKubeAuth KubeAuth represents Kubernetes / kubectl authentication and authorization information.
swagger:model GetKubernetesClusterOKBodyKubeAuth
*/
type GetKubernetesClusterOKBodyKubeAuth struct {
	// Kubeconfig file content.
	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// Validate validates this get kubernetes cluster OK body kube auth
func (o *GetKubernetesClusterOKBodyKubeAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get kubernetes cluster OK body kube auth based on context it is used
func (o *GetKubernetesClusterOKBodyKubeAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetKubernetesClusterOKBodyKubeAuth) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetKubernetesClusterOKBodyKubeAuth) UnmarshalBinary(b []byte) error {
	var res GetKubernetesClusterOKBodyKubeAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}