// Code generated by protoc-gen-go.
// source: echo/echo_service.proto
// DO NOT EDIT!

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	echo/echo_service.proto

It has these top-level messages:
	EchoRequest
	EchoResponse
*/
package echo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoRequest struct {
	Message          *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoRequest) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type EchoResponse struct {
	Message          *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EchoResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "sofa.pbrpc.test.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "sofa.pbrpc.test.EchoResponse")
}

func init() { proto.RegisterFile("echo/echo_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4d, 0xce, 0xc8,
	0xd7, 0x07, 0x11, 0xf1, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xfc, 0xc5, 0xf9, 0x69, 0x89, 0x7a, 0x05, 0x49, 0x45, 0x05, 0xc9, 0x7a, 0x25, 0xa9,
	0xc5, 0x25, 0x4a, 0xea, 0x5c, 0xdc, 0xae, 0xc9, 0x19, 0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x4c,
	0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92, 0x06, 0x17, 0x0f, 0x44, 0x61, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0x2a, 0x6e, 0x95, 0x46, 0x81, 0x5c, 0x5c, 0x20, 0x95, 0xc1, 0xa9, 0x45, 0x65, 0xa9, 0x45, 0x42,
	0xce, 0x5c, 0x2c, 0x20, 0x9e, 0x90, 0x8c, 0x1e, 0x9a, 0xd5, 0x7a, 0x48, 0xf6, 0x4a, 0xc9, 0xe2,
	0x90, 0x85, 0x58, 0xe6, 0xc4, 0xd6, 0xc0, 0xc8, 0xd8, 0xc1, 0xc8, 0x08, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x08, 0xdc, 0xde, 0xc4, 0xd8, 0x00, 0x00, 0x00,
}
