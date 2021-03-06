// Code generated by go-swagger; DO NOT EDIT.

package implementations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteImplementationReader is a Reader for the DeleteImplementation structure.
type DeleteImplementationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImplementationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteImplementationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteImplementationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteImplementationOK creates a DeleteImplementationOK with default headers values
func NewDeleteImplementationOK() *DeleteImplementationOK {
	return &DeleteImplementationOK{}
}

/*DeleteImplementationOK handles this case with default header values.

Object is deleted
*/
type DeleteImplementationOK struct {
	Payload string
}

func (o *DeleteImplementationOK) Error() string {
	return fmt.Sprintf("[DELETE /implementations/{resourceId}][%d] deleteImplementationOK  %+v", 200, o.Payload)
}

func (o *DeleteImplementationOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteImplementationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImplementationUnauthorized creates a DeleteImplementationUnauthorized with default headers values
func NewDeleteImplementationUnauthorized() *DeleteImplementationUnauthorized {
	return &DeleteImplementationUnauthorized{}
}

/*DeleteImplementationUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteImplementationUnauthorized struct {
}

func (o *DeleteImplementationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /implementations/{resourceId}][%d] deleteImplementationUnauthorized ", 401)
}

func (o *DeleteImplementationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
