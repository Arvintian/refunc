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

// FuncdevesGetter has a method to return a FuncdefInterface.
// A group's client should implement this interface.
type FuncdevesGetter interface {
	Funcdeves(namespace string) FuncdefInterface
}

// FuncdefInterface has methods to work with Funcdef resources.
type FuncdefInterface interface {
	Create(*v1beta3.Funcdef) (*v1beta3.Funcdef, error)
	Update(*v1beta3.Funcdef) (*v1beta3.Funcdef, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta3.Funcdef, error)
	List(opts v1.ListOptions) (*v1beta3.FuncdefList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta3.Funcdef, err error)
	FuncdefExpansion
}

// funcdeves implements FuncdefInterface
type funcdeves struct {
	client rest.Interface
	ns     string
}

// newFuncdeves returns a Funcdeves
func newFuncdeves(c *RefuncV1beta3Client, namespace string) *funcdeves {
	return &funcdeves{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the funcdef, and returns the corresponding funcdef object, and an error if there is any.
func (c *funcdeves) Get(name string, options v1.GetOptions) (result *v1beta3.Funcdef, err error) {
	result = &v1beta3.Funcdef{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("funcdeves").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Funcdeves that match those selectors.
func (c *funcdeves) List(opts v1.ListOptions) (result *v1beta3.FuncdefList, err error) {
	result = &v1beta3.FuncdefList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("funcdeves").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested funcdeves.
func (c *funcdeves) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("funcdeves").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a funcdef and creates it.  Returns the server's representation of the funcdef, and an error, if there is any.
func (c *funcdeves) Create(funcdef *v1beta3.Funcdef) (result *v1beta3.Funcdef, err error) {
	result = &v1beta3.Funcdef{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("funcdeves").
		Body(funcdef).
		Do().
		Into(result)
	return
}

// Update takes the representation of a funcdef and updates it. Returns the server's representation of the funcdef, and an error, if there is any.
func (c *funcdeves) Update(funcdef *v1beta3.Funcdef) (result *v1beta3.Funcdef, err error) {
	result = &v1beta3.Funcdef{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("funcdeves").
		Name(funcdef.Name).
		Body(funcdef).
		Do().
		Into(result)
	return
}

// Delete takes name of the funcdef and deletes it. Returns an error if one occurs.
func (c *funcdeves) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("funcdeves").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *funcdeves) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("funcdeves").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched funcdef.
func (c *funcdeves) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta3.Funcdef, err error) {
	result = &v1beta3.Funcdef{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("funcdeves").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
