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

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// NewUpdateLibraryAPIVersionAttributesParams creates a new UpdateLibraryAPIVersionAttributesParams object
// with the default values initialized.
func NewUpdateLibraryAPIVersionAttributesParams() *UpdateLibraryAPIVersionAttributesParams {
	var ()
	return &UpdateLibraryAPIVersionAttributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLibraryAPIVersionAttributesParamsWithTimeout creates a new UpdateLibraryAPIVersionAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLibraryAPIVersionAttributesParamsWithTimeout(timeout time.Duration) *UpdateLibraryAPIVersionAttributesParams {
	var ()
	return &UpdateLibraryAPIVersionAttributesParams{

		timeout: timeout,
	}
}

// NewUpdateLibraryAPIVersionAttributesParamsWithContext creates a new UpdateLibraryAPIVersionAttributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLibraryAPIVersionAttributesParamsWithContext(ctx context.Context) *UpdateLibraryAPIVersionAttributesParams {
	var ()
	return &UpdateLibraryAPIVersionAttributesParams{

		Context: ctx,
	}
}

// NewUpdateLibraryAPIVersionAttributesParamsWithHTTPClient creates a new UpdateLibraryAPIVersionAttributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLibraryAPIVersionAttributesParamsWithHTTPClient(client *http.Client) *UpdateLibraryAPIVersionAttributesParams {
	var ()
	return &UpdateLibraryAPIVersionAttributesParams{
		HTTPClient: client,
	}
}

/*UpdateLibraryAPIVersionAttributesParams contains all the parameters to send to the API endpoint
for the update library Api version attributes operation typically these are written to a http.Request
*/
type UpdateLibraryAPIVersionAttributesParams struct {

	/*Libraryapiversion
	  LibraryApiVersion object to be updated

	*/
	Libraryapiversion *models.LibraryAPIVersion
	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) WithTimeout(timeout time.Duration) *UpdateLibraryAPIVersionAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) WithContext(ctx context.Context) *UpdateLibraryAPIVersionAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) WithHTTPClient(client *http.Client) *UpdateLibraryAPIVersionAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLibraryapiversion adds the libraryapiversion to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) WithLibraryapiversion(libraryapiversion *models.LibraryAPIVersion) *UpdateLibraryAPIVersionAttributesParams {
	o.SetLibraryapiversion(libraryapiversion)
	return o
}

// SetLibraryapiversion adds the libraryapiversion to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) SetLibraryapiversion(libraryapiversion *models.LibraryAPIVersion) {
	o.Libraryapiversion = libraryapiversion
}

// WithResourceID adds the resourceID to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) WithResourceID(resourceID string) *UpdateLibraryAPIVersionAttributesParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update library Api version attributes params
func (o *UpdateLibraryAPIVersionAttributesParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLibraryAPIVersionAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
