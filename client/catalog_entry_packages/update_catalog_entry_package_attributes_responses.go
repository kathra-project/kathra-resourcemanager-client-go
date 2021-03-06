// Code generated by go-swagger; DO NOT EDIT.

package catalog_entry_packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	CatalogEntryPackage "github.com/kathra-project/kathra-core-model-go/models"
)

// UpdateCatalogEntryPackageAttributesReader is a Reader for the UpdateCatalogEntryPackageAttributes structure.
type UpdateCatalogEntryPackageAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCatalogEntryPackageAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCatalogEntryPackageAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateCatalogEntryPackageAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCatalogEntryPackageAttributesOK creates a UpdateCatalogEntryPackageAttributesOK with default headers values
func NewUpdateCatalogEntryPackageAttributesOK() *UpdateCatalogEntryPackageAttributesOK {
	return &UpdateCatalogEntryPackageAttributesOK{}
}

/*UpdateCatalogEntryPackageAttributesOK handles this case with default header values.

Returns the modified object
*/
type UpdateCatalogEntryPackageAttributesOK struct {
	Payload CatalogEntryPackage.CatalogEntryPackage
}

func (o *UpdateCatalogEntryPackageAttributesOK) Error() string {
	return fmt.Sprintf("[PATCH /catalogentrypackages/{resourceId}][%d] updateCatalogEntryPackageAttributesOK  %+v", 200, o.Payload)
}

func (o *UpdateCatalogEntryPackageAttributesOK) GetPayload() CatalogEntryPackage.CatalogEntryPackage {
	return o.Payload
}

func (o *UpdateCatalogEntryPackageAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCatalogEntryPackageAttributesUnauthorized creates a UpdateCatalogEntryPackageAttributesUnauthorized with default headers values
func NewUpdateCatalogEntryPackageAttributesUnauthorized() *UpdateCatalogEntryPackageAttributesUnauthorized {
	return &UpdateCatalogEntryPackageAttributesUnauthorized{}
}

/*UpdateCatalogEntryPackageAttributesUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateCatalogEntryPackageAttributesUnauthorized struct {
}

func (o *UpdateCatalogEntryPackageAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /catalogentrypackages/{resourceId}][%d] updateCatalogEntryPackageAttributesUnauthorized ", 401)
}

func (o *UpdateCatalogEntryPackageAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
