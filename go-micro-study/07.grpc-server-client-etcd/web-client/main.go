package main

import (
	"context"
	"fmt"
	proto_user "go-micro-study/06.grpc-web/proto"
	"log"

	"github.com/micro/go-micro/v2"
)

const (
	serviceName = "study.micro.server"
	clientName  = "study.micro.client"
	address     = ":8000"
	etcdAddress = "192.168.61.100:2379"
)

// 第一步： 连接ETCD
// 第二步： 使用 gin 做 web 框架
// 第三步： 实例 web go-micro
// 第四步： 请求 服务端 grpc 的方法
func main() {
	// 第二步： 实例 proto 客户端，并服务发现
	myService := micro.NewService()
	userService := proto_user.NewUserService("greeter", myService.Client())

	rsp, err := userService.GetInfo(context.Background(), &proto_user.GetInfoRequest{
		UserId: 3,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("name", rsp.GetUserName())
	//
	//// 第三步： 使用 gin 做 web 框架
	//ginHandler := gin.Default()
	//ginHandler.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, "PONG")
	//})
	//ginHandler.GET("/user/:id", func(c *gin.Context) {
	//	id := c.Param("id")
	//	id32 := cast.ToInt32(id)
	//	log.Println("id", id32)
	//	rsp, err := userService.GetInfo(context.Background(), &proto_user.GetInfoRequest{
	//		UserId: id32,
	//	})
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	log.Println("rsp", rsp.GetUserName())
	//	c.JSON(200, rsp)
	//})
	//// 第四步： 实例 web go-micro
	//service := web.NewService(
	//	web.Handler(ginHandler),
	//	web.Address(address),
	//)
	//service.Init()
	//if err := service.Run(); err != nil {
	//	log.Fatal(err)
	//}
}
