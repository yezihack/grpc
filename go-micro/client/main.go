package main

import (
	"context"
	"fmt"
	hello "go-micro-simple/proto"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:2379",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("hello"),
	)
	service.Init()
	helleHandler := hello.NewHelloService("hello", service.Client())
	resp, err := helleHandler.Hello(context.TODO(), &hello.HelloRequest{
		Name:    "sgfoot",
		Content: "go micro 微服务框架",
	})
	if err != nil {
		log.Fatal(err)
	}
	result := resp.GetResult()
	fmt.Println("result", result)
}
