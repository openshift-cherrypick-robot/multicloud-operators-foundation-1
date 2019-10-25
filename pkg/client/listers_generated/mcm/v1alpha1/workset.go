// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/apis/mcm/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkSetLister helps list WorkSets.
type WorkSetLister interface {
	// List lists all WorkSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WorkSet, err error)
	// WorkSets returns an object that can list and get WorkSets.
	WorkSets(namespace string) WorkSetNamespaceLister
	WorkSetListerExpansion
}

// workSetLister implements the WorkSetLister interface.
type workSetLister struct {
	indexer cache.Indexer
}

// NewWorkSetLister returns a new WorkSetLister.
func NewWorkSetLister(indexer cache.Indexer) WorkSetLister {
	return &workSetLister{indexer: indexer}
}

// List lists all WorkSets in the indexer.
func (s *workSetLister) List(selector labels.Selector) (ret []*v1alpha1.WorkSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkSet))
	})
	return ret, err
}

// WorkSets returns an object that can list and get WorkSets.
func (s *workSetLister) WorkSets(namespace string) WorkSetNamespaceLister {
	return workSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkSetNamespaceLister helps list and get WorkSets.
type WorkSetNamespaceLister interface {
	// List lists all WorkSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WorkSet, err error)
	// Get retrieves the WorkSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WorkSet, error)
	WorkSetNamespaceListerExpansion
}

// workSetNamespaceLister implements the WorkSetNamespaceLister
// interface.
type workSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorkSets in the indexer for a given namespace.
func (s workSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WorkSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkSet))
	})
	return ret, err
}

// Get retrieves the WorkSet from the indexer for a given namespace and name.
func (s workSetNamespaceLister) Get(name string) (*v1alpha1.WorkSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workset"), name)
	}
	return obj.(*v1alpha1.WorkSet), nil
}
