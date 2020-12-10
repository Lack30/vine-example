package main

import (
	"context"
	"fmt"

	"github.com/lack-io/vine"
	"github.com/lack-io/vine/client"
)

func main() {
	service := vine.NewService()
	service.Init()
	c := service.Client()

	request := c.NewRequest("greeter", "Greeter.Hello", "john", client.WithContentType("application/json"))
	var response string

	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
