// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "git.woa.com/crane/api/pkg/generated/clientset/versioned/typed/nodeoperation/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNodeoperationV1alpha1 struct {
	*testing.Fake
}

func (c *FakeNodeoperationV1alpha1) NodeOperations() v1alpha1.NodeOperationInterface {
	return &FakeNodeOperations{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNodeoperationV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}