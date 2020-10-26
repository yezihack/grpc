package main

import (
	"log"
	"net/http"

	"github.com/micro/go-micro/v2/web"
)

// 第1课，学习使用 go-micro 自带的 web 框架

func main() {
	// 实例一个 go-micro 服务
	service := web.NewService(
		web.Name("go.micro.web"), // web 服务名称
		web.Address(":8000"),     // 开外一个对外端口
	)
	// 使用 go-micro 自带的 web
	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello go-micro web"))
	})
	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
