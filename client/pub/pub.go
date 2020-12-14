package main

import (
	"context"
	"fmt"

	"github.com/lack-io/vine"

	example "github.com/lack-io/vine-example/server/proto/example"
)

// publishes a message
func pub(i int, p vine.Event) {
	msg := &example.Message{
		Say: fmt.Sprintf("This is an async message %d", i),
	}

	if err := p.Publish(context.TODO(), msg); err != nil {
		fmt.Println("pub err: ", err)
		return
	}

	fmt.Printf("Published %d: %v\n", i, msg)
}

func main() {
	service := vine.NewService()

	vine.Cmd(nil)

	service.Init()

	p := vine.NewEvent("example", service.Client())

	fmt.Println("\n--- Publisher example ---")

	for i := 0; i < 10; i++ {
		pub(i, p)
	}
}