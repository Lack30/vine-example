package main

import (
	"context"
	"fmt"

	"github.com/lack-io/vine"

	pb "github.com/lack-io/vine-example/helloworld/proto"
)

type HelloWorld struct {
}

func (t *HelloWorld) Call(ctx context.Context, req *pb.HelloWorldRequest, rsp *pb.HelloWorldResponse) error {
	fmt.Println("get echo request")
	rsp.Reply = "hello " + req.Name
	return nil
}

func main() {
	service := vine.NewService(
		vine.Name("go.vine.helloworld"),
	)

	service.Init()

	pb.RegisterHelloworldHandler(service.Server(), new(HelloWorld))

	service.Run()
}
