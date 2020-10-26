package main

import (
	"context"
	proto_api "go-micro-simple/02.grpc-etcd/proto"
	"log"

	"github.com/micro/go-micro/v2/client"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

type logWrapper struct {
	client.Client
}

func (w *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Println("wrapper")
	return w.Client.Call(ctx, req, rsp)
}
func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {
	// 注册 etcd 服务
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.61.100:2379"),
	)
	// 使用 gin 做 web 服务
	r := gin.Default()
	r.GET("/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "ok",
		})
	})
	// go-micro 开启一个 web 服务
	webService := web.NewService(
		web.Registry(etcdReg),     // 注册到etcd里
		web.Handler(r),            // 使用 gin 做 web
		web.Name("micro.api.web"), // 起个服务名称
		web.Address(":7000"),      // 开启对外的服务端口
	)
	webService.Init()
	go func() {
		if err := webService.Run(); err != nil {
			log.Fatalln(err)
		}
	}()

	// New Service micro服务
	grpcService := micro.NewService(
		micro.Name("micro.api.grpc"),    // 设置名称
		micro.Address(":8001"),          // 对外提供端口
		micro.Registry(etcdReg),         // 将etcd加载到服务里
		micro.WrapClient(NewLogWrapper), // 装饰器
	)
	grpcService.Init() // 服务初使化
	// 注册 proto 服务
	proto_api.RegisterApiServiceHandler(grpcService.Server(), new(ApiServiceImpl))
	// 运行 micro 服务
	if err := grpcService.Run(); err != nil {
		log.Fatal(err)
	}
}
