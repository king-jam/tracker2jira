// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteUserByIDParams creates a new DeleteUserByIDParams object
// with the default values initialized.
func NewDeleteUserByIDParams() DeleteUserByIDParams {
	var ()
	return DeleteUserByIDParams{}
}

// DeleteUserByIDParams contains all the bound params for the delete user by ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteUserByID
type DeleteUserByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID of user to delete
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteUserByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rUserID, rhkUserID, _ := route.Params.GetOK("userID")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteUserByIDParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.UserID = raw

	return nil
}