///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// DeleteFunctionReader is a Reader for the DeleteFunction structure.
type DeleteFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteFunctionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteFunctionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteFunctionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFunctionOK creates a DeleteFunctionOK with default headers values
func NewDeleteFunctionOK() *DeleteFunctionOK {
	return &DeleteFunctionOK{}
}

/*DeleteFunctionOK handles this case with default header values.

Successful operation
*/
type DeleteFunctionOK struct {
	Payload *v1.Function
}

func (o *DeleteFunctionOK) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunctionOK  %+v", 200, o.Payload)
}

func (o *DeleteFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Function)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFunctionBadRequest creates a DeleteFunctionBadRequest with default headers values
func NewDeleteFunctionBadRequest() *DeleteFunctionBadRequest {
	return &DeleteFunctionBadRequest{}
}

/*DeleteFunctionBadRequest handles this case with default header values.

Invalid Name supplied
*/
type DeleteFunctionBadRequest struct {
	Payload *v1.Error
}

func (o *DeleteFunctionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFunctionUnauthorized creates a DeleteFunctionUnauthorized with default headers values
func NewDeleteFunctionUnauthorized() *DeleteFunctionUnauthorized {
	return &DeleteFunctionUnauthorized{}
}

/*DeleteFunctionUnauthorized handles this case with default header values.

Unauthorized Request
*/
type DeleteFunctionUnauthorized struct {
	Payload *v1.Error
}

func (o *DeleteFunctionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunctionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFunctionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFunctionForbidden creates a DeleteFunctionForbidden with default headers values
func NewDeleteFunctionForbidden() *DeleteFunctionForbidden {
	return &DeleteFunctionForbidden{}
}

/*DeleteFunctionForbidden handles this case with default header values.

access to this resource is forbidden
*/
type DeleteFunctionForbidden struct {
	Payload *v1.Error
}

func (o *DeleteFunctionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunctionForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFunctionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFunctionNotFound creates a DeleteFunctionNotFound with default headers values
func NewDeleteFunctionNotFound() *DeleteFunctionNotFound {
	return &DeleteFunctionNotFound{}
}

/*DeleteFunctionNotFound handles this case with default header values.

Function not found
*/
type DeleteFunctionNotFound struct {
	Payload *v1.Error
}

func (o *DeleteFunctionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunctionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFunctionDefault creates a DeleteFunctionDefault with default headers values
func NewDeleteFunctionDefault(code int) *DeleteFunctionDefault {
	return &DeleteFunctionDefault{
		_statusCode: code,
	}
}

/*DeleteFunctionDefault handles this case with default header values.

Unknown error
*/
type DeleteFunctionDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the delete function default response
func (o *DeleteFunctionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFunctionDefault) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunction default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFunctionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
