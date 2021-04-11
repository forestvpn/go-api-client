// Code generated by go-swagger; DO NOT EDIT.

package wireguard

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

// NewPeersListParams creates a new PeersListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPeersListParams() *PeersListParams {
	return &PeersListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPeersListParamsWithTimeout creates a new PeersListParams object
// with the ability to set a timeout on a request.
func NewPeersListParamsWithTimeout(timeout time.Duration) *PeersListParams {
	return &PeersListParams{
		timeout: timeout,
	}
}

// NewPeersListParamsWithContext creates a new PeersListParams object
// with the ability to set a context for a request.
func NewPeersListParamsWithContext(ctx context.Context) *PeersListParams {
	return &PeersListParams{
		Context: ctx,
	}
}

// NewPeersListParamsWithHTTPClient creates a new PeersListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPeersListParamsWithHTTPClient(client *http.Client) *PeersListParams {
	return &PeersListParams{
		HTTPClient: client,
	}
}

/* PeersListParams contains all the parameters to send to the API endpoint
   for the peers list operation.

   Typically these are written to a http.Request.
*/
type PeersListParams struct {

	// XRelayHostname.
	XRelayHostname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the peers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PeersListParams) WithDefaults() *PeersListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the peers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PeersListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the peers list params
func (o *PeersListParams) WithTimeout(timeout time.Duration) *PeersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peers list params
func (o *PeersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peers list params
func (o *PeersListParams) WithContext(ctx context.Context) *PeersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peers list params
func (o *PeersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peers list params
func (o *PeersListParams) WithHTTPClient(client *http.Client) *PeersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peers list params
func (o *PeersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRelayHostname adds the xRelayHostname to the peers list params
func (o *PeersListParams) WithXRelayHostname(xRelayHostname string) *PeersListParams {
	o.SetXRelayHostname(xRelayHostname)
	return o
}

// SetXRelayHostname adds the xRelayHostname to the peers list params
func (o *PeersListParams) SetXRelayHostname(xRelayHostname string) {
	o.XRelayHostname = xRelayHostname
}

// WriteToRequest writes these params to a swagger request
func (o *PeersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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