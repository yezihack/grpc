package main

import (
	"fmt"
	hello "go-grpc-simple-stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)
const (
	Addr = "localhost:3721"
)

type HelloServiceImpl struct {
}

// 服务端->客户端 推送单向流
func (h *HelloServiceImpl) ServerToClient(req *hello.StreamRequest, server hello.HelloService_ServerToClientServer) error {
	// 模拟一个向客户端推送10次的单向流
	var i int
	for {
		// 打印接受客户端的消息
		log.Printf("接受到客户端的消息:%s\n", req.GetData())
		// 向客户端发送消息
		err := server.Send(&hello.StreamResponse{Data: fmt.Sprintf("来自服务器<server>, %s", time.Now().Format("2006-01-02 15:04:05"))})
		if err != nil {
			break
		}
		i ++
		if i > 10 {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}
// 客户端->服务端 推送单向流
// 需要循环的接受来自客户端的消息，至到 io.EOF
func (h *HelloServiceImpl) ClientToServer(server hello.HelloService_ClientToServerServer) error {
	for {
		// 接受客户端的消息
		data, err := server.Recv()
		if err != nil {
			// 无数据时跳出循环
			if err == io.EOF {
				break
			}
			return err
		}
		log.Printf("接受到客户端的消息:%s", data.GetData())
	}
	err := server.SendAndClose(&hello.StreamResponse{
		Data: "接受完毕",
	})
	if err != nil {
		return err
	}
	return nil
}
// 双向流，即可以从服务端不断发送流数据，也可以不断的接受客户端发送过来的流数据。
// 所以需要处理发送与接受，需要采用两个协程处理。
func (h *HelloServiceImpl) AllStream(server hello.HelloService_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 处理服务端向客户端发送的流数据
	go func() {
		defer wg.Done()
		i := 0
		for {
			err := server.Send(&hello.StreamResponse{
				Data: fmt.Sprintf("%d,来自双向流的服务端:%s\n",i, time.Now().Format("2006-01-02 15:04:05")),
			})
			if err != nil {
				break
			}
			i ++
			if i > 3 {
				break
			}
			time.Sleep(time.Second)
		}

	}()
	// 处理客户端向服务端发送过来的流数据。
	go func() {
		for {
			data, err := server.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalln(err)
			}
			log.Printf("来自客户端的消息：%s\n", data.GetData())
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	log.SetFlags(log.Lshortfile|log.LstdFlags)
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Printf("PID:%d, %s\n", os.Getpid(), Addr)
	}()
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}


