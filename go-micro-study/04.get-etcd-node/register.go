package main

import (
	"log"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// 连接 etcd 服务
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.61.100:2379"),
	)
	service := web.NewService(
		web.Name("go.micro.get.etcd.node"), // go-micro 服务名称
		web.Address(":8000"),               // 开外一个对外端口
		web.Registry(etcdReg),              // etcd 注册到 go-micro 服务里
	)
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// 启动时，多启动几个不同的端口。使用 --server_address 是 go-micro 自带的参数
/**
go run register.go --server_address :8000
go run register.go --server_address :8001
go run register.go --server_address :8002
go run register.go --server_address :8003
*/
