// Code generated by go-swagger; DO NOT EDIT.

package source_repositories

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

// NewUpdateSourceRepositoryAttributesParams creates a new UpdateSourceRepositoryAttributesParams object
// with the default values initialized.
func NewUpdateSourceRepositoryAttributesParams() *UpdateSourceRepositoryAttributesParams {
	var ()
	return &UpdateSourceRepositoryAttributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSourceRepositoryAttributesParamsWithTimeout creates a new UpdateSourceRepositoryAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSourceRepositoryAttributesParamsWithTimeout(timeout time.Duration) *UpdateSourceRepositoryAttributesParams {
	var ()
	return &UpdateSourceRepositoryAttributesParams{

		timeout: timeout,
	}
}

// NewUpdateSourceRepositoryAttributesParamsWithContext creates a new UpdateSourceRepositoryAttributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSourceRepositoryAttributesParamsWithContext(ctx context.Context) *UpdateSourceRepositoryAttributesParams {
	var ()
	return &UpdateSourceRepositoryAttributesParams{

		Context: ctx,
	}
}

// NewUpdateSourceRepositoryAttributesParamsWithHTTPClient creates a new UpdateSourceRepositoryAttributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSourceRepositoryAttributesParamsWithHTTPClient(client *http.Client) *UpdateSourceRepositoryAttributesParams {
	var ()
	return &UpdateSourceRepositoryAttributesParams{
		HTTPClient: client,
	}
}

/*UpdateSourceRepositoryAttributesParams contains all the parameters to send to the API endpoint
for the update source repository attributes operation typically these are written to a http.Request
*/
type UpdateSourceRepositoryAttributesParams struct {

	/*ResourceID
	  resource's id

	*/
	ResourceID string
	/*Sourcerepository
	  SourceRepository object to be updated

	*/
	Sourcerepository *models.SourceRepository

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) WithTimeout(timeout time.Duration) *UpdateSourceRepositoryAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) WithContext(ctx context.Context) *UpdateSourceRepositoryAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) WithHTTPClient(client *http.Client) *UpdateSourceRepositoryAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) WithResourceID(resourceID string) *UpdateSourceRepositoryAttributesParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithSourcerepository adds the sourcerepository to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) WithSourcerepository(sourcerepository *models.SourceRepository) *UpdateSourceRepositoryAttributesParams {
	o.SetSourcerepository(sourcerepository)
	return o
}

// SetSourcerepository adds the sourcerepository to the update source repository attributes params
func (o *UpdateSourceRepositoryAttributesParams) SetSourcerepository(sourcerepository *models.SourceRepository) {
	o.Sourcerepository = sourcerepository
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSourceRepositoryAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID); err != nil {
		return err
	}

	if o.Sourcerepository != nil {
		if err := r.SetBodyParam(o.Sourcerepository); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
