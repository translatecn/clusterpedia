// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "xxxxx/api/cluster/v1alpha2"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ClusterSyncResourcesLister helps list ClusterSyncResources.
// All objects returned here must be treated as read-only.
type ClusterSyncResourcesLister interface {
	// List lists all ClusterSyncResources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.ClusterSyncResources, err error)
	// Get retrieves the ClusterSyncResources from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha2.ClusterSyncResources, error)
	ClusterSyncResourcesListerExpansion
}

// clusterSyncResourcesLister implements the ClusterSyncResourcesLister interface.
type clusterSyncResourcesLister struct {
	listers.ResourceIndexer[*v1alpha2.ClusterSyncResources]
}

// NewClusterSyncResourcesLister returns a new ClusterSyncResourcesLister.
func NewClusterSyncResourcesLister(indexer cache.Indexer) ClusterSyncResourcesLister {
	return &clusterSyncResourcesLister{listers.New[*v1alpha2.ClusterSyncResources](indexer, v1alpha2.Resource("clustersyncresources"))}
}
