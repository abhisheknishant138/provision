package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/rackn/rocket-skates/models"
)

// PutTemplateHandlerFunc turns a function with the right signature into a put template handler
type PutTemplateHandlerFunc func(PutTemplateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PutTemplateHandlerFunc) Handle(params PutTemplateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PutTemplateHandler interface for that can handle valid put template params
type PutTemplateHandler interface {
	Handle(PutTemplateParams, *models.Principal) middleware.Responder
}

// NewPutTemplate creates a new http.Handler for the put template operation
func NewPutTemplate(ctx *middleware.Context, handler PutTemplateHandler) *PutTemplate {
	return &PutTemplate{Context: ctx, Handler: handler}
}

/*PutTemplate swagger:route PUT /templates/{uuid} Templates putTemplate

Update Template

*/
type PutTemplate struct {
	Context *middleware.Context
	Handler PutTemplateHandler
}

func (o *PutTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPutTemplateParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
