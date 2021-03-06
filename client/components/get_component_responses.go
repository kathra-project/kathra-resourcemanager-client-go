// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	Component "github.com/kathra-project/kathra-core-model-go/models"
)

// GetComponentReader is a Reader for the GetComponent structure.
type GetComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComponentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetComponentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetComponentOK creates a GetComponentOK with default headers values
func NewGetComponentOK() *GetComponentOK {
	return &GetComponentOK{}
}

/*GetComponentOK handles this case with default header values.

Returns the object
*/
type GetComponentOK struct {
	Payload Component.Component
}

func (o *GetComponentOK) Error() string {
	return fmt.Sprintf("[GET /components/{resourceId}][%d] getComponentOK  %+v", 200, o.Payload)
}

func (o *GetComponentOK) GetPayload() Component.Component {
	return o.Payload
}

func (o *GetComponentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentUnauthorized creates a GetComponentUnauthorized with default headers values
func NewGetComponentUnauthorized() *GetComponentUnauthorized {
	return &GetComponentUnauthorized{}
}

/*GetComponentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetComponentUnauthorized struct {
}

func (o *GetComponentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /components/{resourceId}][%d] getComponentUnauthorized ", 401)
}

func (o *GetComponentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
