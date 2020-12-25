package main

import (
	"context"
	"fmt"

	"github.com/lack-io/vine/service"

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
	srv := service.NewService(
		service.Name("go.vine.helloworld"),
		service.Address(":59090"),
	)

	srv.Init()
	pb.RegisterHelloworldHandler(srv.Server(), new(HelloWorld))
	srv.Run()
}
