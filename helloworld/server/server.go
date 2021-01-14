package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lack-io/vine/proto/errors"
	"github.com/lack-io/vine/service"
	"github.com/lack-io/vine/service/server"
	"github.com/lack-io/vine/util/context/metadata"

	"github.com/lack-io/vine-example/goproto/api"
	pb "github.com/lack-io/vine-example/helloworld/proto"
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[Log Wrapper] Before serving request method: %v", req.Endpoint())

		m, _ := metadata.FromContext(ctx)
		a, _ := m.Get("Authorization")
		fmt.Println(a)

		err := fn(ctx, req, rsp)
		log.Printf("[Log Wrapper] After serving request")
		return err
	}
}

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
		service.WrapHandler(logWrapper),
	)

	srv.Init()

	fmt.Println(srv.Name())
	pb.RegisterHelloworldHandler(srv.Server(), new(HelloWorld))
	srv.Run()
}
