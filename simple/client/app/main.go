package main

import (
	"context"
	"github.com/yezihack/grpc/simple/server/proto"
	"google.golang.org/grpc"
	"log"
	"os"
)

var (
	addrees = "localhost:8008"
)

func main() {
	log.SetOutput(os.Stdout)
	conn, err := grpc.Dial(addrees, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := chat.NewChatServiceClient(conn)
	req := chat.SendRequest{
		Content: "world 2019",
	}
	reply, err := client.Send(context.Background(), &req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(reply.Msg)

}
