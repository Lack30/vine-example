package main

import (
	"context"
	"fmt"
	"log"
	"time"

	vine "github.com/lack-io/vine/service"

	proto "github.com/lack-io/vine-example/service/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}


// Setup and the client
func runClient(service vine.Function) {
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
	function := vine.NewFunction(
		vine.Name("greeter"),
		vine.Version("latest"),
	)

	function.Init()

	function.Handle(new(Greeter))

	go func() {
		<-time.After(time.Second * 2)
		//runClient(function)
		function.Done()
	}()

	if err := function.Run(); err != nil {
		log.Fatal(err)
	}
}
