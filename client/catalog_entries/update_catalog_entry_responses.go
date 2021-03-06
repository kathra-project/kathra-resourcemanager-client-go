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

// UpdateCatalogEntryReader is a Reader for the UpdateCatalogEntry structure.
type UpdateCatalogEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCatalogEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCatalogEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateCatalogEntryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCatalogEntryOK creates a UpdateCatalogEntryOK with default headers values
func NewUpdateCatalogEntryOK() *UpdateCatalogEntryOK {
	return &UpdateCatalogEntryOK{}
}

/*UpdateCatalogEntryOK handles this case with default header values.

Returns the modified object
*/
type UpdateCatalogEntryOK struct {
	Payload CatalogEntry.CatalogEntry
}

func (o *UpdateCatalogEntryOK) Error() string {
	return fmt.Sprintf("[PUT /catalogentries/{resourceId}][%d] updateCatalogEntryOK  %+v", 200, o.Payload)
}

func (o *UpdateCatalogEntryOK) GetPayload() CatalogEntry.CatalogEntry {
	return o.Payload
}

func (o *UpdateCatalogEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCatalogEntryUnauthorized creates a UpdateCatalogEntryUnauthorized with default headers values
func NewUpdateCatalogEntryUnauthorized() *UpdateCatalogEntryUnauthorized {
	return &UpdateCatalogEntryUnauthorized{}
}

/*UpdateCatalogEntryUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateCatalogEntryUnauthorized struct {
}

func (o *UpdateCatalogEntryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /catalogentries/{resourceId}][%d] updateCatalogEntryUnauthorized ", 401)
}

func (o *UpdateCatalogEntryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
