// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "git.woa.com/crane/api/autoscaling/v1alpha1"
	scheme "git.woa.com/crane/api/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SubstitutesGetter has a method to return a SubstituteInterface.
// A group's client should implement this interface.
type SubstitutesGetter interface {
	Substitutes(namespace string) SubstituteInterface
}

// SubstituteInterface has methods to work with Substitute resources.
type SubstituteInterface interface {
	Create(ctx context.Context, substitute *v1alpha1.Substitute, opts v1.CreateOptions) (*v1alpha1.Substitute, error)
	Update(ctx context.Context, substitute *v1alpha1.Substitute, opts v1.UpdateOptions) (*v1alpha1.Substitute, error)
	UpdateStatus(ctx context.Context, substitute *v1alpha1.Substitute, opts v1.UpdateOptions) (*v1alpha1.Substitute, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Substitute, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SubstituteList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Substitute, err error)
	SubstituteExpansion
}

// substitutes implements SubstituteInterface
type substitutes struct {
	client rest.Interface
	ns     string
}

// newSubstitutes returns a Substitutes
func newSubstitutes(c *AutoscalingV1alpha1Client, namespace string) *substitutes {
	return &substitutes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the substitute, and returns the corresponding substitute object, and an error if there is any.
func (c *substitutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Substitute, err error) {
	result = &v1alpha1.Substitute{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("substitutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Substitutes that match those selectors.
func (c *substitutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubstituteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SubstituteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("substitutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested substitutes.
func (c *substitutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("substitutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a substitute and creates it.  Returns the server's representation of the substitute, and an error, if there is any.
func (c *substitutes) Create(ctx context.Context, substitute *v1alpha1.Substitute, opts v1.CreateOptions) (result *v1alpha1.Substitute, err error) {
	result = &v1alpha1.Substitute{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("substitutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(substitute).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a substitute and updates it. Returns the server's representation of the substitute, and an error, if there is any.
func (c *substitutes) Update(ctx context.Context, substitute *v1alpha1.Substitute, opts v1.UpdateOptions) (result *v1alpha1.Substitute, err error) {
	result = &v1alpha1.Substitute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("substitutes").
		Name(substitute.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(substitute).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *substitutes) UpdateStatus(ctx context.Context, substitute *v1alpha1.Substitute, opts v1.UpdateOptions) (result *v1alpha1.Substitute, err error) {
	result = &v1alpha1.Substitute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("substitutes").
		Name(substitute.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(substitute).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the substitute and deletes it. Returns an error if one occurs.
func (c *substitutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("substitutes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *substitutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("substitutes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched substitute.
func (c *substitutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Substitute, err error) {
	result = &v1alpha1.Substitute{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("substitutes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
