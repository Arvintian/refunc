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

// FuncinstsGetter has a method to return a FuncinstInterface.
// A group's client should implement this interface.
type FuncinstsGetter interface {
	Funcinsts(namespace string) FuncinstInterface
}

// FuncinstInterface has methods to work with Funcinst resources.
type FuncinstInterface interface {
	Create(*v1beta3.Funcinst) (*v1beta3.Funcinst, error)
	Update(*v1beta3.Funcinst) (*v1beta3.Funcinst, error)
	UpdateStatus(*v1beta3.Funcinst) (*v1beta3.Funcinst, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta3.Funcinst, error)
	List(opts v1.ListOptions) (*v1beta3.FuncinstList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta3.Funcinst, err error)
	FuncinstExpansion
}

// funcinsts implements FuncinstInterface
type funcinsts struct {
	client rest.Interface
	ns     string
}

// newFuncinsts returns a Funcinsts
func newFuncinsts(c *RefuncV1beta3Client, namespace string) *funcinsts {
	return &funcinsts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the funcinst, and returns the corresponding funcinst object, and an error if there is any.
func (c *funcinsts) Get(name string, options v1.GetOptions) (result *v1beta3.Funcinst, err error) {
	result = &v1beta3.Funcinst{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("funcinsts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Funcinsts that match those selectors.
func (c *funcinsts) List(opts v1.ListOptions) (result *v1beta3.FuncinstList, err error) {
	result = &v1beta3.FuncinstList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("funcinsts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested funcinsts.
func (c *funcinsts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("funcinsts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a funcinst and creates it.  Returns the server's representation of the funcinst, and an error, if there is any.
func (c *funcinsts) Create(funcinst *v1beta3.Funcinst) (result *v1beta3.Funcinst, err error) {
	result = &v1beta3.Funcinst{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("funcinsts").
		Body(funcinst).
		Do().
		Into(result)
	return
}

// Update takes the representation of a funcinst and updates it. Returns the server's representation of the funcinst, and an error, if there is any.
func (c *funcinsts) Update(funcinst *v1beta3.Funcinst) (result *v1beta3.Funcinst, err error) {
	result = &v1beta3.Funcinst{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("funcinsts").
		Name(funcinst.Name).
		Body(funcinst).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *funcinsts) UpdateStatus(funcinst *v1beta3.Funcinst) (result *v1beta3.Funcinst, err error) {
	result = &v1beta3.Funcinst{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("funcinsts").
		Name(funcinst.Name).
		SubResource("status").
		Body(funcinst).
		Do().
		Into(result)
	return
}

// Delete takes name of the funcinst and deletes it. Returns an error if one occurs.
func (c *funcinsts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("funcinsts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *funcinsts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("funcinsts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched funcinst.
func (c *funcinsts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta3.Funcinst, err error) {
	result = &v1beta3.Funcinst{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("funcinsts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
