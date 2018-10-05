///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteOrganizationParams creates a new DeleteOrganizationParams object
// no default values defined in spec.
func NewDeleteOrganizationParams() DeleteOrganizationParams {

	return DeleteOrganizationParams{}
}

// DeleteOrganizationParams contains all the bound params for the delete organization operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteOrganization
type DeleteOrganizationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	XDispatchOrg string
	/*Name of Organization to work on
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	OrganizationName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteOrganizationParams() beforehand.
func (o *DeleteOrganizationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXDispatchOrg(r.Header[http.CanonicalHeaderKey("X-Dispatch-Org")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rOrganizationName, rhkOrganizationName, _ := route.Params.GetOK("organizationName")
	if err := o.bindOrganizationName(rOrganizationName, rhkOrganizationName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXDispatchOrg binds and validates parameter XDispatchOrg from header.
func (o *DeleteOrganizationParams) bindXDispatchOrg(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindOrganizationName binds and validates parameter OrganizationName from path.
func (o *DeleteOrganizationParams) bindOrganizationName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.OrganizationName = raw

	if err := o.validateOrganizationName(formats); err != nil {
		return err
	}

	return nil
}

// validateOrganizationName carries on validations for parameter OrganizationName
func (o *DeleteOrganizationParams) validateOrganizationName(formats strfmt.Registry) error {

	if err := validate.Pattern("organizationName", "path", o.OrganizationName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}
