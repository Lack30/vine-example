package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/lack-io/vine/service/client"
	"github.com/lack-io/vine/service/client/selector"
	"github.com/lack-io/vine/service/config/cmd"
	"github.com/lack-io/vine/service/registry"
	"github.com/lack-io/vine/util/context/metadata"

	example "github.com/lack-io/vine-example/server/proto/example"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// A Wrapper that creates a Datacenter Selector Option
type dcWrapper struct {
	client.Client
}

func (dc *dcWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)

	filter := func(services []*registry.Service) []*registry.Service {
		for _, service := range services {
			var nodes []*registry.Node
			for _, node := range service.Nodes {
				if node.Metadata["datacenter"] == md["datacenter"] {
					nodes = append(nodes, node)
				}
			}
			service.Nodes = nodes
		}
		return services
	}

	callOptions := append(opts, client.WithSelectOption(
		selector.WithFilter(filter),
	))

	fmt.Printf("[DC Wrapper] filtering for datacenter %s\n", md["datacenter"])
	return dc.Client.Call(ctx, req, rsp, callOptions...)
}

func NewDCWrapper(c client.Client) client.Client {
	return &dcWrapper{c}
}

func call(i int) {
	// Create new request to service go.vine.srv.example, method Example.Call
	req := client.NewRequest("go.vine.srv.example", "Example.Call", &example.Request{
		Name: "John",
	})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"datacenter": "local",
	})

	rsp := &example.Response{}

	// Call service
	if err := client.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp.Msg)
}

func main() {
	cmd.Init()

	client.DefaultClient = client.NewClient(
		client.Wrap(NewDCWrapper),
	)

	fmt.Println("\n--- Call example ---")
	for i := 0; i < 10; i++ {
		call(i)
	}
}
