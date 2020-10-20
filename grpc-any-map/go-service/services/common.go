package services

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
)

var (
	FloatTypeUrl  string
	Int64TypeUrl  string
	BoolTypeUrl   string
	StringTypeUrl string
)

func init() {
	FloatTypeUrl = proto.MessageName(&wrappers.FloatValue{})
	Int64TypeUrl = proto.MessageName(&wrappers.Int64Value{})
	BoolTypeUrl = proto.MessageName(&wrappers.BoolValue{})
	StringTypeUrl = proto.MessageName(&wrappers.StringValue{})
}

func ProtoUnmarshalWithTypeUrl(typeUrl string, value []byte) (columnValue interface{}, err error) {
	fmt.Println("typeurl", typeUrl)
	switch typeUrl {
	case FloatTypeUrl:
		pb := &wrappers.FloatValue{}
		err = proto.Unmarshal(value, pb)
		if err != nil {
			return
		}
		columnValue = pb.Value
	case Int64TypeUrl:
		pb := &wrappers.Int64Value{}
		err = proto.Unmarshal(value, pb)
		if err != nil {
			return
		}
		columnValue = pb.Value
	case BoolTypeUrl:
		pb := &wrappers.BoolValue{}
		err = proto.Unmarshal(value, pb)
		if err != nil {
			return
		}
		columnValue = pb.Value
	case StringTypeUrl:
		pb := &wrappers.StringValue{}
		err = proto.Unmarshal(value, pb)
		if err != nil {
			return
		}
		columnValue = pb.Value
	default:
		// Not Support Type : DoubleValue UInt64Value Int32Value UInt32Value BytesValue
		err = fmt.Errorf("not support type %s", typeUrl)
		return
	}
	return
}
