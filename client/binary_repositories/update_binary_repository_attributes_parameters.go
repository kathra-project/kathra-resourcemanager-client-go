// Code generated by go-swagger; DO NOT EDIT.

package binary_repositories

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

// NewUpdateBinaryRepositoryAttributesParams creates a new UpdateBinaryRepositoryAttributesParams object
// with the default values initialized.
func NewUpdateBinaryRepositoryAttributesParams() *UpdateBinaryRepositoryAttributesParams {
	var ()
	return &UpdateBinaryRepositoryAttributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBinaryRepositoryAttributesParamsWithTimeout creates a new UpdateBinaryRepositoryAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateBinaryRepositoryAttributesParamsWithTimeout(timeout time.Duration) *UpdateBinaryRepositoryAttributesParams {
	var ()
	return &UpdateBinaryRepositoryAttributesParams{

		timeout: timeout,
	}
}

// NewUpdateBinaryRepositoryAttributesParamsWithContext creates a new UpdateBinaryRepositoryAttributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateBinaryRepositoryAttributesParamsWithContext(ctx context.Context) *UpdateBinaryRepositoryAttributesParams {
	var ()
	return &UpdateBinaryRepositoryAttributesParams{

		Context: ctx,
	}
}

// NewUpdateBinaryRepositoryAttributesParamsWithHTTPClient creates a new UpdateBinaryRepositoryAttributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateBinaryRepositoryAttributesParamsWithHTTPClient(client *http.Client) *UpdateBinaryRepositoryAttributesParams {
	var ()
	return &UpdateBinaryRepositoryAttributesParams{
		HTTPClient: client,
	}
}

/*UpdateBinaryRepositoryAttributesParams contains all the parameters to send to the API endpoint
for the update binary repository attributes operation typically these are written to a http.Request
*/
type UpdateBinaryRepositoryAttributesParams struct {

	/*Binaryrepository
	  BinaryRepository object to be updated

	*/
	Binaryrepository models.BinaryRepository
	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) WithTimeout(timeout time.Duration) *UpdateBinaryRepositoryAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) WithContext(ctx context.Context) *UpdateBinaryRepositoryAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) WithHTTPClient(client *http.Client) *UpdateBinaryRepositoryAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBinaryrepository adds the binaryrepository to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) WithBinaryrepository(binaryrepository models.BinaryRepository) *UpdateBinaryRepositoryAttributesParams {
	o.SetBinaryrepository(binaryrepository)
	return o
}

// SetBinaryrepository adds the binaryrepository to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) SetBinaryrepository(binaryrepository models.BinaryRepository) {
	o.Binaryrepository = binaryrepository
}

// WithResourceID adds the resourceID to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) WithResourceID(resourceID string) *UpdateBinaryRepositoryAttributesParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update binary repository attributes params
func (o *UpdateBinaryRepositoryAttributesParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBinaryRepositoryAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Binaryrepository != nil {
		if err := r.SetBodyParam(o.Binaryrepository); err != nil {
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
