module github.com/lack-io/vine-example

go 1.15

require (
	github.com/gogo/googleapis v1.4.0
	github.com/gogo/protobuf v1.3.1
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lack-io/cli v1.0.2
	github.com/lack-io/plugins v0.1.9
	github.com/lack-io/vine v0.2.2
	github.com/nats-io/nats-server/v2 v2.1.9 // indirect
	google.golang.org/grpc v1.34.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
