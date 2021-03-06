// Code generated by go-swagger; DO NOT EDIT.

package key_pairs

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

// NewGetKeyPairsParams creates a new GetKeyPairsParams object
// with the default values initialized.
func NewGetKeyPairsParams() *GetKeyPairsParams {

	return &GetKeyPairsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeyPairsParamsWithTimeout creates a new GetKeyPairsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeyPairsParamsWithTimeout(timeout time.Duration) *GetKeyPairsParams {

	return &GetKeyPairsParams{

		timeout: timeout,
	}
}

// NewGetKeyPairsParamsWithContext creates a new GetKeyPairsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeyPairsParamsWithContext(ctx context.Context) *GetKeyPairsParams {

	return &GetKeyPairsParams{

		Context: ctx,
	}
}

// NewGetKeyPairsParamsWithHTTPClient creates a new GetKeyPairsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeyPairsParamsWithHTTPClient(client *http.Client) *GetKeyPairsParams {

	return &GetKeyPairsParams{
		HTTPClient: client,
	}
}

/*GetKeyPairsParams contains all the parameters to send to the API endpoint
for the get key pairs operation typically these are written to a http.Request
*/
type GetKeyPairsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get key pairs params
func (o *GetKeyPairsParams) WithTimeout(timeout time.Duration) *GetKeyPairsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get key pairs params
func (o *GetKeyPairsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get key pairs params
func (o *GetKeyPairsParams) WithContext(ctx context.Context) *GetKeyPairsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get key pairs params
func (o *GetKeyPairsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get key pairs params
func (o *GetKeyPairsParams) WithHTTPClient(client *http.Client) *GetKeyPairsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get key pairs params
func (o *GetKeyPairsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeyPairsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
