// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package scpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ScClient is the client API for Sc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScClient interface {
	Greeting(ctx context.Context, opts ...grpc.CallOption) (Sc_GreetingClient, error)
}

type scClient struct {
	cc grpc.ClientConnInterface
}

func NewScClient(cc grpc.ClientConnInterface) ScClient {
	return &scClient{cc}
}

func (c *scClient) Greeting(ctx context.Context, opts ...grpc.CallOption) (Sc_GreetingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sc_ServiceDesc.Streams[0], "/sc/greeting", opts...)
	if err != nil {
		return nil, err
	}
	x := &scGreetingClient{stream}
	return x, nil
}

type Sc_GreetingClient interface {
	Send(*X) error
	CloseAndRecv() (*Y, error)
	grpc.ClientStream
}

type scGreetingClient struct {
	grpc.ClientStream
}

func (x *scGreetingClient) Send(m *X) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scGreetingClient) CloseAndRecv() (*Y, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Y)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ScServer is the server API for Sc service.
// All implementations must embed UnimplementedScServer
// for forward compatibility
type ScServer interface {
	Greeting(Sc_GreetingServer) error
	mustEmbedUnimplementedScServer()
}

// UnimplementedScServer must be embedded to have forward compatible implementations.
type UnimplementedScServer struct {
}

func (UnimplementedScServer) Greeting(Sc_GreetingServer) error {
	return status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}
func (UnimplementedScServer) mustEmbedUnimplementedScServer() {}

// UnsafeScServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScServer will
// result in compilation errors.
type UnsafeScServer interface {
	mustEmbedUnimplementedScServer()
}

func RegisterScServer(s grpc.ServiceRegistrar, srv ScServer) {
	s.RegisterService(&Sc_ServiceDesc, srv)
}

func _Sc_Greeting_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScServer).Greeting(&scGreetingServer{stream})
}

type Sc_GreetingServer interface {
	SendAndClose(*Y) error
	Recv() (*X, error)
	grpc.ServerStream
}

type scGreetingServer struct {
	grpc.ServerStream
}

func (x *scGreetingServer) SendAndClose(m *Y) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scGreetingServer) Recv() (*X, error) {
	m := new(X)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Sc_ServiceDesc is the grpc.ServiceDesc for Sc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sc",
	HandlerType: (*ScServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "greeting",
			Handler:       _Sc_Greeting_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "scpb/sc.proto",
}
