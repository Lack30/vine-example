package main

import (
	"fmt"
	"time"

	"github.com/lack-io/vine/service/broker"
	log "github.com/lack-io/vine/service/logger"
)

func main() {
	topic := "go.vine.topic.foo"

	b := broker.NewBroker()

	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	go func() {
		// receive message from broker
		b.Subscribe(topic, func(p broker.Event) error {
			fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
			return nil
		})
	}()

	go func() {
		<-time.After(time.Second * 1)
		// publish message to broker
		b.Publish(topic, &broker.Message{Header: map[string]string{"a": "b"}, Body: []byte("hello world")})
	}()

	time.Sleep(time.Second * 2)

	b.Disconnect()
}
