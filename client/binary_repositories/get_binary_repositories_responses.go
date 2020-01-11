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

// GetBinaryRepositoriesReader is a Reader for the GetBinaryRepositories structure.
type GetBinaryRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBinaryRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBinaryRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBinaryRepositoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBinaryRepositoriesOK creates a GetBinaryRepositoriesOK with default headers values
func NewGetBinaryRepositoriesOK() *GetBinaryRepositoriesOK {
	return &GetBinaryRepositoriesOK{}
}

/*GetBinaryRepositoriesOK handles this case with default header values.

List of accessible binaryrepositories for the authenticated user
*/
type GetBinaryRepositoriesOK struct {
	Payload []models.BinaryRepository
}

func (o *GetBinaryRepositoriesOK) Error() string {
	return fmt.Sprintf("[GET /binaryrepositories][%d] getBinaryRepositoriesOK  %+v", 200, o.Payload)
}

func (o *GetBinaryRepositoriesOK) GetPayload() []models.BinaryRepository {
	return o.Payload
}

func (o *GetBinaryRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBinaryRepositoriesUnauthorized creates a GetBinaryRepositoriesUnauthorized with default headers values
func NewGetBinaryRepositoriesUnauthorized() *GetBinaryRepositoriesUnauthorized {
	return &GetBinaryRepositoriesUnauthorized{}
}

/*GetBinaryRepositoriesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetBinaryRepositoriesUnauthorized struct {
}

func (o *GetBinaryRepositoriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /binaryrepositories][%d] getBinaryRepositoriesUnauthorized ", 401)
}

func (o *GetBinaryRepositoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
