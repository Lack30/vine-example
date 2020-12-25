package main

import (
	"context"
	"log"
	"time"

	"github.com/lack-io/vine/proto/errors"
	"github.com/lack-io/vine/service"

	proto "github.com/lack-io/vine-example/api/rpc/proto"
)

type Example struct{}

type Foo struct{}

// Example.Call is a method which will be served by http request /example/call
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /example/call goes to go.vine.api.example Example.Call
func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Print("Received Example.Call request")

	if len(req.Name) == 0 {
		return errors.BadRequest("go.vine.api.example", "no content")
	}

	rsp.Message = "got your request " + req.Name
	return nil
}

// Foo.Bar is a method which will be served by http request /example/foo/bar
// Because Foo is not the same as the service name it is mapped beyond /example/
func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Print("Received Foo.Bar request")

	// noop

	return nil
}

func main() {
	srv := service.NewService(
		service.Name("go.vine.api.example"),
	)

	srv.Init()

	// register example handler
	proto.RegisterExampleHandler(srv.Server(), new(Example))

	// register foo handler
	proto.RegisterFooHandler(srv.Server(), new(Foo))

	go func() {
		<-time.After(time.Second * 2)
		client := proto.NewExampleService("go.vine.api.example", srv.Client())
		rsp, err := client.Call(context.TODO(), &proto.CallRequest{Name: ""})
		log.Printf("rsp = %v, err = %v", rsp, err)
	}()

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}