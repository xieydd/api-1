// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	scheme "git.woa.com/crane/api/pkg/generated/clientset/versioned/scheme"
	v1alpha1 "git.woa.com/crane/api/scheduling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterNodeResourcePoliciesGetter has a method to return a ClusterNodeResourcePolicyInterface.
// A group's client should implement this interface.
type ClusterNodeResourcePoliciesGetter interface {
	ClusterNodeResourcePolicies() ClusterNodeResourcePolicyInterface
}

// ClusterNodeResourcePolicyInterface has methods to work with ClusterNodeResourcePolicy resources.
type ClusterNodeResourcePolicyInterface interface {
	Create(ctx context.Context, clusterNodeResourcePolicy *v1alpha1.ClusterNodeResourcePolicy, opts v1.CreateOptions) (*v1alpha1.ClusterNodeResourcePolicy, error)
	Update(ctx context.Context, clusterNodeResourcePolicy *v1alpha1.ClusterNodeResourcePolicy, opts v1.UpdateOptions) (*v1alpha1.ClusterNodeResourcePolicy, error)
	UpdateStatus(ctx context.Context, clusterNodeResourcePolicy *v1alpha1.ClusterNodeResourcePolicy, opts v1.UpdateOptions) (*v1alpha1.ClusterNodeResourcePolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterNodeResourcePolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterNodeResourcePolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterNodeResourcePolicy, err error)
	ClusterNodeResourcePolicyExpansion
}

// clusterNodeResourcePolicies implements ClusterNodeResourcePolicyInterface
type clusterNodeResourcePolicies struct {
	client rest.Interface
}

// newClusterNodeResourcePolicies returns a ClusterNodeResourcePolicies
func newClusterNodeResourcePolicies(c *SchedulingV1alpha1Client) *clusterNodeResourcePolicies {
	return &clusterNodeResourcePolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterNodeResourcePolicy, and returns the corresponding clusterNodeResourcePolicy object, and an error if there is any.
func (c *clusterNodeResourcePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterNodeResourcePolicy, err error) {
	result = &v1alpha1.ClusterNodeResourcePolicy{}
	err = c.client.Get().
		Resource("clusternoderesourcepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterNodeResourcePolicies that match those selectors.
func (c *clusterNodeResourcePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterNodeResourcePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterNodeResourcePolicyList{}
	err = c.client.Get().
		Resource("clusternoderesourcepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterNodeResourcePolicies.
func (c *clusterNodeResourcePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusternoderesourcepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterNodeResourcePolicy and creates it.  Returns the server's representation of the clusterNodeResourcePolicy, and an error, if there is any.
func (c *clusterNodeResourcePolicies) Create(ctx context.Context, clusterNodeResourcePolicy *v1alpha1.ClusterNodeResourcePolicy, opts v1.CreateOptions) (result *v1alpha1.ClusterNodeResourcePolicy, err error) {
	result = &v1alpha1.ClusterNodeResourcePolicy{}
	err = c.client.Post().
		Resource("clusternoderesourcepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterNodeResourcePolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterNodeResourcePolicy and updates it. Returns the server's representation of the clusterNodeResourcePolicy, and an error, if there is any.
func (c *clusterNodeResourcePolicies) Update(ctx context.Context, clusterNodeResourcePolicy *v1alpha1.ClusterNodeResourcePolicy, opts v1.UpdateOptions) (result *v1alpha1.ClusterNodeResourcePolicy, err error) {
	result = &v1alpha1.ClusterNodeResourcePolicy{}
	err = c.client.Put().
		Resource("clusternoderesourcepolicies").
		Name(clusterNodeResourcePolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterNodeResourcePolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterNodeResourcePolicies) UpdateStatus(ctx context.Context, clusterNodeResourcePolicy *v1alpha1.ClusterNodeResourcePolicy, opts v1.UpdateOptions) (result *v1alpha1.ClusterNodeResourcePolicy, err error) {
	result = &v1alpha1.ClusterNodeResourcePolicy{}
	err = c.client.Put().
		Resource("clusternoderesourcepolicies").
		Name(clusterNodeResourcePolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterNodeResourcePolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterNodeResourcePolicy and deletes it. Returns an error if one occurs.
func (c *clusterNodeResourcePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusternoderesourcepolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterNodeResourcePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusternoderesourcepolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterNodeResourcePolicy.
func (c *clusterNodeResourcePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterNodeResourcePolicy, err error) {
	result = &v1alpha1.ClusterNodeResourcePolicy{}
	err = c.client.Patch(pt).
		Resource("clusternoderesourcepolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
