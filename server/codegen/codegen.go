package main

import (
	"context"
	"log"

	"github.com/lack-io/vine/service/config/cmd"
	"github.com/lack-io/vine/service/server"

	example "github.com/lack-io/vine-example/server/proto/example"
	"github.com/lack-io/vine-example/server/subscriber"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	log.Print("Received Example.Call request")
	rsp.Msg = server.DefaultOptions().Id + ": Hello " + req.Name
	return nil
}

func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
	log.Printf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Printf("Responding: %d", i)
		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func main() {
	// optionally setup command line usage
	cmd.Init()

	// Initialise Server
	server.Init(
		server.Name("go.vine.srv.example"),
	)

	// Register Subscribers
	server.Subscribe(
		server.NewSubscriber(
			"topic.go.vine.srv.example",
			new(subscriber.Example),
		),
	)

	server.Subscribe(
		server.NewSubscriber(
			"topic.go.vine.srv.example",
			subscriber.Handler,
		),
	)

	// Register Handler
	example.RegisterExampleHandler(
		server.DefaultServer, new(Example),
	)

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
