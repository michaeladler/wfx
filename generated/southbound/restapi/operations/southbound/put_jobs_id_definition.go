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

// PutJobsIDDefinitionHandlerFunc turns a function with the right signature into a put jobs ID definition handler
type PutJobsIDDefinitionHandlerFunc func(PutJobsIDDefinitionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutJobsIDDefinitionHandlerFunc) Handle(params PutJobsIDDefinitionParams) middleware.Responder {
	return fn(params)
}

// PutJobsIDDefinitionHandler interface for that can handle valid put jobs ID definition params
type PutJobsIDDefinitionHandler interface {
	Handle(PutJobsIDDefinitionParams) middleware.Responder
}

// NewPutJobsIDDefinition creates a new http.Handler for the put jobs ID definition operation
func NewPutJobsIDDefinition(ctx *middleware.Context, handler PutJobsIDDefinitionHandler) *PutJobsIDDefinition {
	return &PutJobsIDDefinition{Context: ctx, Handler: handler}
}

/*
	PutJobsIDDefinition swagger:route PUT /jobs/{id}/definition southbound putJobsIdDefinition

# Modify job definition

Modify the job definition of an existing job
*/
type PutJobsIDDefinition struct {
	Context *middleware.Context
	Handler PutJobsIDDefinitionHandler
}

func (o *PutJobsIDDefinition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutJobsIDDefinitionParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
