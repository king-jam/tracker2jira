// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/king-jam/tracker2jira/rest/models"
)

// PostProjectCreatedCode is the HTTP code returned for type PostProjectCreated
const PostProjectCreatedCode int = 201

/*PostProjectCreated Created

swagger:response postProjectCreated
*/
type PostProjectCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Project `json:"body,omitempty"`
}

// NewPostProjectCreated creates PostProjectCreated with default headers values
func NewPostProjectCreated() *PostProjectCreated {

	return &PostProjectCreated{}
}

// WithPayload adds the payload to the post project created response
func (o *PostProjectCreated) WithPayload(payload *models.Project) *PostProjectCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project created response
func (o *PostProjectCreated) SetPayload(payload *models.Project) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostProjectBadRequestCode is the HTTP code returned for type PostProjectBadRequest
const PostProjectBadRequestCode int = 400

/*PostProjectBadRequest Bad Request

swagger:response postProjectBadRequest
*/
type PostProjectBadRequest struct {
}

// NewPostProjectBadRequest creates PostProjectBadRequest with default headers values
func NewPostProjectBadRequest() *PostProjectBadRequest {

	return &PostProjectBadRequest{}
}

// WriteResponse to the client
func (o *PostProjectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
