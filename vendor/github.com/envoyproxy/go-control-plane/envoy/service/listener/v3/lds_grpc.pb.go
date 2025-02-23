// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: envoy/service/listener/v3/lds.proto

package listenerv3

import (
	context "context"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ListenerDiscoveryService_DeltaListeners_FullMethodName  = "/envoy.service.listener.v3.ListenerDiscoveryService/DeltaListeners"
	ListenerDiscoveryService_StreamListeners_FullMethodName = "/envoy.service.listener.v3.ListenerDiscoveryService/StreamListeners"
	ListenerDiscoveryService_FetchListeners_FullMethodName  = "/envoy.service.listener.v3.ListenerDiscoveryService/FetchListeners"
)

// ListenerDiscoveryServiceClient is the client API for ListenerDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListenerDiscoveryServiceClient interface {
	DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error)
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListenerDiscoveryServiceClient(cc grpc.ClientConnInterface) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ListenerDiscoveryService_ServiceDesc.Streams[0], ListenerDiscoveryService_DeltaListeners_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceDeltaListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_DeltaListenersClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceDeltaListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ListenerDiscoveryService_ServiceDesc.Streams[1], ListenerDiscoveryService_StreamListeners_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, ListenerDiscoveryService_FetchListeners_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListenerDiscoveryServiceServer is the server API for ListenerDiscoveryService service.
// All implementations should embed UnimplementedListenerDiscoveryServiceServer
// for forward compatibility
type ListenerDiscoveryServiceServer interface {
	DeltaListeners(ListenerDiscoveryService_DeltaListenersServer) error
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedListenerDiscoveryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedListenerDiscoveryServiceServer struct {
}

func (UnimplementedListenerDiscoveryServiceServer) DeltaListeners(ListenerDiscoveryService_DeltaListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaListeners not implemented")
}
func (UnimplementedListenerDiscoveryServiceServer) StreamListeners(ListenerDiscoveryService_StreamListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamListeners not implemented")
}
func (UnimplementedListenerDiscoveryServiceServer) FetchListeners(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchListeners not implemented")
}

// UnsafeListenerDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListenerDiscoveryServiceServer will
// result in compilation errors.
type UnsafeListenerDiscoveryServiceServer interface {
	mustEmbedUnimplementedListenerDiscoveryServiceServer()
}

func RegisterListenerDiscoveryServiceServer(s grpc.ServiceRegistrar, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&ListenerDiscoveryService_ServiceDesc, srv)
}

func _ListenerDiscoveryService_DeltaListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).DeltaListeners(&listenerDiscoveryServiceDeltaListenersServer{stream})
}

type ListenerDiscoveryService_DeltaListenersServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceDeltaListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerDiscoveryService_FetchListeners_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListenerDiscoveryService_ServiceDesc is the grpc.ServiceDesc for ListenerDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListenerDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.listener.v3.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaListeners",
			Handler:       _ListenerDiscoveryService_DeltaListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/listener/v3/lds.proto",
}
