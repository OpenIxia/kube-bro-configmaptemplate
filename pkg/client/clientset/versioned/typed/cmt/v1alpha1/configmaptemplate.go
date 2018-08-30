/*
Copyright 2018 Keysight Technologies

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
	v1alpha1 "github.com/openixia/cmt-controller/pkg/apis/cmt/v1alpha1"
	scheme "github.com/openixia/cmt-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConfigMapTemplatesGetter has a method to return a ConfigMapTemplateInterface.
// A group's client should implement this interface.
type ConfigMapTemplatesGetter interface {
	ConfigMapTemplates(namespace string) ConfigMapTemplateInterface
}

// ConfigMapTemplateInterface has methods to work with ConfigMapTemplate resources.
type ConfigMapTemplateInterface interface {
	Create(*v1alpha1.ConfigMapTemplate) (*v1alpha1.ConfigMapTemplate, error)
	Update(*v1alpha1.ConfigMapTemplate) (*v1alpha1.ConfigMapTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ConfigMapTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.ConfigMapTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigMapTemplate, err error)
	ConfigMapTemplateExpansion
}

// configMapTemplates implements ConfigMapTemplateInterface
type configMapTemplates struct {
	client rest.Interface
	ns     string
}

// newConfigMapTemplates returns a ConfigMapTemplates
func newConfigMapTemplates(c *CmtV1alpha1Client, namespace string) *configMapTemplates {
	return &configMapTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the configMapTemplate, and returns the corresponding configMapTemplate object, and an error if there is any.
func (c *configMapTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.ConfigMapTemplate, err error) {
	result = &v1alpha1.ConfigMapTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configmaptemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigMapTemplates that match those selectors.
func (c *configMapTemplates) List(opts v1.ListOptions) (result *v1alpha1.ConfigMapTemplateList, err error) {
	result = &v1alpha1.ConfigMapTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configmaptemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configMapTemplates.
func (c *configMapTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configmaptemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a configMapTemplate and creates it.  Returns the server's representation of the configMapTemplate, and an error, if there is any.
func (c *configMapTemplates) Create(configMapTemplate *v1alpha1.ConfigMapTemplate) (result *v1alpha1.ConfigMapTemplate, err error) {
	result = &v1alpha1.ConfigMapTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configmaptemplates").
		Body(configMapTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a configMapTemplate and updates it. Returns the server's representation of the configMapTemplate, and an error, if there is any.
func (c *configMapTemplates) Update(configMapTemplate *v1alpha1.ConfigMapTemplate) (result *v1alpha1.ConfigMapTemplate, err error) {
	result = &v1alpha1.ConfigMapTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configmaptemplates").
		Name(configMapTemplate.Name).
		Body(configMapTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the configMapTemplate and deletes it. Returns an error if one occurs.
func (c *configMapTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configmaptemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configMapTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configmaptemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched configMapTemplate.
func (c *configMapTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigMapTemplate, err error) {
	result = &v1alpha1.ConfigMapTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configmaptemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
