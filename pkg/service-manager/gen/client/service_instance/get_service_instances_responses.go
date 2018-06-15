///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package service_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetServiceInstancesReader is a Reader for the GetServiceInstances structure.
type GetServiceInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetServiceInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetServiceInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetServiceInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetServiceInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceInstancesOK creates a GetServiceInstancesOK with default headers values
func NewGetServiceInstancesOK() *GetServiceInstancesOK {
	return &GetServiceInstancesOK{}
}

/*GetServiceInstancesOK handles this case with default header values.

successful operation
*/
type GetServiceInstancesOK struct {
	Payload []*v1.ServiceInstance
}

func (o *GetServiceInstancesOK) Error() string {
	return fmt.Sprintf("[GET /serviceinstance][%d] getServiceInstancesOK  %+v", 200, o.Payload)
}

func (o *GetServiceInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceInstancesBadRequest creates a GetServiceInstancesBadRequest with default headers values
func NewGetServiceInstancesBadRequest() *GetServiceInstancesBadRequest {
	return &GetServiceInstancesBadRequest{}
}

/*GetServiceInstancesBadRequest handles this case with default header values.

Invalid input
*/
type GetServiceInstancesBadRequest struct {
	Payload *v1.Error
}

func (o *GetServiceInstancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /serviceinstance][%d] getServiceInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *GetServiceInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceInstancesUnauthorized creates a GetServiceInstancesUnauthorized with default headers values
func NewGetServiceInstancesUnauthorized() *GetServiceInstancesUnauthorized {
	return &GetServiceInstancesUnauthorized{}
}

/*GetServiceInstancesUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetServiceInstancesUnauthorized struct {
	Payload *v1.Error
}

func (o *GetServiceInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /serviceinstance][%d] getServiceInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetServiceInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceInstancesForbidden creates a GetServiceInstancesForbidden with default headers values
func NewGetServiceInstancesForbidden() *GetServiceInstancesForbidden {
	return &GetServiceInstancesForbidden{}
}

/*GetServiceInstancesForbidden handles this case with default header values.

access to this resource is forbidden
*/
type GetServiceInstancesForbidden struct {
	Payload *v1.Error
}

func (o *GetServiceInstancesForbidden) Error() string {
	return fmt.Sprintf("[GET /serviceinstance][%d] getServiceInstancesForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceInstancesDefault creates a GetServiceInstancesDefault with default headers values
func NewGetServiceInstancesDefault(code int) *GetServiceInstancesDefault {
	return &GetServiceInstancesDefault{
		_statusCode: code,
	}
}

/*GetServiceInstancesDefault handles this case with default header values.

Generic error response
*/
type GetServiceInstancesDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the get service instances default response
func (o *GetServiceInstancesDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /serviceinstance][%d] getServiceInstances default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
