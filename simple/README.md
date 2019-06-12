#   Grpc


## 资料
- https://grpc.io/docs/quickstart/go/
- https://studygolang.com/articles/16627

## 使用方法

```
make run
```


## 新建一个proto文件
```
syntax = "proto3";

package chat;

service ChatService {
    //发送消息
    rpc Send(SendRequest) returns (SendReply){};
}
//发送消息 请求结构体
message SendRequest {
    string content = 1;//发送内容
}
//发送消息 响应结构体
message SendReply {
    string msg = 1;//返回消息
}
```

## GRPC服务端
```
package main

import (
	"context"
	"github.com/yezihack/grpc/simple/server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	port = ":8008"
)

func main() {
	log.SetOutput(os.Stdout)
	//新建一个tcp监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	//起一个服务
	s := grpc.NewServer()
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	chat.RegisterChatServiceServer(s, &Chats{})
	log.Printf("server port %s start...\n", port)
	//启动服务
	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

//新建一个结构体,实现proto里定义的方法
type Chats struct {
}

//实现proto方法
func (c *Chats) Send(ctx context.Context, in *chat.SendRequest) (*chat.SendReply, error) {
	out := chat.SendReply{
		Msg: "hello " + in.Content,
	}
	return &out, nil
}

```

## 客户端代码
```
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

```
