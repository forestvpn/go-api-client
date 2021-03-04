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

// NewDeviceListParams creates a new DeviceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeviceListParams() *DeviceListParams {
	return &DeviceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeviceListParamsWithTimeout creates a new DeviceListParams object
// with the ability to set a timeout on a request.
func NewDeviceListParamsWithTimeout(timeout time.Duration) *DeviceListParams {
	return &DeviceListParams{
		timeout: timeout,
	}
}

// NewDeviceListParamsWithContext creates a new DeviceListParams object
// with the ability to set a context for a request.
func NewDeviceListParamsWithContext(ctx context.Context) *DeviceListParams {
	return &DeviceListParams{
		Context: ctx,
	}
}

// NewDeviceListParamsWithHTTPClient creates a new DeviceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeviceListParamsWithHTTPClient(client *http.Client) *DeviceListParams {
	return &DeviceListParams{
		HTTPClient: client,
	}
}

/* DeviceListParams contains all the parameters to send to the API endpoint
   for the device list operation.

   Typically these are written to a http.Request.
*/
type DeviceListParams struct {

	// XRelayHostname.
	XRelayHostname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the device list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeviceListParams) WithDefaults() *DeviceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the device list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeviceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the device list params
func (o *DeviceListParams) WithTimeout(timeout time.Duration) *DeviceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the device list params
func (o *DeviceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the device list params
func (o *DeviceListParams) WithContext(ctx context.Context) *DeviceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the device list params
func (o *DeviceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the device list params
func (o *DeviceListParams) WithHTTPClient(client *http.Client) *DeviceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the device list params
func (o *DeviceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRelayHostname adds the xRelayHostname to the device list params
func (o *DeviceListParams) WithXRelayHostname(xRelayHostname string) *DeviceListParams {
	o.SetXRelayHostname(xRelayHostname)
	return o
}

// SetXRelayHostname adds the xRelayHostname to the device list params
func (o *DeviceListParams) SetXRelayHostname(xRelayHostname string) {
	o.XRelayHostname = xRelayHostname
}

// WriteToRequest writes these params to a swagger request
func (o *DeviceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Relay-Hostname
	if err := r.SetHeaderParam("X-Relay-Hostname", o.XRelayHostname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
