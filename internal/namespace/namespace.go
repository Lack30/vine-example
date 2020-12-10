package namespace

import (
	"context"

	"github.com/lack-io/vine/util/metadata"
)

const (
	DefaultNamespace = "go.vine"
	// NamespaceKey is used to set/get the namespace from the context
	NamespaceKey = "Vine-Namespace"
)

// FromContext gets the namespace from the context
func FromContext(ctx context.Context) string {
	// get the namespace which is set at ingress by vine web / api / proxy etc. The go-vine auth
	// wrapper will ensure the account making the request has the necessary issuer.
	ns, _ := metadata.Get(ctx, NamespaceKey)
	return ns
}

// ContextWithNamespace sets the namespace in the context
func ContextWithNamespace(ctx context.Context, ns string) context.Context {
	return metadata.Set(ctx, NamespaceKey, ns)
}
