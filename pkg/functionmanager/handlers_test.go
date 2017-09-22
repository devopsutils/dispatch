///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package functionmanager

import (
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"

	entitystore "gitlab.eng.vmware.com/serverless/serverless/pkg/entity-store"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/models"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/restapi/operations"
	fnstore "gitlab.eng.vmware.com/serverless/serverless/pkg/functionmanager/gen/restapi/operations/store"
	helpers "gitlab.eng.vmware.com/serverless/serverless/pkg/testing/api"
)

type testEntity struct {
	entitystore.BaseEntity
	Value string `json:"value"`
}

func TestStoreAddFunctionHandler(t *testing.T) {
	api := operations.NewFunctionManagerAPI(nil)
	helpers.MakeAPI(t, ConfigureHandlers, api)

	var tags []*models.Tag
	tags = append(tags, &models.Tag{Key: "role", Value: "test"})
	schema := models.Schema{In: "testInSchema", Out: "testOutSchema"}
	reqBody := &models.Function{
		Name:     swag.String("testEntity"),
		Active:   true,
		Schema:   &schema,
		Language: "python3",
		Code:     swag.String("some code"),
		Image:    swag.String("imageID"),
		Tags:     tags,
	}
	r := httptest.NewRequest("POST", "/v1/function", nil)
	params := fnstore.AddFunctionParams{
		HTTPRequest: r,
		Body:        reqBody,
	}
	responder := api.StoreAddFunctionHandler.Handle(params)
	var respBody models.Function
	helpers.HandlerRequest(t, responder, &respBody, 202)

	assert.NotNil(t, respBody.CreatedTime)
	assert.NotEmpty(t, respBody.ID)
	assert.Equal(t, reqBody.Name, respBody.Name)
	assert.Equal(t, reqBody.Active, respBody.Active)
	assert.Equal(t, reqBody.Language, respBody.Language)
	assert.Equal(t, reqBody.Schema, respBody.Schema)
	assert.Equal(t, reqBody.Code, respBody.Code)
	assert.Equal(t, reqBody.Image, respBody.Image)
	assert.Len(t, respBody.Tags, 1)
	assert.Equal(t, "role", respBody.Tags[0].Key)
	assert.Equal(t, "test", respBody.Tags[0].Value)
}

func TestStoreGetFunctionByNameHandler(t *testing.T) {
	api := operations.NewFunctionManagerAPI(nil)
	helpers.MakeAPI(t, ConfigureHandlers, api)

	var tags []*models.Tag
	tags = append(tags, &models.Tag{Key: "role", Value: "test"})
	schema := models.Schema{In: "testInSchema", Out: "testOutSchema"}
	reqBody := &models.Function{
		Name:     swag.String("testEntity"),
		Active:   true,
		Schema:   &schema,
		Language: "python3",
		Code:     swag.String("some code"),
		Image:    swag.String("imageID"),
		Tags:     tags,
	}
	r := httptest.NewRequest("POST", "/v1/function", nil)
	add := fnstore.AddFunctionParams{
		HTTPRequest: r,
		Body:        reqBody,
	}
	addResponder := api.StoreAddFunctionHandler.Handle(add)
	var addBody models.Function
	helpers.HandlerRequest(t, addResponder, &addBody, 202)

	assert.NotEmpty(t, addBody.ID)

	id := addBody.ID
	createdTime := addBody.CreatedTime
	r = httptest.NewRequest("GET", "/v1/function/testEntity", nil)
	get := fnstore.GetFunctionByNameParams{
		HTTPRequest:  r,
		FunctionName: "testEntity",
	}
	getResponder := api.StoreGetFunctionByNameHandler.Handle(get)
	var getBody models.Function
	helpers.HandlerRequest(t, getResponder, &getBody, 200)

	assert.Equal(t, id, getBody.ID)
	assert.Equal(t, createdTime, getBody.CreatedTime)
	assert.Equal(t, reqBody.Name, getBody.Name)
	assert.Equal(t, reqBody.Active, getBody.Active)
	assert.Equal(t, reqBody.Language, getBody.Language)
	assert.Equal(t, reqBody.Schema, getBody.Schema)
	assert.Equal(t, reqBody.Code, getBody.Code)
	assert.Equal(t, reqBody.Schema, getBody.Schema)
	assert.Len(t, getBody.Tags, 1)
	assert.Equal(t, "role", getBody.Tags[0].Key)
	assert.Equal(t, "test", getBody.Tags[0].Value)
}