// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostAdminLoginHandlerFunc turns a function with the right signature into a post admin login handler
type PostAdminLoginHandlerFunc func(PostAdminLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAdminLoginHandlerFunc) Handle(params PostAdminLoginParams) middleware.Responder {
	return fn(params)
}

// PostAdminLoginHandler interface for that can handle valid post admin login params
type PostAdminLoginHandler interface {
	Handle(PostAdminLoginParams) middleware.Responder
}

// NewPostAdminLogin creates a new http.Handler for the post admin login operation
func NewPostAdminLogin(ctx *middleware.Context, handler PostAdminLoginHandler) *PostAdminLogin {
	return &PostAdminLogin{Context: ctx, Handler: handler}
}

/* PostAdminLogin swagger:route POST /admin/login postAdminLogin

Login With Admin

*/
type PostAdminLogin struct {
	Context *middleware.Context
	Handler PostAdminLoginHandler
}

func (o *PostAdminLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAdminLoginParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAdminLoginBody post admin login body
//
// swagger:model PostAdminLoginBody
type PostAdminLoginBody struct {

	// login Id
	// Required: true
	LoginID *string `json:"loginId"`

	// password
	// Required: true
	Password *string `json:"password"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *PostAdminLoginBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// login Id
		// Required: true
		LoginID *string `json:"loginId"`

		// password
		// Required: true
		Password *string `json:"password"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.LoginID = props.LoginID
	o.Password = props.Password
	return nil
}

// Validate validates this post admin login body
func (o *PostAdminLoginBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLoginID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAdminLoginBody) validateLoginID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"loginId", "body", o.LoginID); err != nil {
		return err
	}

	return nil
}

func (o *PostAdminLoginBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post admin login body based on context it is used
func (o *PostAdminLoginBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAdminLoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdminLoginBody) UnmarshalBinary(b []byte) error {
	var res PostAdminLoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAdminLoginOKBody post admin login o k body
//
// swagger:model PostAdminLoginOKBody
type PostAdminLoginOKBody struct {

	// token
	// Required: true
	Token *string `json:"token"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *PostAdminLoginOKBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// token
		// Required: true
		Token *string `json:"token"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Token = props.Token
	return nil
}

// Validate validates this post admin login o k body
func (o *PostAdminLoginOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAdminLoginOKBody) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("postAdminLoginOK"+"."+"token", "body", o.Token); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post admin login o k body based on context it is used
func (o *PostAdminLoginOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAdminLoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdminLoginOKBody) UnmarshalBinary(b []byte) error {
	var res PostAdminLoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}