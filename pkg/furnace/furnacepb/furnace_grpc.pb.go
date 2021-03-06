// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package furnacepb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FurnaceClient is the client API for Furnace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FurnaceClient interface {
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error)
	IsQueued(ctx context.Context, in *IsQueuedRequest, opts ...grpc.CallOption) (*IsQueuedResponse, error)
}

type furnaceClient struct {
	cc grpc.ClientConnInterface
}

func NewFurnaceClient(cc grpc.ClientConnInterface) FurnaceClient {
	return &furnaceClient{cc}
}

func (c *furnaceClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/foundry.Furnace/Build", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *furnaceClient) IsQueued(ctx context.Context, in *IsQueuedRequest, opts ...grpc.CallOption) (*IsQueuedResponse, error) {
	out := new(IsQueuedResponse)
	err := c.cc.Invoke(ctx, "/foundry.Furnace/IsQueued", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FurnaceServer is the server API for Furnace service.
// All implementations must embed UnimplementedFurnaceServer
// for forward compatibility
type FurnaceServer interface {
	Build(context.Context, *BuildRequest) (*BuildResponse, error)
	IsQueued(context.Context, *IsQueuedRequest) (*IsQueuedResponse, error)
	mustEmbedUnimplementedFurnaceServer()
}

// UnimplementedFurnaceServer must be embedded to have forward compatible implementations.
type UnimplementedFurnaceServer struct {
}

func (UnimplementedFurnaceServer) Build(context.Context, *BuildRequest) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}
func (UnimplementedFurnaceServer) IsQueued(context.Context, *IsQueuedRequest) (*IsQueuedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsQueued not implemented")
}
func (UnimplementedFurnaceServer) mustEmbedUnimplementedFurnaceServer() {}

// UnsafeFurnaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FurnaceServer will
// result in compilation errors.
type UnsafeFurnaceServer interface {
	mustEmbedUnimplementedFurnaceServer()
}

func RegisterFurnaceServer(s grpc.ServiceRegistrar, srv FurnaceServer) {
	s.RegisterService(&_Furnace_serviceDesc, srv)
}

func _Furnace_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FurnaceServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foundry.Furnace/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FurnaceServer).Build(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Furnace_IsQueued_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsQueuedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FurnaceServer).IsQueued(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foundry.Furnace/IsQueued",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FurnaceServer).IsQueued(ctx, req.(*IsQueuedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Furnace_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foundry.Furnace",
	HandlerType: (*FurnaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Build",
			Handler:    _Furnace_Build_Handler,
		},
		{
			MethodName: "IsQueued",
			Handler:    _Furnace_IsQueued_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "furnace.proto",
}
