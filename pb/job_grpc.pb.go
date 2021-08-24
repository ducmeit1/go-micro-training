// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// FPTJobClient is the client API for FPTJob service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FPTJobClient interface {
	CreateJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error)
	ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*ListJobResponse, error)
}

type fPTJobClient struct {
	cc grpc.ClientConnInterface
}

func NewFPTJobClient(cc grpc.ClientConnInterface) FPTJobClient {
	return &fPTJobClient{cc}
}

func (c *fPTJobClient) CreateJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/training.FPTJob/CreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTJobClient) ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*ListJobResponse, error) {
	out := new(ListJobResponse)
	err := c.cc.Invoke(ctx, "/training.FPTJob/ListJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FPTJobServer is the server API for FPTJob service.
// All implementations must embed UnimplementedFPTJobServer
// for forward compatibility
type FPTJobServer interface {
	CreateJob(context.Context, *Job) (*Job, error)
	ListJob(context.Context, *ListJobRequest) (*ListJobResponse, error)
	mustEmbedUnimplementedFPTJobServer()
}

// UnimplementedFPTJobServer must be embedded to have forward compatible implementations.
type UnimplementedFPTJobServer struct {
}

func (UnimplementedFPTJobServer) CreateJob(context.Context, *Job) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedFPTJobServer) ListJob(context.Context, *ListJobRequest) (*ListJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJob not implemented")
}
func (UnimplementedFPTJobServer) mustEmbedUnimplementedFPTJobServer() {}

// UnsafeFPTJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FPTJobServer will
// result in compilation errors.
type UnsafeFPTJobServer interface {
	mustEmbedUnimplementedFPTJobServer()
}

func RegisterFPTJobServer(s grpc.ServiceRegistrar, srv FPTJobServer) {
	s.RegisterService(&FPTJob_ServiceDesc, srv)
}

func _FPTJob_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTJobServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTJob/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTJobServer).CreateJob(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTJob_ListJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTJobServer).ListJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTJob/ListJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTJobServer).ListJob(ctx, req.(*ListJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FPTJob_ServiceDesc is the grpc.ServiceDesc for FPTJob service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FPTJob_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "training.FPTJob",
	HandlerType: (*FPTJobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _FPTJob_CreateJob_Handler,
		},
		{
			MethodName: "ListJob",
			Handler:    _FPTJob_ListJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}