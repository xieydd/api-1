// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "git.woa.com/crane/api/autoscaling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEffectiveVerticalPodAutoscalers implements EffectiveVerticalPodAutoscalerInterface
type FakeEffectiveVerticalPodAutoscalers struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var effectiveverticalpodautoscalersResource = schema.GroupVersionResource{Group: "autoscaling.crane.io", Version: "v1alpha1", Resource: "effectiveverticalpodautoscalers"}

var effectiveverticalpodautoscalersKind = schema.GroupVersionKind{Group: "autoscaling.crane.io", Version: "v1alpha1", Kind: "EffectiveVerticalPodAutoscaler"}

// Get takes name of the effectiveVerticalPodAutoscaler, and returns the corresponding effectiveVerticalPodAutoscaler object, and an error if there is any.
func (c *FakeEffectiveVerticalPodAutoscalers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EffectiveVerticalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(effectiveverticalpodautoscalersResource, c.ns, name), &v1alpha1.EffectiveVerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EffectiveVerticalPodAutoscaler), err
}

// List takes label and field selectors, and returns the list of EffectiveVerticalPodAutoscalers that match those selectors.
func (c *FakeEffectiveVerticalPodAutoscalers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EffectiveVerticalPodAutoscalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(effectiveverticalpodautoscalersResource, effectiveverticalpodautoscalersKind, c.ns, opts), &v1alpha1.EffectiveVerticalPodAutoscalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EffectiveVerticalPodAutoscalerList{ListMeta: obj.(*v1alpha1.EffectiveVerticalPodAutoscalerList).ListMeta}
	for _, item := range obj.(*v1alpha1.EffectiveVerticalPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested effectiveVerticalPodAutoscalers.
func (c *FakeEffectiveVerticalPodAutoscalers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(effectiveverticalpodautoscalersResource, c.ns, opts))

}

// Create takes the representation of a effectiveVerticalPodAutoscaler and creates it.  Returns the server's representation of the effectiveVerticalPodAutoscaler, and an error, if there is any.
func (c *FakeEffectiveVerticalPodAutoscalers) Create(ctx context.Context, effectiveVerticalPodAutoscaler *v1alpha1.EffectiveVerticalPodAutoscaler, opts v1.CreateOptions) (result *v1alpha1.EffectiveVerticalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(effectiveverticalpodautoscalersResource, c.ns, effectiveVerticalPodAutoscaler), &v1alpha1.EffectiveVerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EffectiveVerticalPodAutoscaler), err
}

// Update takes the representation of a effectiveVerticalPodAutoscaler and updates it. Returns the server's representation of the effectiveVerticalPodAutoscaler, and an error, if there is any.
func (c *FakeEffectiveVerticalPodAutoscalers) Update(ctx context.Context, effectiveVerticalPodAutoscaler *v1alpha1.EffectiveVerticalPodAutoscaler, opts v1.UpdateOptions) (result *v1alpha1.EffectiveVerticalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(effectiveverticalpodautoscalersResource, c.ns, effectiveVerticalPodAutoscaler), &v1alpha1.EffectiveVerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EffectiveVerticalPodAutoscaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEffectiveVerticalPodAutoscalers) UpdateStatus(ctx context.Context, effectiveVerticalPodAutoscaler *v1alpha1.EffectiveVerticalPodAutoscaler, opts v1.UpdateOptions) (*v1alpha1.EffectiveVerticalPodAutoscaler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(effectiveverticalpodautoscalersResource, "status", c.ns, effectiveVerticalPodAutoscaler), &v1alpha1.EffectiveVerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EffectiveVerticalPodAutoscaler), err
}

// Delete takes name of the effectiveVerticalPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *FakeEffectiveVerticalPodAutoscalers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(effectiveverticalpodautoscalersResource, c.ns, name), &v1alpha1.EffectiveVerticalPodAutoscaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEffectiveVerticalPodAutoscalers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(effectiveverticalpodautoscalersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EffectiveVerticalPodAutoscalerList{})
	return err
}

// Patch applies the patch and returns the patched effectiveVerticalPodAutoscaler.
func (c *FakeEffectiveVerticalPodAutoscalers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EffectiveVerticalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(effectiveverticalpodautoscalersResource, c.ns, name, pt, data, subresources...), &v1alpha1.EffectiveVerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EffectiveVerticalPodAutoscaler), err
}
