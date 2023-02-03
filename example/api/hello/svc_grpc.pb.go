// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: hello/svc.proto

package hello

import (
	context "context"
	v1 "github.com/sarafanfm/mtserver/example/api/hello/v1"
	v2 "github.com/sarafanfm/mtserver/example/api/hello/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// V1Client is the client API for V1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type V1Client interface {
	SayHello(ctx context.Context, in *v1.Request, opts ...grpc.CallOption) (*v1.Response, error)
}

type v1Client struct {
	cc grpc.ClientConnInterface
}

func NewV1Client(cc grpc.ClientConnInterface) V1Client {
	return &v1Client{cc}
}

func (c *v1Client) SayHello(ctx context.Context, in *v1.Request, opts ...grpc.CallOption) (*v1.Response, error) {
	out := new(v1.Response)
	err := c.cc.Invoke(ctx, "/hello.V1/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V1Server is the server API for V1 service.
// All implementations must embed UnimplementedV1Server
// for forward compatibility
type V1Server interface {
	SayHello(context.Context, *v1.Request) (*v1.Response, error)
	mustEmbedUnimplementedV1Server()
}

// UnimplementedV1Server must be embedded to have forward compatible implementations.
type UnimplementedV1Server struct {
}

func (UnimplementedV1Server) SayHello(context.Context, *v1.Request) (*v1.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedV1Server) mustEmbedUnimplementedV1Server() {}

// UnsafeV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to V1Server will
// result in compilation errors.
type UnsafeV1Server interface {
	mustEmbedUnimplementedV1Server()
}

func RegisterV1Server(s grpc.ServiceRegistrar, srv V1Server) {
	s.RegisterService(&V1_ServiceDesc, srv)
}

func _V1_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.V1/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1Server).SayHello(ctx, req.(*v1.Request))
	}
	return interceptor(ctx, in, info, handler)
}

// V1_ServiceDesc is the grpc.ServiceDesc for V1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var V1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.V1",
	HandlerType: (*V1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _V1_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/svc.proto",
}

// V2Client is the client API for V2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type V2Client interface {
	SayHello(ctx context.Context, in *v2.Request, opts ...grpc.CallOption) (*v2.Response, error)
	NotifyHello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (V2_NotifyHelloClient, error)
	ThrowError(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type v2Client struct {
	cc grpc.ClientConnInterface
}

func NewV2Client(cc grpc.ClientConnInterface) V2Client {
	return &v2Client{cc}
}

func (c *v2Client) SayHello(ctx context.Context, in *v2.Request, opts ...grpc.CallOption) (*v2.Response, error) {
	out := new(v2.Response)
	err := c.cc.Invoke(ctx, "/hello.V2/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v2Client) NotifyHello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (V2_NotifyHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &V2_ServiceDesc.Streams[0], "/hello.V2/NotifyHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &v2NotifyHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type V2_NotifyHelloClient interface {
	Recv() (*v2.Response, error)
	grpc.ClientStream
}

type v2NotifyHelloClient struct {
	grpc.ClientStream
}

func (x *v2NotifyHelloClient) Recv() (*v2.Response, error) {
	m := new(v2.Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *v2Client) ThrowError(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hello.V2/ThrowError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V2Server is the server API for V2 service.
// All implementations must embed UnimplementedV2Server
// for forward compatibility
type V2Server interface {
	SayHello(context.Context, *v2.Request) (*v2.Response, error)
	NotifyHello(*emptypb.Empty, V2_NotifyHelloServer) error
	ThrowError(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedV2Server()
}

// UnimplementedV2Server must be embedded to have forward compatible implementations.
type UnimplementedV2Server struct {
}

func (UnimplementedV2Server) SayHello(context.Context, *v2.Request) (*v2.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedV2Server) NotifyHello(*emptypb.Empty, V2_NotifyHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method NotifyHello not implemented")
}
func (UnimplementedV2Server) ThrowError(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThrowError not implemented")
}
func (UnimplementedV2Server) mustEmbedUnimplementedV2Server() {}

// UnsafeV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to V2Server will
// result in compilation errors.
type UnsafeV2Server interface {
	mustEmbedUnimplementedV2Server()
}

func RegisterV2Server(s grpc.ServiceRegistrar, srv V2Server) {
	s.RegisterService(&V2_ServiceDesc, srv)
}

func _V2_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V2Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.V2/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V2Server).SayHello(ctx, req.(*v2.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _V2_NotifyHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(V2Server).NotifyHello(m, &v2NotifyHelloServer{stream})
}

type V2_NotifyHelloServer interface {
	Send(*v2.Response) error
	grpc.ServerStream
}

type v2NotifyHelloServer struct {
	grpc.ServerStream
}

func (x *v2NotifyHelloServer) Send(m *v2.Response) error {
	return x.ServerStream.SendMsg(m)
}

func _V2_ThrowError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V2Server).ThrowError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.V2/ThrowError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V2Server).ThrowError(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// V2_ServiceDesc is the grpc.ServiceDesc for V2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var V2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.V2",
	HandlerType: (*V2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _V2_SayHello_Handler,
		},
		{
			MethodName: "ThrowError",
			Handler:    _V2_ThrowError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NotifyHello",
			Handler:       _V2_NotifyHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hello/svc.proto",
}
