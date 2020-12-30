package main

import (
	"context"
	"fmt"

	"github.com/lack-io/vine/service"

	pb "github.com/lack-io/vine-example/helloworld/proto"
)

func main() {
	srv := service.NewService(service.Name("go.vine.helloworld"))
	service := pb.NewHelloworldService("go.vine.helloworld", srv.Client())

	rsp, err := service.Call(context.TODO(), &pb.HelloWorldRequest{Name: ""})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Call: %v\n", rsp)
	}

	rsp, err = service.Call(context.TODO(), &pb.HelloWorldRequest{Name: "world"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Call: %v\n", rsp)

	rsp1, err := service.MulPath(context.TODO(), &pb.MulPathRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("MulPath: %v\n", rsp1)
}
