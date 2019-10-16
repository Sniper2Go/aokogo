// Code generated by protoc-gen-go.
// source: MSG_Server.proto
// DO NOT EDIT!

/*
Package MSG_Server is a generated protocol buffer package.

It is generated from these files:
	MSG_Server.proto

It has these top-level messages:
	CS_EnterServer_Req
	SC_EnterServer_Rsp
	CS_ServerRegister_Req
	SC_ServerRegister_Rsp
*/
package MSG_Server

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

// add by stefanchen
// server
type SUBMSG int32

const (
	SUBMSG_Begin             SUBMSG = 0
	SUBMSG_CS_EnterServer    SUBMSG = 1
	SUBMSG_SC_EnterServer    SUBMSG = 2
	SUBMSG_CS_ServerRegister SUBMSG = 3
	SUBMSG_SC_ServerRegister SUBMSG = 4
)

var SUBMSG_name = map[int32]string{
	0: "Begin",
	1: "CS_EnterServer",
	2: "SC_EnterServer",
	3: "CS_ServerRegister",
	4: "SC_ServerRegister",
}
var SUBMSG_value = map[string]int32{
	"Begin":             0,
	"CS_EnterServer":    1,
	"SC_EnterServer":    2,
	"CS_ServerRegister": 3,
	"SC_ServerRegister": 4,
}

