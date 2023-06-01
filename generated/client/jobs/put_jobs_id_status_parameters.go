// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/siemens/wfx/generated/model"
)

// NewPutJobsIDStatusParams creates a new PutJobsIDStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutJobsIDStatusParams() *PutJobsIDStatusParams {
	return &PutJobsIDStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutJobsIDStatusParamsWithTimeout creates a new PutJobsIDStatusParams object
// with the ability to set a timeout on a request.
func NewPutJobsIDStatusParamsWithTimeout(timeout time.Duration) *PutJobsIDStatusParams {
	return &PutJobsIDStatusParams{
		timeout: timeout,
	}
}

// NewPutJobsIDStatusParamsWithContext creates a new PutJobsIDStatusParams object
// with the ability to set a context for a request.
func NewPutJobsIDStatusParamsWithContext(ctx context.Context) *PutJobsIDStatusParams {
	return &PutJobsIDStatusParams{
		Context: ctx,
	}
}

// NewPutJobsIDStatusParamsWithHTTPClient creates a new PutJobsIDStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutJobsIDStatusParamsWithHTTPClient(client *http.Client) *PutJobsIDStatusParams {
	return &PutJobsIDStatusParams{
		HTTPClient: client,
	}
}

/*
PutJobsIDStatusParams contains all the parameters to send to the API endpoint

	for the put jobs ID status operation.

	Typically these are written to a http.Request.
*/
type PutJobsIDStatusParams struct {

	/* NewJobStatus.

	   This contains the new job status
	*/
	NewJobStatus *model.JobStatus

	/* ID.

	   Job ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put jobs ID status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutJobsIDStatusParams) WithDefaults() *PutJobsIDStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put jobs ID status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutJobsIDStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put jobs ID status params
func (o *PutJobsIDStatusParams) WithTimeout(timeout time.Duration) *PutJobsIDStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put jobs ID status params
func (o *PutJobsIDStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put jobs ID status params
func (o *PutJobsIDStatusParams) WithContext(ctx context.Context) *PutJobsIDStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put jobs ID status params
func (o *PutJobsIDStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put jobs ID status params
func (o *PutJobsIDStatusParams) WithHTTPClient(client *http.Client) *PutJobsIDStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put jobs ID status params
func (o *PutJobsIDStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewJobStatus adds the newJobStatus to the put jobs ID status params
func (o *PutJobsIDStatusParams) WithNewJobStatus(newJobStatus *model.JobStatus) *PutJobsIDStatusParams {
	o.SetNewJobStatus(newJobStatus)
	return o
}

// SetNewJobStatus adds the newJobStatus to the put jobs ID status params
func (o *PutJobsIDStatusParams) SetNewJobStatus(newJobStatus *model.JobStatus) {
	o.NewJobStatus = newJobStatus
}

// WithID adds the id to the put jobs ID status params
func (o *PutJobsIDStatusParams) WithID(id string) *PutJobsIDStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put jobs ID status params
func (o *PutJobsIDStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutJobsIDStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NewJobStatus != nil {
		if err := r.SetBodyParam(o.NewJobStatus); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
