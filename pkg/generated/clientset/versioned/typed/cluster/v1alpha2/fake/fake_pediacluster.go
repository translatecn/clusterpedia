// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha2 "xxxxx/api/cluster/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePediaClusters implements PediaClusterInterface
type FakePediaClusters struct {
	Fake *FakeClusterV1alpha2
}

var pediaclustersResource = v1alpha2.SchemeGroupVersion.WithResource("pediaclusters")

var pediaclustersKind = v1alpha2.SchemeGroupVersion.WithKind("PediaCluster")

// Get takes name of the pediaCluster, and returns the corresponding pediaCluster object, and an error if there is any.
func (c *FakePediaClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.PediaCluster, err error) {
	emptyResult := &v1alpha2.PediaCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(pediaclustersResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.PediaCluster), err
}

// List takes label and field selectors, and returns the list of PediaClusters that match those selectors.
func (c *FakePediaClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.PediaClusterList, err error) {
	emptyResult := &v1alpha2.PediaClusterList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(pediaclustersResource, pediaclustersKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.PediaClusterList{ListMeta: obj.(*v1alpha2.PediaClusterList).ListMeta}
	for _, item := range obj.(*v1alpha2.PediaClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pediaClusters.
func (c *FakePediaClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(pediaclustersResource, opts))
}

// Create takes the representation of a pediaCluster and creates it.  Returns the server's representation of the pediaCluster, and an error, if there is any.
func (c *FakePediaClusters) Create(ctx context.Context, pediaCluster *v1alpha2.PediaCluster, opts v1.CreateOptions) (result *v1alpha2.PediaCluster, err error) {
	emptyResult := &v1alpha2.PediaCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(pediaclustersResource, pediaCluster, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.PediaCluster), err
}

// Update takes the representation of a pediaCluster and updates it. Returns the server's representation of the pediaCluster, and an error, if there is any.
func (c *FakePediaClusters) Update(ctx context.Context, pediaCluster *v1alpha2.PediaCluster, opts v1.UpdateOptions) (result *v1alpha2.PediaCluster, err error) {
	emptyResult := &v1alpha2.PediaCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(pediaclustersResource, pediaCluster, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.PediaCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePediaClusters) UpdateStatus(ctx context.Context, pediaCluster *v1alpha2.PediaCluster, opts v1.UpdateOptions) (result *v1alpha2.PediaCluster, err error) {
	emptyResult := &v1alpha2.PediaCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(pediaclustersResource, "status", pediaCluster, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.PediaCluster), err
}

// Delete takes name of the pediaCluster and deletes it. Returns an error if one occurs.
func (c *FakePediaClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(pediaclustersResource, name, opts), &v1alpha2.PediaCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePediaClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(pediaclustersResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.PediaClusterList{})
	return err
}

// Patch applies the patch and returns the patched pediaCluster.
func (c *FakePediaClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.PediaCluster, err error) {
	emptyResult := &v1alpha2.PediaCluster{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(pediaclustersResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.PediaCluster), err
}
