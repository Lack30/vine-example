package main

import (
	"context"

	"github.com/lack-io/vine"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello " + *name
	return nil
}

func main() {
	// create new service
	service := vine.NewService(
		vine.Name("greeter"),
	)

	// initialise command line
	service.Init()

	// set the handler
	vine.RegisterHandler(service.Server(), new(Greeter))

	// run service
	service.Run()
}
