///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	baseimagev1 "github.com/vmware/dispatch/pkg/resources/baseimage/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBaseImages implements BaseImageInterface
type FakeBaseImages struct {
	Fake *FakeResourcesV1
	ns   string
}

var baseimagesResource = schema.GroupVersionResource{Group: "resources.dispatchframework.io", Version: "v1", Resource: "baseimages"}

var baseimagesKind = schema.GroupVersionKind{Group: "resources.dispatchframework.io", Version: "v1", Kind: "BaseImage"}

// Get takes name of the baseImage, and returns the corresponding baseImage object, and an error if there is any.
func (c *FakeBaseImages) Get(name string, options v1.GetOptions) (result *baseimagev1.BaseImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(baseimagesResource, c.ns, name), &baseimagev1.BaseImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*baseimagev1.BaseImage), err
}

// List takes label and field selectors, and returns the list of BaseImages that match those selectors.
func (c *FakeBaseImages) List(opts v1.ListOptions) (result *baseimagev1.BaseImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(baseimagesResource, baseimagesKind, c.ns, opts), &baseimagev1.BaseImageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &baseimagev1.BaseImageList{ListMeta: obj.(*baseimagev1.BaseImageList).ListMeta}
	for _, item := range obj.(*baseimagev1.BaseImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested baseImages.
func (c *FakeBaseImages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(baseimagesResource, c.ns, opts))

}

// Create takes the representation of a baseImage and creates it.  Returns the server's representation of the baseImage, and an error, if there is any.
func (c *FakeBaseImages) Create(baseImage *baseimagev1.BaseImage) (result *baseimagev1.BaseImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(baseimagesResource, c.ns, baseImage), &baseimagev1.BaseImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*baseimagev1.BaseImage), err
}

// Update takes the representation of a baseImage and updates it. Returns the server's representation of the baseImage, and an error, if there is any.
func (c *FakeBaseImages) Update(baseImage *baseimagev1.BaseImage) (result *baseimagev1.BaseImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(baseimagesResource, c.ns, baseImage), &baseimagev1.BaseImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*baseimagev1.BaseImage), err
}

// Delete takes name of the baseImage and deletes it. Returns an error if one occurs.
func (c *FakeBaseImages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(baseimagesResource, c.ns, name), &baseimagev1.BaseImage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBaseImages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(baseimagesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &baseimagev1.BaseImageList{})
	return err
}

// Patch applies the patch and returns the patched baseImage.
func (c *FakeBaseImages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *baseimagev1.BaseImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(baseimagesResource, c.ns, name, data, subresources...), &baseimagev1.BaseImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*baseimagev1.BaseImage), err
}
