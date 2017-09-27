///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package openwhisk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functions"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/testing/dev"
)

var driver functions.FaaSDriver

func init() {
	d, err := New(&Config{
		Insecure:  true,
		Host:      "172.17.0.1",
		AuthToken: "23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP",
	})
	if err != nil {
		panic(err)
	}
	driver = d
}

func TestWskDriver_GetRunnable(t *testing.T) {
	dev.EnsureLocal(t)
	f := driver.GetRunnable("hello")
	r, err := f(map[string]interface{}{})

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"myField": "Hello, Noone from Nowhere"}, r)
}

func TestWskDriver_Delete(t *testing.T) {
	dev.EnsureLocal(t)
	err := driver.Delete("hello")
	assert.NoError(t, err)
}

func TestWskDriver_Create(t *testing.T) {
	dev.EnsureLocal(t)
	err := driver.Create("hello", &functions.Exec{
		Image: "imikushin/nodejs6action",
		Code: `
function main(params) {
    var name = "Noone";
    if (params.name) {
        name = params.name;
    }
    var place = "Nowhere";
    if (params.place) {
        place = params.place;
    }
    return {myField:  'Hello, ' + name + ' from ' + place};
}
		`,
	})
	assert.NoError(t, err)
}