// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/ready.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc/ready.proto

It has these top-level messages:
	ReadyRequest
	ReadyResponse
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ReadyRequest struct {
}

func (m *ReadyRequest) Reset()                    { *m = ReadyRequest{} }
func (m *ReadyRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadyRequest) ProtoMessage()               {}
func (*ReadyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReadyResponse struct {
	Ready bool `protobuf:"varint,1,opt,name=ready" json:"ready,omitempty"`
}

func (m *ReadyResponse) Reset()                    { *m = ReadyResponse{} }
func (m *ReadyResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadyResponse) ProtoMessage()               {}
func (*ReadyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadyResponse) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

func init() {
	proto.RegisterType((*ReadyRequest)(nil), "grpc.ReadyRequest")
	proto.RegisterType((*ReadyResponse)(nil), "grpc.ReadyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for SRP service

type SRPClient interface {
	Ready(ctx context.Context, in *ReadyRequest, opts ...grpc1.CallOption) (*ReadyResponse, error)
}

type sRPClient struct {
	cc *grpc1.ClientConn
}

func NewSRPClient(cc *grpc1.ClientConn) SRPClient {
	return &sRPClient{cc}
}

func (c *sRPClient) Ready(ctx context.Context, in *ReadyRequest, opts ...grpc1.CallOption) (*ReadyResponse, error) {
	out := new(ReadyResponse)
	err := grpc1.Invoke(ctx, "/grpc.SRP/Ready", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SRP service

type SRPServer interface {
	Ready(context.Context, *ReadyRequest) (*ReadyResponse, error)
}

func RegisterSRPServer(s *grpc1.Server, srv SRPServer) {
	s.RegisterService(&_SRP_serviceDesc, srv)
}

func _SRP_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SRPServer).Ready(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SRP/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SRPServer).Ready(ctx, req.(*ReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SRP_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.SRP",
	HandlerType: (*SRPServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Ready",
			Handler:    _SRP_Ready_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "grpc/ready.proto",
}

func init() { proto.RegisterFile("grpc/ready.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2f, 0x2a, 0x48,
	0xd6, 0x2f, 0x4a, 0x4d, 0x4c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89,
	0x28, 0xf1, 0x71, 0xf1, 0x04, 0x81, 0x04, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x54,
	0xb9, 0x78, 0xa1, 0xfc, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x11, 0x2e, 0x56, 0xb0, 0x2e,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x08, 0xc7, 0xc8, 0x92, 0x8b, 0x39, 0x38, 0x28, 0x40,
	0xc8, 0x88, 0x8b, 0x15, 0xac, 0x5a, 0x48, 0x48, 0x0f, 0x64, 0x9a, 0x1e, 0xb2, 0x51, 0x52, 0xc2,
	0x28, 0x62, 0x10, 0xe3, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0xd6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x71, 0x70, 0x14, 0x08, 0x92, 0x00, 0x00, 0x00,
}
