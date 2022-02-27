// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/ping.proto

package ping

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Ping service

func NewPingEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Ping service

type PingService interface {
	Ip(ctx context.Context, in *IpRequest, opts ...client.CallOption) (*IpResponse, error)
	Url(ctx context.Context, in *UrlRequest, opts ...client.CallOption) (*UrlResponse, error)
	Tcp(ctx context.Context, in *TcpRequest, opts ...client.CallOption) (*TcpResponse, error)
}

type pingService struct {
	c    client.Client
	name string
}

func NewPingService(name string, c client.Client) PingService {
	return &pingService{
		c:    c,
		name: name,
	}
}

func (c *pingService) Ip(ctx context.Context, in *IpRequest, opts ...client.CallOption) (*IpResponse, error) {
	req := c.c.NewRequest(c.name, "Ping.Ip", in)
	out := new(IpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingService) Url(ctx context.Context, in *UrlRequest, opts ...client.CallOption) (*UrlResponse, error) {
	req := c.c.NewRequest(c.name, "Ping.Url", in)
	out := new(UrlResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingService) Tcp(ctx context.Context, in *TcpRequest, opts ...client.CallOption) (*TcpResponse, error) {
	req := c.c.NewRequest(c.name, "Ping.Tcp", in)
	out := new(TcpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ping service

type PingHandler interface {
	Ip(context.Context, *IpRequest, *IpResponse) error
	Url(context.Context, *UrlRequest, *UrlResponse) error
	Tcp(context.Context, *TcpRequest, *TcpResponse) error
}

func RegisterPingHandler(s server.Server, hdlr PingHandler, opts ...server.HandlerOption) error {
	type ping interface {
		Ip(ctx context.Context, in *IpRequest, out *IpResponse) error
		Url(ctx context.Context, in *UrlRequest, out *UrlResponse) error
		Tcp(ctx context.Context, in *TcpRequest, out *TcpResponse) error
	}
	type Ping struct {
		ping
	}
	h := &pingHandler{hdlr}
	return s.Handle(s.NewHandler(&Ping{h}, opts...))
}

type pingHandler struct {
	PingHandler
}

func (h *pingHandler) Ip(ctx context.Context, in *IpRequest, out *IpResponse) error {
	return h.PingHandler.Ip(ctx, in, out)
}

func (h *pingHandler) Url(ctx context.Context, in *UrlRequest, out *UrlResponse) error {
	return h.PingHandler.Url(ctx, in, out)
}

func (h *pingHandler) Tcp(ctx context.Context, in *TcpRequest, out *TcpResponse) error {
	return h.PingHandler.Tcp(ctx, in, out)
}