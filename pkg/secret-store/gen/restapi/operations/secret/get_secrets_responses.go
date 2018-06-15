///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetSecretsOKCode is the HTTP code returned for type GetSecretsOK
const GetSecretsOKCode int = 200

/*GetSecretsOK An array of registered secrets

swagger:response getSecretsOK
*/
type GetSecretsOK struct {

	/*
	  In: Body
	*/
	Payload []*v1.Secret `json:"body,omitempty"`
}

// NewGetSecretsOK creates GetSecretsOK with default headers values
func NewGetSecretsOK() *GetSecretsOK {

	return &GetSecretsOK{}
}

// WithPayload adds the payload to the get secrets o k response
func (o *GetSecretsOK) WithPayload(payload []*v1.Secret) *GetSecretsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secrets o k response
func (o *GetSecretsOK) SetPayload(payload []*v1.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*v1.Secret, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetSecretsBadRequestCode is the HTTP code returned for type GetSecretsBadRequest
const GetSecretsBadRequestCode int = 400

/*GetSecretsBadRequest Bad Request

swagger:response getSecretsBadRequest
*/
type GetSecretsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetSecretsBadRequest creates GetSecretsBadRequest with default headers values
func NewGetSecretsBadRequest() *GetSecretsBadRequest {

	return &GetSecretsBadRequest{}
}

// WithPayload adds the payload to the get secrets bad request response
func (o *GetSecretsBadRequest) WithPayload(payload *v1.Error) *GetSecretsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secrets bad request response
func (o *GetSecretsBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSecretsUnauthorizedCode is the HTTP code returned for type GetSecretsUnauthorized
const GetSecretsUnauthorizedCode int = 401

/*GetSecretsUnauthorized Unauthorized Request

swagger:response getSecretsUnauthorized
*/
type GetSecretsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetSecretsUnauthorized creates GetSecretsUnauthorized with default headers values
func NewGetSecretsUnauthorized() *GetSecretsUnauthorized {

	return &GetSecretsUnauthorized{}
}

// WithPayload adds the payload to the get secrets unauthorized response
func (o *GetSecretsUnauthorized) WithPayload(payload *v1.Error) *GetSecretsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secrets unauthorized response
func (o *GetSecretsUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSecretsForbiddenCode is the HTTP code returned for type GetSecretsForbidden
const GetSecretsForbiddenCode int = 403

/*GetSecretsForbidden access to this resource is forbidden

swagger:response getSecretsForbidden
*/
type GetSecretsForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetSecretsForbidden creates GetSecretsForbidden with default headers values
func NewGetSecretsForbidden() *GetSecretsForbidden {

	return &GetSecretsForbidden{}
}

// WithPayload adds the payload to the get secrets forbidden response
func (o *GetSecretsForbidden) WithPayload(payload *v1.Error) *GetSecretsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secrets forbidden response
func (o *GetSecretsForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSecretsDefault Standard error

swagger:response getSecretsDefault
*/
type GetSecretsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetSecretsDefault creates GetSecretsDefault with default headers values
func NewGetSecretsDefault(code int) *GetSecretsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSecretsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get secrets default response
func (o *GetSecretsDefault) WithStatusCode(code int) *GetSecretsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get secrets default response
func (o *GetSecretsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get secrets default response
func (o *GetSecretsDefault) WithPayload(payload *v1.Error) *GetSecretsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get secrets default response
func (o *GetSecretsDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSecretsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
