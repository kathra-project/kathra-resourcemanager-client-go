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

// UpdateImplementationReader is a Reader for the UpdateImplementation structure.
type UpdateImplementationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateImplementationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateImplementationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateImplementationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateImplementationOK creates a UpdateImplementationOK with default headers values
func NewUpdateImplementationOK() *UpdateImplementationOK {
	return &UpdateImplementationOK{}
}

/*UpdateImplementationOK handles this case with default header values.

Returns the modified object
*/
type UpdateImplementationOK struct {
	Payload Implementation.Implementation
}

func (o *UpdateImplementationOK) Error() string {
	return fmt.Sprintf("[PUT /implementations/{resourceId}][%d] updateImplementationOK  %+v", 200, o.Payload)
}

func (o *UpdateImplementationOK) GetPayload() Implementation.Implementation {
	return o.Payload
}

func (o *UpdateImplementationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImplementationUnauthorized creates a UpdateImplementationUnauthorized with default headers values
func NewUpdateImplementationUnauthorized() *UpdateImplementationUnauthorized {
	return &UpdateImplementationUnauthorized{}
}

/*UpdateImplementationUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateImplementationUnauthorized struct {
}

func (o *UpdateImplementationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /implementations/{resourceId}][%d] updateImplementationUnauthorized ", 401)
}

func (o *UpdateImplementationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
