package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

// HTTP code for type PutMachineOK
const PutMachineOKCode int = 200

/*PutMachineOK put machine o k

swagger:response putMachineOK
*/
type PutMachineOK struct {

	/*
	  In: Body
	*/
	Payload *models.MachineOutput `json:"body,omitempty"`
}

// NewPutMachineOK creates PutMachineOK with default headers values
func NewPutMachineOK() *PutMachineOK {
	return &PutMachineOK{}
}

// WithPayload adds the payload to the put machine o k response
func (o *PutMachineOK) WithPayload(payload *models.MachineOutput) *PutMachineOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put machine o k response
func (o *PutMachineOK) SetPayload(payload *models.MachineOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMachineOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PutMachineUnauthorized
const PutMachineUnauthorizedCode int = 401

/*PutMachineUnauthorized put machine unauthorized

swagger:response putMachineUnauthorized
*/
type PutMachineUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutMachineUnauthorized creates PutMachineUnauthorized with default headers values
func NewPutMachineUnauthorized() *PutMachineUnauthorized {
	return &PutMachineUnauthorized{}
}

// WithPayload adds the payload to the put machine unauthorized response
func (o *PutMachineUnauthorized) WithPayload(payload *models.Error) *PutMachineUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put machine unauthorized response
func (o *PutMachineUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMachineUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PutMachineNotFound
const PutMachineNotFoundCode int = 404

/*PutMachineNotFound put machine not found

swagger:response putMachineNotFound
*/
type PutMachineNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutMachineNotFound creates PutMachineNotFound with default headers values
func NewPutMachineNotFound() *PutMachineNotFound {
	return &PutMachineNotFound{}
}

// WithPayload adds the payload to the put machine not found response
func (o *PutMachineNotFound) WithPayload(payload *models.Error) *PutMachineNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put machine not found response
func (o *PutMachineNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMachineNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PutMachineConflict
const PutMachineConflictCode int = 409

/*PutMachineConflict put machine conflict

swagger:response putMachineConflict
*/
type PutMachineConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutMachineConflict creates PutMachineConflict with default headers values
func NewPutMachineConflict() *PutMachineConflict {
	return &PutMachineConflict{}
}

// WithPayload adds the payload to the put machine conflict response
func (o *PutMachineConflict) WithPayload(payload *models.Error) *PutMachineConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put machine conflict response
func (o *PutMachineConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMachineConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type PutMachineInternalServerError
const PutMachineInternalServerErrorCode int = 500

/*PutMachineInternalServerError put machine internal server error

swagger:response putMachineInternalServerError
*/
type PutMachineInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutMachineInternalServerError creates PutMachineInternalServerError with default headers values
func NewPutMachineInternalServerError() *PutMachineInternalServerError {
	return &PutMachineInternalServerError{}
}

// WithPayload adds the payload to the put machine internal server error response
func (o *PutMachineInternalServerError) WithPayload(payload *models.Error) *PutMachineInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put machine internal server error response
func (o *PutMachineInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMachineInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
