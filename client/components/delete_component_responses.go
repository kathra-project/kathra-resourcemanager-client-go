// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteComponentReader is a Reader for the DeleteComponent structure.
type DeleteComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteComponentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteComponentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteComponentOK creates a DeleteComponentOK with default headers values
func NewDeleteComponentOK() *DeleteComponentOK {
	return &DeleteComponentOK{}
}

/*DeleteComponentOK handles this case with default header values.

Object is deleted
*/
type DeleteComponentOK struct {
	Payload string
}

func (o *DeleteComponentOK) Error() string {
	return fmt.Sprintf("[DELETE /components/{resourceId}][%d] deleteComponentOK  %+v", 200, o.Payload)
}

func (o *DeleteComponentOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteComponentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteComponentUnauthorized creates a DeleteComponentUnauthorized with default headers values
func NewDeleteComponentUnauthorized() *DeleteComponentUnauthorized {
	return &DeleteComponentUnauthorized{}
}

/*DeleteComponentUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteComponentUnauthorized struct {
}

func (o *DeleteComponentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /components/{resourceId}][%d] deleteComponentUnauthorized ", 401)
}

func (o *DeleteComponentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
