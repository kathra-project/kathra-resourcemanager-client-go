// Code generated by go-swagger; DO NOT EDIT.

package source_repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	SourceRepository "github.com/kathra-project/kathra-core-model-go/models"
)

// GetSourceRepositoryReader is a Reader for the GetSourceRepository structure.
type GetSourceRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSourceRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSourceRepositoryOK creates a GetSourceRepositoryOK with default headers values
func NewGetSourceRepositoryOK() *GetSourceRepositoryOK {
	return &GetSourceRepositoryOK{}
}

/*GetSourceRepositoryOK handles this case with default header values.

Returns the object
*/
type GetSourceRepositoryOK struct {
	Payload SourceRepository.SourceRepository
}

func (o *GetSourceRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /sourcerepositories/{resourceId}][%d] getSourceRepositoryOK  %+v", 200, o.Payload)
}

func (o *GetSourceRepositoryOK) GetPayload() SourceRepository.SourceRepository {
	return o.Payload
}

func (o *GetSourceRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceRepositoryUnauthorized creates a GetSourceRepositoryUnauthorized with default headers values
func NewGetSourceRepositoryUnauthorized() *GetSourceRepositoryUnauthorized {
	return &GetSourceRepositoryUnauthorized{}
}

/*GetSourceRepositoryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSourceRepositoryUnauthorized struct {
}

func (o *GetSourceRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sourcerepositories/{resourceId}][%d] getSourceRepositoryUnauthorized ", 401)
}

func (o *GetSourceRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}