package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/v2/client/selector"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	ServiceName = "go.micro.get.etcd.node"
)

// 轮询获取 Etcd 里的服务端地址 IP:PORT
func GetEtcdNodeByRoundRobin(reg registry.Registry) (string, error) {
	getService, err := reg.GetService(ServiceName)
	if err != nil {
		return "", err
	}
	next := selector.RoundRobin(getService)
	node, err := next()
	if err != nil {
		return "", err
	}
	return node.Address, nil
}

// 随机获取 Etcd 里的服务端地址 IP:PORT
func GetEtcdNodeByRandom(reg registry.Registry) (string, error) {
	getService, err := reg.GetService(ServiceName) // 获取 register.go 里的 web 服务名称
	if err != nil {
		return "", err
	}
	next := selector.Random(getService) //选择操作，使用随机，还支持轮询。
	node, err := next()
	if err != nil {
		return "", err
	}
	return node.Address, nil
}

func main() {
	// 连接 etcd 服务
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.61.100:2379"),
	)
	go func() {
		for i := 0; i < 10; i++ {
			address, err := GetEtcdNodeByRandom(etcdReg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("random", address)
		}
		time.Sleep(time.Millisecond * 50)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			address, err := GetEtcdNodeByRoundRobin(etcdReg)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("roundBobin", address)
		}
		time.Sleep(time.Millisecond * 50)
	}()
	time.Sleep(time.Second)
}
