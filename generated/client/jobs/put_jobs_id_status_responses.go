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

// PutJobsIDStatusReader is a Reader for the PutJobsIDStatus structure.
type PutJobsIDStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutJobsIDStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutJobsIDStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutJobsIDStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutJobsIDStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutJobsIDStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutJobsIDStatusOK creates a PutJobsIDStatusOK with default headers values
func NewPutJobsIDStatusOK() *PutJobsIDStatusOK {
	return &PutJobsIDStatusOK{}
}

/*
PutJobsIDStatusOK describes a response with status code 200, with default header values.

Job modified successfully
*/
type PutJobsIDStatusOK struct {
	Payload *model.JobStatus
}

// IsSuccess returns true when this put jobs Id status o k response has a 2xx status code
func (o *PutJobsIDStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put jobs Id status o k response has a 3xx status code
func (o *PutJobsIDStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put jobs Id status o k response has a 4xx status code
func (o *PutJobsIDStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put jobs Id status o k response has a 5xx status code
func (o *PutJobsIDStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put jobs Id status o k response a status code equal to that given
func (o *PutJobsIDStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put jobs Id status o k response
func (o *PutJobsIDStatusOK) Code() int {
	return 200
}

func (o *PutJobsIDStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /jobs/{id}/status][%d] putJobsIdStatusOK %s", 200, payload)
}

func (o *PutJobsIDStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /jobs/{id}/status][%d] putJobsIdStatusOK %s", 200, payload)
}

func (o *PutJobsIDStatusOK) GetPayload() *model.JobStatus {
	return o.Payload
}

func (o *PutJobsIDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutJobsIDStatusBadRequest creates a PutJobsIDStatusBadRequest with default headers values
func NewPutJobsIDStatusBadRequest() *PutJobsIDStatusBadRequest {
	return &PutJobsIDStatusBadRequest{}
}

/*
PutJobsIDStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutJobsIDStatusBadRequest struct {
	Payload *model.ErrorResponse
}

// IsSuccess returns true when this put jobs Id status bad request response has a 2xx status code
func (o *PutJobsIDStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put jobs Id status bad request response has a 3xx status code
func (o *PutJobsIDStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put jobs Id status bad request response has a 4xx status code
func (o *PutJobsIDStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put jobs Id status bad request response has a 5xx status code
func (o *PutJobsIDStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put jobs Id status bad request response a status code equal to that given
func (o *PutJobsIDStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put jobs Id status bad request response
func (o *PutJobsIDStatusBadRequest) Code() int {
	return 400
}

func (o *PutJobsIDStatusBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /jobs/{id}/status][%d] putJobsIdStatusBadRequest %s", 400, payload)
}

func (o *PutJobsIDStatusBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /jobs/{id}/status][%d] putJobsIdStatusBadRequest %s", 400, payload)
}

func (o *PutJobsIDStatusBadRequest) GetPayload() *model.ErrorResponse {
	return o.Payload
}

func (o *PutJobsIDStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutJobsIDStatusNotFound creates a PutJobsIDStatusNotFound with default headers values
func NewPutJobsIDStatusNotFound() *PutJobsIDStatusNotFound {
	return &PutJobsIDStatusNotFound{}
}

/*
PutJobsIDStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutJobsIDStatusNotFound struct {
	Payload *model.ErrorResponse
}

// IsSuccess returns true when this put jobs Id status not found response has a 2xx status code
func (o *PutJobsIDStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put jobs Id status not found response has a 3xx status code
func (o *PutJobsIDStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put jobs Id status not found response has a 4xx status code
func (o *PutJobsIDStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put jobs Id status not found response has a 5xx status code
func (o *PutJobsIDStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put jobs Id status not found response a status code equal to that given
func (o *PutJobsIDStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put jobs Id status not found response
func (o *PutJobsIDStatusNotFound) Code() int {
	return 404
}

func (o *PutJobsIDStatusNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /jobs/{id}/status][%d] putJobsIdStatusNotFound %s", 404, payload)
}

func (o *PutJobsIDStatusNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /jobs/{id}/status][%d] putJobsIdStatusNotFound %s", 404, payload)
}

func (o *PutJobsIDStatusNotFound) GetPayload() *model.ErrorResponse {
	return o.Payload
}

func (o *PutJobsIDStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutJobsIDStatusDefault creates a PutJobsIDStatusDefault with default headers values
func NewPutJobsIDStatusDefault(code int) *PutJobsIDStatusDefault {
	return &PutJobsIDStatusDefault{
		_statusCode: code,
	}
}

/*
PutJobsIDStatusDefault describes a response with status code -1, with default header values.

Other error with any status code and response body format.
*/
type PutJobsIDStatusDefault struct {
	_statusCode int
}

// IsSuccess returns true when this put jobs ID status default response has a 2xx status code
func (o *PutJobsIDStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put jobs ID status default response has a 3xx status code
func (o *PutJobsIDStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put jobs ID status default response has a 4xx status code
func (o *PutJobsIDStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put jobs ID status default response has a 5xx status code
func (o *PutJobsIDStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put jobs ID status default response a status code equal to that given
func (o *PutJobsIDStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put jobs ID status default response
func (o *PutJobsIDStatusDefault) Code() int {
	return o._statusCode
}

func (o *PutJobsIDStatusDefault) Error() string {
	return fmt.Sprintf("[PUT /jobs/{id}/status][%d] PutJobsIDStatus default", o._statusCode)
}

func (o *PutJobsIDStatusDefault) String() string {
	return fmt.Sprintf("[PUT /jobs/{id}/status][%d] PutJobsIDStatus default", o._statusCode)
}

func (o *PutJobsIDStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
