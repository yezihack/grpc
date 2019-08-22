package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/yezihack/grpc/grpc-any-map/grpc_proto"
	"google.golang.org/grpc"
	"testing"
	"time"
)

var (
	address = "localhost:8282"
)

/**
foo := &pb.Foo{...}
//      any, err := ptypes.MarshalAny(foo)
//      ...
//      foo := &pb.Foo{}
//      if err := ptypes.UnmarshalAny(any, foo); err != nil {
//        ...
//      }
*/
func TestSend(t *testing.T) {
	//拨号
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		//grpc.WithCompressor(grpc.NewGZIPCompressor()),
		//grpc.WithDecompressor(grpc.NewGZIPDecompressor()),
	)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	client := sms.NewSmsServiceClient(conn)
	req := new(sms.SendRequest)
	data := make(map[string]*any.Any)
	int64Value := &wrappers.Int64Value{Value: 1}
	int64ValueBytes, err := proto.Marshal(int64Value)
	if err != nil {
		return
	}
	data["city_id"] = &any.Any{
		TypeUrl: proto.MessageName(int64Value),
		Value:   int64ValueBytes,
	}
	for k, v := range data {
		fmt.Println("url", v.TypeUrl, "key", k)
	}
	req.Data = data
	req.Id = 100
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	reply, err := client.Send(ctx, req)
	if err != nil {
		t.Error(err)
	}
	println(reply)
}
