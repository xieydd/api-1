// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "git.woa.com/crane/api/prediction/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TimeSeriesPredictionLister helps list TimeSeriesPredictions.
// All objects returned here must be treated as read-only.
type TimeSeriesPredictionLister interface {
	// List lists all TimeSeriesPredictions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TimeSeriesPrediction, err error)
	// TimeSeriesPredictions returns an object that can list and get TimeSeriesPredictions.
	TimeSeriesPredictions(namespace string) TimeSeriesPredictionNamespaceLister
	TimeSeriesPredictionListerExpansion
}

// timeSeriesPredictionLister implements the TimeSeriesPredictionLister interface.
type timeSeriesPredictionLister struct {
	indexer cache.Indexer
}

// NewTimeSeriesPredictionLister returns a new TimeSeriesPredictionLister.
func NewTimeSeriesPredictionLister(indexer cache.Indexer) TimeSeriesPredictionLister {
	return &timeSeriesPredictionLister{indexer: indexer}
}

// List lists all TimeSeriesPredictions in the indexer.
func (s *timeSeriesPredictionLister) List(selector labels.Selector) (ret []*v1alpha1.TimeSeriesPrediction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TimeSeriesPrediction))
	})
	return ret, err
}

// TimeSeriesPredictions returns an object that can list and get TimeSeriesPredictions.
func (s *timeSeriesPredictionLister) TimeSeriesPredictions(namespace string) TimeSeriesPredictionNamespaceLister {
	return timeSeriesPredictionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TimeSeriesPredictionNamespaceLister helps list and get TimeSeriesPredictions.
// All objects returned here must be treated as read-only.
type TimeSeriesPredictionNamespaceLister interface {
	// List lists all TimeSeriesPredictions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TimeSeriesPrediction, err error)
	// Get retrieves the TimeSeriesPrediction from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TimeSeriesPrediction, error)
	TimeSeriesPredictionNamespaceListerExpansion
}

// timeSeriesPredictionNamespaceLister implements the TimeSeriesPredictionNamespaceLister
// interface.
type timeSeriesPredictionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TimeSeriesPredictions in the indexer for a given namespace.
func (s timeSeriesPredictionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TimeSeriesPrediction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TimeSeriesPrediction))
	})
	return ret, err
}

// Get retrieves the TimeSeriesPrediction from the indexer for a given namespace and name.
func (s timeSeriesPredictionNamespaceLister) Get(name string) (*v1alpha1.TimeSeriesPrediction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("timeseriesprediction"), name)
	}
	return obj.(*v1alpha1.TimeSeriesPrediction), nil
}
