// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/clusterpedia-io/api/policy/v1alpha1"
	scheme "xxxxx/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// PediaClusterLifecyclesGetter has a method to return a PediaClusterLifecycleInterface.
// A group's client should implement this interface.
type PediaClusterLifecyclesGetter interface {
	PediaClusterLifecycles() PediaClusterLifecycleInterface
}

// PediaClusterLifecycleInterface has methods to work with PediaClusterLifecycle resources.
type PediaClusterLifecycleInterface interface {
	Create(ctx context.Context, pediaClusterLifecycle *v1alpha1.PediaClusterLifecycle, opts v1.CreateOptions) (*v1alpha1.PediaClusterLifecycle, error)
	Update(ctx context.Context, pediaClusterLifecycle *v1alpha1.PediaClusterLifecycle, opts v1.UpdateOptions) (*v1alpha1.PediaClusterLifecycle, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, pediaClusterLifecycle *v1alpha1.PediaClusterLifecycle, opts v1.UpdateOptions) (*v1alpha1.PediaClusterLifecycle, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PediaClusterLifecycle, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PediaClusterLifecycleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PediaClusterLifecycle, err error)
	PediaClusterLifecycleExpansion
}

// pediaClusterLifecycles implements PediaClusterLifecycleInterface
type pediaClusterLifecycles struct {
	*gentype.ClientWithList[*v1alpha1.PediaClusterLifecycle, *v1alpha1.PediaClusterLifecycleList]
}

// newPediaClusterLifecycles returns a PediaClusterLifecycles
func newPediaClusterLifecycles(c *PolicyV1alpha1Client) *pediaClusterLifecycles {
	return &pediaClusterLifecycles{
		gentype.NewClientWithList[*v1alpha1.PediaClusterLifecycle, *v1alpha1.PediaClusterLifecycleList](
			"pediaclusterlifecycles",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.PediaClusterLifecycle { return &v1alpha1.PediaClusterLifecycle{} },
			func() *v1alpha1.PediaClusterLifecycleList { return &v1alpha1.PediaClusterLifecycleList{} }),
	}
}