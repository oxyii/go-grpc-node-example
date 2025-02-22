// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: test.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Test_Unary_FullMethodName        = "/proto.Test/Unary"
	Test_ClientStream_FullMethodName = "/proto.Test/ClientStream"
	Test_ServerStream_FullMethodName = "/proto.Test/ServerStream"
	Test_BidiStream_FullMethodName   = "/proto.Test/BidiStream"
)

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestClient interface {
	Unary(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wrapperspb.StringValue, wrapperspb.StringValue], error)
	ServerStream(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (grpc.ServerStreamingClient[wrapperspb.StringValue], error)
	BidiStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[wrapperspb.StringValue, wrapperspb.StringValue], error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) Unary(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, Test_Unary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wrapperspb.StringValue, wrapperspb.StringValue], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Test_ServiceDesc.Streams[0], Test_ClientStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wrapperspb.StringValue, wrapperspb.StringValue]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Test_ClientStreamClient = grpc.ClientStreamingClient[wrapperspb.StringValue, wrapperspb.StringValue]

func (c *testClient) ServerStream(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (grpc.ServerStreamingClient[wrapperspb.StringValue], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Test_ServiceDesc.Streams[1], Test_ServerStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wrapperspb.StringValue, wrapperspb.StringValue]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Test_ServerStreamClient = grpc.ServerStreamingClient[wrapperspb.StringValue]

func (c *testClient) BidiStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[wrapperspb.StringValue, wrapperspb.StringValue], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Test_ServiceDesc.Streams[2], Test_BidiStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wrapperspb.StringValue, wrapperspb.StringValue]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Test_BidiStreamClient = grpc.BidiStreamingClient[wrapperspb.StringValue, wrapperspb.StringValue]

// TestServer is the server API for Test service.
// All implementations must embed UnimplementedTestServer
// for forward compatibility.
type TestServer interface {
	Unary(context.Context, *wrapperspb.StringValue) (*wrapperspb.StringValue, error)
	ClientStream(grpc.ClientStreamingServer[wrapperspb.StringValue, wrapperspb.StringValue]) error
	ServerStream(*wrapperspb.StringValue, grpc.ServerStreamingServer[wrapperspb.StringValue]) error
	BidiStream(grpc.BidiStreamingServer[wrapperspb.StringValue, wrapperspb.StringValue]) error
	mustEmbedUnimplementedTestServer()
}

// UnimplementedTestServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestServer struct{}

func (UnimplementedTestServer) Unary(context.Context, *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary not implemented")
}
func (UnimplementedTestServer) ClientStream(grpc.ClientStreamingServer[wrapperspb.StringValue, wrapperspb.StringValue]) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (UnimplementedTestServer) ServerStream(*wrapperspb.StringValue, grpc.ServerStreamingServer[wrapperspb.StringValue]) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedTestServer) BidiStream(grpc.BidiStreamingServer[wrapperspb.StringValue, wrapperspb.StringValue]) error {
	return status.Errorf(codes.Unimplemented, "method BidiStream not implemented")
}
func (UnimplementedTestServer) mustEmbedUnimplementedTestServer() {}
func (UnimplementedTestServer) testEmbeddedByValue()              {}

// UnsafeTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServer will
// result in compilation errors.
type UnsafeTestServer interface {
	mustEmbedUnimplementedTestServer()
}

func RegisterTestServer(s grpc.ServiceRegistrar, srv TestServer) {
	// If the following call pancis, it indicates UnimplementedTestServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Test_ServiceDesc, srv)
}

func _Test_Unary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Unary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Test_Unary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Unary(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).ClientStream(&grpc.GenericServerStream[wrapperspb.StringValue, wrapperspb.StringValue]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Test_ClientStreamServer = grpc.ClientStreamingServer[wrapperspb.StringValue, wrapperspb.StringValue]

func _Test_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServer).ServerStream(m, &grpc.GenericServerStream[wrapperspb.StringValue, wrapperspb.StringValue]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Test_ServerStreamServer = grpc.ServerStreamingServer[wrapperspb.StringValue]

func _Test_BidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).BidiStream(&grpc.GenericServerStream[wrapperspb.StringValue, wrapperspb.StringValue]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Test_BidiStreamServer = grpc.BidiStreamingServer[wrapperspb.StringValue, wrapperspb.StringValue]

// Test_ServiceDesc is the grpc.ServiceDesc for Test service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Test_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unary",
			Handler:    _Test_Unary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStream",
			Handler:       _Test_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStream",
			Handler:       _Test_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidiStream",
			Handler:       _Test_BidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}
