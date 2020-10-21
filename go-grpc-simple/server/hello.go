package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello "grpc-simple/proto"
	"log"
	"net"
	"os"
)
const (
	Addr = "localhost:3721"
)

type HelloServiceImpl struct {
}
func (p *HelloServiceImpl) Hello(ctx context.Context, req *hello.StringRequest)(
	resp *hello.StringResponse, err error) {
	resp = new(hello.StringResponse)
	resp.Result = "hello grpc client."
	log.Println("req", req.GetValue())
	return resp, nil
}
func main() {
	// 构造一个 gRPC 服务对象
	grpcServer := grpc.NewServer()
	// 然后使用 protoc 工具生成的 go 代码函数(RegisterHelloServiceServer)  注册我们实现的 HelloServiceImpl 服务
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	// 通过 grpcServer.Serve 在一个监听端口上提供 gRPC 服务
	lst, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		fmt.Printf("服务端的gRPC进程服务 %d %s", os.Getpid(), Addr)
	}()
	log.Fatal(grpcServer.Serve(lst))
}

