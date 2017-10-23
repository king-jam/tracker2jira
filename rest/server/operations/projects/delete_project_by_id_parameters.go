// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteProjectByIDParams creates a new DeleteProjectByIDParams object
// with the default values initialized.
func NewDeleteProjectByIDParams() DeleteProjectByIDParams {
	var ()
	return DeleteProjectByIDParams{}
}

// DeleteProjectByIDParams contains all the bound params for the delete project by ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteProjectByID
type DeleteProjectByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID of project to delete
	  Required: true
	  In: path
	*/
	ProjectID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteProjectByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rProjectID, rhkProjectID, _ := route.Params.GetOK("projectID")
	if err := o.bindProjectID(rProjectID, rhkProjectID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteProjectByIDParams) bindProjectID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ProjectID = raw

	return nil
}