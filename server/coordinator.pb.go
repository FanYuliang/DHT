// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coordinator.proto

//protoc -I . *.proto --go_out=plugins=grpc:.

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

func init() { proto.RegisterFile("coordinator.proto", fileDescriptor_99e779eb11ceee19) }

var fileDescriptor_99e779eb11ceee19 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0xcf, 0x2f,
	0x4a, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03,
	0x53, 0xc5, 0x52, 0x3c, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x10, 0x51, 0xa3, 0x4f, 0x8c, 0x5c,
	0xdc, 0xce, 0x08, 0xb5, 0x42, 0xe6, 0x5c, 0xfc, 0xfe, 0x05, 0xa9, 0x79, 0x21, 0x45, 0x89, 0x79,
	0xc5, 0x89, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0x42, 0xbc, 0x10, 0xa5, 0xc5, 0x7a, 0xae, 0xb9, 0x05,
	0x25, 0x95, 0x52, 0xc2, 0x30, 0x2e, 0x92, 0x1a, 0x25, 0x06, 0x21, 0x6b, 0x2e, 0x01, 0xe7, 0x9c,
	0xfc, 0xe2, 0x54, 0x64, 0x9d, 0xd8, 0x94, 0x4a, 0x09, 0xc0, 0x04, 0xdd, 0x52, 0x53, 0x53, 0x92,
	0x12, 0x93, 0xb3, 0x95, 0x18, 0x84, 0xcc, 0xb8, 0x04, 0x9d, 0xf3, 0x73, 0x73, 0x33, 0x4b, 0xf0,
	0xd8, 0x8b, 0x4d, 0x9f, 0x29, 0x97, 0x80, 0x63, 0x52, 0x7e, 0x11, 0x89, 0xda, 0x00, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x6e, 0x89, 0x88, 0x41, 0x1e, 0x01, 0x00, 0x00,
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
	CommitTransaction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Feedback, error)
	AbortTransaction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Feedback, error)
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

func (c *coordinatorClient) CommitTransaction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Feedback, error) {
	out := new(Feedback)
	err := c.cc.Invoke(ctx, "/protos.Coordinator/CommitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) AbortTransaction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Feedback, error) {
	out := new(Feedback)
	err := c.cc.Invoke(ctx, "/protos.Coordinator/AbortTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

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

func _Coordinator_CommitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CommitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Coordinator/CommitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CommitTransaction(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_AbortTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).AbortTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Coordinator/AbortTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).AbortTransaction(ctx, req.(*Empty))
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
			MethodName: "CommitTransaction",
			Handler:    _Coordinator_CommitTransaction_Handler,
		},
		{
			MethodName: "AbortTransaction",
			Handler:    _Coordinator_AbortTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coordinator.proto",
}