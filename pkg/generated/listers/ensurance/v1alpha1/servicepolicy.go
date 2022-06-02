// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "git.woa.com/crane/api/ensurance/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServicePolicyLister helps list ServicePolicies.
// All objects returned here must be treated as read-only.
type ServicePolicyLister interface {
	// List lists all ServicePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicePolicy, err error)
	// Get retrieves the ServicePolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServicePolicy, error)
	ServicePolicyListerExpansion
}

// servicePolicyLister implements the ServicePolicyLister interface.
type servicePolicyLister struct {
	indexer cache.Indexer
}

// NewServicePolicyLister returns a new ServicePolicyLister.
func NewServicePolicyLister(indexer cache.Indexer) ServicePolicyLister {
	return &servicePolicyLister{indexer: indexer}
}

// List lists all ServicePolicies in the indexer.
func (s *servicePolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ServicePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicePolicy))
	})
	return ret, err
}

// Get retrieves the ServicePolicy from the index for a given name.
func (s *servicePolicyLister) Get(name string) (*v1alpha1.ServicePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicepolicy"), name)
	}
	return obj.(*v1alpha1.ServicePolicy), nil
}
