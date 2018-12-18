/*
Copyright 2018 The refunc Authors

TODO: choose a opensource licence.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta3

import (
	v1beta3 "github.com/refunc/refunc/pkg/apis/refunc/v1beta3"
	scheme "github.com/refunc/refunc/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// XenvsGetter has a method to return a XenvInterface.
// A group's client should implement this interface.
type XenvsGetter interface {
	Xenvs(namespace string) XenvInterface
}

// XenvInterface has methods to work with Xenv resources.
type XenvInterface interface {
	Create(*v1beta3.Xenv) (*v1beta3.Xenv, error)
	Update(*v1beta3.Xenv) (*v1beta3.Xenv, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta3.Xenv, error)
	List(opts v1.ListOptions) (*v1beta3.XenvList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta3.Xenv, err error)
	XenvExpansion
}

// xenvs implements XenvInterface
type xenvs struct {
	client rest.Interface
	ns     string
}

// newXenvs returns a Xenvs
func newXenvs(c *RefuncV1beta3Client, namespace string) *xenvs {
	return &xenvs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the xenv, and returns the corresponding xenv object, and an error if there is any.
func (c *xenvs) Get(name string, options v1.GetOptions) (result *v1beta3.Xenv, err error) {
	result = &v1beta3.Xenv{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("xenvs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Xenvs that match those selectors.
func (c *xenvs) List(opts v1.ListOptions) (result *v1beta3.XenvList, err error) {
	result = &v1beta3.XenvList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("xenvs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested xenvs.
func (c *xenvs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("xenvs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a xenv and creates it.  Returns the server's representation of the xenv, and an error, if there is any.
func (c *xenvs) Create(xenv *v1beta3.Xenv) (result *v1beta3.Xenv, err error) {
	result = &v1beta3.Xenv{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("xenvs").
		Body(xenv).
		Do().
		Into(result)
	return
}

// Update takes the representation of a xenv and updates it. Returns the server's representation of the xenv, and an error, if there is any.
func (c *xenvs) Update(xenv *v1beta3.Xenv) (result *v1beta3.Xenv, err error) {
	result = &v1beta3.Xenv{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("xenvs").
		Name(xenv.Name).
		Body(xenv).
		Do().
		Into(result)
	return
}

// Delete takes name of the xenv and deletes it. Returns an error if one occurs.
func (c *xenvs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("xenvs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *xenvs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("xenvs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched xenv.
func (c *xenvs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta3.Xenv, err error) {
	result = &v1beta3.Xenv{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("xenvs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}