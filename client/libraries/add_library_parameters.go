// Code generated by go-swagger; DO NOT EDIT.

package libraries

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

// NewAddLibraryParams creates a new AddLibraryParams object
// with the default values initialized.
func NewAddLibraryParams() *AddLibraryParams {
	var ()
	return &AddLibraryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddLibraryParamsWithTimeout creates a new AddLibraryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddLibraryParamsWithTimeout(timeout time.Duration) *AddLibraryParams {
	var ()
	return &AddLibraryParams{

		timeout: timeout,
	}
}

// NewAddLibraryParamsWithContext creates a new AddLibraryParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddLibraryParamsWithContext(ctx context.Context) *AddLibraryParams {
	var ()
	return &AddLibraryParams{

		Context: ctx,
	}
}

// NewAddLibraryParamsWithHTTPClient creates a new AddLibraryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddLibraryParamsWithHTTPClient(client *http.Client) *AddLibraryParams {
	var ()
	return &AddLibraryParams{
		HTTPClient: client,
	}
}

/*AddLibraryParams contains all the parameters to send to the API endpoint
for the add library operation typically these are written to a http.Request
*/
type AddLibraryParams struct {

	/*Library
	  Library object to be created

	*/
	Library models.Library

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add library params
func (o *AddLibraryParams) WithTimeout(timeout time.Duration) *AddLibraryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add library params
func (o *AddLibraryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add library params
func (o *AddLibraryParams) WithContext(ctx context.Context) *AddLibraryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add library params
func (o *AddLibraryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add library params
func (o *AddLibraryParams) WithHTTPClient(client *http.Client) *AddLibraryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add library params
func (o *AddLibraryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLibrary adds the library to the add library params
func (o *AddLibraryParams) WithLibrary(library models.Library) *AddLibraryParams {
	o.SetLibrary(library)
	return o
}

// SetLibrary adds the library to the add library params
func (o *AddLibraryParams) SetLibrary(library models.Library) {
	o.Library = library
}

// WriteToRequest writes these params to a swagger request
func (o *AddLibraryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Library != nil {
		if err := r.SetBodyParam(o.Library); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
