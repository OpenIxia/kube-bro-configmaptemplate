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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "keysight.io/cmt-controller/pkg/apis/cmt/v1alpha1"
)

// ConfigMapTemplateLister helps list ConfigMapTemplates.
type ConfigMapTemplateLister interface {
	// List lists all ConfigMapTemplates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigMapTemplate, err error)
	// ConfigMapTemplates returns an object that can list and get ConfigMapTemplates.
	ConfigMapTemplates(namespace string) ConfigMapTemplateNamespaceLister
	ConfigMapTemplateListerExpansion
}

// configMapTemplateLister implements the ConfigMapTemplateLister interface.
type configMapTemplateLister struct {
	indexer cache.Indexer
}

// NewConfigMapTemplateLister returns a new ConfigMapTemplateLister.
func NewConfigMapTemplateLister(indexer cache.Indexer) ConfigMapTemplateLister {
	return &configMapTemplateLister{indexer: indexer}
}

// List lists all ConfigMapTemplates in the indexer.
func (s *configMapTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigMapTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigMapTemplate))
	})
	return ret, err
}

// ConfigMapTemplates returns an object that can list and get ConfigMapTemplates.
func (s *configMapTemplateLister) ConfigMapTemplates(namespace string) ConfigMapTemplateNamespaceLister {
	return configMapTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigMapTemplateNamespaceLister helps list and get ConfigMapTemplates.
type ConfigMapTemplateNamespaceLister interface {
	// List lists all ConfigMapTemplates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigMapTemplate, err error)
	// Get retrieves the ConfigMapTemplate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ConfigMapTemplate, error)
	ConfigMapTemplateNamespaceListerExpansion
}

// configMapTemplateNamespaceLister implements the ConfigMapTemplateNamespaceLister
// interface.
type configMapTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigMapTemplates in the indexer for a given namespace.
func (s configMapTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigMapTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigMapTemplate))
	})
	return ret, err
}

// Get retrieves the ConfigMapTemplate from the indexer for a given namespace and name.
func (s configMapTemplateNamespaceLister) Get(name string) (*v1alpha1.ConfigMapTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("configmaptemplate"), name)
	}
	return obj.(*v1alpha1.ConfigMapTemplate), nil
}