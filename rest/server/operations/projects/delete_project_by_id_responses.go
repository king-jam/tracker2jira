// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteProjectByIDNoContentCode is the HTTP code returned for type DeleteProjectByIDNoContent
const DeleteProjectByIDNoContentCode int = 204

/*DeleteProjectByIDNoContent successful operation

swagger:response deleteProjectByIdNoContent
*/
type DeleteProjectByIDNoContent struct {
}

// NewDeleteProjectByIDNoContent creates DeleteProjectByIDNoContent with default headers values
func NewDeleteProjectByIDNoContent() *DeleteProjectByIDNoContent {
	return &DeleteProjectByIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteProjectByIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteProjectByIDNotFoundCode is the HTTP code returned for type DeleteProjectByIDNotFound
const DeleteProjectByIDNotFoundCode int = 404

/*DeleteProjectByIDNotFound Project not found

swagger:response deleteProjectByIdNotFound
*/
type DeleteProjectByIDNotFound struct {
}

// NewDeleteProjectByIDNotFound creates DeleteProjectByIDNotFound with default headers values
func NewDeleteProjectByIDNotFound() *DeleteProjectByIDNotFound {
	return &DeleteProjectByIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteProjectByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
