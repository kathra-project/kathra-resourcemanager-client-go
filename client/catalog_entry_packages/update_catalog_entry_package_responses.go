// Code generated by go-swagger; DO NOT EDIT.

package catalog_entry_packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// UpdateCatalogEntryPackageReader is a Reader for the UpdateCatalogEntryPackage structure.
type UpdateCatalogEntryPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCatalogEntryPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCatalogEntryPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateCatalogEntryPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCatalogEntryPackageOK creates a UpdateCatalogEntryPackageOK with default headers values
func NewUpdateCatalogEntryPackageOK() *UpdateCatalogEntryPackageOK {
	return &UpdateCatalogEntryPackageOK{}
}

/*UpdateCatalogEntryPackageOK handles this case with default header values.

Returns the modified object
*/
type UpdateCatalogEntryPackageOK struct {
	Payload *models.CatalogEntryPackage
}

func (o *UpdateCatalogEntryPackageOK) Error() string {
	return fmt.Sprintf("[PUT /catalogentrypackages/{resourceId}][%d] updateCatalogEntryPackageOK  %+v", 200, o.Payload)
}

func (o *UpdateCatalogEntryPackageOK) GetPayload() *models.CatalogEntryPackage {
	return o.Payload
}

func (o *UpdateCatalogEntryPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogEntryPackage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCatalogEntryPackageUnauthorized creates a UpdateCatalogEntryPackageUnauthorized with default headers values
func NewUpdateCatalogEntryPackageUnauthorized() *UpdateCatalogEntryPackageUnauthorized {
	return &UpdateCatalogEntryPackageUnauthorized{}
}

/*UpdateCatalogEntryPackageUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateCatalogEntryPackageUnauthorized struct {
}

func (o *UpdateCatalogEntryPackageUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /catalogentrypackages/{resourceId}][%d] updateCatalogEntryPackageUnauthorized ", 401)
}

func (o *UpdateCatalogEntryPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
