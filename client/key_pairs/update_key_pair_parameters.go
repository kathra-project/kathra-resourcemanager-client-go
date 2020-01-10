// Code generated by go-swagger; DO NOT EDIT.

package key_pairs

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

	KeyPair "github.com/kathra-project/kathra-core-model-go/models"
)

// NewUpdateKeyPairParams creates a new UpdateKeyPairParams object
// with the default values initialized.
func NewUpdateKeyPairParams() *UpdateKeyPairParams {
	var ()
	return &UpdateKeyPairParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateKeyPairParamsWithTimeout creates a new UpdateKeyPairParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateKeyPairParamsWithTimeout(timeout time.Duration) *UpdateKeyPairParams {
	var ()
	return &UpdateKeyPairParams{

		timeout: timeout,
	}
}

// NewUpdateKeyPairParamsWithContext creates a new UpdateKeyPairParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateKeyPairParamsWithContext(ctx context.Context) *UpdateKeyPairParams {
	var ()
	return &UpdateKeyPairParams{

		Context: ctx,
	}
}

// NewUpdateKeyPairParamsWithHTTPClient creates a new UpdateKeyPairParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateKeyPairParamsWithHTTPClient(client *http.Client) *UpdateKeyPairParams {
	var ()
	return &UpdateKeyPairParams{
		HTTPClient: client,
	}
}

/*UpdateKeyPairParams contains all the parameters to send to the API endpoint
for the update key pair operation typically these are written to a http.Request
*/
type UpdateKeyPairParams struct {

	/*Keypair
	  KeyPair object to be updated

	*/
	Keypair KeyPair.KeyPair
	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update key pair params
func (o *UpdateKeyPairParams) WithTimeout(timeout time.Duration) *UpdateKeyPairParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update key pair params
func (o *UpdateKeyPairParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update key pair params
func (o *UpdateKeyPairParams) WithContext(ctx context.Context) *UpdateKeyPairParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update key pair params
func (o *UpdateKeyPairParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update key pair params
func (o *UpdateKeyPairParams) WithHTTPClient(client *http.Client) *UpdateKeyPairParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update key pair params
func (o *UpdateKeyPairParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeypair adds the keypair to the update key pair params
func (o *UpdateKeyPairParams) WithKeypair(keypair KeyPair.KeyPair) *UpdateKeyPairParams {
	o.SetKeypair(keypair)
	return o
}

// SetKeypair adds the keypair to the update key pair params
func (o *UpdateKeyPairParams) SetKeypair(keypair KeyPair.KeyPair) {
	o.Keypair = keypair
}

// WithResourceID adds the resourceID to the update key pair params
func (o *UpdateKeyPairParams) WithResourceID(resourceID string) *UpdateKeyPairParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update key pair params
func (o *UpdateKeyPairParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateKeyPairParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Keypair != nil {
		if err := r.SetBodyParam(o.Keypair); err != nil {
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
