// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	Group "github.com/kathra-project/kathra-core-model-go/models"
)

// UpdateGroupAttributesReader is a Reader for the UpdateGroupAttributes structure.
type UpdateGroupAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGroupAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGroupAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateGroupAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateGroupAttributesOK creates a UpdateGroupAttributesOK with default headers values
func NewUpdateGroupAttributesOK() *UpdateGroupAttributesOK {
	return &UpdateGroupAttributesOK{}
}

/*UpdateGroupAttributesOK handles this case with default header values.

Returns the modified object
*/
type UpdateGroupAttributesOK struct {
	Payload Group.Group
}

func (o *UpdateGroupAttributesOK) Error() string {
	return fmt.Sprintf("[PATCH /groups/{resourceId}][%d] updateGroupAttributesOK  %+v", 200, o.Payload)
}

func (o *UpdateGroupAttributesOK) GetPayload() Group.Group {
	return o.Payload
}

func (o *UpdateGroupAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupAttributesUnauthorized creates a UpdateGroupAttributesUnauthorized with default headers values
func NewUpdateGroupAttributesUnauthorized() *UpdateGroupAttributesUnauthorized {
	return &UpdateGroupAttributesUnauthorized{}
}

/*UpdateGroupAttributesUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateGroupAttributesUnauthorized struct {
}

func (o *UpdateGroupAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /groups/{resourceId}][%d] updateGroupAttributesUnauthorized ", 401)
}

func (o *UpdateGroupAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
