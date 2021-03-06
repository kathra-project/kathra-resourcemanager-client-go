// Code generated by go-swagger; DO NOT EDIT.

package catalog_entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog entries API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog entries API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddCatalogEntry adds a new catalogentry
*/
func (a *Client) AddCatalogEntry(params *AddCatalogEntryParams, authInfo runtime.ClientAuthInfoWriter) (*AddCatalogEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddCatalogEntryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addCatalogEntry",
		Method:             "POST",
		PathPattern:        "/catalogentries",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddCatalogEntryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddCatalogEntryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addCatalogEntry: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteCatalogEntry deletes a registered catalogentry
*/
func (a *Client) DeleteCatalogEntry(params *DeleteCatalogEntryParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCatalogEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCatalogEntryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCatalogEntry",
		Method:             "DELETE",
		PathPattern:        "/catalogentries/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCatalogEntryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCatalogEntryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCatalogEntry: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCatalogEntries retrieves a list of accessible catalogentries for authenticated user
*/
func (a *Client) GetCatalogEntries(params *GetCatalogEntriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogEntriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogEntriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCatalogEntries",
		Method:             "GET",
		PathPattern:        "/catalogentries",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogEntriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogEntriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalogEntries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCatalogEntry retrieves a specific catalogentry object
*/
func (a *Client) GetCatalogEntry(params *GetCatalogEntryParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogEntryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCatalogEntry",
		Method:             "GET",
		PathPattern:        "/catalogentries/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogEntryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogEntryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalogEntry: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCatalogEntry fullies update a registered catalogentry
*/
func (a *Client) UpdateCatalogEntry(params *UpdateCatalogEntryParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCatalogEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCatalogEntryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCatalogEntry",
		Method:             "PUT",
		PathPattern:        "/catalogentries/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCatalogEntryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCatalogEntryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCatalogEntry: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCatalogEntryAttributes partiallies update a registered catalogentry
*/
func (a *Client) UpdateCatalogEntryAttributes(params *UpdateCatalogEntryAttributesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCatalogEntryAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCatalogEntryAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCatalogEntryAttributes",
		Method:             "PATCH",
		PathPattern:        "/catalogentries/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCatalogEntryAttributesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCatalogEntryAttributesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCatalogEntryAttributes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
