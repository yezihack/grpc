// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package proto_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetNameRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNameRequest) Reset()         { *m = GetNameRequest{} }
func (m *GetNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetNameRequest) ProtoMessage()    {}
func (*GetNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *GetNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNameRequest.Unmarshal(m, b)
}
func (m *GetNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNameRequest.Marshal(b, m, deterministic)
}
func (m *GetNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNameRequest.Merge(m, src)
}
func (m *GetNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetNameRequest.Size(m)
}
func (m *GetNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNameRequest proto.InternalMessageInfo

func (m *GetNameRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *GetNameRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GetNameResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNameResponse) Reset()         { *m = GetNameResponse{} }
func (m *GetNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetNameResponse) ProtoMessage()    {}
func (*GetNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *GetNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNameResponse.Unmarshal(m, b)
}
func (m *GetNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNameResponse.Marshal(b, m, deterministic)
}
func (m *GetNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNameResponse.Merge(m, src)
}
func (m *GetNameResponse) XXX_Size() int {
	return xxx_messageInfo_GetNameResponse.Size(m)
}
func (m *GetNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNameResponse proto.InternalMessageInfo

func (m *GetNameResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*GetNameRequest)(nil), "proto.api.GetNameRequest")
	proto.RegisterType((*GetNameResponse)(nil), "proto.api.GetNameResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x53, 0x7a, 0x89, 0x05, 0x99, 0x4a, 0x3e, 0x5c,
	0x7c, 0xee, 0xa9, 0x25, 0x7e, 0x89, 0xb9, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0xb2, 0x5c, 0x5c, 0x69, 0x99, 0x45, 0xc5, 0x25, 0xf1, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x9c, 0x60, 0x11, 0x90, 0x2a, 0x21, 0x69, 0x2e, 0xce, 0x9c, 0x44, 0x98,
	0x2c, 0x13, 0x58, 0x96, 0x03, 0x24, 0x00, 0x92, 0x54, 0xd2, 0xe4, 0xe2, 0x87, 0x9b, 0x56, 0x5c,
	0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0x35,
	0x0a, 0xca, 0x33, 0x0a, 0xe0, 0xe2, 0x72, 0x2c, 0xc8, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x15, 0x72, 0xe2, 0x62, 0x87, 0x6a, 0x14, 0x92, 0xd4, 0x83, 0xbb, 0x4e, 0x0f, 0xd5, 0x69, 0x52,
	0x52, 0xd8, 0xa4, 0x20, 0xf6, 0x28, 0x31, 0x24, 0xb1, 0x81, 0x25, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x36, 0x34, 0xe1, 0x0d, 0xe9, 0x00, 0x00, 0x00,
}
