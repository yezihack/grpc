// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api.proto

package proto_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for ApiService service

func NewApiServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ApiService service

type ApiService interface {
	GetName(ctx context.Context, in *GetNameRequest, opts ...client.CallOption) (*GetNameResponse, error)
}

type apiService struct {
	c    client.Client
	name string
}

func NewApiService(name string, c client.Client) ApiService {
	return &apiService{
		c:    c,
		name: name,
	}
}

func (c *apiService) GetName(ctx context.Context, in *GetNameRequest, opts ...client.CallOption) (*GetNameResponse, error) {
	req := c.c.NewRequest(c.name, "ApiService.GetName", in)
	out := new(GetNameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApiService service

type ApiServiceHandler interface {
	GetName(context.Context, *GetNameRequest, *GetNameResponse) error
}

func RegisterApiServiceHandler(s server.Server, hdlr ApiServiceHandler, opts ...server.HandlerOption) error {
	type apiService interface {
		GetName(ctx context.Context, in *GetNameRequest, out *GetNameResponse) error
	}
	type ApiService struct {
		apiService
	}
	h := &apiServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ApiService{h}, opts...))
}

type apiServiceHandler struct {
	ApiServiceHandler
}

func (h *apiServiceHandler) GetName(ctx context.Context, in *GetNameRequest, out *GetNameResponse) error {
	return h.ApiServiceHandler.GetName(ctx, in, out)
}
