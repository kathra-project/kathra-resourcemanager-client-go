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

// GetCatalogEntryPackageReader is a Reader for the GetCatalogEntryPackage structure.
type GetCatalogEntryPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogEntryPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCatalogEntryPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCatalogEntryPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCatalogEntryPackageOK creates a GetCatalogEntryPackageOK with default headers values
func NewGetCatalogEntryPackageOK() *GetCatalogEntryPackageOK {
	return &GetCatalogEntryPackageOK{}
}

/*GetCatalogEntryPackageOK handles this case with default header values.

Returns the object
*/
type GetCatalogEntryPackageOK struct {
	Payload CatalogEntryPackage.CatalogEntryPackage
}

func (o *GetCatalogEntryPackageOK) Error() string {
	return fmt.Sprintf("[GET /catalogentrypackages/{resourceId}][%d] getCatalogEntryPackageOK  %+v", 200, o.Payload)
}

func (o *GetCatalogEntryPackageOK) GetPayload() CatalogEntryPackage.CatalogEntryPackage {
	return o.Payload
}

func (o *GetCatalogEntryPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogEntryPackageUnauthorized creates a GetCatalogEntryPackageUnauthorized with default headers values
func NewGetCatalogEntryPackageUnauthorized() *GetCatalogEntryPackageUnauthorized {
	return &GetCatalogEntryPackageUnauthorized{}
}

/*GetCatalogEntryPackageUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCatalogEntryPackageUnauthorized struct {
}

func (o *GetCatalogEntryPackageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /catalogentrypackages/{resourceId}][%d] getCatalogEntryPackageUnauthorized ", 401)
}

func (o *GetCatalogEntryPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
