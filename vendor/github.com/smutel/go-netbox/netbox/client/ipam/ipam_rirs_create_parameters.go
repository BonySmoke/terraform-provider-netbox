// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package ipam

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

	"github.com/smutel/go-netbox/netbox/models"
)

// NewIpamRirsCreateParams creates a new IpamRirsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRirsCreateParams() *IpamRirsCreateParams {
	return &IpamRirsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRirsCreateParamsWithTimeout creates a new IpamRirsCreateParams object
// with the ability to set a timeout on a request.
func NewIpamRirsCreateParamsWithTimeout(timeout time.Duration) *IpamRirsCreateParams {
	return &IpamRirsCreateParams{
		timeout: timeout,
	}
}

// NewIpamRirsCreateParamsWithContext creates a new IpamRirsCreateParams object
// with the ability to set a context for a request.
func NewIpamRirsCreateParamsWithContext(ctx context.Context) *IpamRirsCreateParams {
	return &IpamRirsCreateParams{
		Context: ctx,
	}
}

// NewIpamRirsCreateParamsWithHTTPClient creates a new IpamRirsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRirsCreateParamsWithHTTPClient(client *http.Client) *IpamRirsCreateParams {
	return &IpamRirsCreateParams{
		HTTPClient: client,
	}
}

/* IpamRirsCreateParams contains all the parameters to send to the API endpoint
   for the ipam rirs create operation.

   Typically these are written to a http.Request.
*/
type IpamRirsCreateParams struct {

	// Data.
	Data *models.RIR

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam rirs create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsCreateParams) WithDefaults() *IpamRirsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam rirs create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam rirs create params
func (o *IpamRirsCreateParams) WithTimeout(timeout time.Duration) *IpamRirsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs create params
func (o *IpamRirsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs create params
func (o *IpamRirsCreateParams) WithContext(ctx context.Context) *IpamRirsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs create params
func (o *IpamRirsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs create params
func (o *IpamRirsCreateParams) WithHTTPClient(client *http.Client) *IpamRirsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs create params
func (o *IpamRirsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam rirs create params
func (o *IpamRirsCreateParams) WithData(data *models.RIR) *IpamRirsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam rirs create params
func (o *IpamRirsCreateParams) SetData(data *models.RIR) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRirsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
