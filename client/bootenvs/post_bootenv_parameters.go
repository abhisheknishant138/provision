package bootenvs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/galthaus/swagger-test/models"
)

// NewPostBootenvParams creates a new PostBootenvParams object
// with the default values initialized.
func NewPostBootenvParams() *PostBootenvParams {
	var ()
	return &PostBootenvParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostBootenvParamsWithTimeout creates a new PostBootenvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostBootenvParamsWithTimeout(timeout time.Duration) *PostBootenvParams {
	var ()
	return &PostBootenvParams{

		timeout: timeout,
	}
}

// NewPostBootenvParamsWithContext creates a new PostBootenvParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostBootenvParamsWithContext(ctx context.Context) *PostBootenvParams {
	var ()
	return &PostBootenvParams{

		Context: ctx,
	}
}

/*PostBootenvParams contains all the parameters to send to the API endpoint
for the post bootenv operation typically these are written to a http.Request
*/
type PostBootenvParams struct {

	/*Body*/
	Body *models.BootenvInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post bootenv params
func (o *PostBootenvParams) WithTimeout(timeout time.Duration) *PostBootenvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post bootenv params
func (o *PostBootenvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post bootenv params
func (o *PostBootenvParams) WithContext(ctx context.Context) *PostBootenvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post bootenv params
func (o *PostBootenvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post bootenv params
func (o *PostBootenvParams) WithBody(body *models.BootenvInput) *PostBootenvParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post bootenv params
func (o *PostBootenvParams) SetBody(body *models.BootenvInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostBootenvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.BootenvInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
