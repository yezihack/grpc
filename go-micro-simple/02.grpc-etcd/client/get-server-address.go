package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
	"time"
)
// 获取服务端注册的 IP:PORT
func GetServerAddress(reg registry.Registry, serviceName string) (string, error) {
	// 获取 Etcd 注册的服务名称。也就是服务端的名称。
	// 在 etcd 中可以使用 etcdctl get / --prefix --keys-only 查看所有的key
	server, err :=reg.GetService(serviceName)
	if err != nil {
		return "", err
	}
	// 随机获取服务端的地址。
	//next := selector.Random(server) // 随机算法
	next := selector.RoundRobin(server)// 轮询算法
	node, err := next()
	if err != nil {
		return "", err
	}
	return node.Address, nil
}

func main() {
	// 获取服务端的IP:PORT
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 创建一个 etcd 对象，并连接etcd
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.61.100:2379"),
	)
	for {
		// 获取 Etcd 注册的服务名称。也就是服务端的名称。
		addr, err := GetServerAddress(etcdReg, "micro.api")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(addr)
		time.Sleep(time.Second)
	}
}
