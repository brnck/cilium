// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cilium/cilium/api/v1/models"
)

// DeleteFqdnCacheReader is a Reader for the DeleteFqdnCache structure.
type DeleteFqdnCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFqdnCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteFqdnCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteFqdnCacheBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFqdnCacheOK creates a DeleteFqdnCacheOK with default headers values
func NewDeleteFqdnCacheOK() *DeleteFqdnCacheOK {
	return &DeleteFqdnCacheOK{}
}

/*DeleteFqdnCacheOK handles this case with default header values.

Success
*/
type DeleteFqdnCacheOK struct {
}

func (o *DeleteFqdnCacheOK) Error() string {
	return fmt.Sprintf("[DELETE /fqdn/cache][%d] deleteFqdnCacheOK ", 200)
}

func (o *DeleteFqdnCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFqdnCacheBadRequest creates a DeleteFqdnCacheBadRequest with default headers values
func NewDeleteFqdnCacheBadRequest() *DeleteFqdnCacheBadRequest {
	return &DeleteFqdnCacheBadRequest{}
}

/*DeleteFqdnCacheBadRequest handles this case with default header values.

Invalid request (error parsing parameters)
*/
type DeleteFqdnCacheBadRequest struct {
	Payload models.Error
}

func (o *DeleteFqdnCacheBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fqdn/cache][%d] deleteFqdnCacheBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFqdnCacheBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
