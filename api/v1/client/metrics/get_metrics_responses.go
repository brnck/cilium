// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cilium/cilium/api/v1/models"
)

// GetMetricsReader is a Reader for the GetMetrics structure.
type GetMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMetricsOK creates a GetMetricsOK with default headers values
func NewGetMetricsOK() *GetMetricsOK {
	return &GetMetricsOK{}
}

/*GetMetricsOK handles this case with default header values.

Success
*/
type GetMetricsOK struct {
	Payload []*models.Metric
}

func (o *GetMetricsOK) Error() string {
	return fmt.Sprintf("[GET /metrics/][%d] getMetricsOK  %+v", 200, o.Payload)
}

func (o *GetMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsInternalServerError creates a GetMetricsInternalServerError with default headers values
func NewGetMetricsInternalServerError() *GetMetricsInternalServerError {
	return &GetMetricsInternalServerError{}
}

/*GetMetricsInternalServerError handles this case with default header values.

Metrics cannot be retrieved
*/
type GetMetricsInternalServerError struct {
}

func (o *GetMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /metrics/][%d] getMetricsInternalServerError ", 500)
}

func (o *GetMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
