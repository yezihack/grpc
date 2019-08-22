package main

import (
	"context"
	"github.com/yezihack/grpc/simple/server/proto"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

var (
	address = "localhost:8008"
)

func main() {
	log.SetOutput(os.Stdout)
	//拨号GRPC服务器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := chat.NewChatServiceClient(conn)
	req := chat.SendRequest{
		Content: "world 2019",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	reply, err := client.Send(ctx, &req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(reply.Msg)

	//测试joinSQL

}
