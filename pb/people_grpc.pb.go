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

// FPTPeopleClient is the client API for FPTPeople service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FPTPeopleClient interface {
	CreatePeople(ctx context.Context, in *People, opts ...grpc.CallOption) (*People, error)
	UpdatePeople(ctx context.Context, in *People, opts ...grpc.CallOption) (*People, error)
	FindPeople(ctx context.Context, in *FindPeopleRequest, opts ...grpc.CallOption) (*People, error)
	ListPeoples(ctx context.Context, in *ListPeopleRequest, opts ...grpc.CallOption) (*ListPeopleResponse, error)
	DeletePeople(ctx context.Context, in *DeletePeopleRequest, opts ...grpc.CallOption) (*Empty, error)
}

type fPTPeopleClient struct {
	cc grpc.ClientConnInterface
}

func NewFPTPeopleClient(cc grpc.ClientConnInterface) FPTPeopleClient {
	return &fPTPeopleClient{cc}
}

func (c *fPTPeopleClient) CreatePeople(ctx context.Context, in *People, opts ...grpc.CallOption) (*People, error) {
	out := new(People)
	err := c.cc.Invoke(ctx, "/training.FPTPeople/CreatePeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTPeopleClient) UpdatePeople(ctx context.Context, in *People, opts ...grpc.CallOption) (*People, error) {
	out := new(People)
	err := c.cc.Invoke(ctx, "/training.FPTPeople/UpdatePeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTPeopleClient) FindPeople(ctx context.Context, in *FindPeopleRequest, opts ...grpc.CallOption) (*People, error) {
	out := new(People)
	err := c.cc.Invoke(ctx, "/training.FPTPeople/FindPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTPeopleClient) ListPeoples(ctx context.Context, in *ListPeopleRequest, opts ...grpc.CallOption) (*ListPeopleResponse, error) {
	out := new(ListPeopleResponse)
	err := c.cc.Invoke(ctx, "/training.FPTPeople/ListPeoples", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTPeopleClient) DeletePeople(ctx context.Context, in *DeletePeopleRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/training.FPTPeople/DeletePeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FPTPeopleServer is the server API for FPTPeople service.
// All implementations must embed UnimplementedFPTPeopleServer
// for forward compatibility
type FPTPeopleServer interface {
	CreatePeople(context.Context, *People) (*People, error)
	UpdatePeople(context.Context, *People) (*People, error)
	FindPeople(context.Context, *FindPeopleRequest) (*People, error)
	ListPeoples(context.Context, *ListPeopleRequest) (*ListPeopleResponse, error)
	DeletePeople(context.Context, *DeletePeopleRequest) (*Empty, error)
	mustEmbedUnimplementedFPTPeopleServer()
}

// UnimplementedFPTPeopleServer must be embedded to have forward compatible implementations.
type UnimplementedFPTPeopleServer struct {
}

func (UnimplementedFPTPeopleServer) CreatePeople(context.Context, *People) (*People, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeople not implemented")
}
func (UnimplementedFPTPeopleServer) UpdatePeople(context.Context, *People) (*People, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePeople not implemented")
}
func (UnimplementedFPTPeopleServer) FindPeople(context.Context, *FindPeopleRequest) (*People, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPeople not implemented")
}
func (UnimplementedFPTPeopleServer) ListPeoples(context.Context, *ListPeopleRequest) (*ListPeopleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPeoples not implemented")
}
func (UnimplementedFPTPeopleServer) DeletePeople(context.Context, *DeletePeopleRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePeople not implemented")
}
func (UnimplementedFPTPeopleServer) mustEmbedUnimplementedFPTPeopleServer() {}

// UnsafeFPTPeopleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FPTPeopleServer will
// result in compilation errors.
type UnsafeFPTPeopleServer interface {
	mustEmbedUnimplementedFPTPeopleServer()
}

func RegisterFPTPeopleServer(s grpc.ServiceRegistrar, srv FPTPeopleServer) {
	s.RegisterService(&FPTPeople_ServiceDesc, srv)
}

func _FPTPeople_CreatePeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(People)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTPeopleServer).CreatePeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTPeople/CreatePeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTPeopleServer).CreatePeople(ctx, req.(*People))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTPeople_UpdatePeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(People)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTPeopleServer).UpdatePeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTPeople/UpdatePeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTPeopleServer).UpdatePeople(ctx, req.(*People))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTPeople_FindPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPeopleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTPeopleServer).FindPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTPeople/FindPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTPeopleServer).FindPeople(ctx, req.(*FindPeopleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTPeople_ListPeoples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPeopleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTPeopleServer).ListPeoples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTPeople/ListPeoples",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTPeopleServer).ListPeoples(ctx, req.(*ListPeopleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTPeople_DeletePeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePeopleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTPeopleServer).DeletePeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTPeople/DeletePeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTPeopleServer).DeletePeople(ctx, req.(*DeletePeopleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FPTPeople_ServiceDesc is the grpc.ServiceDesc for FPTPeople service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FPTPeople_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "training.FPTPeople",
	HandlerType: (*FPTPeopleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePeople",
			Handler:    _FPTPeople_CreatePeople_Handler,
		},
		{
			MethodName: "UpdatePeople",
			Handler:    _FPTPeople_UpdatePeople_Handler,
		},
		{
			MethodName: "FindPeople",
			Handler:    _FPTPeople_FindPeople_Handler,
		},
		{
			MethodName: "ListPeoples",
			Handler:    _FPTPeople_ListPeoples_Handler,
		},
		{
			MethodName: "DeletePeople",
			Handler:    _FPTPeople_DeletePeople_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "people.proto",
}