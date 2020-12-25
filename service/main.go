package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lack-io/cli"
	"github.com/lack-io/vine/service"

	proto "github.com/lack-io/vine-example/service/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

// Setup and the client
func runClient(service service.Service) {
	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Greeting)
}

func main() {
	// Create a new service. Optionally include some options here.
	srv := service.NewService(
		service.Name("greeter"),
		service.Version("latest"),
		service.Metadata(map[string]string{
			"type": "helloworld",
		}),

		// Setup some flags. Specify --run-client to run the client

		// Add runtime flags
		// We could do this below too
		service.Flags(&cli.BoolFlag{
			Name:  "run-client",
			Usage: "Launch the client",
		}),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	srv.Init(
		// Add runtime action
		// We could actually do this above
		service.Action(func(c *cli.Context) error {
			if c.Bool("run-client") {
				runClient(srv)
				os.Exit(0)
			}
			return nil
		}),
	)

	// By default we'll run the server unless the flags catch us

	// Setup the server

	// Register handler
	proto.RegisterGreeterHandler(srv.Server(), new(Greeter))

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
