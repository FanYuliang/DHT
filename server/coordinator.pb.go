// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coordinator.proto

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

type TryLockParam struct {
	TransactionID        *string  `protobuf:"bytes,1,req,name=transactionID" json:"transactionID,omitempty"`
	ServerIdentifier     *string  `protobuf:"bytes,2,req,name=server_identifier,json=serverIdentifier" json:"server_identifier,omitempty"`
	Object               *string  `protobuf:"bytes,3,req,name=object" json:"object,omitempty"`
	LockType             *string  `protobuf:"bytes,4,req,name=lockType" json:"lockType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TryLockParam) Reset()         { *m = TryLockParam{} }
func (m *TryLockParam) String() string { return proto.CompactTextString(m) }
func (*TryLockParam) ProtoMessage()    {}
func (*TryLockParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e779eb11ceee19, []int{0}
}

func (m *TryLockParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TryLockParam.Unmarshal(m, b)
}
func (m *TryLockParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TryLockParam.Marshal(b, m, deterministic)
}
func (m *TryLockParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TryLockParam.Merge(m, src)
}
func (m *TryLockParam) XXX_Size() int {
	return xxx_messageInfo_TryLockParam.Size(m)
}
func (m *TryLockParam) XXX_DiscardUnknown() {
	xxx_messageInfo_TryLockParam.DiscardUnknown(m)
}

var xxx_messageInfo_TryLockParam proto.InternalMessageInfo

func (m *TryLockParam) GetTransactionID() string {
	if m != nil && m.TransactionID != nil {
		return *m.TransactionID
	}
	return ""
}

func (m *TryLockParam) GetServerIdentifier() string {
	if m != nil && m.ServerIdentifier != nil {
		return *m.ServerIdentifier
	}
	return ""
}

func (m *TryLockParam) GetObject() string {
	if m != nil && m.Object != nil {
		return *m.Object
	}
	return ""
}

func (m *TryLockParam) GetLockType() string {
	if m != nil && m.LockType != nil {
		return *m.LockType
	}
	return ""
}

type ReportUnLockParam struct {
	TransactionID        *string  `protobuf:"bytes,1,req,name=transactionID" json:"transactionID,omitempty"`
	ServerIdentifier     *string  `protobuf:"bytes,2,req,name=server_identifier,json=serverIdentifier" json:"server_identifier,omitempty"`
	Object               *string  `protobuf:"bytes,3,req,name=object" json:"object,omitempty"`
	LockType             *string  `protobuf:"bytes,4,req,name=lockType" json:"lockType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportUnLockParam) Reset()         { *m = ReportUnLockParam{} }
func (m *ReportUnLockParam) String() string { return proto.CompactTextString(m) }
func (*ReportUnLockParam) ProtoMessage()    {}
func (*ReportUnLockParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e779eb11ceee19, []int{1}
}

func (m *ReportUnLockParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportUnLockParam.Unmarshal(m, b)
}
func (m *ReportUnLockParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportUnLockParam.Marshal(b, m, deterministic)
}
func (m *ReportUnLockParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportUnLockParam.Merge(m, src)
}
func (m *ReportUnLockParam) XXX_Size() int {
	return xxx_messageInfo_ReportUnLockParam.Size(m)
}
func (m *ReportUnLockParam) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportUnLockParam.DiscardUnknown(m)
}

var xxx_messageInfo_ReportUnLockParam proto.InternalMessageInfo

func (m *ReportUnLockParam) GetTransactionID() string {
	if m != nil && m.TransactionID != nil {
		return *m.TransactionID
	}
	return ""
}

func (m *ReportUnLockParam) GetServerIdentifier() string {
	if m != nil && m.ServerIdentifier != nil {
		return *m.ServerIdentifier
	}
	return ""
}

func (m *ReportUnLockParam) GetObject() string {
	if m != nil && m.Object != nil {
		return *m.Object
	}
	return ""
}

func (m *ReportUnLockParam) GetLockType() string {
	if m != nil && m.LockType != nil {
		return *m.LockType
	}
	return ""
}

func init() {
	proto.RegisterType((*TryLockParam)(nil), "protos.TryLockParam")
	proto.RegisterType((*ReportUnLockParam)(nil), "protos.ReportUnLockParam")
}

