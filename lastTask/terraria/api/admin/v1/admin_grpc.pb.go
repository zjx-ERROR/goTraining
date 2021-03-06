// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: admin/v1/admin.proto

package v1

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

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	CreateMonster(ctx context.Context, in *CreateMonsterRequest, opts ...grpc.CallOption) (*CreateMonsterReply, error)
	UpdateMonster(ctx context.Context, in *UpdateMonsterRequest, opts ...grpc.CallOption) (*UpdateMonsterReply, error)
	DeleteMonster(ctx context.Context, in *DeleteMonsterRequest, opts ...grpc.CallOption) (*DeleteMonsterReply, error)
	GetMonster(ctx context.Context, in *GetMonsterRequest, opts ...grpc.CallOption) (*GetMonsterReply, error)
	ListMonster(ctx context.Context, in *ListMonsterRequest, opts ...grpc.CallOption) (*ListMonsterReply, error)
	CreateWeapon(ctx context.Context, in *CreateWeaponRequest, opts ...grpc.CallOption) (*CreateWeaponReply, error)
	UpdateWeapon(ctx context.Context, in *UpdateWeaponRequest, opts ...grpc.CallOption) (*UpdateWeaponReply, error)
	DeleteWeapon(ctx context.Context, in *DeleteWeaponRequest, opts ...grpc.CallOption) (*DeleteWeaponReply, error)
	GetWeapon(ctx context.Context, in *GetWeaponRequest, opts ...grpc.CallOption) (*GetWeaponReply, error)
	ListWeapon(ctx context.Context, in *ListWeaponRequest, opts ...grpc.CallOption) (*ListWeaponReply, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) CreateMonster(ctx context.Context, in *CreateMonsterRequest, opts ...grpc.CallOption) (*CreateMonsterReply, error) {
	out := new(CreateMonsterReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/CreateMonster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpdateMonster(ctx context.Context, in *UpdateMonsterRequest, opts ...grpc.CallOption) (*UpdateMonsterReply, error) {
	out := new(UpdateMonsterReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/UpdateMonster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeleteMonster(ctx context.Context, in *DeleteMonsterRequest, opts ...grpc.CallOption) (*DeleteMonsterReply, error) {
	out := new(DeleteMonsterReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/DeleteMonster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetMonster(ctx context.Context, in *GetMonsterRequest, opts ...grpc.CallOption) (*GetMonsterReply, error) {
	out := new(GetMonsterReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/GetMonster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListMonster(ctx context.Context, in *ListMonsterRequest, opts ...grpc.CallOption) (*ListMonsterReply, error) {
	out := new(ListMonsterReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/ListMonster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateWeapon(ctx context.Context, in *CreateWeaponRequest, opts ...grpc.CallOption) (*CreateWeaponReply, error) {
	out := new(CreateWeaponReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/CreateWeapon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpdateWeapon(ctx context.Context, in *UpdateWeaponRequest, opts ...grpc.CallOption) (*UpdateWeaponReply, error) {
	out := new(UpdateWeaponReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/UpdateWeapon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeleteWeapon(ctx context.Context, in *DeleteWeaponRequest, opts ...grpc.CallOption) (*DeleteWeaponReply, error) {
	out := new(DeleteWeaponReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/DeleteWeapon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetWeapon(ctx context.Context, in *GetWeaponRequest, opts ...grpc.CallOption) (*GetWeaponReply, error) {
	out := new(GetWeaponReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/GetWeapon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListWeapon(ctx context.Context, in *ListWeaponRequest, opts ...grpc.CallOption) (*ListWeaponReply, error) {
	out := new(ListWeaponReply)
	err := c.cc.Invoke(ctx, "/api.admin.v1.Admin/ListWeapon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	CreateMonster(context.Context, *CreateMonsterRequest) (*CreateMonsterReply, error)
	UpdateMonster(context.Context, *UpdateMonsterRequest) (*UpdateMonsterReply, error)
	DeleteMonster(context.Context, *DeleteMonsterRequest) (*DeleteMonsterReply, error)
	GetMonster(context.Context, *GetMonsterRequest) (*GetMonsterReply, error)
	ListMonster(context.Context, *ListMonsterRequest) (*ListMonsterReply, error)
	CreateWeapon(context.Context, *CreateWeaponRequest) (*CreateWeaponReply, error)
	UpdateWeapon(context.Context, *UpdateWeaponRequest) (*UpdateWeaponReply, error)
	DeleteWeapon(context.Context, *DeleteWeaponRequest) (*DeleteWeaponReply, error)
	GetWeapon(context.Context, *GetWeaponRequest) (*GetWeaponReply, error)
	ListWeapon(context.Context, *ListWeaponRequest) (*ListWeaponReply, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) CreateMonster(context.Context, *CreateMonsterRequest) (*CreateMonsterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMonster not implemented")
}
func (UnimplementedAdminServer) UpdateMonster(context.Context, *UpdateMonsterRequest) (*UpdateMonsterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMonster not implemented")
}
func (UnimplementedAdminServer) DeleteMonster(context.Context, *DeleteMonsterRequest) (*DeleteMonsterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMonster not implemented")
}
func (UnimplementedAdminServer) GetMonster(context.Context, *GetMonsterRequest) (*GetMonsterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonster not implemented")
}
func (UnimplementedAdminServer) ListMonster(context.Context, *ListMonsterRequest) (*ListMonsterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMonster not implemented")
}
func (UnimplementedAdminServer) CreateWeapon(context.Context, *CreateWeaponRequest) (*CreateWeaponReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWeapon not implemented")
}
func (UnimplementedAdminServer) UpdateWeapon(context.Context, *UpdateWeaponRequest) (*UpdateWeaponReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWeapon not implemented")
}
func (UnimplementedAdminServer) DeleteWeapon(context.Context, *DeleteWeaponRequest) (*DeleteWeaponReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWeapon not implemented")
}
func (UnimplementedAdminServer) GetWeapon(context.Context, *GetWeaponRequest) (*GetWeaponReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeapon not implemented")
}
func (UnimplementedAdminServer) ListWeapon(context.Context, *ListWeaponRequest) (*ListWeaponReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWeapon not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_CreateMonster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMonsterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateMonster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/CreateMonster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateMonster(ctx, req.(*CreateMonsterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpdateMonster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMonsterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpdateMonster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/UpdateMonster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpdateMonster(ctx, req.(*UpdateMonsterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeleteMonster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMonsterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeleteMonster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/DeleteMonster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeleteMonster(ctx, req.(*DeleteMonsterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetMonster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonsterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetMonster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/GetMonster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetMonster(ctx, req.(*GetMonsterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListMonster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMonsterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListMonster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/ListMonster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListMonster(ctx, req.(*ListMonsterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateWeapon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWeaponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateWeapon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/CreateWeapon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateWeapon(ctx, req.(*CreateWeaponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpdateWeapon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWeaponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpdateWeapon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/UpdateWeapon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpdateWeapon(ctx, req.(*UpdateWeaponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeleteWeapon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWeaponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeleteWeapon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/DeleteWeapon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeleteWeapon(ctx, req.(*DeleteWeaponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetWeapon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeaponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetWeapon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/GetWeapon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetWeapon(ctx, req.(*GetWeaponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListWeapon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWeaponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListWeapon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.admin.v1.Admin/ListWeapon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListWeapon(ctx, req.(*ListWeaponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.v1.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMonster",
			Handler:    _Admin_CreateMonster_Handler,
		},
		{
			MethodName: "UpdateMonster",
			Handler:    _Admin_UpdateMonster_Handler,
		},
		{
			MethodName: "DeleteMonster",
			Handler:    _Admin_DeleteMonster_Handler,
		},
		{
			MethodName: "GetMonster",
			Handler:    _Admin_GetMonster_Handler,
		},
		{
			MethodName: "ListMonster",
			Handler:    _Admin_ListMonster_Handler,
		},
		{
			MethodName: "CreateWeapon",
			Handler:    _Admin_CreateWeapon_Handler,
		},
		{
			MethodName: "UpdateWeapon",
			Handler:    _Admin_UpdateWeapon_Handler,
		},
		{
			MethodName: "DeleteWeapon",
			Handler:    _Admin_DeleteWeapon_Handler,
		},
		{
			MethodName: "GetWeapon",
			Handler:    _Admin_GetWeapon_Handler,
		},
		{
			MethodName: "ListWeapon",
			Handler:    _Admin_ListWeapon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/v1/admin.proto",
}
