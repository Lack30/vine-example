package main

import (
	"context"
	"log"

	"github.com/lack-io/vine/config/cmd"
	"github.com/lack-io/vine/server"

	"github.com/lack-io/vine-example/server/handler"
	"github.com/lack-io/vine-example/server/subscriber"
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[Log Wrapper] Before serving request method: %v", req.Endpoint())
		err := fn(ctx, req, rsp)
		log.Printf("[Log Wrapper] After serving request")
		return err
	}
}

func logSubWrapper(fn server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, req server.Message) error {
		log.Printf("[Log Sub Wrapper] Before serving publication topic: %v", req.Topic())
		err := fn(ctx, req)
		log.Printf("[Log Sub Wrapper] After serving publication")
		return err
	}
}

func main() {
	// optionally setup command line usage
	cmd.Init()

	md := server.DefaultOptions().Metadata
	md["datacenter"] = "local"

	server.DefaultServer = server.NewServer(
		server.WrapHandler(logWrapper),
		server.WrapSubscriber(logSubWrapper),
		server.Metadata(md),
	)

	// Initialise Server
	server.Init(
		server.Name("go.vine.srv.example"),
	)

	// Register Handlers
	server.Handle(
		server.NewHandler(
			new(handler.Example),
		),
	)

	// Register Subscribers
	if err := server.Subscribe(
		server.NewSubscriber(
			"topic.go.vine.srv.example",
			new(subscriber.Example),
		),
	); err != nil {
		log.Fatal(err)
	}

	if err := server.Subscribe(
		server.NewSubscriber(
			"topic.go.vine.srv.example",
			subscriber.Handler,
		),
	); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}