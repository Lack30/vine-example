package version

import (
	"github.com/lack-io/vine/client"
	"github.com/lack-io/vine/client/selector"
	"github.com/lack-io/vine/registry"
)

// Filter will filter the version of the service
func Filter(v string) client.CallOption {
	filter := func(services []*registry.Service) []*registry.Service {
		var filtered []*registry.Service

		for _, service := range services {
			if service.Version == v {
				filtered = append(filtered, service)
			}
		}

		return filtered
	}

	return client.WithSelectOption(selector.WithFilter(filter))
}

