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

	CatalogEntryPackage "github.com/kathra-project/kathra-core-model-go/models"
)

// NewAddCatalogEntryPackageParams creates a new AddCatalogEntryPackageParams object
// with the default values initialized.
func NewAddCatalogEntryPackageParams() *AddCatalogEntryPackageParams {
	var ()
	return &AddCatalogEntryPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCatalogEntryPackageParamsWithTimeout creates a new AddCatalogEntryPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCatalogEntryPackageParamsWithTimeout(timeout time.Duration) *AddCatalogEntryPackageParams {
	var ()
	return &AddCatalogEntryPackageParams{

		timeout: timeout,
	}
}

// NewAddCatalogEntryPackageParamsWithContext creates a new AddCatalogEntryPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCatalogEntryPackageParamsWithContext(ctx context.Context) *AddCatalogEntryPackageParams {
	var ()
	return &AddCatalogEntryPackageParams{

		Context: ctx,
	}
}

// NewAddCatalogEntryPackageParamsWithHTTPClient creates a new AddCatalogEntryPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCatalogEntryPackageParamsWithHTTPClient(client *http.Client) *AddCatalogEntryPackageParams {
	var ()
	return &AddCatalogEntryPackageParams{
		HTTPClient: client,
	}
}

/*AddCatalogEntryPackageParams contains all the parameters to send to the API endpoint
for the add catalog entry package operation typically these are written to a http.Request
*/
type AddCatalogEntryPackageParams struct {

	/*Catalogentrypackage
	  CatalogEntryPackage object to be created

	*/
	Catalogentrypackage CatalogEntryPackage.CatalogEntryPackage

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add catalog entry package params
func (o *AddCatalogEntryPackageParams) WithTimeout(timeout time.Duration) *AddCatalogEntryPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add catalog entry package params
func (o *AddCatalogEntryPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add catalog entry package params
func (o *AddCatalogEntryPackageParams) WithContext(ctx context.Context) *AddCatalogEntryPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add catalog entry package params
func (o *AddCatalogEntryPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add catalog entry package params
func (o *AddCatalogEntryPackageParams) WithHTTPClient(client *http.Client) *AddCatalogEntryPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add catalog entry package params
func (o *AddCatalogEntryPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogentrypackage adds the catalogentrypackage to the add catalog entry package params
func (o *AddCatalogEntryPackageParams) WithCatalogentrypackage(catalogentrypackage CatalogEntryPackage.CatalogEntryPackage) *AddCatalogEntryPackageParams {
	o.SetCatalogentrypackage(catalogentrypackage)
	return o
}

// SetCatalogentrypackage adds the catalogentrypackage to the add catalog entry package params
func (o *AddCatalogEntryPackageParams) SetCatalogentrypackage(catalogentrypackage CatalogEntryPackage.CatalogEntryPackage) {
	o.Catalogentrypackage = catalogentrypackage
}

// WriteToRequest writes these params to a swagger request
func (o *AddCatalogEntryPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Catalogentrypackage != nil {
		if err := r.SetBodyParam(o.Catalogentrypackage); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
