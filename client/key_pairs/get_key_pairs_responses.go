// Code generated by go-swagger; DO NOT EDIT.

package key_pairs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// GetKeyPairsReader is a Reader for the GetKeyPairs structure.
type GetKeyPairsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeyPairsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKeyPairsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetKeyPairsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKeyPairsOK creates a GetKeyPairsOK with default headers values
func NewGetKeyPairsOK() *GetKeyPairsOK {
	return &GetKeyPairsOK{}
}

/*GetKeyPairsOK handles this case with default header values.

List of accessible keypairs for the authenticated user
*/
type GetKeyPairsOK struct {
	Payload []models.KeyPair
}

func (o *GetKeyPairsOK) Error() string {
	return fmt.Sprintf("[GET /keypairs][%d] getKeyPairsOK  %+v", 200, o.Payload)
}

func (o *GetKeyPairsOK) GetPayload() []models.KeyPair {
	return o.Payload
}

func (o *GetKeyPairsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeyPairsUnauthorized creates a GetKeyPairsUnauthorized with default headers values
func NewGetKeyPairsUnauthorized() *GetKeyPairsUnauthorized {
	return &GetKeyPairsUnauthorized{}
}

/*GetKeyPairsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetKeyPairsUnauthorized struct {
}

func (o *GetKeyPairsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /keypairs][%d] getKeyPairsUnauthorized ", 401)
}

func (o *GetKeyPairsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
