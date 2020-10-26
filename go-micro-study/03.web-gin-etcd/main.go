package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

// 第3课，将服务注册到 etcd 中，实现服务注册
// etcd docker 安装

func main() {
	// 连接 etcd 服务
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.61.100:2379"),
	)
	// 使用 gin web 框架
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello go-micro web-gin.etcd")
	})
	// 实例 go-micro 服务
	service := web.NewService(
		web.Name("go.micro.web.gin.etcd"), // go-micro 服务名称
		web.Registry(etcdReg),             // etcd 注册到 go-micro 服务里
		web.Handler(r),                    // 将 gin　句柄注册到 go-micro 服务里
		web.Address(":8000"),              // 开外一个对外端口
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
