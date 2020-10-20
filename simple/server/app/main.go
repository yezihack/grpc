package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yezihack/grpc/simple/server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var (
	port        = ":8008"
	mysqlMaster *sql.DB
)

func init() {
	var err error
	source := "test_order:61NNT9RJSLwGelGy@tcp(rm-bp1y256043o82d4wa.mysql.rds.aliyuncs.com:3306)/swift_nuochou_com?charset=utf8"
	mysqlMaster, err = sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
}
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

func (c *Chats) JoinSql(ctx context.Context, request *chat.JoinSqlRequest) (reply *chat.JoinSqlReply, err error) {
	reply = new(chat.JoinSqlReply)
	projectId := request.ProjectId
	nodes := request.Nodes
	requestMap := make(map[string]interface{})
	for _, node := range nodes {
		var (
			column      string
			columnValue interface{}
		)
		column = strings.TrimSpace(node.Field)
		//判断是否触发黑名单字段
		if exist := forbiddenColumn(column); exist {
			err = fmt.Errorf("forbid update column %s", column)
			return
		}
		//判断类型是否合法.并转换类型
		switch node.FieldType {
		case chat.SqlTypeEnum_INT64:
			columnValue, err = strconv.ParseInt(strings.TrimSpace(node.Value), 10, 64)
			if err != nil {
				err = fmt.Errorf("field value conversion error %s", node.Value)
				return
			}
		case chat.SqlTypeEnum_DOUBLE:
			columnValue, err = strconv.ParseFloat(strings.TrimSpace(node.Value), 64)
			if err != nil {
				err = fmt.Errorf("field value conversion error %s", node.Value)
				return
			}
		case chat.SqlTypeEnum_STRING:
			columnValue = strings.TrimSpace(node.Value)
		default:
			err = fmt.Errorf("not support type %s", node.Field)
			return
		}
		//判断是否有重复字段
		if _, exist := requestMap[column]; exist {
			err = fmt.Errorf("repeatable column %s", column)
			return
		}
		//加入到map里
		requestMap[column] = columnValue
	}
	//判断请求字段是否为空
	if len(requestMap) == 0 {
		err = fmt.Errorf("empty request data")
		return
	}
	var (
		columnList []string
		paramList  []interface{}
	)
	//将字段与值分离出来
	for k, v := range requestMap {
		columnList = append(columnList, k)
		paramList = append(paramList, v)
	}
	cmd := fmt.Sprintf("UPDATE raisefunds SET %s WHERE id = %d", columns2condition(columnList), projectId)
	stmt, err := mysqlMaster.Prepare(cmd)
	if err != nil {
		err = fmt.Errorf("sql prepare err %s", cmd)
		return
	}
	res, err := stmt.Exec(paramList...)
	if err != nil {
		err = fmt.Errorf("exec sql err %s", err)
		return
	}
	reply.Sql = cmd
	reply.Count, err = res.RowsAffected()
	return
}

// name conf => name= ? , conf = ?
func columns2condition(columns []string) (condition string) {
	buf := bytes.Buffer{}
	for index, item := range columns {
		buf.WriteString(item)
		buf.WriteString(" = ? ")
		if index != len(columns)-1 {
			buf.WriteString(",")
		}
	}
	condition = buf.String()
	return
}

//黑名单
func forbiddenColumn(field string) bool {
	var forbiddenColumn map[string]struct{}
	forbiddenColumn = make(map[string]struct{})
	forbiddenColumn["id"] = struct{}{}
	forbiddenColumn["uuid"] = struct{}{}
	forbiddenColumn["user_id"] = struct{}{}
	forbiddenColumn["active"] = struct{}{}
	forbiddenColumn["private"] = struct{}{}
	forbiddenColumn["updated_at"] = struct{}{}
	forbiddenColumn["state"] = struct{}{}
	forbiddenColumn["stopped"] = struct{}{}
	forbiddenColumn["over_timestamp"] = struct{}{}
	forbiddenColumn["startDate"] = struct{}{}
	forbiddenColumn["endDate"] = struct{}{}
	forbiddenColumn["created"] = struct{}{}
	forbiddenColumn["created_at"] = struct{}{}
	forbiddenColumn["updated"] = struct{}{}
	if _, ok := forbiddenColumn[field]; ok {
		return true
	}
	return false
}
