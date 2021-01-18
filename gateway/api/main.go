package main

import (
	"net/http"

	"github.com/gorilla/mux"
	rrvine "github.com/lack-io/vine/client/resolver/api"
	"github.com/lack-io/vine/service"
	ahandler "github.com/lack-io/vine/service/api/handler"
	arpc "github.com/lack-io/vine/service/api/handler/rpc"
	"github.com/lack-io/vine/service/api/resolver"
	"github.com/lack-io/vine/service/api/router"
	regRouter "github.com/lack-io/vine/service/api/router/registry"
	httpapi "github.com/lack-io/vine/service/api/server/http"
	log "github.com/lack-io/vine/service/logger"
)

func main() {
	Handler := "rpc"

	// apiNamespace has the format: "go.vine.api"
	apiNamespace := "go.vine.api"
	APIPath := "/"
	Address := ":8080"

	srv := service.NewService(service.Name("go.vine.api"))

	// create the router
	r := mux.NewRouter()

	// return version and list of services
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}

		w.Write([]byte(""))
	})

	// resolver options
	ropts := []resolver.Option{
		resolver.WithHandler(Handler),
	}

	// default resolver
	rr := rrvine.NewResolver(ropts...)

	log.Infof("Registering API RPC Handler at %s", APIPath)
	rt := regRouter.NewRouter(
		router.WithHandler(arpc.Handler),
		router.WithResolver(rr),
		router.WithRegistry(srv.Options().Registry),
	)
	rp := arpc.NewHandler(
		ahandler.WithNamespace(apiNamespace),
		ahandler.WithRouter(rt),
		ahandler.WithClient(srv.Client()),
	)
	r.PathPrefix(APIPath).Handler(rp)

	api := httpapi.NewServer(Address)

	srv.Init()
	api.Init()
	api.Handle("/", r)

	// Start API
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

	// Stop API
	if err := api.Stop(); err != nil {
		log.Fatal(err)
	}
}
