/*
Copyright 2020 The Jenkins X Authors.

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
	v1alpha1 "github.com/jenkins-x/jx-preview/pkg/apis/preview/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PreviewLister helps list Previews.
// All objects returned here must be treated as read-only.
type PreviewLister interface {
	// List lists all Previews in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Preview, err error)
	// Previews returns an object that can list and get Previews.
	Previews(namespace string) PreviewNamespaceLister
	PreviewListerExpansion
}

// previewLister implements the PreviewLister interface.
type previewLister struct {
	indexer cache.Indexer
}

// NewPreviewLister returns a new PreviewLister.
func NewPreviewLister(indexer cache.Indexer) PreviewLister {
	return &previewLister{indexer: indexer}
}

// List lists all Previews in the indexer.
func (s *previewLister) List(selector labels.Selector) (ret []*v1alpha1.Preview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Preview))
	})
	return ret, err
}

// Previews returns an object that can list and get Previews.
func (s *previewLister) Previews(namespace string) PreviewNamespaceLister {
	return previewNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PreviewNamespaceLister helps list and get Previews.
// All objects returned here must be treated as read-only.
type PreviewNamespaceLister interface {
	// List lists all Previews in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Preview, err error)
	// Get retrieves the Preview from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Preview, error)
	PreviewNamespaceListerExpansion
}

// previewNamespaceLister implements the PreviewNamespaceLister
// interface.
type previewNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Previews in the indexer for a given namespace.
func (s previewNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Preview, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Preview))
	})
	return ret, err
}

// Get retrieves the Preview from the indexer for a given namespace and name.
func (s previewNamespaceLister) Get(name string) (*v1alpha1.Preview, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("preview"), name)
	}
	return obj.(*v1alpha1.Preview), nil
}
