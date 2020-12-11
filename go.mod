module github.com/lack-io/vine-example

go 1.15

require (
	github.com/gogo/googleapis v1.4.0
	github.com/gogo/protobuf v1.3.1
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lack-io/cli v0.1.1
	github.com/lack-io/plugins v0.1.8
	github.com/lack-io/vine v0.2.1
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb
	google.golang.org/grpc v1.34.0
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
