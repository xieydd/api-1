// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "git.woa.com/crane/api/prediction/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterNodePredictions implements ClusterNodePredictionInterface
type FakeClusterNodePredictions struct {
	Fake *FakePredictionV1alpha1
	ns   string
}

var clusternodepredictionsResource = schema.GroupVersionResource{Group: "prediction.crane.io", Version: "v1alpha1", Resource: "clusternodepredictions"}

var clusternodepredictionsKind = schema.GroupVersionKind{Group: "prediction.crane.io", Version: "v1alpha1", Kind: "ClusterNodePrediction"}

// Get takes name of the clusterNodePrediction, and returns the corresponding clusterNodePrediction object, and an error if there is any.
func (c *FakeClusterNodePredictions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterNodePrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusternodepredictionsResource, c.ns, name), &v1alpha1.ClusterNodePrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterNodePrediction), err
}

// List takes label and field selectors, and returns the list of ClusterNodePredictions that match those selectors.
func (c *FakeClusterNodePredictions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterNodePredictionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusternodepredictionsResource, clusternodepredictionsKind, c.ns, opts), &v1alpha1.ClusterNodePredictionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterNodePredictionList{ListMeta: obj.(*v1alpha1.ClusterNodePredictionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterNodePredictionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterNodePredictions.
func (c *FakeClusterNodePredictions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusternodepredictionsResource, c.ns, opts))

}

// Create takes the representation of a clusterNodePrediction and creates it.  Returns the server's representation of the clusterNodePrediction, and an error, if there is any.
func (c *FakeClusterNodePredictions) Create(ctx context.Context, clusterNodePrediction *v1alpha1.ClusterNodePrediction, opts v1.CreateOptions) (result *v1alpha1.ClusterNodePrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusternodepredictionsResource, c.ns, clusterNodePrediction), &v1alpha1.ClusterNodePrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterNodePrediction), err
}

// Update takes the representation of a clusterNodePrediction and updates it. Returns the server's representation of the clusterNodePrediction, and an error, if there is any.
func (c *FakeClusterNodePredictions) Update(ctx context.Context, clusterNodePrediction *v1alpha1.ClusterNodePrediction, opts v1.UpdateOptions) (result *v1alpha1.ClusterNodePrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusternodepredictionsResource, c.ns, clusterNodePrediction), &v1alpha1.ClusterNodePrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterNodePrediction), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterNodePredictions) UpdateStatus(ctx context.Context, clusterNodePrediction *v1alpha1.ClusterNodePrediction, opts v1.UpdateOptions) (*v1alpha1.ClusterNodePrediction, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusternodepredictionsResource, "status", c.ns, clusterNodePrediction), &v1alpha1.ClusterNodePrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterNodePrediction), err
}

// Delete takes name of the clusterNodePrediction and deletes it. Returns an error if one occurs.
func (c *FakeClusterNodePredictions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusternodepredictionsResource, c.ns, name), &v1alpha1.ClusterNodePrediction{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterNodePredictions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusternodepredictionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterNodePredictionList{})
	return err
}

// Patch applies the patch and returns the patched clusterNodePrediction.
func (c *FakeClusterNodePredictions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterNodePrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusternodepredictionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterNodePrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterNodePrediction), err
}
