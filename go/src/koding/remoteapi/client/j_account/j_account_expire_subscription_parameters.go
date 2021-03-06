package j_account

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

// NewJAccountExpireSubscriptionParams creates a new JAccountExpireSubscriptionParams object
// with the default values initialized.
func NewJAccountExpireSubscriptionParams() *JAccountExpireSubscriptionParams {
	var ()
	return &JAccountExpireSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJAccountExpireSubscriptionParamsWithTimeout creates a new JAccountExpireSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJAccountExpireSubscriptionParamsWithTimeout(timeout time.Duration) *JAccountExpireSubscriptionParams {
	var ()
	return &JAccountExpireSubscriptionParams{

		timeout: timeout,
	}
}

// NewJAccountExpireSubscriptionParamsWithContext creates a new JAccountExpireSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewJAccountExpireSubscriptionParamsWithContext(ctx context.Context) *JAccountExpireSubscriptionParams {
	var ()
	return &JAccountExpireSubscriptionParams{

		Context: ctx,
	}
}

/*JAccountExpireSubscriptionParams contains all the parameters to send to the API endpoint
for the j account expire subscription operation typically these are written to a http.Request
*/
type JAccountExpireSubscriptionParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j account expire subscription params
func (o *JAccountExpireSubscriptionParams) WithTimeout(timeout time.Duration) *JAccountExpireSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j account expire subscription params
func (o *JAccountExpireSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j account expire subscription params
func (o *JAccountExpireSubscriptionParams) WithContext(ctx context.Context) *JAccountExpireSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j account expire subscription params
func (o *JAccountExpireSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j account expire subscription params
func (o *JAccountExpireSubscriptionParams) WithBody(body models.DefaultSelector) *JAccountExpireSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j account expire subscription params
func (o *JAccountExpireSubscriptionParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j account expire subscription params
func (o *JAccountExpireSubscriptionParams) WithID(id string) *JAccountExpireSubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j account expire subscription params
func (o *JAccountExpireSubscriptionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JAccountExpireSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
