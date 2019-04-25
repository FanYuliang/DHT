// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"math"
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

type GetParams struct {
	TransactionID 		 *string
	ServerIdentifier     *string  `protobuf:"bytes,1,req,name=server_identifier,json=serverIdentifier" json:"server_identifier,omitempty"`
	ObjectName           *string  `protobuf:"bytes,2,req,name=object_name,json=objectName" json:"object_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParams) Reset()         { *m = GetParams{} }
func (m *GetParams) String() string { return proto.CompactTextString(m) }
func (*GetParams) ProtoMessage()    {}
func (*GetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

func (m *GetParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParams.Unmarshal(m, b)
}
func (m *GetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParams.Marshal(b, m, deterministic)
}
func (m *GetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParams.Merge(m, src)
}
func (m *GetParams) XXX_Size() int {
	return xxx_messageInfo_GetParams.Size(m)
}
func (m *GetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetParams proto.InternalMessageInfo

func (m *GetParams) GetServerIdentifier() string {
	if m != nil && m.ServerIdentifier != nil {
		return *m.ServerIdentifier
	}
	return ""
}

func (m *GetParams) GetObjectName() string {
	if m != nil && m.ObjectName != nil {
		return *m.ObjectName
	}
	return ""
}

type SetParams struct {
	ServerIdentifier     *string  `protobuf:"bytes,1,req,name=server_identifier,json=serverIdentifier" json:"server_identifier,omitempty"`
	ObjectName           *string  `protobuf:"bytes,2,req,name=object_name,json=objectName" json:"object_name,omitempty"`
	Value                *string  `protobuf:"bytes,3,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetParams) Reset()         { *m = SetParams{} }
func (m *SetParams) String() string { return proto.CompactTextString(m) }
func (*SetParams) ProtoMessage()    {}
func (*SetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

func (m *SetParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetParams.Unmarshal(m, b)
}
func (m *SetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetParams.Marshal(b, m, deterministic)
}
func (m *SetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetParams.Merge(m, src)
}
func (m *SetParams) XXX_Size() int {
	return xxx_messageInfo_SetParams.Size(m)
}
func (m *SetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SetParams.DiscardUnknown(m)
}

var xxx_messageInfo_SetParams proto.InternalMessageInfo

func (m *SetParams) GetServerIdentifier() string {
	if m != nil && m.ServerIdentifier != nil {
		return *m.ServerIdentifier
	}
	return ""
}

func (m *SetParams) GetObjectName() string {
	if m != nil && m.ObjectName != nil {
		return *m.ObjectName
	}
	return ""
}

func (m *SetParams) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GetParams)(nil), "protos.GetParams")
	proto.RegisterType((*SetParams)(nil), "protos.SetParams")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x90, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0xaf, 0xd7, 0x1f, 0xc8, 0xe8, 0xe2, 0x36, 0xba, 0x28, 0xdd, 0x28, 0x59, 0x09, 0x42,
	0x17, 0x05, 0x9f, 0x40, 0xb0, 0xb8, 0x29, 0x62, 0xdd, 0xb8, 0x2a, 0xd3, 0x64, 0x84, 0x68, 0x93,
	0xd1, 0x24, 0xf6, 0xf9, 0x85, 0x86, 0x16, 0x5f, 0xc0, 0xd5, 0x30, 0xdf, 0x99, 0x33, 0xcc, 0x1c,
	0x00, 0xcf, 0x86, 0xea, 0xaf, 0xc0, 0x89, 0xe5, 0xd9, 0x52, 0x62, 0x75, 0xa1, 0xd9, 0x39, 0xf6,
	0x99, 0xaa, 0x37, 0x10, 0x2d, 0xa5, 0x67, 0x0c, 0xe8, 0xa2, 0xbc, 0x83, 0x22, 0x52, 0x98, 0x29,
	0x0c, 0xd6, 0x90, 0x4f, 0xf6, 0xdd, 0x52, 0x28, 0x8f, 0x6e, 0xf6, 0xb7, 0xe2, 0xe5, 0x90, 0x85,
	0xa7, 0x8d, 0xcb, 0x6b, 0x38, 0xe7, 0xf1, 0x83, 0x74, 0x1a, 0x3c, 0x3a, 0x2a, 0xf7, 0xcb, 0x18,
	0x64, 0xd4, 0xa1, 0x23, 0xc5, 0x20, 0xfa, 0xff, 0x59, 0x2d, 0xaf, 0xe0, 0x74, 0xc6, 0xe9, 0x87,
	0xca, 0xe3, 0x45, 0xca, 0x4d, 0xf3, 0x0d, 0x27, 0x1d, 0x1b, 0x92, 0x0d, 0x88, 0x87, 0xc9, 0x92,
	0x4f, 0x3d, 0x25, 0x59, 0xe4, 0x47, 0x63, 0xbd, 0xdd, 0x52, 0x1d, 0x56, 0xf4, 0x48, 0x64, 0x46,
	0xd4, 0x9f, 0x6a, 0x27, 0xef, 0x57, 0x4f, 0xfb, 0xd7, 0xb3, 0x45, 0x53, 0x5d, 0xae, 0xe8, 0x35,
	0xa0, 0x8f, 0xa8, 0x93, 0x65, 0xaf, 0x76, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0x83, 0xb5,
	0x50, 0x61, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	ClientSet(ctx context.Context, in *SetParams, opts ...grpc.CallOption) (*Feedback, error)
	ClientGet(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Transaction, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) ClientSet(ctx context.Context, in *SetParams, opts ...grpc.CallOption) (*Feedback, error) {
	out := new(Feedback)
	err := c.cc.Invoke(ctx, "/protos.Node/ClientSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ClientGet(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/protos.Node/ClientGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_ClientSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ClientSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Node/ClientSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ClientSet(ctx, req.(*SetParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ClientGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ClientGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Node/ClientGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ClientGet(ctx, req.(*GetParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientSet",
			Handler:    _Node_ClientSet_Handler,
		},
		{
			MethodName: "ClientGet",
			Handler:    _Node_ClientGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}
