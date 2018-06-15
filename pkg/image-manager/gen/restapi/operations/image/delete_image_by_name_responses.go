///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// DeleteImageByNameOKCode is the HTTP code returned for type DeleteImageByNameOK
const DeleteImageByNameOKCode int = 200

/*DeleteImageByNameOK successful operation

swagger:response deleteImageByNameOK
*/
type DeleteImageByNameOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Image `json:"body,omitempty"`
}

// NewDeleteImageByNameOK creates DeleteImageByNameOK with default headers values
func NewDeleteImageByNameOK() *DeleteImageByNameOK {

	return &DeleteImageByNameOK{}
}

// WithPayload adds the payload to the delete image by name o k response
func (o *DeleteImageByNameOK) WithPayload(payload *v1.Image) *DeleteImageByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete image by name o k response
func (o *DeleteImageByNameOK) SetPayload(payload *v1.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteImageByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteImageByNameBadRequestCode is the HTTP code returned for type DeleteImageByNameBadRequest
const DeleteImageByNameBadRequestCode int = 400

/*DeleteImageByNameBadRequest Invalid ID supplied

swagger:response deleteImageByNameBadRequest
*/
type DeleteImageByNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteImageByNameBadRequest creates DeleteImageByNameBadRequest with default headers values
func NewDeleteImageByNameBadRequest() *DeleteImageByNameBadRequest {

	return &DeleteImageByNameBadRequest{}
}

// WithPayload adds the payload to the delete image by name bad request response
func (o *DeleteImageByNameBadRequest) WithPayload(payload *v1.Error) *DeleteImageByNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete image by name bad request response
func (o *DeleteImageByNameBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteImageByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteImageByNameUnauthorizedCode is the HTTP code returned for type DeleteImageByNameUnauthorized
const DeleteImageByNameUnauthorizedCode int = 401

/*DeleteImageByNameUnauthorized Unauthorized Request

swagger:response deleteImageByNameUnauthorized
*/
type DeleteImageByNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteImageByNameUnauthorized creates DeleteImageByNameUnauthorized with default headers values
func NewDeleteImageByNameUnauthorized() *DeleteImageByNameUnauthorized {

	return &DeleteImageByNameUnauthorized{}
}

// WithPayload adds the payload to the delete image by name unauthorized response
func (o *DeleteImageByNameUnauthorized) WithPayload(payload *v1.Error) *DeleteImageByNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete image by name unauthorized response
func (o *DeleteImageByNameUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteImageByNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteImageByNameForbiddenCode is the HTTP code returned for type DeleteImageByNameForbidden
const DeleteImageByNameForbiddenCode int = 403

/*DeleteImageByNameForbidden access to this resource is forbidden

swagger:response deleteImageByNameForbidden
*/
type DeleteImageByNameForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteImageByNameForbidden creates DeleteImageByNameForbidden with default headers values
func NewDeleteImageByNameForbidden() *DeleteImageByNameForbidden {

	return &DeleteImageByNameForbidden{}
}

// WithPayload adds the payload to the delete image by name forbidden response
func (o *DeleteImageByNameForbidden) WithPayload(payload *v1.Error) *DeleteImageByNameForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete image by name forbidden response
func (o *DeleteImageByNameForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteImageByNameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteImageByNameNotFoundCode is the HTTP code returned for type DeleteImageByNameNotFound
const DeleteImageByNameNotFoundCode int = 404

/*DeleteImageByNameNotFound Image not found

swagger:response deleteImageByNameNotFound
*/
type DeleteImageByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteImageByNameNotFound creates DeleteImageByNameNotFound with default headers values
func NewDeleteImageByNameNotFound() *DeleteImageByNameNotFound {

	return &DeleteImageByNameNotFound{}
}

// WithPayload adds the payload to the delete image by name not found response
func (o *DeleteImageByNameNotFound) WithPayload(payload *v1.Error) *DeleteImageByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete image by name not found response
func (o *DeleteImageByNameNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteImageByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteImageByNameDefault Generic error response

swagger:response deleteImageByNameDefault
*/
type DeleteImageByNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteImageByNameDefault creates DeleteImageByNameDefault with default headers values
func NewDeleteImageByNameDefault(code int) *DeleteImageByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteImageByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete image by name default response
func (o *DeleteImageByNameDefault) WithStatusCode(code int) *DeleteImageByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete image by name default response
func (o *DeleteImageByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete image by name default response
func (o *DeleteImageByNameDefault) WithPayload(payload *v1.Error) *DeleteImageByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete image by name default response
func (o *DeleteImageByNameDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteImageByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
