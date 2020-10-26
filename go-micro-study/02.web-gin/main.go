package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

// 第2课，学习使用 gin 嵌入到 go-micro 使用。

func main() {
	// 使用 gin 做为 web 框架
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello go-micro web-gin")
	})
	// 使用 go-micro 自带的 web
	service := web.NewService(
		web.Name("go.micro.web.gin"), // go-micro 服务名称
		web.Handler(r),               // 将 gin　句柄注册到 go-micro 服务里
		web.Address(":8000"),         // 开外一个对外端口
	)
	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
