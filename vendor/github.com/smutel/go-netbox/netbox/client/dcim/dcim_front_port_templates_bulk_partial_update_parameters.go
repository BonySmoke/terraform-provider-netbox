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

package dcim

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

// NewDcimFrontPortTemplatesBulkPartialUpdateParams creates a new DcimFrontPortTemplatesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortTemplatesBulkPartialUpdateParams() *DcimFrontPortTemplatesBulkPartialUpdateParams {
	return &DcimFrontPortTemplatesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesBulkPartialUpdateParamsWithTimeout creates a new DcimFrontPortTemplatesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortTemplatesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesBulkPartialUpdateParams {
	return &DcimFrontPortTemplatesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesBulkPartialUpdateParamsWithContext creates a new DcimFrontPortTemplatesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimFrontPortTemplatesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesBulkPartialUpdateParams {
	return &DcimFrontPortTemplatesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesBulkPartialUpdateParamsWithHTTPClient creates a new DcimFrontPortTemplatesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortTemplatesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesBulkPartialUpdateParams {
	return &DcimFrontPortTemplatesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortTemplatesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim front port templates bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortTemplatesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableFrontPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front port templates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) WithDefaults() *DcimFrontPortTemplatesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front port templates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front port templates bulk partial update params
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates bulk partial update params
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates bulk partial update params
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates bulk partial update params
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates bulk partial update params
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates bulk partial update params
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front port templates bulk partial update params
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) WithData(data *models.WritableFrontPortTemplate) *DcimFrontPortTemplatesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front port templates bulk partial update params
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) SetData(data *models.WritableFrontPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
