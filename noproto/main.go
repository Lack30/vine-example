package main

import (
	"context"

	"github.com/lack-io/vine/service"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello " + *name
	return nil
}

func main() {
	// create new service
	server := service.NewService(
		service.Name("greeter"),
	)

	// initialise command line
	server.Init()

	// set the handler
	service.RegisterHandler(server.Server(), new(Greeter))

	// run service
	server.Run()
}
