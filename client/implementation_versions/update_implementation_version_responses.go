// Code generated by go-swagger; DO NOT EDIT.

package implementation_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	ImplementationVersion "github.com/kathra-project/kathra-core-model-go/models"
)

// UpdateImplementationVersionReader is a Reader for the UpdateImplementationVersion structure.
type UpdateImplementationVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateImplementationVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateImplementationVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateImplementationVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateImplementationVersionOK creates a UpdateImplementationVersionOK with default headers values
func NewUpdateImplementationVersionOK() *UpdateImplementationVersionOK {
	return &UpdateImplementationVersionOK{}
}

/*UpdateImplementationVersionOK handles this case with default header values.

Returns the modified object
*/
type UpdateImplementationVersionOK struct {
	Payload ImplementationVersion.ImplementationVersion
}

func (o *UpdateImplementationVersionOK) Error() string {
	return fmt.Sprintf("[PUT /implementationversions/{resourceId}][%d] updateImplementationVersionOK  %+v", 200, o.Payload)
}

func (o *UpdateImplementationVersionOK) GetPayload() ImplementationVersion.ImplementationVersion {
	return o.Payload
}

func (o *UpdateImplementationVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImplementationVersionUnauthorized creates a UpdateImplementationVersionUnauthorized with default headers values
func NewUpdateImplementationVersionUnauthorized() *UpdateImplementationVersionUnauthorized {
	return &UpdateImplementationVersionUnauthorized{}
}

/*UpdateImplementationVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateImplementationVersionUnauthorized struct {
}

func (o *UpdateImplementationVersionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /implementationversions/{resourceId}][%d] updateImplementationVersionUnauthorized ", 401)
}

func (o *UpdateImplementationVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
