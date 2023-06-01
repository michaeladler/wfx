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

// PutJobsIDDefinitionOKCode is the HTTP code returned for type PutJobsIDDefinitionOK
const PutJobsIDDefinitionOKCode int = 200

/*
PutJobsIDDefinitionOK Job modified successfully

swagger:response putJobsIdDefinitionOK
*/
type PutJobsIDDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload map[string]interface{} `json:"body,omitempty"`
}

// NewPutJobsIDDefinitionOK creates PutJobsIDDefinitionOK with default headers values
func NewPutJobsIDDefinitionOK() *PutJobsIDDefinitionOK {

	return &PutJobsIDDefinitionOK{}
}

// WithPayload adds the payload to the put jobs Id definition o k response
func (o *PutJobsIDDefinitionOK) WithPayload(payload map[string]interface{}) *PutJobsIDDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put jobs Id definition o k response
func (o *PutJobsIDDefinitionOK) SetPayload(payload map[string]interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutJobsIDDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = make(map[string]interface{}, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutJobsIDDefinitionBadRequestCode is the HTTP code returned for type PutJobsIDDefinitionBadRequest
const PutJobsIDDefinitionBadRequestCode int = 400

/*
PutJobsIDDefinitionBadRequest Bad Request

swagger:response putJobsIdDefinitionBadRequest
*/
type PutJobsIDDefinitionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *model.ErrorResponse `json:"body,omitempty"`
}

// NewPutJobsIDDefinitionBadRequest creates PutJobsIDDefinitionBadRequest with default headers values
func NewPutJobsIDDefinitionBadRequest() *PutJobsIDDefinitionBadRequest {

	return &PutJobsIDDefinitionBadRequest{}
}

// WithPayload adds the payload to the put jobs Id definition bad request response
func (o *PutJobsIDDefinitionBadRequest) WithPayload(payload *model.ErrorResponse) *PutJobsIDDefinitionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put jobs Id definition bad request response
func (o *PutJobsIDDefinitionBadRequest) SetPayload(payload *model.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutJobsIDDefinitionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutJobsIDDefinitionNotFoundCode is the HTTP code returned for type PutJobsIDDefinitionNotFound
const PutJobsIDDefinitionNotFoundCode int = 404

/*
PutJobsIDDefinitionNotFound Not Found

swagger:response putJobsIdDefinitionNotFound
*/
type PutJobsIDDefinitionNotFound struct {

	/*
	  In: Body
	*/
	Payload *model.ErrorResponse `json:"body,omitempty"`
}

// NewPutJobsIDDefinitionNotFound creates PutJobsIDDefinitionNotFound with default headers values
func NewPutJobsIDDefinitionNotFound() *PutJobsIDDefinitionNotFound {

	return &PutJobsIDDefinitionNotFound{}
}

// WithPayload adds the payload to the put jobs Id definition not found response
func (o *PutJobsIDDefinitionNotFound) WithPayload(payload *model.ErrorResponse) *PutJobsIDDefinitionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put jobs Id definition not found response
func (o *PutJobsIDDefinitionNotFound) SetPayload(payload *model.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutJobsIDDefinitionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PutJobsIDDefinitionDefault Other error with any status code and response body format.

swagger:response putJobsIdDefinitionDefault
*/
type PutJobsIDDefinitionDefault struct {
	_statusCode int
}

// NewPutJobsIDDefinitionDefault creates PutJobsIDDefinitionDefault with default headers values
func NewPutJobsIDDefinitionDefault(code int) *PutJobsIDDefinitionDefault {
	if code <= 0 {
		code = 500
	}

	return &PutJobsIDDefinitionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put jobs ID definition default response
func (o *PutJobsIDDefinitionDefault) WithStatusCode(code int) *PutJobsIDDefinitionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put jobs ID definition default response
func (o *PutJobsIDDefinitionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PutJobsIDDefinitionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
