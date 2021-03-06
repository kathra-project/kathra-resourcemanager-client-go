// Code generated by go-swagger; DO NOT EDIT.

package key_pairs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	KeyPair "github.com/kathra-project/kathra-core-model-go/models"
)

// UpdateKeyPairAttributesReader is a Reader for the UpdateKeyPairAttributes structure.
type UpdateKeyPairAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateKeyPairAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateKeyPairAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateKeyPairAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateKeyPairAttributesOK creates a UpdateKeyPairAttributesOK with default headers values
func NewUpdateKeyPairAttributesOK() *UpdateKeyPairAttributesOK {
	return &UpdateKeyPairAttributesOK{}
}

/*UpdateKeyPairAttributesOK handles this case with default header values.

Returns the modified object
*/
type UpdateKeyPairAttributesOK struct {
	Payload KeyPair.KeyPair
}

func (o *UpdateKeyPairAttributesOK) Error() string {
	return fmt.Sprintf("[PATCH /keypairs/{resourceId}][%d] updateKeyPairAttributesOK  %+v", 200, o.Payload)
}

func (o *UpdateKeyPairAttributesOK) GetPayload() KeyPair.KeyPair {
	return o.Payload
}

func (o *UpdateKeyPairAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateKeyPairAttributesUnauthorized creates a UpdateKeyPairAttributesUnauthorized with default headers values
func NewUpdateKeyPairAttributesUnauthorized() *UpdateKeyPairAttributesUnauthorized {
	return &UpdateKeyPairAttributesUnauthorized{}
}

/*UpdateKeyPairAttributesUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateKeyPairAttributesUnauthorized struct {
}

func (o *UpdateKeyPairAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /keypairs/{resourceId}][%d] updateKeyPairAttributesUnauthorized ", 401)
}

func (o *UpdateKeyPairAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
