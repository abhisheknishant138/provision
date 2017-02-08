package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/galthaus/swagger-test/models"
)

// GetFileReader is a Reader for the GetFile structure.
type GetFileReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFileOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFileOK creates a GetFileOK with default headers values
func NewGetFileOK(writer io.Writer) *GetFileOK {
	return &GetFileOK{
		Payload: writer,
	}
}

/*GetFileOK handles this case with default header values.

GetFileOK get file o k
*/
type GetFileOK struct {
	Payload io.Writer
}

func (o *GetFileOK) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFileOK  %+v", 200, o.Payload)
}

func (o *GetFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileUnauthorized creates a GetFileUnauthorized with default headers values
func NewGetFileUnauthorized() *GetFileUnauthorized {
	return &GetFileUnauthorized{}
}

/*GetFileUnauthorized handles this case with default header values.

GetFileUnauthorized get file unauthorized
*/
type GetFileUnauthorized struct {
	Payload *models.Result
}

func (o *GetFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFileUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileNotFound creates a GetFileNotFound with default headers values
func NewGetFileNotFound() *GetFileNotFound {
	return &GetFileNotFound{}
}

/*GetFileNotFound handles this case with default header values.

GetFileNotFound get file not found
*/
type GetFileNotFound struct {
	Payload *models.Result
}

func (o *GetFileNotFound) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFileNotFound  %+v", 404, o.Payload)
}

func (o *GetFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileInternalServerError creates a GetFileInternalServerError with default headers values
func NewGetFileInternalServerError() *GetFileInternalServerError {
	return &GetFileInternalServerError{}
}

/*GetFileInternalServerError handles this case with default header values.

GetFileInternalServerError get file internal server error
*/
type GetFileInternalServerError struct {
	Payload *models.Result
}

func (o *GetFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
