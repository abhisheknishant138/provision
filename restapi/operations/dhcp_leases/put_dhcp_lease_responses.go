package dhcp_leases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

// HTTP code for type PutDhcpLeaseOK
const PutDhcpLeaseOKCode int = 200

/*PutDhcpLeaseOK put dhcp lease o k

swagger:response putDhcpLeaseOK
*/
type PutDhcpLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.DhcpLeaseInput `json:"body,omitempty"`
}

// NewPutDhcpLeaseOK creates PutDhcpLeaseOK with default headers values
func NewPutDhcpLeaseOK() *PutDhcpLeaseOK {
	return &PutDhcpLeaseOK{}
}

// WithPayload adds the payload to the put dhcp lease o k response
func (o *PutDhcpLeaseOK) WithPayload(payload *models.DhcpLeaseInput) *PutDhcpLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put dhcp lease o k response
func (o *PutDhcpLeaseOK) SetPayload(payload *models.DhcpLeaseInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDhcpLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PutDhcpLeaseUnauthorized
const PutDhcpLeaseUnauthorizedCode int = 401

/*PutDhcpLeaseUnauthorized put dhcp lease unauthorized

swagger:response putDhcpLeaseUnauthorized
*/
type PutDhcpLeaseUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutDhcpLeaseUnauthorized creates PutDhcpLeaseUnauthorized with default headers values
func NewPutDhcpLeaseUnauthorized() *PutDhcpLeaseUnauthorized {
	return &PutDhcpLeaseUnauthorized{}
}

// WithPayload adds the payload to the put dhcp lease unauthorized response
func (o *PutDhcpLeaseUnauthorized) WithPayload(payload *models.Error) *PutDhcpLeaseUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put dhcp lease unauthorized response
func (o *PutDhcpLeaseUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDhcpLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PutDhcpLeaseNotFound
const PutDhcpLeaseNotFoundCode int = 404

/*PutDhcpLeaseNotFound put dhcp lease not found

swagger:response putDhcpLeaseNotFound
*/
type PutDhcpLeaseNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutDhcpLeaseNotFound creates PutDhcpLeaseNotFound with default headers values
func NewPutDhcpLeaseNotFound() *PutDhcpLeaseNotFound {
	return &PutDhcpLeaseNotFound{}
}

// WithPayload adds the payload to the put dhcp lease not found response
func (o *PutDhcpLeaseNotFound) WithPayload(payload *models.Error) *PutDhcpLeaseNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put dhcp lease not found response
func (o *PutDhcpLeaseNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDhcpLeaseNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PutDhcpLeaseConflict
const PutDhcpLeaseConflictCode int = 409

/*PutDhcpLeaseConflict put dhcp lease conflict

swagger:response putDhcpLeaseConflict
*/
type PutDhcpLeaseConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutDhcpLeaseConflict creates PutDhcpLeaseConflict with default headers values
func NewPutDhcpLeaseConflict() *PutDhcpLeaseConflict {
	return &PutDhcpLeaseConflict{}
}

// WithPayload adds the payload to the put dhcp lease conflict response
func (o *PutDhcpLeaseConflict) WithPayload(payload *models.Error) *PutDhcpLeaseConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put dhcp lease conflict response
func (o *PutDhcpLeaseConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDhcpLeaseConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PutDhcpLeaseInternalServerError
const PutDhcpLeaseInternalServerErrorCode int = 500

/*PutDhcpLeaseInternalServerError put dhcp lease internal server error

swagger:response putDhcpLeaseInternalServerError
*/
type PutDhcpLeaseInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutDhcpLeaseInternalServerError creates PutDhcpLeaseInternalServerError with default headers values
func NewPutDhcpLeaseInternalServerError() *PutDhcpLeaseInternalServerError {
	return &PutDhcpLeaseInternalServerError{}
}

// WithPayload adds the payload to the put dhcp lease internal server error response
func (o *PutDhcpLeaseInternalServerError) WithPayload(payload *models.Error) *PutDhcpLeaseInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put dhcp lease internal server error response
func (o *PutDhcpLeaseInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDhcpLeaseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
