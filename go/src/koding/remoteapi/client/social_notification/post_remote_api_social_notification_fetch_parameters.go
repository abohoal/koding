package social_notification

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

// NewPostRemoteAPISocialNotificationFetchParams creates a new PostRemoteAPISocialNotificationFetchParams object
// with the default values initialized.
func NewPostRemoteAPISocialNotificationFetchParams() *PostRemoteAPISocialNotificationFetchParams {
	var ()
	return &PostRemoteAPISocialNotificationFetchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialNotificationFetchParamsWithTimeout creates a new PostRemoteAPISocialNotificationFetchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialNotificationFetchParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialNotificationFetchParams {
	var ()
	return &PostRemoteAPISocialNotificationFetchParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialNotificationFetchParamsWithContext creates a new PostRemoteAPISocialNotificationFetchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialNotificationFetchParamsWithContext(ctx context.Context) *PostRemoteAPISocialNotificationFetchParams {
	var ()
	return &PostRemoteAPISocialNotificationFetchParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialNotificationFetchParams contains all the parameters to send to the API endpoint
for the post remote API social notification fetch operation typically these are written to a http.Request
*/
type PostRemoteAPISocialNotificationFetchParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API social notification fetch params
func (o *PostRemoteAPISocialNotificationFetchParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialNotificationFetchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social notification fetch params
func (o *PostRemoteAPISocialNotificationFetchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social notification fetch params
func (o *PostRemoteAPISocialNotificationFetchParams) WithContext(ctx context.Context) *PostRemoteAPISocialNotificationFetchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social notification fetch params
func (o *PostRemoteAPISocialNotificationFetchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social notification fetch params
func (o *PostRemoteAPISocialNotificationFetchParams) WithBody(body models.DefaultSelector) *PostRemoteAPISocialNotificationFetchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social notification fetch params
func (o *PostRemoteAPISocialNotificationFetchParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialNotificationFetchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
