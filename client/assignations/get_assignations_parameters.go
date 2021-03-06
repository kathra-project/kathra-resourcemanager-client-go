// Code generated by go-swagger; DO NOT EDIT.

package assignations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAssignationsParams creates a new GetAssignationsParams object
// with the default values initialized.
func NewGetAssignationsParams() *GetAssignationsParams {

	return &GetAssignationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssignationsParamsWithTimeout creates a new GetAssignationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAssignationsParamsWithTimeout(timeout time.Duration) *GetAssignationsParams {

	return &GetAssignationsParams{

		timeout: timeout,
	}
}

// NewGetAssignationsParamsWithContext creates a new GetAssignationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAssignationsParamsWithContext(ctx context.Context) *GetAssignationsParams {

	return &GetAssignationsParams{

		Context: ctx,
	}
}

// NewGetAssignationsParamsWithHTTPClient creates a new GetAssignationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAssignationsParamsWithHTTPClient(client *http.Client) *GetAssignationsParams {

	return &GetAssignationsParams{
		HTTPClient: client,
	}
}

/*GetAssignationsParams contains all the parameters to send to the API endpoint
for the get assignations operation typically these are written to a http.Request
*/
type GetAssignationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get assignations params
func (o *GetAssignationsParams) WithTimeout(timeout time.Duration) *GetAssignationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get assignations params
func (o *GetAssignationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get assignations params
func (o *GetAssignationsParams) WithContext(ctx context.Context) *GetAssignationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get assignations params
func (o *GetAssignationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get assignations params
func (o *GetAssignationsParams) WithHTTPClient(client *http.Client) *GetAssignationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get assignations params
func (o *GetAssignationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssignationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
