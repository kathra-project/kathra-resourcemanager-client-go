// Code generated by go-swagger; DO NOT EDIT.

package implementation_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// AddImplementationVersionReader is a Reader for the AddImplementationVersion structure.
type AddImplementationVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddImplementationVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddImplementationVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddImplementationVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddImplementationVersionOK creates a AddImplementationVersionOK with default headers values
func NewAddImplementationVersionOK() *AddImplementationVersionOK {
	return &AddImplementationVersionOK{}
}

/*AddImplementationVersionOK handles this case with default header values.

Returns the created object
*/
type AddImplementationVersionOK struct {
	Payload *models.ImplementationVersion
}

func (o *AddImplementationVersionOK) Error() string {
	return fmt.Sprintf("[POST /implementationversions][%d] addImplementationVersionOK  %+v", 200, o.Payload)
}

func (o *AddImplementationVersionOK) GetPayload() *models.ImplementationVersion {
	return o.Payload
}

func (o *AddImplementationVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImplementationVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddImplementationVersionUnauthorized creates a AddImplementationVersionUnauthorized with default headers values
func NewAddImplementationVersionUnauthorized() *AddImplementationVersionUnauthorized {
	return &AddImplementationVersionUnauthorized{}
}

/*AddImplementationVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type AddImplementationVersionUnauthorized struct {
}

func (o *AddImplementationVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /implementationversions][%d] addImplementationVersionUnauthorized ", 401)
}

func (o *AddImplementationVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
