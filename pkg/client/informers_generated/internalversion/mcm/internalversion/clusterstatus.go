// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	mcm "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/apis/mcm"
	internalclientset "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/client/listers_generated/mcm/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterStatusInformer provides access to a shared informer and lister for
// ClusterStatuses.
type ClusterStatusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ClusterStatusLister
}

type clusterStatusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterStatusInformer constructs a new informer for ClusterStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterStatusInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterStatusInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterStatusInformer constructs a new informer for ClusterStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterStatusInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Mcm().ClusterStatuses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Mcm().ClusterStatuses(namespace).Watch(options)
			},
		},
		&mcm.ClusterStatus{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterStatusInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterStatusInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterStatusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mcm.ClusterStatus{}, f.defaultInformer)
}

func (f *clusterStatusInformer) Lister() internalversion.ClusterStatusLister {
	return internalversion.NewClusterStatusLister(f.Informer().GetIndexer())
}
