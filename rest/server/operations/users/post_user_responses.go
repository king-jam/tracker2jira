// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/king-jam/tracker2jira/rest/models"
)

// PostUserCreatedCode is the HTTP code returned for type PostUserCreated
const PostUserCreatedCode int = 201

/*PostUserCreated Created

swagger:response postUserCreated
*/
type PostUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPostUserCreated creates PostUserCreated with default headers values
func NewPostUserCreated() *PostUserCreated {
	return &PostUserCreated{}
}

// WithPayload adds the payload to the post user created response
func (o *PostUserCreated) WithPayload(payload *models.User) *PostUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user created response
func (o *PostUserCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUserBadRequestCode is the HTTP code returned for type PostUserBadRequest
const PostUserBadRequestCode int = 400

/*PostUserBadRequest Bad Request

swagger:response postUserBadRequest
*/
type PostUserBadRequest struct {
}

// NewPostUserBadRequest creates PostUserBadRequest with default headers values
func NewPostUserBadRequest() *PostUserBadRequest {
	return &PostUserBadRequest{}
}

// WriteResponse to the client
func (o *PostUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
