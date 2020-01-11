// Code generated by go-swagger; DO NOT EDIT.

package binary_repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// AddBinaryRepositoryReader is a Reader for the AddBinaryRepository structure.
type AddBinaryRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddBinaryRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddBinaryRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddBinaryRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddBinaryRepositoryOK creates a AddBinaryRepositoryOK with default headers values
func NewAddBinaryRepositoryOK() *AddBinaryRepositoryOK {
	return &AddBinaryRepositoryOK{}
}

/*AddBinaryRepositoryOK handles this case with default header values.

Returns the created object
*/
type AddBinaryRepositoryOK struct {
	Payload *models.BinaryRepository
}

func (o *AddBinaryRepositoryOK) Error() string {
	return fmt.Sprintf("[POST /binaryrepositories][%d] addBinaryRepositoryOK  %+v", 200, o.Payload)
}

func (o *AddBinaryRepositoryOK) GetPayload() *models.BinaryRepository {
	return o.Payload
}

func (o *AddBinaryRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BinaryRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddBinaryRepositoryUnauthorized creates a AddBinaryRepositoryUnauthorized with default headers values
func NewAddBinaryRepositoryUnauthorized() *AddBinaryRepositoryUnauthorized {
	return &AddBinaryRepositoryUnauthorized{}
}

/*AddBinaryRepositoryUnauthorized handles this case with default header values.

Unauthorized
*/
type AddBinaryRepositoryUnauthorized struct {
}

func (o *AddBinaryRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /binaryrepositories][%d] addBinaryRepositoryUnauthorized ", 401)
}

func (o *AddBinaryRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
