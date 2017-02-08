package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/galthaus/swagger-test/models"
	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMachineHandlerFunc turns a function with the right signature into a get machine handler
type GetMachineHandlerFunc func(GetMachineParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMachineHandlerFunc) Handle(params GetMachineParams) middleware.Responder {
	return fn(params)
}

// GetMachineHandler interface for that can handle valid get machine params
type GetMachineHandler interface {
	Handle(GetMachineParams) middleware.Responder
}

// NewGetMachine creates a new http.Handler for the get machine operation
func NewGetMachine(ctx *middleware.Context, handler GetMachineHandler) *GetMachine {
	return &GetMachine{Context: ctx, Handler: handler}
}

/*GetMachine swagger:route GET /machines/{uuid} Machines getMachine

Get Machine

*/
type GetMachine struct {
	Context *middleware.Context
	Handler GetMachineHandler
}

func (o *GetMachine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetMachineParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetMachineOKBody get machine o k body
// swagger:model GetMachineOKBody
type GetMachineOKBody struct {

	// data
	// Required: true
	Data *models.MachineOutput `json:"Data"`

	// result
	// Required: true
	Result *models.Result `json:"Result"`
}
