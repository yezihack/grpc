package main

import (
	"github.com/yezihack/grpc/grpc-any-map/go-service/services"
	"github.com/yezihack/grpc/grpc-any-map/grpc_proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	port = ":8282"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lst, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}
	g := grpc.NewServer()
	s := new(services.SmsService)
	log.Printf("start service port%s\n", port)
	sms.RegisterSmsServiceServer(g, s)
	if err = g.Serve(lst); err != nil {
		log.Fatalln(err)
	}
}
