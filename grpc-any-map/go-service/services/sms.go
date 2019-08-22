package services

import (
	"context"
	"fmt"
	"github.com/yezihack/grpc/grpc-any-map/grpc_proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type SmsService struct {
}

func (s *SmsService) Send(
	ctx context.Context,
	req *sms.SendRequest) (reply *sms.SendReply, err error) {
	fmt.Println("send", "in")
	reply = new(sms.SendReply)
	if req == nil || req.Id == 0 {
		err = status.Errorf(codes.InvalidArgument, "invalid arg")
		return
	}
	fmt.Println("send", "in2")
	requestMap := make(map[string]interface{})
	fmt.Println("data", req.Data)
	for k, v := range req.Data {
		var (
			column      string
			columnValue interface{}
		)
		column = strings.TrimSpace(k)
		// unmarshal
		columnValue, err = ProtoUnmarshalWithTypeUrl(v.TypeUrl, v.Value)
		if err != nil {
			return
		}
		// repeatable
		if _, exist := requestMap[column]; exist {
			err = fmt.Errorf("repeatable column %s", column)
			return
		}
		requestMap[column] = columnValue
	}
	if len(requestMap) == 0 {
		err = status.Errorf(codes.InvalidArgument, "empty request data")
		return
	}
	fmt.Println(requestMap)
	return
}
