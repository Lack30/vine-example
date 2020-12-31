package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lack-io/vine/service/registry"
)

func main() {
	// create new registry
	r := registry.NewRegistry()

	// initialize registry
	r.Init()

	s := &registry.Service{
		Name:     "go.vine.srv",
		Version:  "1.0.0",
		Metadata: map[string]string{},
		Endpoints: []*registry.Endpoint{{
			Request: &registry.Value{},
		}},
		Nodes: []*registry.Node{&registry.Node{
			Id:      uuid.New().String(),
			Address: "127.0.0.1:1111",
		}},
	}

	// registry service
	r.Register(s)

	// watch registry
	w, err := r.Watch(registry.WatchService("go.vine.srv"))
	if err != nil {
		return
	}
	defer w.Stop()
	go func() {
		for {
			// blocking
			r, err := w.Next()
			if err != nil {
				return
			}

			log.Printf("watch [%v] %v", r.Action, r.Service)
		}
	}()

	time.Sleep(time.Second)
	// get service
	service, err := r.GetService("go.vine.srv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", service)

	// destroy registry
	r.Deregister(s)

	time.Sleep(time.Second)
}