func init() { proto.RegisterFile("coordinator.proto", fileDescriptor_99e779eb11ceee19) }

var fileDescriptor_99e779eb11ceee19 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdb, 0x2a, 0x55, 0xc7, 0x14, 0x93, 0x6d, 0x91, 0x98, 0x93, 0x04, 0x0f, 0x82, 0xd0,
	0x83, 0x1e, 0x04, 0x05, 0x25, 0x44, 0x85, 0x82, 0xa0, 0x94, 0x78, 0x96, 0xcd, 0x66, 0x85, 0x35,
	0xdd, 0x9d, 0xb0, 0xbb, 0x08, 0xf9, 0x21, 0x5e, 0xfc, 0x9f, 0xde, 0xc5, 0xa4, 0x69, 0x5a, 0xec,
	0xa9, 0x27, 0x4f, 0xcb, 0xbc, 0xf9, 0xde, 0x32, 0xbc, 0x19, 0xf0, 0x18, 0xa2, 0xce, 0x84, 0xa2,
	0x16, 0xf5, 0xb8, 0xd0, 0x68, 0x91, 0xf4, 0xab, 0xc7, 0x04, 0x0e, 0x43, 0x29, 0x51, 0xd5, 0x6a,
	0xf8, 0xd9, 0x05, 0x27, 0xd1, 0xe5, 0x23, 0xb2, 0xfc, 0x99, 0x6a, 0x2a, 0xc9, 0x09, 0x0c, 0xac,
	0xa6, 0xca, 0x50, 0x66, 0x05, 0xaa, 0xc9, 0x9d, 0xdf, 0x3d, 0xee, 0x9d, 0xee, 0x4d, 0x57, 0x45,
	0x72, 0x06, 0x9e, 0xe1, 0xfa, 0x83, 0xeb, 0x57, 0x91, 0x71, 0x65, 0xc5, 0x9b, 0xe0, 0xda, 0xef,
	0x55, 0xa4, 0x5b, 0x37, 0x26, 0x0b, 0x9d, 0x1c, 0x42, 0x1f, 0xd3, 0x77, 0xce, 0xac, 0xbf, 0x55,
	0x11, 0xf3, 0x8a, 0x04, 0xb0, 0x3b, 0x43, 0x96, 0x27, 0x65, 0xc1, 0xfd, 0xed, 0xaa, 0xb3, 0xa8,
	0xc3, 0xaf, 0x2e, 0x78, 0x53, 0x5e, 0xa0, 0xb6, 0x2f, 0xea, 0xbf, 0x0d, 0x77, 0xfe, 0xdd, 0x83,
	0xfd, 0xb8, 0x0d, 0x98, 0x5c, 0xc2, 0xc1, 0x53, 0xc1, 0x55, 0xd2, 0x4e, 0x41, 0x06, 0x75, 0xbe,
	0x66, 0x7c, 0x2f, 0x0b, 0x5b, 0x06, 0xc3, 0xa6, 0x5c, 0x62, 0xc2, 0x0e, 0xb9, 0x06, 0x37, 0x9e,
	0xa1, 0xe1, 0xcb, 0xce, 0x75, 0x68, 0xe0, 0x36, 0xe2, 0x03, 0xe7, 0x59, 0x4a, 0x59, 0x1e, 0x76,
	0xc8, 0x2d, 0x8c, 0x22, 0x93, 0xc7, 0x28, 0xa5, 0xb0, 0x1b, 0x7d, 0x70, 0x03, 0xc3, 0xc8, 0xe4,
	0x51, 0x8a, 0x7a, 0x33, 0xff, 0x05, 0xec, 0xcc, 0x4f, 0x87, 0x8c, 0x5a, 0x4f, 0x7b, 0x4b, 0x6b,
	0x4d, 0x57, 0xe0, 0x34, 0x7b, 0xfd, 0xcd, 0x93, 0x1c, 0x35, 0xcc, 0x9f, 0x6d, 0x07, 0xab, 0x19,
	0x86, 0x9d, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xd3, 0x16, 0x3e, 0xd6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinatorClient interface {
	OpenTransaction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Transaction, error)
	CloseTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Feedback, error)
	AskCommitTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Feedback, error)
	AskAbortTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Feedback, error)
	TryLock(ctx context.Context, in *TryLockParam, opts ...grpc.CallOption) (*Feedback, error)
	ReportUnlock(ctx context.Context, in *ReportUnLockParam, opts ...grpc.CallOption) (*Empty, error)
}

