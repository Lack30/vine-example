package testdata

import (
	"github.com/lack-io/vine/service/registry"
	"github.com/lack-io/vine/service/registry/memory"
)

var (
	// mock registry data
	Data = map[string][]*registry.Service{
		"foo": {
			{
				Name:    "foo",
				Version: "1.0.0",
				Nodes: []*registry.Node{
					{
						Id:      "foo-1.0.0-123",
						Address: "localhost:9999",
					},
				},
			},
		},
	}
)


var RR = memory.NewRegistry(memory.Services(Data))

