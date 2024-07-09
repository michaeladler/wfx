// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/siemens/wfx/generated/model"
)

// GetJobsIDStatusReader is a Reader for the GetJobsIDStatus structure.
type GetJobsIDStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobsIDStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobsIDStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetJobsIDStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetJobsIDStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetJobsIDStatusOK creates a GetJobsIDStatusOK with default headers values
func NewGetJobsIDStatusOK() *GetJobsIDStatusOK {
	return &GetJobsIDStatusOK{}
}

/*
GetJobsIDStatusOK describes a response with status code 200, with default header values.

Job status
*/
type GetJobsIDStatusOK struct {
	Payload *model.JobStatus
}

// IsSuccess returns true when this get jobs Id status o k response has a 2xx status code
func (o *GetJobsIDStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get jobs Id status o k response has a 3xx status code
func (o *GetJobsIDStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get jobs Id status o k response has a 4xx status code
func (o *GetJobsIDStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get jobs Id status o k response has a 5xx status code
func (o *GetJobsIDStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get jobs Id status o k response a status code equal to that given
func (o *GetJobsIDStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get jobs Id status o k response
func (o *GetJobsIDStatusOK) Code() int {
	return 200
}

func (o *GetJobsIDStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /jobs/{id}/status][%d] getJobsIdStatusOK %s", 200, payload)
}

func (o *GetJobsIDStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /jobs/{id}/status][%d] getJobsIdStatusOK %s", 200, payload)
}

func (o *GetJobsIDStatusOK) GetPayload() *model.JobStatus {
	return o.Payload
}

func (o *GetJobsIDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsIDStatusNotFound creates a GetJobsIDStatusNotFound with default headers values
func NewGetJobsIDStatusNotFound() *GetJobsIDStatusNotFound {
	return &GetJobsIDStatusNotFound{}
}

/*
GetJobsIDStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetJobsIDStatusNotFound struct {
	Payload *model.ErrorResponse
}

// IsSuccess returns true when this get jobs Id status not found response has a 2xx status code
func (o *GetJobsIDStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get jobs Id status not found response has a 3xx status code
func (o *GetJobsIDStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get jobs Id status not found response has a 4xx status code
func (o *GetJobsIDStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get jobs Id status not found response has a 5xx status code
func (o *GetJobsIDStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get jobs Id status not found response a status code equal to that given
func (o *GetJobsIDStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get jobs Id status not found response
func (o *GetJobsIDStatusNotFound) Code() int {
	return 404
}

func (o *GetJobsIDStatusNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /jobs/{id}/status][%d] getJobsIdStatusNotFound %s", 404, payload)
}

func (o *GetJobsIDStatusNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /jobs/{id}/status][%d] getJobsIdStatusNotFound %s", 404, payload)
}

func (o *GetJobsIDStatusNotFound) GetPayload() *model.ErrorResponse {
	return o.Payload
}

func (o *GetJobsIDStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsIDStatusDefault creates a GetJobsIDStatusDefault with default headers values
func NewGetJobsIDStatusDefault(code int) *GetJobsIDStatusDefault {
	return &GetJobsIDStatusDefault{
		_statusCode: code,
	}
}

/*
GetJobsIDStatusDefault describes a response with status code -1, with default header values.

Other error with any status code and response body format.
*/
type GetJobsIDStatusDefault struct {
	_statusCode int
}

// IsSuccess returns true when this get jobs ID status default response has a 2xx status code
func (o *GetJobsIDStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get jobs ID status default response has a 3xx status code
func (o *GetJobsIDStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get jobs ID status default response has a 4xx status code
func (o *GetJobsIDStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get jobs ID status default response has a 5xx status code
func (o *GetJobsIDStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get jobs ID status default response a status code equal to that given
func (o *GetJobsIDStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get jobs ID status default response
func (o *GetJobsIDStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetJobsIDStatusDefault) Error() string {
	return fmt.Sprintf("[GET /jobs/{id}/status][%d] GetJobsIDStatus default", o._statusCode)
}

func (o *GetJobsIDStatusDefault) String() string {
	return fmt.Sprintf("[GET /jobs/{id}/status][%d] GetJobsIDStatus default", o._statusCode)
}

func (o *GetJobsIDStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}