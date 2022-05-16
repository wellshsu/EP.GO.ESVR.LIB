// Code generated by protoc-gen-go. DO NOT EDIT.
// source: epb.proto

package protocol

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

// 请求结果通知
type GM_Common struct {
	Result               *int32   `protobuf:"varint,1,req,name=Result" json:"Result,omitempty"`
	Params               []string `protobuf:"bytes,2,rep,name=Params" json:"Params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GM_Common) Reset()         { *m = GM_Common{} }
func (m *GM_Common) String() string { return proto.CompactTextString(m) }
func (*GM_Common) ProtoMessage()    {}
func (*GM_Common) Descriptor() ([]byte, []int) {
	return fileDescriptor_8633840523c387f8, []int{0}
}

func (m *GM_Common) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GM_Common.Unmarshal(m, b)
}
func (m *GM_Common) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GM_Common.Marshal(b, m, deterministic)
}
func (m *GM_Common) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GM_Common.Merge(m, src)
}
func (m *GM_Common) XXX_Size() int {
	return xxx_messageInfo_GM_Common.Size(m)
}
func (m *GM_Common) XXX_DiscardUnknown() {
	xxx_messageInfo_GM_Common.DiscardUnknown(m)
}

var xxx_messageInfo_GM_Common proto.InternalMessageInfo

func (m *GM_Common) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *GM_Common) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

// 请求登录
type GM_LoginReq struct {
	Account              *string  `protobuf:"bytes,1,req,name=Account" json:"Account,omitempty"`
	Password             *string  `protobuf:"bytes,2,req,name=Password" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GM_LoginReq) Reset()         { *m = GM_LoginReq{} }
func (m *GM_LoginReq) String() string { return proto.CompactTextString(m) }
func (*GM_LoginReq) ProtoMessage()    {}
func (*GM_LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8633840523c387f8, []int{1}
}

func (m *GM_LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GM_LoginReq.Unmarshal(m, b)
}
func (m *GM_LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GM_LoginReq.Marshal(b, m, deterministic)
}
func (m *GM_LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GM_LoginReq.Merge(m, src)
}
func (m *GM_LoginReq) XXX_Size() int {
	return xxx_messageInfo_GM_LoginReq.Size(m)
}
func (m *GM_LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GM_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_GM_LoginReq proto.InternalMessageInfo

func (m *GM_LoginReq) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *GM_LoginReq) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

// 响应登录
type GM_LoginResp struct {
	ID                   *int32   `protobuf:"varint,1,req,name=ID" json:"ID,omitempty"`
	Status               *int32   `protobuf:"varint,2,opt,name=Status" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GM_LoginResp) Reset()         { *m = GM_LoginResp{} }
func (m *GM_LoginResp) String() string { return proto.CompactTextString(m) }
func (*GM_LoginResp) ProtoMessage()    {}
func (*GM_LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8633840523c387f8, []int{2}
}

func (m *GM_LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GM_LoginResp.Unmarshal(m, b)
}
func (m *GM_LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GM_LoginResp.Marshal(b, m, deterministic)
}
func (m *GM_LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GM_LoginResp.Merge(m, src)
}
func (m *GM_LoginResp) XXX_Size() int {
	return xxx_messageInfo_GM_LoginResp.Size(m)
}
func (m *GM_LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GM_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_GM_LoginResp proto.InternalMessageInfo

func (m *GM_LoginResp) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *GM_LoginResp) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*GM_Common)(nil), "protocol.GM_Common")
	proto.RegisterType((*GM_LoginReq)(nil), "protocol.GM_LoginReq")
	proto.RegisterType((*GM_LoginResp)(nil), "protocol.GM_LoginResp")
}

func init() { proto.RegisterFile("epb.proto", fileDescriptor_8633840523c387f8) }

var fileDescriptor_8633840523c387f8 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2d, 0x48, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0xda, 0x5c, 0x9c,
	0xee, 0xbe, 0xf1, 0xce, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x42, 0x7c, 0x5c, 0x6c, 0x41, 0xa9, 0xc5,
	0xa5, 0x39, 0x25, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0xac, 0x20, 0x7e, 0x40, 0x62, 0x51, 0x62, 0x6e,
	0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xa7, 0x92, 0x01, 0x17, 0xb7, 0xbb, 0x6f, 0xbc, 0x4f, 0x7e,
	0x7a, 0x66, 0x5e, 0x50, 0x6a, 0xa1, 0x10, 0x3f, 0x17, 0xbb, 0x63, 0x72, 0x72, 0x7e, 0x69, 0x1e,
	0x44, 0x3d, 0xa7, 0x90, 0x00, 0x17, 0x47, 0x40, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04,
	0x13, 0x48, 0x44, 0x49, 0x8b, 0x8b, 0x07, 0xa1, 0xa3, 0xb8, 0x40, 0x88, 0x8b, 0x8b, 0xc9, 0xd3,
	0x05, 0x61, 0x7a, 0x70, 0x49, 0x62, 0x49, 0x29, 0xc8, 0x74, 0x46, 0x0d, 0x56, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x33, 0x43, 0x7e, 0x0c, 0xa0, 0x00, 0x00, 0x00,
}
