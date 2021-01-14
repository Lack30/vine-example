package main

import (
	"log"
	"net/http"

	"github.com/lack-io/vine/service/web"
)

func main() {
	service := web.NewService(
		web.Name("go.vine.web.helloworld"),
	)

	//service.Handle("/", http.RedirectHandler("/index.html", 301))
	service.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "index.html")
	})
	//service.HandleFunc("/", helloWorldHandler)
	service.Handle("/assert/", http.StripPrefix("/assert/", http.FileServer(http.Dir("./static"))))

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
