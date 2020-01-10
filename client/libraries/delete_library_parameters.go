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
)

// NewDeleteLibraryParams creates a new DeleteLibraryParams object
// with the default values initialized.
func NewDeleteLibraryParams() *DeleteLibraryParams {
	var ()
	return &DeleteLibraryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLibraryParamsWithTimeout creates a new DeleteLibraryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLibraryParamsWithTimeout(timeout time.Duration) *DeleteLibraryParams {
	var ()
	return &DeleteLibraryParams{

		timeout: timeout,
	}
}

// NewDeleteLibraryParamsWithContext creates a new DeleteLibraryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLibraryParamsWithContext(ctx context.Context) *DeleteLibraryParams {
	var ()
	return &DeleteLibraryParams{

		Context: ctx,
	}
}

// NewDeleteLibraryParamsWithHTTPClient creates a new DeleteLibraryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLibraryParamsWithHTTPClient(client *http.Client) *DeleteLibraryParams {
	var ()
	return &DeleteLibraryParams{
		HTTPClient: client,
	}
}

/*DeleteLibraryParams contains all the parameters to send to the API endpoint
for the delete library operation typically these are written to a http.Request
*/
type DeleteLibraryParams struct {

	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete library params
func (o *DeleteLibraryParams) WithTimeout(timeout time.Duration) *DeleteLibraryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete library params
func (o *DeleteLibraryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete library params
func (o *DeleteLibraryParams) WithContext(ctx context.Context) *DeleteLibraryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete library params
func (o *DeleteLibraryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete library params
func (o *DeleteLibraryParams) WithHTTPClient(client *http.Client) *DeleteLibraryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete library params
func (o *DeleteLibraryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the delete library params
func (o *DeleteLibraryParams) WithResourceID(resourceID string) *DeleteLibraryParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the delete library params
func (o *DeleteLibraryParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLibraryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}