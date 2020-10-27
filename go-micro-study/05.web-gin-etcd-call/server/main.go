package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/micro/go-micro/v2/web"

	"github.com/gin-gonic/gin"
)

const (
	serviceName = "go.micro.web.gin.etcd.call"
	etcdAddress = "192.168.61.100:2379"
	addr        = ":8000"
)

// 实现将 web 服务注册到服务里，使用服务发现去调用它们。
// 服务端启动三个服务(--server_address是 go-micro 自带的参数)
// go run . --server_address :8000
// go run . --server_address :8001
// go run . --server_address :8002

func GetRandomID(c *gin.Context) {
	rand.NewSource(time.Now().UnixNano()) // 种子，保证每次随机不一样。
	ID := rand.Intn(100)                  // 随机数
	c.JSON(200, gin.H{
		"random-number": ID,
	})
}

func main() {
	// 第一步：使用gin
	r := gin.Default()
	r.GET("/v1/rand", GetRandomID)
	// 第二步：连接 etcd
	etcdReg := etcd.NewRegistry(
		registry.Addrs(etcdAddress),
	)
	// 第三步：实现 go-micro web 实例
	service := web.NewService(
		web.Name(serviceName), // go-micro 服务名称
		web.Registry(etcdReg), // etcd 注册到 go-micro 服务里
		web.Handler(r),        // Gin 句柄注册到 go-micro 服务里
		web.Address(addr),     // 开外一个对外端口
	)
	// 第四步：启动命令行
	if err := service.Init(); err != nil {
		log.Fatalln(err)
	}
	// 第五步：启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
