package main

import (
	"log"

	proto "github.com/lack-io/vine-example/api/meta/proto"
	"github.com/lack-io/vine"
	"github.com/lack-io/vine/api"
	rapi "github.com/lack-io/vine/api/handler/api"
	"github.com/lack-io/vine/api/handler/rpc"
	"github.com/lack-io/vine/errors"

	"context"
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
	service := vine.NewService(
		vine.Name("go.vine.api.example"),
	)

	service.Init()

	// register example handler
	proto.RegisterExampleHandler(service.Server(), new(Example), api.WithEndpoint(&api.Endpoint{
		// The RPC method
		Name: "Example.Call",
		// The HTTP paths. This can be a POSIX regex
		Path: []string{"/example"},
		// The HTTP Methods for this endpoint
		Method: []string{"POST"},
		// The API handler to use
		Handler: rpc.Handler,
	}))

	// register foo handler
	proto.RegisterFooHandler(service.Server(), new(Foo), api.WithEndpoint(&api.Endpoint{
		// The RPC method
		Name: "Foo.Bar",
		// The HTTP paths. This can be a POSIX regex
		Path: []string{"/foo/bar"},
		// The HTTP Methods for this endpoint
		Method: []string{"POST"},
		// The API handler to use
		Handler: rapi.Handler,
	}))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}