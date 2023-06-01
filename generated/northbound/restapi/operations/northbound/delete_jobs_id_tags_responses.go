// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package northbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/siemens/wfx/generated/model"
)

// DeleteJobsIDTagsOKCode is the HTTP code returned for type DeleteJobsIDTagsOK
const DeleteJobsIDTagsOKCode int = 200

/*
DeleteJobsIDTagsOK Successfully deleted tag

swagger:response deleteJobsIdTagsOK
*/
type DeleteJobsIDTagsOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewDeleteJobsIDTagsOK creates DeleteJobsIDTagsOK with default headers values
func NewDeleteJobsIDTagsOK() *DeleteJobsIDTagsOK {

	return &DeleteJobsIDTagsOK{}
}

// WithPayload adds the payload to the delete jobs Id tags o k response
func (o *DeleteJobsIDTagsOK) WithPayload(payload []string) *DeleteJobsIDTagsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete jobs Id tags o k response
func (o *DeleteJobsIDTagsOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteJobsIDTagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteJobsIDTagsBadRequestCode is the HTTP code returned for type DeleteJobsIDTagsBadRequest
const DeleteJobsIDTagsBadRequestCode int = 400

/*
DeleteJobsIDTagsBadRequest Bad Request

swagger:response deleteJobsIdTagsBadRequest
*/
type DeleteJobsIDTagsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *model.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteJobsIDTagsBadRequest creates DeleteJobsIDTagsBadRequest with default headers values
func NewDeleteJobsIDTagsBadRequest() *DeleteJobsIDTagsBadRequest {

	return &DeleteJobsIDTagsBadRequest{}
}

// WithPayload adds the payload to the delete jobs Id tags bad request response
func (o *DeleteJobsIDTagsBadRequest) WithPayload(payload *model.ErrorResponse) *DeleteJobsIDTagsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete jobs Id tags bad request response
func (o *DeleteJobsIDTagsBadRequest) SetPayload(payload *model.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteJobsIDTagsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteJobsIDTagsNotFoundCode is the HTTP code returned for type DeleteJobsIDTagsNotFound
const DeleteJobsIDTagsNotFoundCode int = 404

/*
DeleteJobsIDTagsNotFound Not Found

swagger:response deleteJobsIdTagsNotFound
*/
type DeleteJobsIDTagsNotFound struct {

	/*
	  In: Body
	*/
	Payload *model.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteJobsIDTagsNotFound creates DeleteJobsIDTagsNotFound with default headers values
func NewDeleteJobsIDTagsNotFound() *DeleteJobsIDTagsNotFound {

	return &DeleteJobsIDTagsNotFound{}
}

// WithPayload adds the payload to the delete jobs Id tags not found response
func (o *DeleteJobsIDTagsNotFound) WithPayload(payload *model.ErrorResponse) *DeleteJobsIDTagsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete jobs Id tags not found response
func (o *DeleteJobsIDTagsNotFound) SetPayload(payload *model.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteJobsIDTagsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
DeleteJobsIDTagsDefault Other error with any status code and response body format.

swagger:response deleteJobsIdTagsDefault
*/
type DeleteJobsIDTagsDefault struct {
	_statusCode int
}

// NewDeleteJobsIDTagsDefault creates DeleteJobsIDTagsDefault with default headers values
func NewDeleteJobsIDTagsDefault(code int) *DeleteJobsIDTagsDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteJobsIDTagsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete jobs ID tags default response
func (o *DeleteJobsIDTagsDefault) WithStatusCode(code int) *DeleteJobsIDTagsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete jobs ID tags default response
func (o *DeleteJobsIDTagsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *DeleteJobsIDTagsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
