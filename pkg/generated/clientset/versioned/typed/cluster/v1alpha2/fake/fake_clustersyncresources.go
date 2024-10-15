// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha2 "github.com/clusterpedia-io/api/cluster/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterSyncResources implements ClusterSyncResourcesInterface
type FakeClusterSyncResources struct {
	Fake *FakeClusterV1alpha2
}

var clustersyncresourcesResource = v1alpha2.SchemeGroupVersion.WithResource("clustersyncresources")

var clustersyncresourcesKind = v1alpha2.SchemeGroupVersion.WithKind("ClusterSyncResources")

// Get takes name of the clusterSyncResources, and returns the corresponding clusterSyncResources object, and an error if there is any.
func (c *FakeClusterSyncResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ClusterSyncResources, err error) {
	emptyResult := &v1alpha2.ClusterSyncResources{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(clustersyncresourcesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.ClusterSyncResources), err
}

// List takes label and field selectors, and returns the list of ClusterSyncResources that match those selectors.
func (c *FakeClusterSyncResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ClusterSyncResourcesList, err error) {
	emptyResult := &v1alpha2.ClusterSyncResourcesList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(clustersyncresourcesResource, clustersyncresourcesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ClusterSyncResourcesList{ListMeta: obj.(*v1alpha2.ClusterSyncResourcesList).ListMeta}
	for _, item := range obj.(*v1alpha2.ClusterSyncResourcesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterSyncResources.
func (c *FakeClusterSyncResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(clustersyncresourcesResource, opts))
}

// Create takes the representation of a clusterSyncResources and creates it.  Returns the server's representation of the clusterSyncResources, and an error, if there is any.
func (c *FakeClusterSyncResources) Create(ctx context.Context, clusterSyncResources *v1alpha2.ClusterSyncResources, opts v1.CreateOptions) (result *v1alpha2.ClusterSyncResources, err error) {
	emptyResult := &v1alpha2.ClusterSyncResources{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(clustersyncresourcesResource, clusterSyncResources, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.ClusterSyncResources), err
}

// Update takes the representation of a clusterSyncResources and updates it. Returns the server's representation of the clusterSyncResources, and an error, if there is any.
func (c *FakeClusterSyncResources) Update(ctx context.Context, clusterSyncResources *v1alpha2.ClusterSyncResources, opts v1.UpdateOptions) (result *v1alpha2.ClusterSyncResources, err error) {
	emptyResult := &v1alpha2.ClusterSyncResources{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(clustersyncresourcesResource, clusterSyncResources, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.ClusterSyncResources), err
}

// Delete takes name of the clusterSyncResources and deletes it. Returns an error if one occurs.
func (c *FakeClusterSyncResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clustersyncresourcesResource, name, opts), &v1alpha2.ClusterSyncResources{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterSyncResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(clustersyncresourcesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.ClusterSyncResourcesList{})
	return err
}

// Patch applies the patch and returns the patched clusterSyncResources.
func (c *FakeClusterSyncResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ClusterSyncResources, err error) {
	emptyResult := &v1alpha2.ClusterSyncResources{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clustersyncresourcesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.ClusterSyncResources), err
}