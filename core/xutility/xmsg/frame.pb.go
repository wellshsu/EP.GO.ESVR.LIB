//go:binary-only-package
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frame.proto
package xmsg

import (
	_ "fmt"
	_ "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
// please upgrade the proto package
type Frame struct {
	SrcUrl               *string  `protobuf:"bytes,1,req,name=SrcUrl" json:"SrcUrl,omitempty"`
	DstUrl               *string  `protobuf:"bytes,2,req,name=DstUrl" json:"DstUrl,omitempty"`
	MapIDs               *MapID   `protobuf:"bytes,3,opt,name=MapIDs" json:"MapIDs,omitempty"`
	MapUrls              *MapUrl  `protobuf:"bytes,4,opt,name=MapUrls" json:"MapUrls,omitempty"`
	UID                  *int32   `protobuf:"varint,5,opt,name=UID,def=-1" json:"UID,omitempty"`
	GoID                 *int32   `protobuf:"varint,6,opt,name=GoID,def=-1" json:"GoID,omitempty"`
	RawData              []byte   `protobuf:"bytes,7,opt,name=RawData" json:"RawData,omitempty"`
	RpcID                *int64   `protobuf:"varint,8,opt,name=RpcID,def=-1" json:"RpcID,omitempty"`
	RpcGo                *int32   `protobuf:"varint,9,opt,name=RpcGo,def=-1" json:"RpcGo,omitempty"`
	Rpc                  *bool    `protobuf:"varint,10,opt,name=Rpc,def=0" json:"Rpc,omitempty"`
	Next                 *bool    `protobuf:"varint,11,opt,name=Next,def=0" json:"Next,omitempty"`
	RW                   *bool    `protobuf:"varint,12,opt,name=RW,def=1" json:"RW,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Frame) Reset()         { return }
func (m *Frame) String() string { return "" }
func (*Frame) ProtoMessage()    { return }
func (*Frame) Descriptor() ([]byte, []int) {
	return []byte{}, []int{}
}
func (m *Frame) XXX_Unmarshal(b []byte) error {
	return nil
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return []byte{}, nil
}
func (m *Frame) XXX_Merge(src proto.Message) {
	return
}
func (m *Frame) XXX_Size() int {
	return 0
}
func (m *Frame) XXX_DiscardUnknown() {
	return
}

const Default_Frame_UID int32 = -1
const Default_Frame_GoID int32 = -1
const Default_Frame_RpcID int64 = -1
const Default_Frame_RpcGo int32 = -1
const Default_Frame_Rpc bool = false
const Default_Frame_Next bool = false
const Default_Frame_RW bool = true

func (m *Frame) GetSrcUrl() string {
	return ""
}
func (m *Frame) GetDstUrl() string {
	return ""
}
func (m *Frame) GetMapIDs() *MapID {
	return nil
}
func (m *Frame) GetMapUrls() *MapUrl {
	return nil
}
func (m *Frame) GetUID() int32 {
	return 0
}
func (m *Frame) GetGoID() int32 {
	return 0
}
func (m *Frame) GetRawData() []byte {
	return []byte{}
}
func (m *Frame) GetRpcID() int64 {
	return 0
}
func (m *Frame) GetRpcGo() int32 {
	return 0
}
func (m *Frame) GetRpc() bool {
	return false
}
func (m *Frame) GetNext() bool {
	return false
}
func (m *Frame) GetRW() bool {
	return false
}

type MapID struct {
	ID1                  *int32   `protobuf:"varint,1,opt,name=ID1" json:"ID1,omitempty"`
	ID2                  *int32   `protobuf:"varint,2,opt,name=ID2" json:"ID2,omitempty"`
	ID3                  *int32   `protobuf:"varint,3,opt,name=ID3" json:"ID3,omitempty"`
	ID4                  *int32   `protobuf:"varint,4,opt,name=ID4" json:"ID4,omitempty"`
	ID5                  *int32   `protobuf:"varint,5,opt,name=ID5" json:"ID5,omitempty"`
	ID6                  *int32   `protobuf:"varint,6,opt,name=ID6" json:"ID6,omitempty"`
	ID7                  *int32   `protobuf:"varint,7,opt,name=ID7" json:"ID7,omitempty"`
	ID8                  *int32   `protobuf:"varint,8,opt,name=ID8" json:"ID8,omitempty"`
	ID9                  *int32   `protobuf:"varint,9,opt,name=ID9" json:"ID9,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapID) Reset()         { return }
func (m *MapID) String() string { return "" }
func (*MapID) ProtoMessage()    { return }
func (*MapID) Descriptor() ([]byte, []int) {
	return []byte{}, []int{}
}
func (m *MapID) XXX_Unmarshal(b []byte) error {
	return nil
}
func (m *MapID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return []byte{}, nil
}
func (m *MapID) XXX_Merge(src proto.Message) {
	return
}
func (m *MapID) XXX_Size() int {
	return 0
}
func (m *MapID) XXX_DiscardUnknown() {
	return
}
func (m *MapID) GetID1() int32 {
	return 0
}
func (m *MapID) GetID2() int32 {
	return 0
}
func (m *MapID) GetID3() int32 {
	return 0
}
func (m *MapID) GetID4() int32 {
	return 0
}
func (m *MapID) GetID5() int32 {
	return 0
}
func (m *MapID) GetID6() int32 {
	return 0
}
func (m *MapID) GetID7() int32 {
	return 0
}
func (m *MapID) GetID8() int32 {
	return 0
}
func (m *MapID) GetID9() int32 {
	return 0
}

type MapUrl struct {
	Url1                 *string  `protobuf:"bytes,1,opt,name=Url1" json:"Url1,omitempty"`
	Url2                 *string  `protobuf:"bytes,2,opt,name=Url2" json:"Url2,omitempty"`
	Url3                 *string  `protobuf:"bytes,3,opt,name=Url3" json:"Url3,omitempty"`
	Url4                 *string  `protobuf:"bytes,4,opt,name=Url4" json:"Url4,omitempty"`
	Url5                 *string  `protobuf:"bytes,5,opt,name=Url5" json:"Url5,omitempty"`
	Url6                 *string  `protobuf:"bytes,6,opt,name=Url6" json:"Url6,omitempty"`
	Url7                 *string  `protobuf:"bytes,7,opt,name=Url7" json:"Url7,omitempty"`
	Url8                 *string  `protobuf:"bytes,8,opt,name=Url8" json:"Url8,omitempty"`
	Url9                 *string  `protobuf:"bytes,9,opt,name=Url9" json:"Url9,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapUrl) Reset()         { return }
func (m *MapUrl) String() string { return "" }
func (*MapUrl) ProtoMessage()    { return }
func (*MapUrl) Descriptor() ([]byte, []int) {
	return []byte{}, []int{}
}
func (m *MapUrl) XXX_Unmarshal(b []byte) error {
	return nil
}
func (m *MapUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return []byte{}, nil
}
func (m *MapUrl) XXX_Merge(src proto.Message) {
	return
}
func (m *MapUrl) XXX_Size() int {
	return 0
}
func (m *MapUrl) XXX_DiscardUnknown() {
	return
}
func (m *MapUrl) GetUrl1() string {
	return ""
}
func (m *MapUrl) GetUrl2() string {
	return ""
}
func (m *MapUrl) GetUrl3() string {
	return ""
}
func (m *MapUrl) GetUrl4() string {
	return ""
}
func (m *MapUrl) GetUrl5() string {
	return ""
}
func (m *MapUrl) GetUrl6() string {
	return ""
}
func (m *MapUrl) GetUrl7() string {
	return ""
}
func (m *MapUrl) GetUrl8() string {
	return ""
}
func (m *MapUrl) GetUrl9() string {
	return ""
}
func init() {
	return
}
func init() { return }