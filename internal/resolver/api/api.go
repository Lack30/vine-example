// Package vine provides a vine rpc resolver which prefixes a namespace
package vine

import (
	"net/http"

	"github.com/lack-io/vine/api/resolver"
)

// default resolver for legacy purposes
// it uses proxy routing to resolve names
// /foo becomes namespace.foo
// /v1/foo becomes namespace.v1.foo
type Resolver struct {
	Options resolver.Options
}

func (r *Resolver) Resolve(req *http.Request) (*resolver.Endpoint, error) {
	var name, method string

	switch r.Options.Handler {
	// internal handlers
	case "meta", "api", "rpc", "vine":
		name, method = apiRoute(req.URL.Path)
	default:
		method = req.Method
		name = proxyRoute(req.URL.Path)
	}

	ns := r.Options.Namespace(req)
	return &resolver.Endpoint{
		Name:   ns + "." + name,
		Method: method,
	}, nil
}

func (r *Resolver) String() string {
	return "vine"
}

// NewResolver creates a new vine resolver
func NewResolver(opts ...resolver.Option) resolver.Resolver {
	return &Resolver{
		Options: resolver.NewOptions(opts...),
	}
}
