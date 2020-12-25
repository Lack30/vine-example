package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/lack-io/vine/proto/api"
	"github.com/lack-io/vine/proto/errors"
	"github.com/lack-io/vine/service"

	proto "github.com/lack-io/vine-example/api/api/proto"
)

type Example struct{}

type Foo struct{}

// Example.Call is a method which will be served by http request /example/call
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /example/call goes to go.vine.api.example Example.Call
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Example.Call request")

	// parse values from the get request
	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.vine.api.example", "no content")
	}

	// set response status
	rsp.StatusCode = 200

	// respond with some json
	b, _ := json.Marshal(map[string]string{
		"message": "got your request " + strings.Join(name.Values, " "),
	})

	// set json body
	rsp.Body = string(b)

	return nil
}

// Foo.Bar is a method which will be served by http request /example/foo/bar
// Because Foo is not the same as the service name it is mapped beyond /example/
func (f *Foo) Bar(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Foo.Bar request")

	fmt.Println("header = ", req.Header)
	fmt.Println("body = ", req.Body)
	defer fmt.Println("rsp = ", rsp)
	// check method
	if req.Method != "POST" {
		return errors.BadRequest("go.vine.api.example", "require post")
	}

	// let's make sure we get json
	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.vine.api.example", "need content-type")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.vine.api.example", "expect application/json")
	}

	// parse body
	var body map[string]interface{}
	json.Unmarshal([]byte(req.Body), &body)

	// do something with parsed body
	data, _ := json.Marshal(body)
	rsp.Body = string(data)

	return nil
}

func main() {
	srv := service.NewService(
		service.Name("go.vine.api.example"),
	)

	srv.Init()

	// register example handler
	proto.RegisterExampleHandler(srv.Server(), new(Example))

	// register foo handler
	proto.RegisterFooHandler(srv.Server(), new(Foo))

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
