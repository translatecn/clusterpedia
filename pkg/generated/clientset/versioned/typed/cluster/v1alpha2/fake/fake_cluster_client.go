// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "xxxxx/pkg/generated/clientset/versioned/typed/cluster/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeClusterV1alpha2 struct {
	*testing.Fake
}

func (c *FakeClusterV1alpha2) ClusterSyncResources() v1alpha2.ClusterSyncResourcesInterface {
	return &FakeClusterSyncResources{c}
}

func (c *FakeClusterV1alpha2) PediaClusters() v1alpha2.PediaClusterInterface {
	return &FakePediaClusters{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeClusterV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
