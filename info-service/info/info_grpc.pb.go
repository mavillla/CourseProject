// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package info

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

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoServiceClient interface {
	Create(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Give(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Home(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
}

type infoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoServiceClient(cc grpc.ClientConnInterface) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) Create(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.info.InfoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.info.InfoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) Give(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.info.InfoService/Give", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) Home(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.info.InfoService/Home", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
// All implementations must embed UnimplementedInfoServiceServer
// for forward compatibility
type InfoServiceServer interface {
	Create(context.Context, *Empty) (*Response, error)
	Get(context.Context, *Empty) (*Response, error)
	Give(context.Context, *Empty) (*Response, error)
	Home(context.Context, *Empty) (*Response, error)
	mustEmbedUnimplementedInfoServiceServer()
}

// UnimplementedInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInfoServiceServer struct {
}

func (UnimplementedInfoServiceServer) Create(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInfoServiceServer) Get(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedInfoServiceServer) Give(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Give not implemented")
}
func (UnimplementedInfoServiceServer) Home(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Home not implemented")
}
func (UnimplementedInfoServiceServer) mustEmbedUnimplementedInfoServiceServer() {}

// UnsafeInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoServiceServer will
// result in compilation errors.
type UnsafeInfoServiceServer interface {
	mustEmbedUnimplementedInfoServiceServer()
}

func RegisterInfoServiceServer(s grpc.ServiceRegistrar, srv InfoServiceServer) {
	s.RegisterService(&InfoService_ServiceDesc, srv)
}

func _InfoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.info.InfoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).Create(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.info.InfoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).Get(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_Give_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).Give(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.info.InfoService/Give",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).Give(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_Home_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).Home(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.info.InfoService/Home",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).Home(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// InfoService_ServiceDesc is the grpc.ServiceDesc for InfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.info.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _InfoService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _InfoService_Get_Handler,
		},
		{
			MethodName: "Give",
			Handler:    _InfoService_Give_Handler,
		},
		{
			MethodName: "Home",
			Handler:    _InfoService_Home_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info.proto",
}
