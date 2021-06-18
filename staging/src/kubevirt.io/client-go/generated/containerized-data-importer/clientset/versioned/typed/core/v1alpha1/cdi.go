/*
Copyright 2021 The KubeVirt Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	scheme "kubevirt.io/client-go/generated/containerized-data-importer/clientset/versioned/scheme"
	v1alpha1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1alpha1"
)

// CDIsGetter has a method to return a CDIInterface.
// A group's client should implement this interface.
type CDIsGetter interface {
	CDIs() CDIInterface
}

// CDIInterface has methods to work with CDI resources.
type CDIInterface interface {
	Create(ctx context.Context, cDI *v1alpha1.CDI, opts v1.CreateOptions) (*v1alpha1.CDI, error)
	Update(ctx context.Context, cDI *v1alpha1.CDI, opts v1.UpdateOptions) (*v1alpha1.CDI, error)
	UpdateStatus(ctx context.Context, cDI *v1alpha1.CDI, opts v1.UpdateOptions) (*v1alpha1.CDI, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CDI, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CDIList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CDI, err error)
	CDIExpansion
}

// cDIs implements CDIInterface
type cDIs struct {
	client rest.Interface
}

// newCDIs returns a CDIs
func newCDIs(c *CdiV1alpha1Client) *cDIs {
	return &cDIs{
		client: c.RESTClient(),
	}
}

// Get takes name of the cDI, and returns the corresponding cDI object, and an error if there is any.
func (c *cDIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CDI, err error) {
	result = &v1alpha1.CDI{}
	err = c.client.Get().
		Resource("cdis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CDIs that match those selectors.
func (c *cDIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CDIList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CDIList{}
	err = c.client.Get().
		Resource("cdis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cDIs.
func (c *cDIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cdis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cDI and creates it.  Returns the server's representation of the cDI, and an error, if there is any.
func (c *cDIs) Create(ctx context.Context, cDI *v1alpha1.CDI, opts v1.CreateOptions) (result *v1alpha1.CDI, err error) {
	result = &v1alpha1.CDI{}
	err = c.client.Post().
		Resource("cdis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cDI).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cDI and updates it. Returns the server's representation of the cDI, and an error, if there is any.
func (c *cDIs) Update(ctx context.Context, cDI *v1alpha1.CDI, opts v1.UpdateOptions) (result *v1alpha1.CDI, err error) {
	result = &v1alpha1.CDI{}
	err = c.client.Put().
		Resource("cdis").
		Name(cDI.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cDI).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cDIs) UpdateStatus(ctx context.Context, cDI *v1alpha1.CDI, opts v1.UpdateOptions) (result *v1alpha1.CDI, err error) {
	result = &v1alpha1.CDI{}
	err = c.client.Put().
		Resource("cdis").
		Name(cDI.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cDI).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cDI and deletes it. Returns an error if one occurs.
func (c *cDIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cdis").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cDIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cdis").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cDI.
func (c *cDIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CDI, err error) {
	result = &v1alpha1.CDI{}
	err = c.client.Patch(pt).
		Resource("cdis").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
