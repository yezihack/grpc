package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	serviceName = "go.micro.web.gin.etcd.call" // 来自 server/main.go 定义的服务名称。
	etcdAddress = "192.168.61.100:2379"        // etcd 的服务地址
	urlPath     = "/v1/rand"                   // 需要访问的服务端的 web path 地址
)

// 通过服务名获取服务注册的IP:PORT
func GetServerNodeAddress(reg registry.Registry) (string, error) {
	// 第一步：通过服务名称实现服务发现
	getService, err := reg.GetService(serviceName)
	if err != nil {
		return "", err
	}
	// 第二步：通过不同策略(随机/轮询)获取服务地址 IP:PORT
	next := selector.Random(getService)
	node, err := next()
	if err != nil {
		return "", err
	}
	return node.Address, nil
}

// 通过 URL 获取返回的结果
func GetRand(uri string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return 0, err
	}
	client := http.DefaultClient
	rsp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer rsp.Body.Close()
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 0, err
	}
	type foo struct {
		Number int `json:"random-number"`
	}
	f := foo{}
	if err = json.Unmarshal(data, &f); err != nil {
		return 0, err
	}
	return f.Number, nil
}

func main() {
	// 第一步：连接 etcd
	etcdReg := etcd.NewRegistry(
		registry.Addrs(etcdAddress),
	)
	// 第二步：服务发现 IP:PORT
	address, err := GetServerNodeAddress(etcdReg)
	if err != nil {
		log.Fatal(err)
	}
	// 第三步：获取结果
	uri := "http://" + address + urlPath
	num, err := GetRand(uri)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("url:%s, result:%d", uri, num)
	// 会通过不同的服务端地址获取结果
	// url:http://10.232.9.217:8001/v1/rand, result:45
	// url:http://10.232.9.217:8002/v1/rand, result:40
	// url:http://10.232.9.217:8000/v1/rand, result:95
}
