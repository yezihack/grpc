package main

import (
	"context"
	"fmt"
	hello "go-grpc-simple-stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync"
	"time"
)

const (
	Addr = "localhost:3721"
)

func main() {
	log.SetFlags(log.Lshortfile|log.LstdFlags)
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := hello.NewHelloServiceClient(conn)
	//GetServerStream(client)
	//PutServerStream(client)
	AllStream(client)
}
// 双向流
func AllStream(client hello.HelloServiceClient) {
	all, err := client.AllStream(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 处理服务端->客户端的流
	go func() {
		defer wg.Done()
		for {
			resp, err := all.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalln(err)
			}
			log.Printf("双向流，接受到服务的数据:%s\n", resp.GetData())
		}
	}()
	// 处理 客户端->服务端的流
	go func() {
		defer wg.Done()
		i := 0
		for {
			err = all.Send(&hello.StreamRequest{
				Data: fmt.Sprintf("%d. 来自客户端:%s\n", i, time.Now().Format("15:04:05")),
			})
			if err != nil {
				log.Println(err)
				break
			}
			i ++
			if i > 3 {
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
// 客户端不断的向服务器推送单向流
func PutServerStream(client hello.HelloServiceClient) {
	i := 0
	// 向服务端推送流
	put, err := client.ClientToServer(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for {
		err = put.Send(&hello.StreamRequest{
			Data: fmt.Sprintf("%d.来自客户端<client>，%s", i, time.Now().Format("2006-01-02 15:04:05")),
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
	// 接受
	resp, err := put.CloseAndRecv()
	if err != nil {
		if err != io.EOF {
			log.Fatalln(err)
		}
	}
	// 接受服务端响应的数据
	if resp != nil {
		log.Println(resp.GetData()) // 接受完毕
	}

}

// 获取服务端向客户端不断推送的数据流。
func GetServerStream(client hello.HelloServiceClient) {
	// 向服务器发一个数据标识
	req := hello.StreamRequest{Data: "客户端"}
	// 调用 ServerToClient 函数，准备接受服务端单向流
	resp, err := client.ServerToClient(context.Background(), &req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		data, err := resp.Recv()
		if err != nil {
			// 遇到 io.EOF 表示服务端流关闭
			if err == io.EOF {
				break
			}
			log.Println(err)
			break
		}
		log.Printf("服务端推送的单向流:%s\n", data)
	}
}
