// Code generated by go-swagger; DO NOT EDIT.

package api_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	ImplementationVersion "github.com/kathra-project/kathra-core-model-go/models"
)

// UpdateAPIVersionReader is a Reader for the UpdateAPIVersion structure.
type UpdateAPIVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAPIVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAPIVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateAPIVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAPIVersionOK creates a UpdateAPIVersionOK with default headers values
func NewUpdateAPIVersionOK() *UpdateAPIVersionOK {
	return &UpdateAPIVersionOK{}
}

/*UpdateAPIVersionOK handles this case with default header values.

Returns the modified object
*/
type UpdateAPIVersionOK struct {
	Payload ImplementationVersion.ImplementationVersion
}

func (o *UpdateAPIVersionOK) Error() string {
	return fmt.Sprintf("[PUT /apiversions/{resourceId}][%d] updateApiVersionOK  %+v", 200, o.Payload)
}

func (o *UpdateAPIVersionOK) GetPayload() ImplementationVersion.ImplementationVersion {
	return o.Payload
}

func (o *UpdateAPIVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPIVersionUnauthorized creates a UpdateAPIVersionUnauthorized with default headers values
func NewUpdateAPIVersionUnauthorized() *UpdateAPIVersionUnauthorized {
	return &UpdateAPIVersionUnauthorized{}
}

/*UpdateAPIVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateAPIVersionUnauthorized struct {
}

func (o *UpdateAPIVersionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apiversions/{resourceId}][%d] updateApiVersionUnauthorized ", 401)
}

func (o *UpdateAPIVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}