type coordinatorClient struct {
	cc *grpc.ClientConn
}

func NewCoordinatorClient(cc *grpc.ClientConn) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) OpenTransaction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/protos.Coordinator/OpenTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CloseTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Feedback, error) {
	out := new(Feedback)
	err := c.cc.Invoke(ctx, "/protos.Coordinator/CloseTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) AskCommitTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Feedback, error) {
	out := new(Feedback)
	err := c.cc.Invoke(ctx, "/protos.Coordinator/AskCommitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) AskAbortTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Feedback, error) {
	out := new(Feedback)
	err := c.cc.Invoke(ctx, "/protos.Coordinator/AskAbortTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) TryLock(ctx context.Context, in *TryLockParam, opts ...grpc.CallOption) (*Feedback, error) {
	out := new(Feedback)
	err := c.cc.Invoke(ctx, "/protos.Coordinator/TryLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) ReportUnlock(ctx context.Context, in *ReportUnLockParam, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protos.Coordinator/ReportUnlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
type CoordinatorServer interface {
	OpenTransaction(context.Context, *Empty) (*Transaction, error)
	CloseTransaction(context.Context, *Transaction) (*Feedback, error)
	AskCommitTransaction(context.Context, *Transaction) (*Feedback, error)
	AskAbortTransaction(context.Context, *Transaction) (*Feedback, error)
	TryLock(context.Context, *TryLockParam) (*Feedback, error)
	ReportUnlock(context.Context, *ReportUnLockParam) (*Empty, error)
}

//// UnimplementedCoordinatorServer can be embedded to have forward compatible implementations.
//type UnimplementedCoordinatorServer struct {
//}
//
//func (*UnimplementedCoordinatorServer) OpenTransaction(ctx context.Context, req *Empty) (*Transaction, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method OpenTransaction not implemented")
//}
//func (*UnimplementedCoordinatorServer) CloseTransaction(ctx context.Context, req *Transaction) (*Feedback, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method CloseTransaction not implemented")
//}
//func (*UnimplementedCoordinatorServer) AskCommitTransaction(ctx context.Context, req *Transaction) (*Feedback, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method AskCommitTransaction not implemented")
//}
//func (*UnimplementedCoordinatorServer) AskAbortTransaction(ctx context.Context, req *Transaction) (*Feedback, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method AskAbortTransaction not implemented")
//}
//func (*UnimplementedCoordinatorServer) TryLock(ctx context.Context, req *TryLockParam) (*Feedback, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method TryLock not implemented")
//}
//func (*UnimplementedCoordinatorServer) ReportUnlock(ctx context.Context, req *ReportUnLockParam) (*Empty, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ReportUnlock not implemented")
//}

func RegisterCoordinatorServer(s *grpc.Server, srv CoordinatorServer) {
	s.RegisterService(&_Coordinator_serviceDesc, srv)
}

func _Coordinator_OpenTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).OpenTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Coordinator/OpenTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).OpenTransaction(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CloseTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CloseTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Coordinator/CloseTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CloseTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_AskCommitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).AskCommitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Coordinator/AskCommitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).AskCommitTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_AskAbortTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).AskAbortTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Coordinator/AskAbortTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).AskAbortTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_TryLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TryLockParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).TryLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Coordinator/TryLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).TryLock(ctx, req.(*TryLockParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_ReportUnlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportUnLockParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).ReportUnlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Coordinator/ReportUnlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).ReportUnlock(ctx, req.(*ReportUnLockParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coordinator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenTransaction",
			Handler:    _Coordinator_OpenTransaction_Handler,
		},
		{
			MethodName: "CloseTransaction",
			Handler:    _Coordinator_CloseTransaction_Handler,
		},
		{
			MethodName: "AskCommitTransaction",
			Handler:    _Coordinator_AskCommitTransaction_Handler,
		},
		{
			MethodName: "AskAbortTransaction",
			Handler:    _Coordinator_AskAbortTransaction_Handler,
		},
		{
			MethodName: "TryLock",
			Handler:    _Coordinator_TryLock_Handler,
		},
		{
			MethodName: "ReportUnlock",
			Handler:    _Coordinator_ReportUnlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coordinator.proto",
}
