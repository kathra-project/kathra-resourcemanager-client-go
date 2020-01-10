// Code generated by go-swagger; DO NOT EDIT.

package catalog_entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	CatalogEntry "github.com/kathra-project/kathra-core-model-go/models"
)

// UpdateCatalogEntryAttributesReader is a Reader for the UpdateCatalogEntryAttributes structure.
type UpdateCatalogEntryAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCatalogEntryAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCatalogEntryAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateCatalogEntryAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCatalogEntryAttributesOK creates a UpdateCatalogEntryAttributesOK with default headers values
func NewUpdateCatalogEntryAttributesOK() *UpdateCatalogEntryAttributesOK {
	return &UpdateCatalogEntryAttributesOK{}
}

/*UpdateCatalogEntryAttributesOK handles this case with default header values.

Returns the modified object
*/
type UpdateCatalogEntryAttributesOK struct {
	Payload CatalogEntry.CatalogEntry
}

func (o *UpdateCatalogEntryAttributesOK) Error() string {
	return fmt.Sprintf("[PATCH /catalogentries/{resourceId}][%d] updateCatalogEntryAttributesOK  %+v", 200, o.Payload)
}

func (o *UpdateCatalogEntryAttributesOK) GetPayload() CatalogEntry.CatalogEntry {
	return o.Payload
}

func (o *UpdateCatalogEntryAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCatalogEntryAttributesUnauthorized creates a UpdateCatalogEntryAttributesUnauthorized with default headers values
func NewUpdateCatalogEntryAttributesUnauthorized() *UpdateCatalogEntryAttributesUnauthorized {
	return &UpdateCatalogEntryAttributesUnauthorized{}
}

/*UpdateCatalogEntryAttributesUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateCatalogEntryAttributesUnauthorized struct {
}

func (o *UpdateCatalogEntryAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /catalogentries/{resourceId}][%d] updateCatalogEntryAttributesUnauthorized ", 401)
}

func (o *UpdateCatalogEntryAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
