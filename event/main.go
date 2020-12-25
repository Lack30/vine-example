package main

import (
	"context"

	pb "github.com/lack-io/vine/proto/api"
	"github.com/lack-io/vine/service"
	log "github.com/lack-io/vine/service/logger"
)

// All methods of Event will be executed when a message is received
type Event struct{}

// Method can be of any name
func (e *Event) Process(ctx context.Context, event *pb.Event) error {
	log.Infof("Received event %+v\n", event)
	// do something with event
	return nil
}

func main() {
	srv := service.NewService(
		service.Name("user"),
	)
	srv.Init()

	// register subscriber
	service.RegisterSubscriber("go.vine.evt.user", srv.Server(), new(Event))

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
