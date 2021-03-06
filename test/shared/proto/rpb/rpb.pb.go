// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpb.proto

package rpb

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

type RPC_ConnNotifyOfflineReq struct {
	UID                  *int32   `protobuf:"varint,1,req,name=UID" json:"UID,omitempty"`
	Url                  *string  `protobuf:"bytes,2,req,name=Url" json:"Url,omitempty"`
	CID                  *int64   `protobuf:"varint,3,req,name=CID" json:"CID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPC_ConnNotifyOfflineReq) Reset()         { *m = RPC_ConnNotifyOfflineReq{} }
func (m *RPC_ConnNotifyOfflineReq) String() string { return proto.CompactTextString(m) }
func (*RPC_ConnNotifyOfflineReq) ProtoMessage()    {}
func (*RPC_ConnNotifyOfflineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_133edf69513fc562, []int{0}
}

func (m *RPC_ConnNotifyOfflineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPC_ConnNotifyOfflineReq.Unmarshal(m, b)
}
func (m *RPC_ConnNotifyOfflineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPC_ConnNotifyOfflineReq.Marshal(b, m, deterministic)
}
func (m *RPC_ConnNotifyOfflineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPC_ConnNotifyOfflineReq.Merge(m, src)
}
func (m *RPC_ConnNotifyOfflineReq) XXX_Size() int {
	return xxx_messageInfo_RPC_ConnNotifyOfflineReq.Size(m)
}
func (m *RPC_ConnNotifyOfflineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RPC_ConnNotifyOfflineReq.DiscardUnknown(m)
}

var xxx_messageInfo_RPC_ConnNotifyOfflineReq proto.InternalMessageInfo

func (m *RPC_ConnNotifyOfflineReq) GetUID() int32 {
	if m != nil && m.UID != nil {
		return *m.UID
	}
	return 0
}

func (m *RPC_ConnNotifyOfflineReq) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *RPC_ConnNotifyOfflineReq) GetCID() int64 {
	if m != nil && m.CID != nil {
		return *m.CID
	}
	return 0
}

type RPC_GetOnlineFromConnResp struct {
	ID                   []int32  `protobuf:"varint,1,rep,name=ID" json:"ID"`
	Url                  []string `protobuf:"bytes,2,rep,name=Url" json:"Url"`
	CID                  []int64  `protobuf:"varint,3,rep,name=CID" json:"CID"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPC_GetOnlineFromConnResp) Reset()         { *m = RPC_GetOnlineFromConnResp{} }
func (m *RPC_GetOnlineFromConnResp) String() string { return proto.CompactTextString(m) }
func (*RPC_GetOnlineFromConnResp) ProtoMessage()    {}
func (*RPC_GetOnlineFromConnResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_133edf69513fc562, []int{1}
}

func (m *RPC_GetOnlineFromConnResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPC_GetOnlineFromConnResp.Unmarshal(m, b)
}
func (m *RPC_GetOnlineFromConnResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPC_GetOnlineFromConnResp.Marshal(b, m, deterministic)
}
func (m *RPC_GetOnlineFromConnResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPC_GetOnlineFromConnResp.Merge(m, src)
}
func (m *RPC_GetOnlineFromConnResp) XXX_Size() int {
	return xxx_messageInfo_RPC_GetOnlineFromConnResp.Size(m)
}
func (m *RPC_GetOnlineFromConnResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RPC_GetOnlineFromConnResp.DiscardUnknown(m)
}

var xxx_messageInfo_RPC_GetOnlineFromConnResp proto.InternalMessageInfo

func (m *RPC_GetOnlineFromConnResp) GetID() []int32 {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *RPC_GetOnlineFromConnResp) GetUrl() []string {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *RPC_GetOnlineFromConnResp) GetCID() []int64 {
	if m != nil {
		return m.CID
	}
	return nil
}

func init() {
	proto.RegisterType((*RPC_ConnNotifyOfflineReq)(nil), "rpb.RPC_ConnNotifyOfflineReq")
	proto.RegisterType((*RPC_GetOnlineFromConnResp)(nil), "rpb.RPC_GetOnlineFromConnResp")
}

func init() { proto.RegisterFile("rpb.proto", fileDescriptor_133edf69513fc562) }

var fileDescriptor_133edf69513fc562 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x52, 0x72, 0xe4, 0x92, 0x08, 0x0a,
	0x70, 0x8e, 0x77, 0xce, 0xcf, 0xcb, 0xf3, 0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0xf4, 0x4f, 0x4b, 0xcb,
	0xc9, 0xcc, 0x4b, 0x0d, 0x4a, 0x2d, 0x14, 0xe2, 0xe6, 0x62, 0x0e, 0xf5, 0x74, 0x91, 0x60, 0x54,
	0x60, 0xd2, 0x60, 0x05, 0x73, 0x8a, 0x72, 0x24, 0x98, 0x14, 0x98, 0x34, 0x38, 0x41, 0x1c, 0x67,
	0x4f, 0x17, 0x09, 0x66, 0x05, 0x26, 0x0d, 0x66, 0x25, 0x47, 0x2e, 0x49, 0x90, 0x11, 0xee, 0xa9,
	0x25, 0xfe, 0x79, 0x20, 0xad, 0x6e, 0x45, 0xf9, 0xb9, 0x20, 0xf3, 0x82, 0x52, 0x8b, 0x0b, 0x84,
	0xb8, 0xb8, 0x98, 0xc0, 0x46, 0x30, 0x23, 0x1b, 0xc1, 0x8c, 0x6c, 0x04, 0xb3, 0x06, 0x33, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x01, 0x19, 0x7b, 0x85, 0x96, 0x00, 0x00, 0x00,
}

