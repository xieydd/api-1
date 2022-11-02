// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	nodeoperationv1alpha1 "git.woa.com/crane/api/nodeoperation/v1alpha1"
	versioned "git.woa.com/crane/api/pkg/generated/clientset/versioned"
	internalinterfaces "git.woa.com/crane/api/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "git.woa.com/crane/api/pkg/generated/listers/nodeoperation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodeOperationInformer provides access to a shared informer and lister for
// NodeOperations.
type NodeOperationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NodeOperationLister
}

type nodeOperationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNodeOperationInformer constructs a new informer for NodeOperation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeOperationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeOperationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNodeOperationInformer constructs a new informer for NodeOperation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeOperationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NodeoperationV1alpha1().NodeOperations().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NodeoperationV1alpha1().NodeOperations().Watch(context.TODO(), options)
			},
		},
		&nodeoperationv1alpha1.NodeOperation{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeOperationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeOperationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeOperationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&nodeoperationv1alpha1.NodeOperation{}, f.defaultInformer)
}

func (f *nodeOperationInformer) Lister() v1alpha1.NodeOperationLister {
	return v1alpha1.NewNodeOperationLister(f.Informer().GetIndexer())
}
