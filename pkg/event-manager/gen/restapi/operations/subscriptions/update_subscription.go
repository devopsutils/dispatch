///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateSubscriptionHandlerFunc turns a function with the right signature into a update subscription handler
type UpdateSubscriptionHandlerFunc func(UpdateSubscriptionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateSubscriptionHandlerFunc) Handle(params UpdateSubscriptionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateSubscriptionHandler interface for that can handle valid update subscription params
type UpdateSubscriptionHandler interface {
	Handle(UpdateSubscriptionParams, interface{}) middleware.Responder
}

// NewUpdateSubscription creates a new http.Handler for the update subscription operation
func NewUpdateSubscription(ctx *middleware.Context, handler UpdateSubscriptionHandler) *UpdateSubscription {
	return &UpdateSubscription{Context: ctx, Handler: handler}
}

/*UpdateSubscription swagger:route PUT /subscriptions/{subscriptionName} subscriptions updateSubscription

Update subscription by Name

Updates a single subscription

*/
type UpdateSubscription struct {
	Context *middleware.Context
	Handler UpdateSubscriptionHandler
}

func (o *UpdateSubscription) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateSubscriptionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}