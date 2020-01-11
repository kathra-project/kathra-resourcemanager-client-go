// Code generated by go-swagger; DO NOT EDIT.

package libraries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// UpdateLibraryReader is a Reader for the UpdateLibrary structure.
type UpdateLibraryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLibraryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLibraryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateLibraryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateLibraryOK creates a UpdateLibraryOK with default headers values
func NewUpdateLibraryOK() *UpdateLibraryOK {
	return &UpdateLibraryOK{}
}

/*UpdateLibraryOK handles this case with default header values.

Returns the modified object
*/
type UpdateLibraryOK struct {
	Payload *models.Library
}

func (o *UpdateLibraryOK) Error() string {
	return fmt.Sprintf("[PUT /libraries/{resourceId}][%d] updateLibraryOK  %+v", 200, o.Payload)
}

func (o *UpdateLibraryOK) GetPayload() *models.Library {
	return o.Payload
}

func (o *UpdateLibraryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Library)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLibraryUnauthorized creates a UpdateLibraryUnauthorized with default headers values
func NewUpdateLibraryUnauthorized() *UpdateLibraryUnauthorized {
	return &UpdateLibraryUnauthorized{}
}

/*UpdateLibraryUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateLibraryUnauthorized struct {
}

func (o *UpdateLibraryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /libraries/{resourceId}][%d] updateLibraryUnauthorized ", 401)
}

func (o *UpdateLibraryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
