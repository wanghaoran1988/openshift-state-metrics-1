package hostsubnet

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/api/legacyscheme"

	networkapi "github.com/openshift/origin/pkg/network/apis/network"
	"github.com/openshift/origin/pkg/network/apis/network/validation"
)

// sdnStrategy implements behavior for HostSubnets
type sdnStrategy struct {
	runtime.ObjectTyper
}

// Strategy is the default logic that applies when creating and updating HostSubnet
// objects via the REST API.
var Strategy = sdnStrategy{legacyscheme.Scheme}

var _ rest.GarbageCollectionDeleteStrategy = sdnStrategy{}

func (sdnStrategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	return rest.Unsupported
}

func (sdnStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {}

// Canonicalize normalizes the object after validation.
func (sdnStrategy) Canonicalize(obj runtime.Object) {
}

// NamespaceScoped is false for sdns
func (sdnStrategy) NamespaceScoped() bool {
	return false
}

func (sdnStrategy) GenerateName(base string) string {
	return base
}

func (sdnStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

// Validate validates a new sdn
func (sdnStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return validation.ValidateHostSubnet(obj.(*networkapi.HostSubnet))
}

// AllowCreateOnUpdate is false for sdns
func (sdnStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (sdnStrategy) AllowUnconditionalUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for a HostSubnet
func (sdnStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return validation.ValidateHostSubnetUpdate(obj.(*networkapi.HostSubnet), old.(*networkapi.HostSubnet))
}
