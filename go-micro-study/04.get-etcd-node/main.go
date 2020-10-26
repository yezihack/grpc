package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro/v2/client/selector"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func GetServerAddress(reg etcd.) string {

}

func main() {
	// 连接 etcd 服务
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.61.100:2379"),
	)
	getService, err := etcdReg.GetService("go.micro.get.etcd.node") // 获取 register.go 里的 web 服务名称
	if err != nil {
		log.Fatal(err)
	}
	next := selector.Random(getService) //选择操作，使用随机，还支持轮询。
	//next := selector.RoundRobin(getService)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node.Address)

}
