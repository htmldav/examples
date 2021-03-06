// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yandex-cloud/examples/serverless/alice-shareable-todolist/app/generated/openapi/models"
)

// RevokeInvitationNoContentCode is the HTTP code returned for type RevokeInvitationNoContent
const RevokeInvitationNoContentCode int = 204

/*RevokeInvitationNoContent OK

swagger:response revokeInvitationNoContent
*/
type RevokeInvitationNoContent struct {
}

// NewRevokeInvitationNoContent creates RevokeInvitationNoContent with default headers values
func NewRevokeInvitationNoContent() *RevokeInvitationNoContent {

	return &RevokeInvitationNoContent{}
}

// WriteResponse to the client
func (o *RevokeInvitationNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*RevokeInvitationDefault error

swagger:response revokeInvitationDefault
*/
type RevokeInvitationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRevokeInvitationDefault creates RevokeInvitationDefault with default headers values
func NewRevokeInvitationDefault(code int) *RevokeInvitationDefault {
	if code <= 0 {
		code = 500
	}

	return &RevokeInvitationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the revoke invitation default response
func (o *RevokeInvitationDefault) WithStatusCode(code int) *RevokeInvitationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the revoke invitation default response
func (o *RevokeInvitationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the revoke invitation default response
func (o *RevokeInvitationDefault) WithPayload(payload *models.Error) *RevokeInvitationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the revoke invitation default response
func (o *RevokeInvitationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RevokeInvitationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
