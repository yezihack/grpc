package main

import (
	proto_api "go-micro-simple/02.grpc-etcd/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.61.100:2379"),
	)
	myClient := web.NewService(
		web.Registry(etcdReg),
	)

	service := micro.NewService(
		micro.Name("hello"),
	)
	proto_api.NewApiService("micro.api.proto", service.Client())
}
