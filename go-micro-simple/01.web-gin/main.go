package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// 使用 gin web 框架
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "hello go-micro",
		})
	})
	// 使用 go-micro 里的 web 服务
	service := web.NewService(
		web.Handler(r),              // 注册句柄
		web.Name("go-micro-simple"), // 设置 web 名称
		web.Address(":8000"),        // 定义一个对外端口

	)
	// 原始写法，需要注释掉 gin 的句柄
	//service.HandleFunc("/ok", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("ok"))
	//})
	err := service.Run() // 运行
	if err != nil {
		log.Fatal(err)
	}
}
