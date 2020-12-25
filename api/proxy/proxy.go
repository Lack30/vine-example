package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lack-io/vine/proto/errors"
	"github.com/lack-io/vine/service/web"
)

// exampleCall will handle /example/call
func exampleCall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call ")
	r.ParseForm()
	// get name
	name := r.Form.Get("name")

	if len(name) == 0 {
		http.Error(
			w,
			errors.BadRequest("go.vine.api.example", "no content").Error(),
			400,
		)
		return
	}

	// marshal response
	b, _ := json.Marshal(map[string]interface{}{
		"message": "got your message " + name,
	})

	// write response
	w.Write(b)
}

// exampleFooBar will handle /example/foo/bar
func exampleFooBar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(
			w,
			errors.BadRequest("go.vine.api.example", "require post").Error(),
			400,
		)
		return
	}

	if len(r.Header.Get("Content-Type")) == 0 {
		http.Error(
			w,
			errors.BadRequest("go.vine.api.example", "need content-type").Error(),
			400,
		)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(
			w,
			errors.BadRequest("go.vine.api.example", "expect application/json").Error(),
			400,
		)
		return
	}

	w.Write([]byte("hello world"))
	// do something
}

func main() {
	// we're using go-web for convenience since it registers with discovery
	srv := web.NewService(
		web.Name("go.vine.api.example"),
	)

	srv.HandleFunc("/example/call", exampleCall)
	srv.HandleFunc("/example/foo/bar", exampleFooBar)

	if err := srv.Init(); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}