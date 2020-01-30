/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SecretBindingLister helps list SecretBindings.
type SecretBindingLister interface {
	// List lists all SecretBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecretBinding, err error)
	// SecretBindings returns an object that can list and get SecretBindings.
	SecretBindings(namespace string) SecretBindingNamespaceLister
	SecretBindingListerExpansion
}

// secretBindingLister implements the SecretBindingLister interface.
type secretBindingLister struct {
	indexer cache.Indexer
}

// NewSecretBindingLister returns a new SecretBindingLister.
func NewSecretBindingLister(indexer cache.Indexer) SecretBindingLister {
	return &secretBindingLister{indexer: indexer}
}

// List lists all SecretBindings in the indexer.
func (s *secretBindingLister) List(selector labels.Selector) (ret []*v1alpha1.SecretBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretBinding))
	})
	return ret, err
}

// SecretBindings returns an object that can list and get SecretBindings.
func (s *secretBindingLister) SecretBindings(namespace string) SecretBindingNamespaceLister {
	return secretBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecretBindingNamespaceLister helps list and get SecretBindings.
type SecretBindingNamespaceLister interface {
	// List lists all SecretBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecretBinding, err error)
	// Get retrieves the SecretBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecretBinding, error)
	SecretBindingNamespaceListerExpansion
}

// secretBindingNamespaceLister implements the SecretBindingNamespaceLister
// interface.
type secretBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecretBindings in the indexer for a given namespace.
func (s secretBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecretBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretBinding))
	})
	return ret, err
}

// Get retrieves the SecretBinding from the indexer for a given namespace and name.
func (s secretBindingNamespaceLister) Get(name string) (*v1alpha1.SecretBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("secretbinding"), name)
	}
	return obj.(*v1alpha1.SecretBinding), nil
}
