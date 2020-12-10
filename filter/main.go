package main

import (
	"context"
	"fmt"

	"github.com/lack-io/vine"

	"github.com/lack-io/vine-example/filter/version"
	proto "github.com/lack-io/vine-example/service/proto"
)

func main() {
	service := vine.NewService()
	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(
		// provide a context
		context.TODO(),
		// provide the request
		&proto.Request{Name: "John"},
		// set the filter
		version.Filter("latest"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
