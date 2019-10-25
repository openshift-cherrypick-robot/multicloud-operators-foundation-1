// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	mcmv1alpha1 "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/apis/mcm/v1alpha1"
	clientset "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/client/listers_generated/mcm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceViewInformer provides access to a shared informer and lister for
// ResourceViews.
type ResourceViewInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceViewLister
}

type resourceViewInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceViewInformer constructs a new informer for ResourceView type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceViewInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceViewInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceViewInformer constructs a new informer for ResourceView type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceViewInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.McmV1alpha1().ResourceViews(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.McmV1alpha1().ResourceViews(namespace).Watch(options)
			},
		},
		&mcmv1alpha1.ResourceView{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceViewInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceViewInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceViewInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mcmv1alpha1.ResourceView{}, f.defaultInformer)
}

func (f *resourceViewInformer) Lister() v1alpha1.ResourceViewLister {
	return v1alpha1.NewResourceViewLister(f.Informer().GetIndexer())
}
