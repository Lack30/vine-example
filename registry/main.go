package main

import (
	"context"
	"log"
	"time"

	_ "github.com/lack-io/plugins/registry/etcd"
	"github.com/lack-io/vine"

	pb "github.com/lack-io/vine-example/registry/proto"
)

type Echo struct{}

func (e *Echo) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Reply = "ok"
	return nil
}

func main() {
	service := vine.NewService(
		vine.Name("go.vine.srv.example"),
	)

	service.Init()

	pb.RegisterEchoHandler(service.Server(), new(Echo))

	close := make(chan struct{}, 1)

	go func() {
		defer func() {
			close <- struct{}{}
		}()
		<-time.After(time.Second)
		ss := vine.NewService(
			vine.Name("go.vine.srv.example"),
			vine.Registry(service.Options().Registry),
		)
		srv := pb.NewEchoService("go.vine.srv.example", ss.Client())

		rsp, err := srv.Call(context.Background(), &pb.Request{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(rsp)
	}()

	go service.Run()

	select {
	case <-close:
		service.Server().Stop()
	}
}
