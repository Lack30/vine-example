package main

import (
	"context"
	"fmt"

	"github.com/lack-io/vine/service"

	example "github.com/lack-io/vine-example/server/proto/example"
)

// publishes a message
func pub(i int, p service.Event) {
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
	srv := service.NewService()

	service.Cmd(nil)

	srv.Init()

	p := service.NewEvent("example", srv.Client())

	fmt.Println("\n--- Publisher example ---")

	for i := 0; i < 10; i++ {
		pub(i, p)
	}
}