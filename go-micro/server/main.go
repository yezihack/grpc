package main

import (
	"context"
	hello "go-micro-simple/proto"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, req *hello.HelloRequest, resp *hello.HelloResponse) error {
	resp = new(hello.HelloResponse)
	resp.Result = req.Name + ":" + req.Content
	return nil
}

func main() {
	// 注册 etcd
	reg := etcd.NewRegistry(func(opt *registry.Options) {
		opt.Addrs = []string{
			"127.0.0.1:2379", // etcd 对外访问的地址
		}
	})

	service := micro.NewService(
		micro.Registry(reg), // 注册 etcd 对象
		micro.Name("hello"), // 注册名称
	)
	service.Init()

	err := hello.RegisterHelloServiceHandler(service.Server(), new(HelloServiceImpl))
	if err != nil {
		log.Fatal(err)
	}
	if err = service.Run(); err != nil {
		log.Fatalln(err)
	}
}
