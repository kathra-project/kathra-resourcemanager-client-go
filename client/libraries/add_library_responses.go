// Code generated by go-swagger; DO NOT EDIT.

package libraries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	Library "github.com/kathra-project/kathra-core-model-go/models"
)

// AddLibraryReader is a Reader for the AddLibrary structure.
type AddLibraryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddLibraryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddLibraryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddLibraryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddLibraryOK creates a AddLibraryOK with default headers values
func NewAddLibraryOK() *AddLibraryOK {
	return &AddLibraryOK{}
}

/*AddLibraryOK handles this case with default header values.

Returns the created object
*/
type AddLibraryOK struct {
	Payload Library.Library
}

func (o *AddLibraryOK) Error() string {
	return fmt.Sprintf("[POST /libraries][%d] addLibraryOK  %+v", 200, o.Payload)
}

func (o *AddLibraryOK) GetPayload() Library.Library {
	return o.Payload
}

func (o *AddLibraryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLibraryUnauthorized creates a AddLibraryUnauthorized with default headers values
func NewAddLibraryUnauthorized() *AddLibraryUnauthorized {
	return &AddLibraryUnauthorized{}
}

/*AddLibraryUnauthorized handles this case with default header values.

Unauthorized
*/
type AddLibraryUnauthorized struct {
}

func (o *AddLibraryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /libraries][%d] addLibraryUnauthorized ", 401)
}

func (o *AddLibraryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
