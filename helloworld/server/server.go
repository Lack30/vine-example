package main

import (
	"context"
	"fmt"

	"github.com/lack-io/vine/proto/errors"
	"github.com/lack-io/vine/service"

	"github.com/lack-io/vine-example/goproto/api"
	pb "github.com/lack-io/vine-example/helloworld/proto"
)

type HelloWorld struct {
}

func (t *HelloWorld) Call(ctx context.Context, req *pb.HelloWorldRequest, rsp *pb.HelloWorldResponse) error {
	fmt.Println("called")
	if req.Name == "" {
		return errors.New("go.vine.client", "231", 400).
			WithChild(10001, "stack 111")
	}
	rsp.Reply = "hello " + req.Name + " age " + fmt.Sprintf("%d", req.Age)
	return nil
}

func (t *HelloWorld) MulPath(ctx context.Context, req *pb.MulPathRequest, rsp *pb.MulPathResponse) error {

	rsp.Data = []*api.App{&api.App{
		Name: "app1",
		Type: 1,
	}}
	return nil
}

func main() {

	srv := service.NewService(
		service.Name("go.vine.helloworld"),
		service.Address(":59090"),
	)

	srv.Init()

	fmt.Println(srv.Name())
	pb.RegisterHelloworldHandler(srv.Server(), new(HelloWorld))
	srv.Run()
}
