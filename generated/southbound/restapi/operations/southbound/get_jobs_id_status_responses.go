// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package southbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/siemens/wfx/generated/model"
)

// GetJobsIDStatusOKCode is the HTTP code returned for type GetJobsIDStatusOK
const GetJobsIDStatusOKCode int = 200

/*
GetJobsIDStatusOK Job status

swagger:response getJobsIdStatusOK
*/
type GetJobsIDStatusOK struct {

	/*
	  In: Body
	*/
	Payload *model.JobStatus `json:"body,omitempty"`
}

// NewGetJobsIDStatusOK creates GetJobsIDStatusOK with default headers values
func NewGetJobsIDStatusOK() *GetJobsIDStatusOK {

	return &GetJobsIDStatusOK{}
}

// WithPayload adds the payload to the get jobs Id status o k response
func (o *GetJobsIDStatusOK) WithPayload(payload *model.JobStatus) *GetJobsIDStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get jobs Id status o k response
func (o *GetJobsIDStatusOK) SetPayload(payload *model.JobStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobsIDStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetJobsIDStatusNotFoundCode is the HTTP code returned for type GetJobsIDStatusNotFound
const GetJobsIDStatusNotFoundCode int = 404

/*
GetJobsIDStatusNotFound Not Found

swagger:response getJobsIdStatusNotFound
*/
type GetJobsIDStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *model.ErrorResponse `json:"body,omitempty"`
}

// NewGetJobsIDStatusNotFound creates GetJobsIDStatusNotFound with default headers values
func NewGetJobsIDStatusNotFound() *GetJobsIDStatusNotFound {

	return &GetJobsIDStatusNotFound{}
}

// WithPayload adds the payload to the get jobs Id status not found response
func (o *GetJobsIDStatusNotFound) WithPayload(payload *model.ErrorResponse) *GetJobsIDStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get jobs Id status not found response
func (o *GetJobsIDStatusNotFound) SetPayload(payload *model.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJobsIDStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetJobsIDStatusDefault Other error with any status code and response body format.

swagger:response getJobsIdStatusDefault
*/
type GetJobsIDStatusDefault struct {
	_statusCode int
}

// NewGetJobsIDStatusDefault creates GetJobsIDStatusDefault with default headers values
func NewGetJobsIDStatusDefault(code int) *GetJobsIDStatusDefault {
	if code <= 0 {
		code = 500
	}

	return &GetJobsIDStatusDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get jobs ID status default response
func (o *GetJobsIDStatusDefault) WithStatusCode(code int) *GetJobsIDStatusDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get jobs ID status default response
func (o *GetJobsIDStatusDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *GetJobsIDStatusDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
