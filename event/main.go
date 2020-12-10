package main

import (
	"context"

	"github.com/lack-io/vine"
	proto "github.com/lack-io/vine/api/proto"
	"github.com/lack-io/vine/log"
)

// All methods of Event will be executed when a message is received
type Event struct{}

// Method can be of any name
func (e *Event) Process(ctx context.Context, event *proto.Event) error {
	log.Infof("Received event %+v\n", event)
	// do something with event
	return nil
}

func main() {
	service := vine.NewService(
		vine.Name("user"),
	)
	service.Init()

	// register subscriber
	vine.RegisterSubscriber("go.vine.evt.user", service.Server(), new(Event))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}