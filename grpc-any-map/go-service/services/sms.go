package services

import (
	"context"
	"fmt"
	"github.com/yezihack/grpc/grpc-any-map/grpc_proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
)

type SmsService struct {
}

func (s *SmsService) Send(ctx context.Context,
	req *sms.SendRequest) (reply *sms.SendReply, err error) {
	reply = new(sms.SendReply)
	if req == nil || req.Id == 0 {
		err = status.Errorf(codes.InvalidArgument, "invalid arg")
		return
	}
	requestMap := make(map[string]interface{})
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
	log.Print("requestMap", requestMap)
	msg := strings.Builder{}
	msg.WriteString(fmt.Sprintf("id:%d,", req.Id))
	for field, val := range requestMap {
		msg.WriteString(fmt.Sprintf("field:%s, value: %s,", field, val))
	}
	reply.Msg = msg.String()
	return
}

func (s *SmsService) HelloWorld(ctx context.Context,
	req *sms.HelloWorldRequest) (reply *sms.HelloWorldReply, err error) {
	msg := strings.Builder{}
	for k, v := range req.Data {
		msg.WriteString(fmt.Sprintf("k:%s,v:%s", k, v))
	}
	for _, v := range req.Items {
		msg.WriteString(fmt.Sprintf("id:%d, name:%s, type:%s", v.Id, v.Name, v.Type))
	}
	reply.Msg = msg.String()
	return
}
