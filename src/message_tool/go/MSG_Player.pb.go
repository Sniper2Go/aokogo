// Code generated by protoc-gen-go.
// source: MSG_Player.proto
// DO NOT EDIT!

/*
Package MSG_Player is a generated protocol buffer package.

It is generated from these files:
	MSG_Player.proto

It has these top-level messages:
	BasePlayerInfo_Repeat
	BaseMoney_Repeat
	CS_EnterServer_Req
	SC_EnterServer_Rsp
	CS_PlayerInfo_Req
	SC_PlayerInfo_Rsp
*/
package MSG_Player

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
	SUBMSG_Begin          SUBMSG = 0
	SUBMSG_CS_EnterServer SUBMSG = 1
	SUBMSG_SC_EnterServer SUBMSG = 2
	SUBMSG_CS_PlayerInfo  SUBMSG = 3
	SUBMSG_SC_PlayerInfo  SUBMSG = 4
)

var SUBMSG_name = map[int32]string{
	0: "Begin",
	1: "CS_EnterServer",
	2: "SC_EnterServer",
	3: "CS_PlayerInfo",
	4: "SC_PlayerInfo",
}
var SUBMSG_value = map[string]int32{
	"Begin":          0,
	"CS_EnterServer": 1,
	"SC_EnterServer": 2,
	"CS_PlayerInfo":  3,
	"SC_PlayerInfo":  4,
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

// 玩家基础信息枚举
type EmBaseInfo int32

const (
	EmBaseInfo_baseInvalid EmBaseInfo = 0
	EmBaseInfo_Name        EmBaseInfo = 1
	EmBaseInfo_Level       EmBaseInfo = 2
	EmBaseInfo_HeadIcon    EmBaseInfo = 3
	EmBaseInfo_DBID        EmBaseInfo = 4
)

var EmBaseInfo_name = map[int32]string{
	0: "baseInvalid",
	1: "Name",
	2: "Level",
	3: "HeadIcon",
	4: "DBID",
}
var EmBaseInfo_value = map[string]int32{
	"baseInvalid": 0,
	"Name":        1,
	"Level":       2,
	"HeadIcon":    3,
	"DBID":        4,
}

func (x EmBaseInfo) String() string {
	return proto.EnumName(EmBaseInfo_name, int32(x))
}
func (EmBaseInfo) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 玩家基础金币类枚举
type EmBaseMoney int32

const (
	EmBaseMoney_moneyInvalid EmBaseMoney = 0
	EmBaseMoney_Coin         EmBaseMoney = 1
)

var EmBaseMoney_name = map[int32]string{
	0: "moneyInvalid",
	1: "Coin",
}
var EmBaseMoney_value = map[string]int32{
	"moneyInvalid": 0,
	"Coin":         1,
}

func (x EmBaseMoney) String() string {
	return proto.EnumName(EmBaseMoney_name, int32(x))
}
func (EmBaseMoney) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 玩家基础信息结构
type BasePlayerInfo_Repeat struct {
	Name     string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Level    int32  `protobuf:"varint,2,opt,name=Level" json:"Level,omitempty"`
	HeadIcon string `protobuf:"bytes,3,opt,name=HeadIcon" json:"HeadIcon,omitempty"`
	DBID     int64  `protobuf:"varint,4,opt,name=DBID" json:"DBID,omitempty"`
}

func (m *BasePlayerInfo_Repeat) Reset()                    { *m = BasePlayerInfo_Repeat{} }
func (m *BasePlayerInfo_Repeat) String() string            { return proto.CompactTextString(m) }
func (*BasePlayerInfo_Repeat) ProtoMessage()               {}
func (*BasePlayerInfo_Repeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 玩家基础金币类结构
type BaseMoney_Repeat struct {
	Coin int64 `protobuf:"varint,1,opt,name=Coin" json:"Coin,omitempty"`
}

func (m *BaseMoney_Repeat) Reset()                    { *m = BaseMoney_Repeat{} }
func (m *BaseMoney_Repeat) String() string            { return proto.CompactTextString(m) }
func (*BaseMoney_Repeat) ProtoMessage()               {}
func (*BaseMoney_Repeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// CS_EnterServer
type CS_EnterServer_Req struct {
}

func (m *CS_EnterServer_Req) Reset()                    { *m = CS_EnterServer_Req{} }
func (m *CS_EnterServer_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_EnterServer_Req) ProtoMessage()               {}
func (*CS_EnterServer_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// SC_EnterServer
type SC_EnterServer_Rsp struct {
	Ret ErrorCode `protobuf:"varint,1,opt,name=Ret,enum=MSG_Player.ErrorCode" json:"Ret,omitempty"`
}

func (m *SC_EnterServer_Rsp) Reset()                    { *m = SC_EnterServer_Rsp{} }
func (m *SC_EnterServer_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_EnterServer_Rsp) ProtoMessage()               {}
func (*SC_EnterServer_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// CS_PlayerInfo
type CS_PlayerInfo_Req struct {
}

func (m *CS_PlayerInfo_Req) Reset()                    { *m = CS_PlayerInfo_Req{} }
func (m *CS_PlayerInfo_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_PlayerInfo_Req) ProtoMessage()               {}
func (*CS_PlayerInfo_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// SC_PlayerInfo
type SC_PlayerInfo_Rsp struct {
	BaseInfo  *BasePlayerInfo_Repeat `protobuf:"bytes,1,opt,name=baseInfo" json:"baseInfo,omitempty"`
	BaseMoney *BaseMoney_Repeat      `protobuf:"bytes,2,opt,name=baseMoney" json:"baseMoney,omitempty"`
	Ret       ErrorCode              `protobuf:"varint,3,opt,name=Ret,enum=MSG_Player.ErrorCode" json:"Ret,omitempty"`
}

func (m *SC_PlayerInfo_Rsp) Reset()                    { *m = SC_PlayerInfo_Rsp{} }
func (m *SC_PlayerInfo_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_PlayerInfo_Rsp) ProtoMessage()               {}
func (*SC_PlayerInfo_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SC_PlayerInfo_Rsp) GetBaseInfo() *BasePlayerInfo_Repeat {
	if m != nil {
		return m.BaseInfo
	}
	return nil
}

func (m *SC_PlayerInfo_Rsp) GetBaseMoney() *BaseMoney_Repeat {
	if m != nil {
		return m.BaseMoney
	}
	return nil
}

func init() {
	proto.RegisterType((*BasePlayerInfo_Repeat)(nil), "MSG_Player.BasePlayerInfo_Repeat")
	proto.RegisterType((*BaseMoney_Repeat)(nil), "MSG_Player.BaseMoney_Repeat")
	proto.RegisterType((*CS_EnterServer_Req)(nil), "MSG_Player.CS_EnterServer_Req")
	proto.RegisterType((*SC_EnterServer_Rsp)(nil), "MSG_Player.SC_EnterServer_Rsp")
	proto.RegisterType((*CS_PlayerInfo_Req)(nil), "MSG_Player.CS_PlayerInfo_Req")
	proto.RegisterType((*SC_PlayerInfo_Rsp)(nil), "MSG_Player.SC_PlayerInfo_Rsp")
	proto.RegisterEnum("MSG_Player.SUBMSG", SUBMSG_name, SUBMSG_value)
	proto.RegisterEnum("MSG_Player.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("MSG_Player.EmBaseInfo", EmBaseInfo_name, EmBaseInfo_value)
	proto.RegisterEnum("MSG_Player.EmBaseMoney", EmBaseMoney_name, EmBaseMoney_value)
}

func init() { proto.RegisterFile("MSG_Player.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8f, 0x12, 0x41,
	0x10, 0x85, 0xb7, 0x99, 0x61, 0x85, 0x9a, 0x75, 0x2d, 0xca, 0xdd, 0x84, 0x18, 0x0f, 0x38, 0x07,
	0x45, 0x0e, 0x6b, 0x82, 0x37, 0x93, 0xbd, 0xcc, 0xec, 0xba, 0x62, 0xc4, 0x98, 0xee, 0x78, 0x9e,
	0x34, 0x50, 0x1a, 0x12, 0x98, 0x86, 0x1e, 0x24, 0xd9, 0x5f, 0xe6, 0xdf, 0x33, 0xdd, 0x13, 0x60,
	0xda, 0x98, 0x78, 0xeb, 0xaa, 0x54, 0xbd, 0xea, 0xf7, 0xe5, 0x01, 0x4e, 0xd5, 0x43, 0xf1, 0x6d,
	0xa5, 0x1f, 0xd9, 0xde, 0x6c, 0xac, 0xd9, 0x19, 0x82, 0x53, 0x27, 0xdd, 0xc2, 0x75, 0xa6, 0x2b,
	0xae, 0xab, 0x49, 0xf9, 0xc3, 0x14, 0x92, 0x37, 0xac, 0x77, 0x44, 0x10, 0x7f, 0xd5, 0x6b, 0xee,
	0x8b, 0x81, 0x18, 0x76, 0xa5, 0x7f, 0xd3, 0x15, 0xb4, 0xbf, 0xf0, 0x9e, 0x57, 0xfd, 0xd6, 0x40,
	0x0c, 0xdb, 0xb2, 0x2e, 0xe8, 0x05, 0x74, 0x3e, 0xb1, 0x5e, 0x4c, 0xe6, 0xa6, 0xec, 0x47, 0x7e,
	0xfa, 0x58, 0x3b, 0x95, 0xbb, 0x6c, 0x72, 0xd7, 0x8f, 0x07, 0x62, 0x18, 0x49, 0xff, 0x4e, 0x5f,
	0x03, 0xba, 0x93, 0x53, 0x53, 0xf2, 0x63, 0xe3, 0x5a, 0x6e, 0x96, 0xa5, 0xbf, 0x16, 0x49, 0xff,
	0x4e, 0xaf, 0x80, 0x72, 0x55, 0xdc, 0x97, 0x3b, 0xb6, 0x8a, 0xed, 0x9e, 0x6d, 0x21, 0x79, 0x9b,
	0xde, 0x02, 0xa9, 0x3c, 0xec, 0x56, 0x1b, 0x7a, 0x03, 0x91, 0xe4, 0x9d, 0x5f, 0xbf, 0x1c, 0x5f,
	0xdf, 0x34, 0x2c, 0xdf, 0x5b, 0x6b, 0x6c, 0x6e, 0x16, 0x2c, 0xdd, 0x44, 0xfa, 0x1c, 0x7a, 0xb9,
	0x2a, 0x02, 0xbb, 0xdb, 0xf4, 0xb7, 0x80, 0x9e, 0xca, 0x83, 0x6e, 0xb5, 0xa1, 0x5b, 0xe8, 0xcc,
	0x74, 0xc5, 0xae, 0xf6, 0xc2, 0xc9, 0xf8, 0x55, 0x53, 0xf8, 0x9f, 0xd8, 0xe4, 0x71, 0x85, 0x3e,
	0x40, 0x77, 0x76, 0xb0, 0xe9, 0x81, 0x25, 0xe3, 0x97, 0x7f, 0xef, 0x37, 0x19, 0xc8, 0xd3, 0xf8,
	0xc1, 0x4e, 0xf4, 0x3f, 0x3b, 0x23, 0x0d, 0xe7, 0xea, 0x7b, 0x36, 0x55, 0x0f, 0xd4, 0x85, 0x76,
	0xc6, 0x3f, 0x97, 0x25, 0x9e, 0x11, 0xc1, 0x65, 0x08, 0x0e, 0x85, 0xeb, 0x85, 0xd8, 0xb0, 0x45,
	0x3d, 0x78, 0x1a, 0xb0, 0xc0, 0xc8, 0xb5, 0x02, 0x10, 0x18, 0x8f, 0xde, 0x41, 0xf7, 0x78, 0x94,
	0x12, 0x78, 0x32, 0x29, 0xf7, 0x7a, 0xb5, 0x5c, 0xe0, 0x99, 0x2b, 0xd4, 0xaf, 0xf9, 0x9c, 0xab,
	0x0a, 0x05, 0x75, 0x20, 0xfe, 0xa8, 0x97, 0x2b, 0x6c, 0x8d, 0x3e, 0x03, 0xf0, 0x3a, 0x3b, 0x60,
	0x78, 0x06, 0x49, 0x8d, 0xe4, 0xb0, 0xd5, 0xa9, 0x83, 0x85, 0xc2, 0x7d, 0xd9, 0x27, 0x08, 0x5b,
	0x74, 0x71, 0xca, 0x10, 0x46, 0x6e, 0xc4, 0x25, 0x05, 0xe3, 0xd1, 0x5b, 0x48, 0x6a, 0xad, 0x9a,
	0x0b, 0xc2, 0xc5, 0xda, 0x3d, 0x02, 0x35, 0x17, 0x16, 0x14, 0xb3, 0x73, 0x1f, 0xee, 0xf7, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x99, 0x1b, 0xd2, 0xae, 0xf0, 0x02, 0x00, 0x00,
}
