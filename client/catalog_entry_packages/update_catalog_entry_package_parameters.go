// Code generated by go-swagger; DO NOT EDIT.

package catalog_entry_packages

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

// NewUpdateCatalogEntryPackageParams creates a new UpdateCatalogEntryPackageParams object
// with the default values initialized.
func NewUpdateCatalogEntryPackageParams() *UpdateCatalogEntryPackageParams {
	var ()
	return &UpdateCatalogEntryPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCatalogEntryPackageParamsWithTimeout creates a new UpdateCatalogEntryPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCatalogEntryPackageParamsWithTimeout(timeout time.Duration) *UpdateCatalogEntryPackageParams {
	var ()
	return &UpdateCatalogEntryPackageParams{

		timeout: timeout,
	}
}

// NewUpdateCatalogEntryPackageParamsWithContext creates a new UpdateCatalogEntryPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCatalogEntryPackageParamsWithContext(ctx context.Context) *UpdateCatalogEntryPackageParams {
	var ()
	return &UpdateCatalogEntryPackageParams{

		Context: ctx,
	}
}

// NewUpdateCatalogEntryPackageParamsWithHTTPClient creates a new UpdateCatalogEntryPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCatalogEntryPackageParamsWithHTTPClient(client *http.Client) *UpdateCatalogEntryPackageParams {
	var ()
	return &UpdateCatalogEntryPackageParams{
		HTTPClient: client,
	}
}

/*UpdateCatalogEntryPackageParams contains all the parameters to send to the API endpoint
for the update catalog entry package operation typically these are written to a http.Request
*/
type UpdateCatalogEntryPackageParams struct {

	/*Catalogentrypackage
	  CatalogEntryPackage object to be updated

	*/
	Catalogentrypackage models.CatalogEntryPackage
	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) WithTimeout(timeout time.Duration) *UpdateCatalogEntryPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) WithContext(ctx context.Context) *UpdateCatalogEntryPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) WithHTTPClient(client *http.Client) *UpdateCatalogEntryPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogentrypackage adds the catalogentrypackage to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) WithCatalogentrypackage(catalogentrypackage models.CatalogEntryPackage) *UpdateCatalogEntryPackageParams {
	o.SetCatalogentrypackage(catalogentrypackage)
	return o
}

// SetCatalogentrypackage adds the catalogentrypackage to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) SetCatalogentrypackage(catalogentrypackage models.CatalogEntryPackage) {
	o.Catalogentrypackage = catalogentrypackage
}

// WithResourceID adds the resourceID to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) WithResourceID(resourceID string) *UpdateCatalogEntryPackageParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update catalog entry package params
func (o *UpdateCatalogEntryPackageParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCatalogEntryPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Catalogentrypackage != nil {
		if err := r.SetBodyParam(o.Catalogentrypackage); err != nil {
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
