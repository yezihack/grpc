package main

import (
	"context"
	"fmt"
	proto_user "go-micro-study/06.grpc-web/proto"
	"log"
	"math/rand"
	"time"

	"github.com/micro/go-micro/v2"
)

const (
	serviceName = "study.micro.server"
	address     = ":8080"
	etcdAddress = "192.168.61.100:2379"
)

type UserModel struct {
}

func (u UserModel) GetInfo(ctx context.Context, req *proto_user.GetInfoRequest, rsp *proto_user.GetInfoResponse) error {
	log.Println("ID", req.UserId)
	rsp = &proto_user.GetInfoResponse{
		UserName: "jack-" + fmt.Sprint(1),
		Age:      int32(rand.Intn(30)),
		Like:     "羽毛球",
		Sex:      "男",
	}
	return nil
}

func UserList() map[int32]*proto_user.GetInfoResponse {
	lst := make(map[int32]*proto_user.GetInfoResponse)
	rand.NewSource(time.Now().UnixNano())
	var i int32
	for i = 0; i < 10; i++ {
		lst[i] = &proto_user.GetInfoResponse{
			UserName: "jack-" + fmt.Sprint(i),
			Age:      int32(rand.Intn(30)),
			Like:     "羽毛球",
			Sex:      "男",
		}
	}
	return lst
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)
	service.Init()
	if err := proto_user.RegisterUserServiceHandler(service.Server(), new(UserModel)); err != nil {
		log.Fatalln(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
