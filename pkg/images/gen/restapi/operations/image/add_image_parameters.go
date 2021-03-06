///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewAddImageParams creates a new AddImageParams object
// with the default values initialized.
func NewAddImageParams() AddImageParams {

	var (
		// initialize parameters with default values

		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)

	return AddImageParams{
		XDispatchOrg: &xDispatchOrgDefault,

		XDispatchProject: &xDispatchProjectDefault,
	}
}

// AddImageParams contains all the bound params for the add image operation
// typically these are obtained from a http.Request
//
// swagger:parameters addImage
type AddImageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Pattern: ^[\w\d][\w\d\-]*[\w\d]|[\w\d]+$
	  In: header
	  Default: "default"
	*/
	XDispatchOrg *string
	/*
	  Pattern: ^[\w\d][\w\d\-]*[\w\d]|[\w\d]+$
	  In: header
	  Default: "default"
	*/
	XDispatchProject *string
	/*Image object
	  Required: true
	  In: body
	*/
	Body *v1.Image
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddImageParams() beforehand.
func (o *AddImageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXDispatchOrg(r.Header[http.CanonicalHeaderKey("X-Dispatch-Org")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXDispatchProject(r.Header[http.CanonicalHeaderKey("X-Dispatch-Project")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body v1.Image
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXDispatchOrg binds and validates parameter XDispatchOrg from header.
func (o *AddImageParams) bindXDispatchOrg(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewAddImageParams()
		return nil
	}

	o.XDispatchOrg = &raw

	if err := o.validateXDispatchOrg(formats); err != nil {
		return err
	}

	return nil
}

// validateXDispatchOrg carries on validations for parameter XDispatchOrg
func (o *AddImageParams) validateXDispatchOrg(formats strfmt.Registry) error {

	if err := validate.Pattern("X-Dispatch-Org", "header", (*o.XDispatchOrg), `^[\w\d][\w\d\-]*[\w\d]|[\w\d]+$`); err != nil {
		return err
	}

	return nil
}

// bindXDispatchProject binds and validates parameter XDispatchProject from header.
func (o *AddImageParams) bindXDispatchProject(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewAddImageParams()
		return nil
	}

	o.XDispatchProject = &raw

	if err := o.validateXDispatchProject(formats); err != nil {
		return err
	}

	return nil
}

// validateXDispatchProject carries on validations for parameter XDispatchProject
func (o *AddImageParams) validateXDispatchProject(formats strfmt.Registry) error {

	if err := validate.Pattern("X-Dispatch-Project", "header", (*o.XDispatchProject), `^[\w\d][\w\d\-]*[\w\d]|[\w\d]+$`); err != nil {
		return err
	}

	return nil
}
