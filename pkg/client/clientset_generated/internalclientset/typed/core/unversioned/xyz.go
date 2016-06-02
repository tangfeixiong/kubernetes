package unversioned

import (
	api "k8s.io/kubernetes/pkg/api"
	watch "k8s.io/kubernetes/pkg/watch"
)

// SecurityContextConstraintsGetter has a method to return a SecurityContextConstraintsInterface.
// A group's client should implement this interface.
type SecurityContextConstraintsGetter interface {
    SecurityContextConstraints() SecurityContextConstraintsInterface
}


// SecurityContextConstraintsInterface has methods to work with SecurityContextConstraints resources.
type SecurityContextConstraintsInterface interface {
	Create(*api.SecurityContextConstraints) (*api.SecurityContextConstraints, error)
	Update(*api.SecurityContextConstraints) (*api.SecurityContextConstraints, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*api.SecurityContextConstraints, error)
	List(opts api.ListOptions) (*api.SecurityContextConstraintsList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
	SecurityContextConstraintsExpansion
}

// securityContextConstraints implements SecurityContextConstraintsInterface
type securityContextConstraints struct {
	client *CoreClient
}

// newSecurityContextConstraints returns a SecurityContextConstraints
func newSecurityContextConstraints(c *CoreClient) *securityContextConstraints {
	return &securityContextConstraints{
		client: c,
	}
}

// Create takes the representation of a securityContextConstraints and creates it.  Returns the server's representation of the securityContextConstraints, and an error, if there is any.
func (c *securityContextConstraints) Create(securityContextConstraints *api.SecurityContextConstraints) (result *api.SecurityContextConstraints, err error) {
	result = &api.SecurityContextConstraints{}
	err = c.client.Post().
		Resource("securitycontextconstraints").
		Body(securityContextConstraints).
		Do().
		Into(result)
	return
}

// Update takes the representation of a securityContextConstraints and updates it. Returns the server's representation of the securityContextConstraints, and an error, if there is any.
func (c *securityContextConstraints) Update(securityContextConstraints *api.SecurityContextConstraints) (result *api.SecurityContextConstraints, err error) {
	result = &api.SecurityContextConstraints{}
	err = c.client.Put().
		Resource("securitycontextconstraints").
		Name(securityContextConstraints.Name).
		Body(securityContextConstraints).
		Do().
		Into(result)
	return
}

// Delete takes name of the securityContextConstraints and deletes it. Returns an error if one occurs.
func (c *securityContextConstraints) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Resource("securitycontextconstraints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *securityContextConstraints) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		Resource("securitycontextconstraints").
		VersionedParams(&listOptions, api.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the securityContextConstraints, and returns the corresponding securityContextConstraints object, and an error if there is any.
func (c *securityContextConstraints) Get(name string) (result *api.SecurityContextConstraints, err error) {
	result = &api.SecurityContextConstraints{}
	err = c.client.Get().
		Resource("securitycontextconstraints").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecurityContextConstraints that match those selectors.
func (c *securityContextConstraints) List(opts api.ListOptions) (result *api.SecurityContextConstraintsList, err error) {
	result = &api.SecurityContextConstraintsList{}
	err = c.client.Get().
		Resource("securitycontextconstraints").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested securityContextConstraints.
func (c *securityContextConstraints) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Resource("securitycontextconstraints").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}


func (c *CoreClient) SecurityContextConstraints() SecurityContextConstraintsInterface {
	return newSecurityContextConstraints(c)
}

type SecurityContextConstraintsExpansion interface{}

