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

package v1

import (
	v1 "github.com/anuruddhal/k8s-controller-core-resource/pkg/apis/myresource/v1"
	scheme "github.com/anuruddhal/k8s-controller-core-resource/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MyResourcesGetter has a method to return a MyResourceInterface.
// A group's client should implement this interface.
type MyResourcesGetter interface {
	MyResources(namespace string) MyResourceInterface
}

// MyResourceInterface has methods to work with MyResource resources.
type MyResourceInterface interface {
	Create(*v1.MyResource) (*v1.MyResource, error)
	Update(*v1.MyResource) (*v1.MyResource, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MyResource, error)
	List(opts metav1.ListOptions) (*v1.MyResourceList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MyResource, err error)
	MyResourceExpansion
}

// myResources implements MyResourceInterface
type myResources struct {
	client rest.Interface
	ns     string
}

// newMyResources returns a MyResources
func newMyResources(c *AnuruddhalV1Client, namespace string) *myResources {
	return &myResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the myResource, and returns the corresponding myResource object, and an error if there is any.
func (c *myResources) Get(name string, options metav1.GetOptions) (result *v1.MyResource, err error) {
	result = &v1.MyResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("myresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MyResources that match those selectors.
func (c *myResources) List(opts metav1.ListOptions) (result *v1.MyResourceList, err error) {
	result = &v1.MyResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("myresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested myResources.
func (c *myResources) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("myresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a myResource and creates it.  Returns the server's representation of the myResource, and an error, if there is any.
func (c *myResources) Create(myResource *v1.MyResource) (result *v1.MyResource, err error) {
	result = &v1.MyResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("myresources").
		Body(myResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a myResource and updates it. Returns the server's representation of the myResource, and an error, if there is any.
func (c *myResources) Update(myResource *v1.MyResource) (result *v1.MyResource, err error) {
	result = &v1.MyResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("myresources").
		Name(myResource.Name).
		Body(myResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the myResource and deletes it. Returns an error if one occurs.
func (c *myResources) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("myresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *myResources) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("myresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched myResource.
func (c *myResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MyResource, err error) {
	result = &v1.MyResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("myresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
