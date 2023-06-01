// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package southbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetJobsIDTagsHandlerFunc turns a function with the right signature into a get jobs ID tags handler
type GetJobsIDTagsHandlerFunc func(GetJobsIDTagsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetJobsIDTagsHandlerFunc) Handle(params GetJobsIDTagsParams) middleware.Responder {
	return fn(params)
}

// GetJobsIDTagsHandler interface for that can handle valid get jobs ID tags params
type GetJobsIDTagsHandler interface {
	Handle(GetJobsIDTagsParams) middleware.Responder
}

// NewGetJobsIDTags creates a new http.Handler for the get jobs ID tags operation
func NewGetJobsIDTags(ctx *middleware.Context, handler GetJobsIDTagsHandler) *GetJobsIDTags {
	return &GetJobsIDTags{Context: ctx, Handler: handler}
}

/*
	GetJobsIDTags swagger:route GET /jobs/{id}/tags southbound getJobsIdTags

# Get tags

Get the tags of a job
*/
type GetJobsIDTags struct {
	Context *middleware.Context
	Handler GetJobsIDTagsHandler
}

func (o *GetJobsIDTags) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetJobsIDTagsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
