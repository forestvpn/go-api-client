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

	"github.com/forestvpn/go-api-client/node/models"
)

// NewDevicesStatsUpdateParams creates a new DevicesStatsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDevicesStatsUpdateParams() *DevicesStatsUpdateParams {
	return &DevicesStatsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesStatsUpdateParamsWithTimeout creates a new DevicesStatsUpdateParams object
// with the ability to set a timeout on a request.
func NewDevicesStatsUpdateParamsWithTimeout(timeout time.Duration) *DevicesStatsUpdateParams {
	return &DevicesStatsUpdateParams{
		timeout: timeout,
	}
}

// NewDevicesStatsUpdateParamsWithContext creates a new DevicesStatsUpdateParams object
// with the ability to set a context for a request.
func NewDevicesStatsUpdateParamsWithContext(ctx context.Context) *DevicesStatsUpdateParams {
	return &DevicesStatsUpdateParams{
		Context: ctx,
	}
}

// NewDevicesStatsUpdateParamsWithHTTPClient creates a new DevicesStatsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDevicesStatsUpdateParamsWithHTTPClient(client *http.Client) *DevicesStatsUpdateParams {
	return &DevicesStatsUpdateParams{
		HTTPClient: client,
	}
}

/* DevicesStatsUpdateParams contains all the parameters to send to the API endpoint
   for the devices stats update operation.

   Typically these are written to a http.Request.
*/
type DevicesStatsUpdateParams struct {

	// XRelayHostname.
	XRelayHostname string

	// Request.
	Request models.DeviceStatsMap

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the devices stats update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DevicesStatsUpdateParams) WithDefaults() *DevicesStatsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the devices stats update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DevicesStatsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the devices stats update params
func (o *DevicesStatsUpdateParams) WithTimeout(timeout time.Duration) *DevicesStatsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices stats update params
func (o *DevicesStatsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices stats update params
func (o *DevicesStatsUpdateParams) WithContext(ctx context.Context) *DevicesStatsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices stats update params
func (o *DevicesStatsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices stats update params
func (o *DevicesStatsUpdateParams) WithHTTPClient(client *http.Client) *DevicesStatsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices stats update params
func (o *DevicesStatsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRelayHostname adds the xRelayHostname to the devices stats update params
func (o *DevicesStatsUpdateParams) WithXRelayHostname(xRelayHostname string) *DevicesStatsUpdateParams {
	o.SetXRelayHostname(xRelayHostname)
	return o
}

// SetXRelayHostname adds the xRelayHostname to the devices stats update params
func (o *DevicesStatsUpdateParams) SetXRelayHostname(xRelayHostname string) {
	o.XRelayHostname = xRelayHostname
}

// WithRequest adds the request to the devices stats update params
func (o *DevicesStatsUpdateParams) WithRequest(request models.DeviceStatsMap) *DevicesStatsUpdateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the devices stats update params
func (o *DevicesStatsUpdateParams) SetRequest(request models.DeviceStatsMap) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesStatsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Relay-Hostname
	if err := r.SetHeaderParam("X-Relay-Hostname", o.XRelayHostname); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
