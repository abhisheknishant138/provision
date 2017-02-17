package isos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

// HTTP code for type ListIsosOK
const ListIsosOKCode int = 200

/*ListIsosOK list isos o k

swagger:response listIsosOK
*/
type ListIsosOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewListIsosOK creates ListIsosOK with default headers values
func NewListIsosOK() *ListIsosOK {
	return &ListIsosOK{}
}

// WithPayload adds the payload to the list isos o k response
func (o *ListIsosOK) WithPayload(payload []string) *ListIsosOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list isos o k response
func (o *ListIsosOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIsosOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// HTTP code for type ListIsosUnauthorized
const ListIsosUnauthorizedCode int = 401

/*ListIsosUnauthorized list isos unauthorized

swagger:response listIsosUnauthorized
*/
type ListIsosUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIsosUnauthorized creates ListIsosUnauthorized with default headers values
func NewListIsosUnauthorized() *ListIsosUnauthorized {
	return &ListIsosUnauthorized{}
}

// WithPayload adds the payload to the list isos unauthorized response
func (o *ListIsosUnauthorized) WithPayload(payload *models.Error) *ListIsosUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list isos unauthorized response
func (o *ListIsosUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIsosUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type ListIsosNotFound
const ListIsosNotFoundCode int = 404

/*ListIsosNotFound list isos not found

swagger:response listIsosNotFound
*/
type ListIsosNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIsosNotFound creates ListIsosNotFound with default headers values
func NewListIsosNotFound() *ListIsosNotFound {
	return &ListIsosNotFound{}
}

// WithPayload adds the payload to the list isos not found response
func (o *ListIsosNotFound) WithPayload(payload *models.Error) *ListIsosNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list isos not found response
func (o *ListIsosNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIsosNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type ListIsosInternalServerError
const ListIsosInternalServerErrorCode int = 500

/*ListIsosInternalServerError list isos internal server error

swagger:response listIsosInternalServerError
*/
type ListIsosInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIsosInternalServerError creates ListIsosInternalServerError with default headers values
func NewListIsosInternalServerError() *ListIsosInternalServerError {
	return &ListIsosInternalServerError{}
}

// WithPayload adds the payload to the list isos internal server error response
func (o *ListIsosInternalServerError) WithPayload(payload *models.Error) *ListIsosInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list isos internal server error response
func (o *ListIsosInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIsosInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
