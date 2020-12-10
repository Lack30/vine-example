package subscriber

import (
	"context"
	"log"

	example "github.com/lack-io/vine-example/server/proto/example"
)
type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Print("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *example.Message) error {
	log.Print("Function Received message: ", msg.Say)
	return nil
}