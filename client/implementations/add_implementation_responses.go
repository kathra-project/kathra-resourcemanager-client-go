// Code generated by go-swagger; DO NOT EDIT.

package implementations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	Implementation "github.com/kathra-project/kathra-core-model-go/models"
)

// AddImplementationReader is a Reader for the AddImplementation structure.
type AddImplementationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddImplementationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddImplementationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddImplementationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddImplementationOK creates a AddImplementationOK with default headers values
func NewAddImplementationOK() *AddImplementationOK {
	return &AddImplementationOK{}
}

/*AddImplementationOK handles this case with default header values.

Returns the created object
*/
type AddImplementationOK struct {
	Payload Implementation.Implementation
}

func (o *AddImplementationOK) Error() string {
	return fmt.Sprintf("[POST /implementations][%d] addImplementationOK  %+v", 200, o.Payload)
}

func (o *AddImplementationOK) GetPayload() Implementation.Implementation {
	return o.Payload
}

func (o *AddImplementationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddImplementationUnauthorized creates a AddImplementationUnauthorized with default headers values
func NewAddImplementationUnauthorized() *AddImplementationUnauthorized {
	return &AddImplementationUnauthorized{}
}

/*AddImplementationUnauthorized handles this case with default header values.

Unauthorized
*/
type AddImplementationUnauthorized struct {
}

func (o *AddImplementationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /implementations][%d] addImplementationUnauthorized ", 401)
}

func (o *AddImplementationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
