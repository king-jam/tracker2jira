// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/king-jam/tracker2jira/rest/models"
)

// PostTaskAcceptedCode is the HTTP code returned for type PostTaskAccepted
const PostTaskAcceptedCode int = 202

/*PostTaskAccepted Accepted

swagger:response postTaskAccepted
*/
type PostTaskAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.Task `json:"body,omitempty"`
}

// NewPostTaskAccepted creates PostTaskAccepted with default headers values
func NewPostTaskAccepted() *PostTaskAccepted {

	return &PostTaskAccepted{}
}

// WithPayload adds the payload to the post task accepted response
func (o *PostTaskAccepted) WithPayload(payload *models.Task) *PostTaskAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post task accepted response
func (o *PostTaskAccepted) SetPayload(payload *models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTaskAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTaskBadRequestCode is the HTTP code returned for type PostTaskBadRequest
const PostTaskBadRequestCode int = 400

/*PostTaskBadRequest Bad Request

swagger:response postTaskBadRequest
*/
type PostTaskBadRequest struct {
}

// NewPostTaskBadRequest creates PostTaskBadRequest with default headers values
func NewPostTaskBadRequest() *PostTaskBadRequest {

	return &PostTaskBadRequest{}
}

// WriteResponse to the client
func (o *PostTaskBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
