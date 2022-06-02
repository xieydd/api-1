// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "git.woa.com/crane/api/analysis/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RecommendationLister helps list Recommendations.
// All objects returned here must be treated as read-only.
type RecommendationLister interface {
	// List lists all Recommendations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Recommendation, err error)
	// Recommendations returns an object that can list and get Recommendations.
	Recommendations(namespace string) RecommendationNamespaceLister
	RecommendationListerExpansion
}

// recommendationLister implements the RecommendationLister interface.
type recommendationLister struct {
	indexer cache.Indexer
}

// NewRecommendationLister returns a new RecommendationLister.
func NewRecommendationLister(indexer cache.Indexer) RecommendationLister {
	return &recommendationLister{indexer: indexer}
}

// List lists all Recommendations in the indexer.
func (s *recommendationLister) List(selector labels.Selector) (ret []*v1alpha1.Recommendation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Recommendation))
	})
	return ret, err
}

// Recommendations returns an object that can list and get Recommendations.
func (s *recommendationLister) Recommendations(namespace string) RecommendationNamespaceLister {
	return recommendationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RecommendationNamespaceLister helps list and get Recommendations.
// All objects returned here must be treated as read-only.
type RecommendationNamespaceLister interface {
	// List lists all Recommendations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Recommendation, err error)
	// Get retrieves the Recommendation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Recommendation, error)
	RecommendationNamespaceListerExpansion
}

// recommendationNamespaceLister implements the RecommendationNamespaceLister
// interface.
type recommendationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Recommendations in the indexer for a given namespace.
func (s recommendationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Recommendation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Recommendation))
	})
	return ret, err
}

// Get retrieves the Recommendation from the indexer for a given namespace and name.
func (s recommendationNamespaceLister) Get(name string) (*v1alpha1.Recommendation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("recommendation"), name)
	}
	return obj.(*v1alpha1.Recommendation), nil
}
