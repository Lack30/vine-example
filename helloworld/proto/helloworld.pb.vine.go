// Code generated by proto-gen-vine. DO NOT EDIT.
// source: github.com/lack-io/vine-example/helloworld/proto/helloworld.proto

package testdata

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/lack-io/vine-example/goproto/api"
	math "math"
)

import (
	context "context"
	api "github.com/lack-io/vine/service/api"
	client "github.com/lack-io/vine/service/client"
	server "github.com/lack-io/vine/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Helloworld service
func NewHelloworldEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "Helloworld.Call",
			Path:    []string{"/api/v1/call"},
			Method:  []string{"GET"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for Helloworld service
type HelloworldService interface {
	// +gen:get=/api/v1/call
	// +gen:body=*
	Call(ctx context.Context, in *HelloWorldRequest, opts ...client.CallOption) (*HelloWorldResponse, error)
	MulPath(ctx context.Context, in *MulPathRequest, opts ...client.CallOption) (*MulPathResponse, error)
}

type helloworldService struct {
	c    client.Client
	name string
}

func NewHelloworldService(name string, c client.Client) HelloworldService {
	return &helloworldService{
		c:    c,
		name: name,
	}
}

func (c *helloworldService) Call(ctx context.Context, in *HelloWorldRequest, opts ...client.CallOption) (*HelloWorldResponse, error) {
	req := c.c.NewRequest(c.name, "Helloworld.Call", in)
	out := new(HelloWorldResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldService) MulPath(ctx context.Context, in *MulPathRequest, opts ...client.CallOption) (*MulPathResponse, error) {
	req := c.c.NewRequest(c.name, "Helloworld.MulPath", in)
	out := new(MulPathResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Helloworld service
type HelloworldHandler interface {
	// +gen:get=/api/v1/call
	// +gen:body=*
	Call(context.Context, *HelloWorldRequest, *HelloWorldResponse) error
	MulPath(context.Context, *MulPathRequest, *MulPathResponse) error
}

func RegisterHelloworldHandler(s server.Server, hdlr HelloworldHandler, opts ...server.HandlerOption) error {
	type helloworldImpl interface {
		Call(ctx context.Context, in *HelloWorldRequest, out *HelloWorldResponse) error
		MulPath(ctx context.Context, in *MulPathRequest, out *MulPathResponse) error
	}
	type Helloworld struct {
		helloworldImpl
	}
	h := &helloworldHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Helloworld.Call",
		Path:    []string{"/api/v1/call"},
		Method:  []string{"GET"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Helloworld{h}, opts...))
}

type helloworldHandler struct {
	HelloworldHandler
}

func (h *helloworldHandler) Call(ctx context.Context, in *HelloWorldRequest, out *HelloWorldResponse) error {
	return h.HelloworldHandler.Call(ctx, in, out)
}

func (h *helloworldHandler) MulPath(ctx context.Context, in *MulPathRequest, out *MulPathResponse) error {
	return h.HelloworldHandler.MulPath(ctx, in, out)
}