func (x SUBMSG) String() string {
	return proto.EnumName(SUBMSG_name, int32(x))
}
func (SUBMSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ErrorCode int32

const (
	ErrorCode_Invalid ErrorCode = 0
	ErrorCode_Success ErrorCode = 1
	ErrorCode_Fail    ErrorCode = 2
)

var ErrorCode_name = map[int32]string{
	0: "Invalid",
	1: "Success",
	2: "Fail",
}
var ErrorCode_value = map[string]int32{
	"Invalid": 0,
	"Success": 1,
	"Fail":    2,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// CS_EnterServer
type CS_EnterServer_Req struct {
	Enter int32 `protobuf:"varint,1,opt,name=Enter" json:"Enter,omitempty"`
}

func (m *CS_EnterServer_Req) Reset()                    { *m = CS_EnterServer_Req{} }
func (m *CS_EnterServer_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_EnterServer_Req) ProtoMessage()               {}
func (*CS_EnterServer_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// SC_EnterServer
type SC_EnterServer_Rsp struct {
	Ret ErrorCode `protobuf:"varint,1,opt,name=Ret,enum=MSG_Server.ErrorCode" json:"Ret,omitempty"`
}

func (m *SC_EnterServer_Rsp) Reset()                    { *m = SC_EnterServer_Rsp{} }
func (m *SC_EnterServer_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_EnterServer_Rsp) ProtoMessage()               {}
func (*SC_EnterServer_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// CS_ServerRegister
type CS_ServerRegister_Req struct {
	ServerType int32    `protobuf:"varint,1,opt,name=ServerType" json:"ServerType,omitempty"`
	Msgs       []uint32 `protobuf:"varint,2,rep,name=Msgs" json:"Msgs,omitempty"`
}

func (m *CS_ServerRegister_Req) Reset()                    { *m = CS_ServerRegister_Req{} }
func (m *CS_ServerRegister_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_ServerRegister_Req) ProtoMessage()               {}
func (*CS_ServerRegister_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// SC_ServerRegister
type SC_ServerRegister_Rsp struct {
	Ret ErrorCode `protobuf:"varint,1,opt,name=Ret,enum=MSG_Server.ErrorCode" json:"Ret,omitempty"`
}

func (m *SC_ServerRegister_Rsp) Reset()                    { *m = SC_ServerRegister_Rsp{} }
func (m *SC_ServerRegister_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_ServerRegister_Rsp) ProtoMessage()               {}
func (*SC_ServerRegister_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*CS_EnterServer_Req)(nil), "MSG_Server.CS_EnterServer_Req")
	proto.RegisterType((*SC_EnterServer_Rsp)(nil), "MSG_Server.SC_EnterServer_Rsp")
	proto.RegisterType((*CS_ServerRegister_Req)(nil), "MSG_Server.CS_ServerRegister_Req")
	proto.RegisterType((*SC_ServerRegister_Rsp)(nil), "MSG_Server.SC_ServerRegister_Rsp")
	proto.RegisterEnum("MSG_Server.SUBMSG", SUBMSG_name, SUBMSG_value)
	proto.RegisterEnum("MSG_Server.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("MSG_Server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0xcb, 0x4b, 0xc3, 0x40,
	0x10, 0x87, 0x9b, 0x57, 0xb5, 0x23, 0x96, 0x75, 0x30, 0xd0, 0x93, 0x94, 0x5c, 0x2c, 0x39, 0x54,
	0xd0, 0xb3, 0x20, 0x5d, 0x6a, 0x11, 0xc9, 0x65, 0x47, 0xcf, 0xa1, 0xb6, 0x43, 0x58, 0x28, 0x49,
	0xdc, 0x8d, 0x05, 0xff, 0x7b, 0xc9, 0x03, 0x6d, 0xc8, 0xc9, 0xdb, 0xce, 0xb7, 0x3f, 0xf6, 0x9b,
	0xd9, 0x01, 0x91, 0xd0, 0x26, 0x25, 0x36, 0x47, 0x36, 0xcb, 0xd2, 0x14, 0x55, 0x81, 0xf0, 0x47,
	0xa2, 0x18, 0x50, 0x52, 0xba, 0xce, 0x2b, 0x36, 0x2d, 0x49, 0x15, 0x7f, 0xe2, 0x35, 0x04, 0x0d,
	0x9a, 0x39, 0x73, 0x67, 0x11, 0xa8, 0xb6, 0x88, 0x1e, 0x01, 0x49, 0xf6, 0xb3, 0xb6, 0xc4, 0x5b,
	0xf0, 0x14, 0x57, 0x4d, 0x72, 0x7a, 0x1f, 0x2e, 0x4f, 0x6c, 0x6b, 0x63, 0x0a, 0x23, 0x8b, 0x3d,
	0xab, 0x3a, 0x11, 0xbd, 0x42, 0x28, 0xa9, 0xbb, 0x53, 0x9c, 0x69, 0x5b, 0x75, 0xb6, 0x1b, 0x80,
	0x96, 0xbe, 0x7d, 0x97, 0xdc, 0x29, 0x4f, 0x08, 0x22, 0xf8, 0x89, 0xcd, 0xec, 0xcc, 0x9d, 0x7b,
	0x8b, 0x4b, 0xd5, 0x9c, 0xa3, 0x27, 0x08, 0x49, 0x0e, 0x1e, 0xfb, 0x47, 0x3b, 0xb1, 0x86, 0x31,
	0xbd, 0xaf, 0x12, 0xda, 0xe0, 0x04, 0x82, 0x15, 0x67, 0x3a, 0x17, 0x23, 0x44, 0x98, 0xf6, 0xbf,
	0x43, 0x38, 0x35, 0xeb, 0x8f, 0x2d, 0x5c, 0x0c, 0xe1, 0x6a, 0x30, 0x8b, 0xf0, 0x6a, 0x3c, 0xe8,
	0x4a, 0xf8, 0xf1, 0x1d, 0x4c, 0x7e, 0xe5, 0x78, 0x01, 0x67, 0x2f, 0xf9, 0x71, 0x7b, 0xd0, 0x7b,
	0x31, 0xaa, 0x0b, 0xfa, 0xda, 0xed, 0xd8, 0x5a, 0xe1, 0xe0, 0x39, 0xf8, 0xcf, 0x5b, 0x7d, 0x10,
	0xee, 0xc7, 0xb8, 0x59, 0xd4, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0x9d, 0xf2, 0x73,
	0xbc, 0x01, 0x00, 0x00,
}
