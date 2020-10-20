package main

import (
	"context"
	"google.golang.org/grpc"
	hello "grpc-simple/proto"
	"log"
)

func main() {
	// 1. grpc.Dial 负责和 gRPC 服务建立链接
	conn, err := grpc.Dial("localhost:3721", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 2. 使用NewHelloServiceClient() 函数基于链接构建 helloServiceClient 对象。
	// 然后就可以调用 helloServiceClient{} 结构体定义的方法啦，如 Hello 方法
	client := hello.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &hello.StringRequest{
		Value: "hello grpc server",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("grpc client result:%v\n", resp.GetResult())
}
