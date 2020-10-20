package tests

import (
	"context"
	"fmt"
	"github.com/yezihack/grpc/simple/server/proto"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestJoinSQL(t *testing.T) {
	const address = "localhost:8008"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	client := chat.NewChatServiceClient(conn)
	request := chat.JoinSqlRequest{
		ProjectId: 1,
	}
	request.Nodes = make([]*chat.JoinNode, 0)
	request.Nodes = append(request.Nodes, &chat.JoinNode{
		Field:     "name",
		FieldType: chat.SqlTypeEnum_STRING,
		Value:     "项目名称222",
	})
	request.Nodes = append(request.Nodes, &chat.JoinNode{
		Field:     "accumulativeNumber",
		FieldType: chat.SqlTypeEnum_DOUBLE,
		Value:     "aabc",
	})
	request.Nodes = append(request.Nodes, &chat.JoinNode{
		Field:     "raiseDays",
		FieldType: chat.SqlTypeEnum_INT64,
		Value:     "88",
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	reply, err := client.JoinSql(ctx, &request)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("SQL:", reply.Sql)
	fmt.Println("Count:", reply.Count)
}
