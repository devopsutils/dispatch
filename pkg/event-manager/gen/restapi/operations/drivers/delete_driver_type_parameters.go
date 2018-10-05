///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteDriverTypeParams creates a new DeleteDriverTypeParams object
// no default values defined in spec.
func NewDeleteDriverTypeParams() DeleteDriverTypeParams {

	return DeleteDriverTypeParams{}
}

// DeleteDriverTypeParams contains all the bound params for the delete driver type operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteDriverType
type DeleteDriverTypeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	XDispatchOrg string
	/*Name of the driver type to work on
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	DriverTypeName string
	/*Filter based on tags
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteDriverTypeParams() beforehand.
func (o *DeleteDriverTypeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXDispatchOrg(r.Header[http.CanonicalHeaderKey("X-Dispatch-Org")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rDriverTypeName, rhkDriverTypeName, _ := route.Params.GetOK("driverTypeName")
	if err := o.bindDriverTypeName(rDriverTypeName, rhkDriverTypeName, route.Formats); err != nil {
		res = append(res, err)
	}

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXDispatchOrg binds and validates parameter XDispatchOrg from header.
func (o *DeleteDriverTypeParams) bindXDispatchOrg(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Dispatch-Org", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Dispatch-Org", "header", raw); err != nil {
		return err
	}

	o.XDispatchOrg = raw

	return nil
}

// bindDriverTypeName binds and validates parameter DriverTypeName from path.
func (o *DeleteDriverTypeParams) bindDriverTypeName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.DriverTypeName = raw

	if err := o.validateDriverTypeName(formats); err != nil {
		return err
	}

	return nil
}

// validateDriverTypeName carries on validations for parameter DriverTypeName
func (o *DeleteDriverTypeParams) validateDriverTypeName(formats strfmt.Registry) error {

	if err := validate.Pattern("driverTypeName", "path", o.DriverTypeName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

// bindTags binds and validates array parameter Tags from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *DeleteDriverTypeParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	tagsIC := rawData

	if len(tagsIC) == 0 {
		return nil
	}

	var tagsIR []string
	for _, tagsIV := range tagsIC {
		tagsI := tagsIV

		tagsIR = append(tagsIR, tagsI)
	}

	o.Tags = tagsIR

	return nil
}
