// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sspb

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

// TransportClient is the client API for Transport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportClient interface {
	ListFeatures(ctx context.Context, in *X, opts ...grpc.CallOption) (Transport_ListFeaturesClient, error)
}

type transportClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportClient(cc grpc.ClientConnInterface) TransportClient {
	return &transportClient{cc}
}

func (c *transportClient) ListFeatures(ctx context.Context, in *X, opts ...grpc.CallOption) (Transport_ListFeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transport_ServiceDesc.Streams[0], "/transport/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &transportListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Transport_ListFeaturesClient interface {
	Recv() (*Y, error)
	grpc.ClientStream
}

type transportListFeaturesClient struct {
	grpc.ClientStream
}

func (x *transportListFeaturesClient) Recv() (*Y, error) {
	m := new(Y)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransportServer is the server API for Transport service.
// All implementations must embed UnimplementedTransportServer
// for forward compatibility
type TransportServer interface {
	ListFeatures(*X, Transport_ListFeaturesServer) error
	mustEmbedUnimplementedTransportServer()
}

// UnimplementedTransportServer must be embedded to have forward compatible implementations.
type UnimplementedTransportServer struct {
}

func (UnimplementedTransportServer) ListFeatures(*X, Transport_ListFeaturesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}
func (UnimplementedTransportServer) mustEmbedUnimplementedTransportServer() {}

// UnsafeTransportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServer will
// result in compilation errors.
type UnsafeTransportServer interface {
	mustEmbedUnimplementedTransportServer()
}

func RegisterTransportServer(s grpc.ServiceRegistrar, srv TransportServer) {
	s.RegisterService(&Transport_ServiceDesc, srv)
}

func _Transport_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(X)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransportServer).ListFeatures(m, &transportListFeaturesServer{stream})
}

type Transport_ListFeaturesServer interface {
	Send(*Y) error
	grpc.ServerStream
}

type transportListFeaturesServer struct {
	grpc.ServerStream
}

func (x *transportListFeaturesServer) Send(m *Y) error {
	return x.ServerStream.SendMsg(m)
}

// Transport_ServiceDesc is the grpc.ServiceDesc for Transport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transport",
	HandlerType: (*TransportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _Transport_ListFeatures_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ss.proto",
}
