// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package protos

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Transaction struct {
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type Feedback struct {
	Message              *string  `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feedback) Reset()         { *m = Feedback{} }
func (m *Feedback) String() string { return proto.CompactTextString(m) }
func (*Feedback) ProtoMessage()    {}
func (*Feedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *Feedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feedback.Unmarshal(m, b)
}
func (m *Feedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feedback.Marshal(b, m, deterministic)
}
func (m *Feedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feedback.Merge(m, src)
}
func (m *Feedback) XXX_Size() int {
	return xxx_messageInfo_Feedback.Size(m)
}
func (m *Feedback) XXX_DiscardUnknown() {
	xxx_messageInfo_Feedback.DiscardUnknown(m)
}

var xxx_messageInfo_Feedback proto.InternalMessageInfo

func (m *Feedback) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "protos.Empty")
	proto.RegisterType((*Transaction)(nil), "protos.Transaction")
	proto.RegisterType((*Feedback)(nil), "protos.Feedback")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xec, 0x5c,
	0xac, 0xae, 0xb9, 0x05, 0x25, 0x95, 0x4a, 0xb2, 0x5c, 0xdc, 0x21, 0x45, 0x89, 0x79, 0xc5, 0x89,
	0xc9, 0x25, 0x99, 0xf9, 0x79, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x4c, 0x1a,
	0x9c, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0x2a, 0x5c, 0x1c, 0x6e, 0xa9, 0xa9, 0x29, 0x49, 0x89, 0xc9,
	0xd9, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x50, 0x05, 0x30, 0x2e,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xca, 0xa5, 0xd5, 0x52, 0x64, 0x00, 0x00, 0x00,
}
