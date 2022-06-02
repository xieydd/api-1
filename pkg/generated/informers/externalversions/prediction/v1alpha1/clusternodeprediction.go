// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "git.woa.com/crane/api/pkg/generated/clientset/versioned"
	internalinterfaces "git.woa.com/crane/api/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "git.woa.com/crane/api/pkg/generated/listers/prediction/v1alpha1"
	predictionv1alpha1 "git.woa.com/crane/api/prediction/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterNodePredictionInformer provides access to a shared informer and lister for
// ClusterNodePredictions.
type ClusterNodePredictionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterNodePredictionLister
}

type clusterNodePredictionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterNodePredictionInformer constructs a new informer for ClusterNodePrediction type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterNodePredictionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterNodePredictionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterNodePredictionInformer constructs a new informer for ClusterNodePrediction type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterNodePredictionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PredictionV1alpha1().ClusterNodePredictions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PredictionV1alpha1().ClusterNodePredictions(namespace).Watch(context.TODO(), options)
			},
		},
		&predictionv1alpha1.ClusterNodePrediction{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterNodePredictionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterNodePredictionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterNodePredictionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&predictionv1alpha1.ClusterNodePrediction{}, f.defaultInformer)
}

func (f *clusterNodePredictionInformer) Lister() v1alpha1.ClusterNodePredictionLister {
	return v1alpha1.NewClusterNodePredictionLister(f.Informer().GetIndexer())
}
