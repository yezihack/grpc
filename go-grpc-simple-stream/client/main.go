package main

import (
	"context"
	hello "go-grpc-simple-stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile|log.LstdFlags)
	conn, err := grpc.Dial("localhost:3721", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := hello.NewHelloServiceClient(conn)
	stream, err := client.Say(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			req := hello.SayRequest{
				Value: "hello stream",
			}
			if err := stream.Send(&req); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.CloseAndRecv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		result := reply.GetResult()
		log.Println("reply:", result)
	}
}
