package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/galthaus/swagger-test/models"
	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchMachineHandlerFunc turns a function with the right signature into a patch machine handler
type PatchMachineHandlerFunc func(PatchMachineParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchMachineHandlerFunc) Handle(params PatchMachineParams) middleware.Responder {
	return fn(params)
}

// PatchMachineHandler interface for that can handle valid patch machine params
type PatchMachineHandler interface {
	Handle(PatchMachineParams) middleware.Responder
}

// NewPatchMachine creates a new http.Handler for the patch machine operation
func NewPatchMachine(ctx *middleware.Context, handler PatchMachineHandler) *PatchMachine {
	return &PatchMachine{Context: ctx, Handler: handler}
}

/*PatchMachine swagger:route PATCH /machines/{uuid} Machines patchMachine

Patch Machine

*/
type PatchMachine struct {
	Context *middleware.Context
	Handler PatchMachineHandler
}

func (o *PatchMachine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPatchMachineParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PatchMachineAcceptedBody patch machine accepted body
// swagger:model PatchMachineAcceptedBody
type PatchMachineAcceptedBody struct {

	// data
	// Required: true
	Data *models.MachineOutput `json:"Data"`

	// result
	// Required: true
	Result *models.Result `json:"Result"`
}
