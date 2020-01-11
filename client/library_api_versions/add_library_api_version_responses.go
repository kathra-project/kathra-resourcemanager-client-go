// Code generated by go-swagger; DO NOT EDIT.

package library_api_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// AddLibraryAPIVersionReader is a Reader for the AddLibraryAPIVersion structure.
type AddLibraryAPIVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddLibraryAPIVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddLibraryAPIVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddLibraryAPIVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddLibraryAPIVersionOK creates a AddLibraryAPIVersionOK with default headers values
func NewAddLibraryAPIVersionOK() *AddLibraryAPIVersionOK {
	return &AddLibraryAPIVersionOK{}
}

/*AddLibraryAPIVersionOK handles this case with default header values.

Returns the created object
*/
type AddLibraryAPIVersionOK struct {
	Payload models.LibraryAPIVersion
}

func (o *AddLibraryAPIVersionOK) Error() string {
	return fmt.Sprintf("[POST /libraryapiversions][%d] addLibraryApiVersionOK  %+v", 200, o.Payload)
}

func (o *AddLibraryAPIVersionOK) GetPayload() models.LibraryAPIVersion {
	return o.Payload
}

func (o *AddLibraryAPIVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLibraryAPIVersionUnauthorized creates a AddLibraryAPIVersionUnauthorized with default headers values
func NewAddLibraryAPIVersionUnauthorized() *AddLibraryAPIVersionUnauthorized {
	return &AddLibraryAPIVersionUnauthorized{}
}

/*AddLibraryAPIVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type AddLibraryAPIVersionUnauthorized struct {
}

func (o *AddLibraryAPIVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /libraryapiversions][%d] addLibraryApiVersionUnauthorized ", 401)
}

func (o *AddLibraryAPIVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
