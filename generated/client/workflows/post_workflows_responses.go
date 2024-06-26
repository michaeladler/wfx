// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package workflows

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

// PostWorkflowsReader is a Reader for the PostWorkflows structure.
type PostWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostWorkflowsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkflowsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostWorkflowsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWorkflowsCreated creates a PostWorkflowsCreated with default headers values
func NewPostWorkflowsCreated() *PostWorkflowsCreated {
	return &PostWorkflowsCreated{}
}

/*
PostWorkflowsCreated describes a response with status code 201, with default header values.

Workflow was created
*/
type PostWorkflowsCreated struct {
	Payload *model.Workflow
}

// IsSuccess returns true when this post workflows created response has a 2xx status code
func (o *PostWorkflowsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workflows created response has a 3xx status code
func (o *PostWorkflowsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workflows created response has a 4xx status code
func (o *PostWorkflowsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workflows created response has a 5xx status code
func (o *PostWorkflowsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post workflows created response a status code equal to that given
func (o *PostWorkflowsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post workflows created response
func (o *PostWorkflowsCreated) Code() int {
	return 201
}

func (o *PostWorkflowsCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /workflows][%d] postWorkflowsCreated %s", 201, payload)
}

func (o *PostWorkflowsCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /workflows][%d] postWorkflowsCreated %s", 201, payload)
}

func (o *PostWorkflowsCreated) GetPayload() *model.Workflow {
	return o.Payload
}

func (o *PostWorkflowsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Workflow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkflowsBadRequest creates a PostWorkflowsBadRequest with default headers values
func NewPostWorkflowsBadRequest() *PostWorkflowsBadRequest {
	return &PostWorkflowsBadRequest{}
}

/*
PostWorkflowsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostWorkflowsBadRequest struct {
	Payload *model.ErrorResponse
}

// IsSuccess returns true when this post workflows bad request response has a 2xx status code
func (o *PostWorkflowsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workflows bad request response has a 3xx status code
func (o *PostWorkflowsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workflows bad request response has a 4xx status code
func (o *PostWorkflowsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workflows bad request response has a 5xx status code
func (o *PostWorkflowsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post workflows bad request response a status code equal to that given
func (o *PostWorkflowsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post workflows bad request response
func (o *PostWorkflowsBadRequest) Code() int {
	return 400
}

func (o *PostWorkflowsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /workflows][%d] postWorkflowsBadRequest %s", 400, payload)
}

func (o *PostWorkflowsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /workflows][%d] postWorkflowsBadRequest %s", 400, payload)
}

func (o *PostWorkflowsBadRequest) GetPayload() *model.ErrorResponse {
	return o.Payload
}

func (o *PostWorkflowsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkflowsDefault creates a PostWorkflowsDefault with default headers values
func NewPostWorkflowsDefault(code int) *PostWorkflowsDefault {
	return &PostWorkflowsDefault{
		_statusCode: code,
	}
}

/*
PostWorkflowsDefault describes a response with status code -1, with default header values.

Other error with any status code and response body format.
*/
type PostWorkflowsDefault struct {
	_statusCode int
}

// IsSuccess returns true when this post workflows default response has a 2xx status code
func (o *PostWorkflowsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post workflows default response has a 3xx status code
func (o *PostWorkflowsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post workflows default response has a 4xx status code
func (o *PostWorkflowsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post workflows default response has a 5xx status code
func (o *PostWorkflowsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post workflows default response a status code equal to that given
func (o *PostWorkflowsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post workflows default response
func (o *PostWorkflowsDefault) Code() int {
	return o._statusCode
}

func (o *PostWorkflowsDefault) Error() string {
	return fmt.Sprintf("[POST /workflows][%d] PostWorkflows default", o._statusCode)
}

func (o *PostWorkflowsDefault) String() string {
	return fmt.Sprintf("[POST /workflows][%d] PostWorkflows default", o._statusCode)
}

func (o *PostWorkflowsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
