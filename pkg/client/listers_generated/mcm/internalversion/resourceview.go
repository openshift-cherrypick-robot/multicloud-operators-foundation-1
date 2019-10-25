// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	mcm "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/apis/mcm"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceViewLister helps list ResourceViews.
type ResourceViewLister interface {
	// List lists all ResourceViews in the indexer.
	List(selector labels.Selector) (ret []*mcm.ResourceView, err error)
	// ResourceViews returns an object that can list and get ResourceViews.
	ResourceViews(namespace string) ResourceViewNamespaceLister
	ResourceViewListerExpansion
}

// resourceViewLister implements the ResourceViewLister interface.
type resourceViewLister struct {
	indexer cache.Indexer
}

// NewResourceViewLister returns a new ResourceViewLister.
func NewResourceViewLister(indexer cache.Indexer) ResourceViewLister {
	return &resourceViewLister{indexer: indexer}
}

// List lists all ResourceViews in the indexer.
func (s *resourceViewLister) List(selector labels.Selector) (ret []*mcm.ResourceView, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*mcm.ResourceView))
	})
	return ret, err
}

// ResourceViews returns an object that can list and get ResourceViews.
func (s *resourceViewLister) ResourceViews(namespace string) ResourceViewNamespaceLister {
	return resourceViewNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceViewNamespaceLister helps list and get ResourceViews.
type ResourceViewNamespaceLister interface {
	// List lists all ResourceViews in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*mcm.ResourceView, err error)
	// Get retrieves the ResourceView from the indexer for a given namespace and name.
	Get(name string) (*mcm.ResourceView, error)
	ResourceViewNamespaceListerExpansion
}

// resourceViewNamespaceLister implements the ResourceViewNamespaceLister
// interface.
type resourceViewNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceViews in the indexer for a given namespace.
func (s resourceViewNamespaceLister) List(selector labels.Selector) (ret []*mcm.ResourceView, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*mcm.ResourceView))
	})
	return ret, err
}

// Get retrieves the ResourceView from the indexer for a given namespace and name.
func (s resourceViewNamespaceLister) Get(name string) (*mcm.ResourceView, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(mcm.Resource("resourceview"), name)
	}
	return obj.(*mcm.ResourceView), nil
}
