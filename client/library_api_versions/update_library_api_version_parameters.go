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

// NewUpdateLibraryAPIVersionParams creates a new UpdateLibraryAPIVersionParams object
// with the default values initialized.
func NewUpdateLibraryAPIVersionParams() *UpdateLibraryAPIVersionParams {
	var ()
	return &UpdateLibraryAPIVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLibraryAPIVersionParamsWithTimeout creates a new UpdateLibraryAPIVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLibraryAPIVersionParamsWithTimeout(timeout time.Duration) *UpdateLibraryAPIVersionParams {
	var ()
	return &UpdateLibraryAPIVersionParams{

		timeout: timeout,
	}
}

// NewUpdateLibraryAPIVersionParamsWithContext creates a new UpdateLibraryAPIVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLibraryAPIVersionParamsWithContext(ctx context.Context) *UpdateLibraryAPIVersionParams {
	var ()
	return &UpdateLibraryAPIVersionParams{

		Context: ctx,
	}
}

// NewUpdateLibraryAPIVersionParamsWithHTTPClient creates a new UpdateLibraryAPIVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLibraryAPIVersionParamsWithHTTPClient(client *http.Client) *UpdateLibraryAPIVersionParams {
	var ()
	return &UpdateLibraryAPIVersionParams{
		HTTPClient: client,
	}
}

/*UpdateLibraryAPIVersionParams contains all the parameters to send to the API endpoint
for the update library Api version operation typically these are written to a http.Request
*/
type UpdateLibraryAPIVersionParams struct {

	/*Libraryapiversion
	  LibraryApiVersion object to be updated

	*/
	Libraryapiversion LibraryApiVersion.LibraryApiVersion
	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) WithTimeout(timeout time.Duration) *UpdateLibraryAPIVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) WithContext(ctx context.Context) *UpdateLibraryAPIVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) WithHTTPClient(client *http.Client) *UpdateLibraryAPIVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLibraryapiversion adds the libraryapiversion to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) WithLibraryapiversion(libraryapiversion LibraryApiVersion.LibraryApiVersion) *UpdateLibraryAPIVersionParams {
	o.SetLibraryapiversion(libraryapiversion)
	return o
}

// SetLibraryapiversion adds the libraryapiversion to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) SetLibraryapiversion(libraryapiversion LibraryApiVersion.LibraryApiVersion) {
	o.Libraryapiversion = libraryapiversion
}

// WithResourceID adds the resourceID to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) WithResourceID(resourceID string) *UpdateLibraryAPIVersionParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update library Api version params
func (o *UpdateLibraryAPIVersionParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLibraryAPIVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Libraryapiversion != nil {
		if err := r.SetBodyParam(o.Libraryapiversion); err != nil {
			return err
		}
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
