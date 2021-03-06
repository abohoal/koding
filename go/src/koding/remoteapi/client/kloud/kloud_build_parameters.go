package kloud

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

	"koding/remoteapi/models"
)

// NewKloudBuildParams creates a new KloudBuildParams object
// with the default values initialized.
func NewKloudBuildParams() *KloudBuildParams {
	var ()
	return &KloudBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKloudBuildParamsWithTimeout creates a new KloudBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKloudBuildParamsWithTimeout(timeout time.Duration) *KloudBuildParams {
	var ()
	return &KloudBuildParams{

		timeout: timeout,
	}
}

// NewKloudBuildParamsWithContext creates a new KloudBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewKloudBuildParamsWithContext(ctx context.Context) *KloudBuildParams {
	var ()
	return &KloudBuildParams{

		Context: ctx,
	}
}

/*KloudBuildParams contains all the parameters to send to the API endpoint
for the kloud build operation typically these are written to a http.Request
*/
type KloudBuildParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the kloud build params
func (o *KloudBuildParams) WithTimeout(timeout time.Duration) *KloudBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kloud build params
func (o *KloudBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kloud build params
func (o *KloudBuildParams) WithContext(ctx context.Context) *KloudBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kloud build params
func (o *KloudBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the kloud build params
func (o *KloudBuildParams) WithBody(body models.DefaultSelector) *KloudBuildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kloud build params
func (o *KloudBuildParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *KloudBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
