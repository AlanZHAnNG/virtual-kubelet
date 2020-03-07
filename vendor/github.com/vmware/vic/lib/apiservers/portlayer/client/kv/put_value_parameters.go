package kv

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

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// NewPutValueParams creates a new PutValueParams object
// with the default values initialized.
func NewPutValueParams() *PutValueParams {
	var ()
	return &PutValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutValueParamsWithTimeout creates a new PutValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutValueParamsWithTimeout(timeout time.Duration) *PutValueParams {
	var ()
	return &PutValueParams{

		timeout: timeout,
	}
}

// NewPutValueParamsWithContext creates a new PutValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutValueParamsWithContext(ctx context.Context) *PutValueParams {
	var ()
	return &PutValueParams{

		Context: ctx,
	}
}

// NewPutValueParamsWithHTTPClient creates a new PutValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutValueParamsWithHTTPClient(client *http.Client) *PutValueParams {
	var ()
	return &PutValueParams{
		HTTPClient: client,
	}
}

/*PutValueParams contains all the parameters to send to the API endpoint
for the put value operation typically these are written to a http.Request
*/
type PutValueParams struct {

	/*OpID*/
	OpID *string
	/*Key*/
	Key string
	/*KeyValue*/
	KeyValue *models.KeyValue

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put value params
func (o *PutValueParams) WithTimeout(timeout time.Duration) *PutValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put value params
func (o *PutValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put value params
func (o *PutValueParams) WithContext(ctx context.Context) *PutValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put value params
func (o *PutValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put value params
func (o *PutValueParams) WithHTTPClient(client *http.Client) *PutValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put value params
func (o *PutValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpID adds the opID to the put value params
func (o *PutValueParams) WithOpID(opID *string) *PutValueParams {
	o.SetOpID(opID)
	return o
}

// SetOpID adds the opId to the put value params
func (o *PutValueParams) SetOpID(opID *string) {
	o.OpID = opID
}

// WithKey adds the key to the put value params
func (o *PutValueParams) WithKey(key string) *PutValueParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the put value params
func (o *PutValueParams) SetKey(key string) {
	o.Key = key
}

// WithKeyValue adds the keyValue to the put value params
func (o *PutValueParams) WithKeyValue(keyValue *models.KeyValue) *PutValueParams {
	o.SetKeyValue(keyValue)
	return o
}

// SetKeyValue adds the keyValue to the put value params
func (o *PutValueParams) SetKeyValue(keyValue *models.KeyValue) {
	o.KeyValue = keyValue
}

// WriteToRequest writes these params to a swagger request
func (o *PutValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.OpID != nil {

		// header param Op-ID
		if err := r.SetHeaderParam("Op-ID", *o.OpID); err != nil {
			return err
		}

	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if o.KeyValue == nil {
		o.KeyValue = new(models.KeyValue)
	}

	if err := r.SetBodyParam(o.KeyValue); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
