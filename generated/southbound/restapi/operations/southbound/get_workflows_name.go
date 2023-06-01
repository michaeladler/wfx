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

// GetWorkflowsNameHandlerFunc turns a function with the right signature into a get workflows name handler
type GetWorkflowsNameHandlerFunc func(GetWorkflowsNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWorkflowsNameHandlerFunc) Handle(params GetWorkflowsNameParams) middleware.Responder {
	return fn(params)
}

// GetWorkflowsNameHandler interface for that can handle valid get workflows name params
type GetWorkflowsNameHandler interface {
	Handle(GetWorkflowsNameParams) middleware.Responder
}

// NewGetWorkflowsName creates a new http.Handler for the get workflows name operation
func NewGetWorkflowsName(ctx *middleware.Context, handler GetWorkflowsNameHandler) *GetWorkflowsName {
	return &GetWorkflowsName{Context: ctx, Handler: handler}
}

/*
	GetWorkflowsName swagger:route GET /workflows/{name} southbound getWorkflowsName

# Workflow description for a given name

Workflow description for a given name
*/
type GetWorkflowsName struct {
	Context *middleware.Context
	Handler GetWorkflowsNameHandler
}

func (o *GetWorkflowsName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetWorkflowsNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
