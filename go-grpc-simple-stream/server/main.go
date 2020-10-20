package main

import (
	"fmt"
	hello "go-grpc-simple-stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type HelloServiceImpl struct {
}

// 客户端向服务端推送的单向流
func (h HelloServiceImpl) ClientToServer(request *hello.StreamRequest, server hello.HelloService_ClientToServerServer) error {
	i := 0
	for {
		i ++
		err := server.Send(&hello.StreamResponse{
			Data: fmt.Sprintf("单向流发送:%v",  time.Now().Format("2006-01-02 15:04:05")),
		})
		if err != nil {
			if err == io.EOF{
				return nil
			}
			return err
		}
		if i > 10 {
			break
		}
	}
	return nil
}

// 服务端向客户端推送的单向流
func (h HelloServiceImpl) ServerToClient(server hello.HelloService_ServerToClientServer) error {
	for {
		data, err := server.Recv()
		if err != nil {
			if err == io.EOF{
				return nil
			}
			return err
		}
		log.Printf("recv:%v", data.GetData())
	}
	return nil
}
// 双向流
func (h HelloServiceImpl) AllStream(server hello.HelloService_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 接受客户的流
	go func() {
		defer wg.Done()
		for {
			data, err := server.Recv()
			if err != nil {
				if err == io.EOF{
					break
				}
				log.Fatalln(err)
			}
			log.Printf("双向流的接受:%s\n", data.GetData())
		}
	}()
	go func() {
		defer wg.Done()
		for {
			err := server.Send(&hello.StreamResponse{Data: fmt.Sprintf("双向流发送:%v", time.Now().Format("2006-01-02 15:04:05"))})
			if err != nil {
				break
			}
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", ":3721")
	if err != nil {
		log.Fatal(err)
	}
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}


