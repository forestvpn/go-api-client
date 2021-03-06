// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewBindingListParams creates a new BindingListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBindingListParams() *BindingListParams {
	return &BindingListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBindingListParamsWithTimeout creates a new BindingListParams object
// with the ability to set a timeout on a request.
func NewBindingListParamsWithTimeout(timeout time.Duration) *BindingListParams {
	return &BindingListParams{
		timeout: timeout,
	}
}

// NewBindingListParamsWithContext creates a new BindingListParams object
// with the ability to set a context for a request.
func NewBindingListParamsWithContext(ctx context.Context) *BindingListParams {
	return &BindingListParams{
		Context: ctx,
	}
}

// NewBindingListParamsWithHTTPClient creates a new BindingListParams object
// with the ability to set a custom HTTPClient for a request.
func NewBindingListParamsWithHTTPClient(client *http.Client) *BindingListParams {
	return &BindingListParams{
		HTTPClient: client,
	}
}

/* BindingListParams contains all the parameters to send to the API endpoint
   for the binding list operation.

   Typically these are written to a http.Request.
*/
type BindingListParams struct {

	/* DeviceID.

	   Device ID


	   Format: uuid
	*/
	DeviceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the binding list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BindingListParams) WithDefaults() *BindingListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the binding list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BindingListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the binding list params
func (o *BindingListParams) WithTimeout(timeout time.Duration) *BindingListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the binding list params
func (o *BindingListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the binding list params
func (o *BindingListParams) WithContext(ctx context.Context) *BindingListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the binding list params
func (o *BindingListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the binding list params
func (o *BindingListParams) WithHTTPClient(client *http.Client) *BindingListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the binding list params
func (o *BindingListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the binding list params
func (o *BindingListParams) WithDeviceID(deviceID strfmt.UUID) *BindingListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the binding list params
func (o *BindingListParams) SetDeviceID(deviceID strfmt.UUID) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *BindingListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceID
	if err := r.SetPathParam("deviceID", o.DeviceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
