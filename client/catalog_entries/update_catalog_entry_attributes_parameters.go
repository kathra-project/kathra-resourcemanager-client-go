// Code generated by go-swagger; DO NOT EDIT.

package catalog_entries

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

	CatalogEntry "github.com/kathra-project/kathra-core-model-go/models"
)

// NewUpdateCatalogEntryAttributesParams creates a new UpdateCatalogEntryAttributesParams object
// with the default values initialized.
func NewUpdateCatalogEntryAttributesParams() *UpdateCatalogEntryAttributesParams {
	var ()
	return &UpdateCatalogEntryAttributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCatalogEntryAttributesParamsWithTimeout creates a new UpdateCatalogEntryAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCatalogEntryAttributesParamsWithTimeout(timeout time.Duration) *UpdateCatalogEntryAttributesParams {
	var ()
	return &UpdateCatalogEntryAttributesParams{

		timeout: timeout,
	}
}

// NewUpdateCatalogEntryAttributesParamsWithContext creates a new UpdateCatalogEntryAttributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCatalogEntryAttributesParamsWithContext(ctx context.Context) *UpdateCatalogEntryAttributesParams {
	var ()
	return &UpdateCatalogEntryAttributesParams{

		Context: ctx,
	}
}

// NewUpdateCatalogEntryAttributesParamsWithHTTPClient creates a new UpdateCatalogEntryAttributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCatalogEntryAttributesParamsWithHTTPClient(client *http.Client) *UpdateCatalogEntryAttributesParams {
	var ()
	return &UpdateCatalogEntryAttributesParams{
		HTTPClient: client,
	}
}

/*UpdateCatalogEntryAttributesParams contains all the parameters to send to the API endpoint
for the update catalog entry attributes operation typically these are written to a http.Request
*/
type UpdateCatalogEntryAttributesParams struct {

	/*Catalogentry
	  CatalogEntry object to be updated

	*/
	Catalogentry CatalogEntry.CatalogEntry
	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) WithTimeout(timeout time.Duration) *UpdateCatalogEntryAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) WithContext(ctx context.Context) *UpdateCatalogEntryAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) WithHTTPClient(client *http.Client) *UpdateCatalogEntryAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogentry adds the catalogentry to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) WithCatalogentry(catalogentry CatalogEntry.CatalogEntry) *UpdateCatalogEntryAttributesParams {
	o.SetCatalogentry(catalogentry)
	return o
}

// SetCatalogentry adds the catalogentry to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) SetCatalogentry(catalogentry CatalogEntry.CatalogEntry) {
	o.Catalogentry = catalogentry
}

// WithResourceID adds the resourceID to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) WithResourceID(resourceID string) *UpdateCatalogEntryAttributesParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update catalog entry attributes params
func (o *UpdateCatalogEntryAttributesParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCatalogEntryAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Catalogentry != nil {
		if err := r.SetBodyParam(o.Catalogentry); err != nil {
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
