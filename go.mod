module github.com/lack-io/vine-example

go 1.15

require (
	github.com/gogo/googleapis v1.4.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/uuid v1.1.5
	github.com/gorilla/mux v1.8.0
	github.com/lack-io/cli v1.0.2
	github.com/lack-io/plugins v0.2.1
	github.com/lack-io/vine v0.9.4
	google.golang.org/grpc v1.35.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
