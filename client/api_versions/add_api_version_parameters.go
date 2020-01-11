// Code generated by go-swagger; DO NOT EDIT.

package api_versions

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

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// NewAddAPIVersionParams creates a new AddAPIVersionParams object
// with the default values initialized.
func NewAddAPIVersionParams() *AddAPIVersionParams {
	var ()
	return &AddAPIVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAPIVersionParamsWithTimeout creates a new AddAPIVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAPIVersionParamsWithTimeout(timeout time.Duration) *AddAPIVersionParams {
	var ()
	return &AddAPIVersionParams{

		timeout: timeout,
	}
}

// NewAddAPIVersionParamsWithContext creates a new AddAPIVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAPIVersionParamsWithContext(ctx context.Context) *AddAPIVersionParams {
	var ()
	return &AddAPIVersionParams{

		Context: ctx,
	}
}

// NewAddAPIVersionParamsWithHTTPClient creates a new AddAPIVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAPIVersionParamsWithHTTPClient(client *http.Client) *AddAPIVersionParams {
	var ()
	return &AddAPIVersionParams{
		HTTPClient: client,
	}
}

/*AddAPIVersionParams contains all the parameters to send to the API endpoint
for the add Api version operation typically these are written to a http.Request
*/
type AddAPIVersionParams struct {

	/*Apiversion
	  ApiVersion object to be created

	*/
	Apiversion *models.APIVersion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add Api version params
func (o *AddAPIVersionParams) WithTimeout(timeout time.Duration) *AddAPIVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Api version params
func (o *AddAPIVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Api version params
func (o *AddAPIVersionParams) WithContext(ctx context.Context) *AddAPIVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Api version params
func (o *AddAPIVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Api version params
func (o *AddAPIVersionParams) WithHTTPClient(client *http.Client) *AddAPIVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Api version params
func (o *AddAPIVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApiversion adds the apiversion to the add Api version params
func (o *AddAPIVersionParams) WithApiversion(apiversion *models.APIVersion) *AddAPIVersionParams {
	o.SetApiversion(apiversion)
	return o
}

// SetApiversion adds the apiversion to the add Api version params
func (o *AddAPIVersionParams) SetApiversion(apiversion *models.APIVersion) {
	o.Apiversion = apiversion
}

// WriteToRequest writes these params to a swagger request
func (o *AddAPIVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Apiversion != nil {
		if err := r.SetBodyParam(o.Apiversion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
