// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

// 定义一个包名

package proto_user

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

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	// 以 rpc 开头，定义未实现的方法名称。 必须要有一个请求 message 和响应 message
	// rpc, returns 都是关键字
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetInfo", in)
	out := new(GetInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	// 以 rpc 开头，定义未实现的方法名称。 必须要有一个请求 message 和响应 message
	// rpc, returns 都是关键字
	GetInfo(context.Context, *GetInfoRequest, *GetInfoResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		GetInfo(ctx context.Context, in *GetInfoRequest, out *GetInfoResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) GetInfo(ctx context.Context, in *GetInfoRequest, out *GetInfoResponse) error {
	return h.UserServiceHandler.GetInfo(ctx, in, out)
}
