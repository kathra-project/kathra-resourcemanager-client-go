// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*GetUsersOK handles this case with default header values.

List of accessible users for the authenticated user
*/
type GetUsersOK struct {
	Payload []*models.User
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUnauthorized creates a GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

/*GetUsersUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUsersUnauthorized struct {
}

func (o *GetUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersUnauthorized ", 401)
}

func (o *GetUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
