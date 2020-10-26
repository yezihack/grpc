package main

import (
	"context"
	proto_api "go-micro-simple/02.grpc-etcd/proto"
)

type ApiServiceImpl struct {
}

func (a ApiServiceImpl) GetName(ctx context.Context, request *proto_api.GetNameRequest, response *proto_api.GetNameResponse) error {
	result := request.FirstName + ":" + request.LastName
	response = new(proto_api.GetNameResponse)
	response.Result = result
	return nil
}
