// Code generated by go-swagger; DO NOT EDIT.

package library_api_versions

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

	LibraryApiVersion "github.com/kathra-project/kathra-core-model-go/models"
)

// NewAddLibraryAPIVersionParams creates a new AddLibraryAPIVersionParams object
// with the default values initialized.
func NewAddLibraryAPIVersionParams() *AddLibraryAPIVersionParams {
	var ()
	return &AddLibraryAPIVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddLibraryAPIVersionParamsWithTimeout creates a new AddLibraryAPIVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddLibraryAPIVersionParamsWithTimeout(timeout time.Duration) *AddLibraryAPIVersionParams {
	var ()
	return &AddLibraryAPIVersionParams{

		timeout: timeout,
	}
}

// NewAddLibraryAPIVersionParamsWithContext creates a new AddLibraryAPIVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddLibraryAPIVersionParamsWithContext(ctx context.Context) *AddLibraryAPIVersionParams {
	var ()
	return &AddLibraryAPIVersionParams{

		Context: ctx,
	}
}

// NewAddLibraryAPIVersionParamsWithHTTPClient creates a new AddLibraryAPIVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddLibraryAPIVersionParamsWithHTTPClient(client *http.Client) *AddLibraryAPIVersionParams {
	var ()
	return &AddLibraryAPIVersionParams{
		HTTPClient: client,
	}
}

/*AddLibraryAPIVersionParams contains all the parameters to send to the API endpoint
for the add library Api version operation typically these are written to a http.Request
*/
type AddLibraryAPIVersionParams struct {

	/*Libraryapiversion
	  LibraryApiVersion object to be created

	*/
	Libraryapiversion LibraryApiVersion.LibraryApiVersion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add library Api version params
func (o *AddLibraryAPIVersionParams) WithTimeout(timeout time.Duration) *AddLibraryAPIVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add library Api version params
func (o *AddLibraryAPIVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add library Api version params
func (o *AddLibraryAPIVersionParams) WithContext(ctx context.Context) *AddLibraryAPIVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add library Api version params
func (o *AddLibraryAPIVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add library Api version params
func (o *AddLibraryAPIVersionParams) WithHTTPClient(client *http.Client) *AddLibraryAPIVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add library Api version params
func (o *AddLibraryAPIVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLibraryapiversion adds the libraryapiversion to the add library Api version params
func (o *AddLibraryAPIVersionParams) WithLibraryapiversion(libraryapiversion LibraryApiVersion.LibraryApiVersion) *AddLibraryAPIVersionParams {
	o.SetLibraryapiversion(libraryapiversion)
	return o
}

// SetLibraryapiversion adds the libraryapiversion to the add library Api version params
func (o *AddLibraryAPIVersionParams) SetLibraryapiversion(libraryapiversion LibraryApiVersion.LibraryApiVersion) {
	o.Libraryapiversion = libraryapiversion
}

// WriteToRequest writes these params to a swagger request
func (o *AddLibraryAPIVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Libraryapiversion != nil {
		if err := r.SetBodyParam(o.Libraryapiversion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}