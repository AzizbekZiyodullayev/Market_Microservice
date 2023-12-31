// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: staff_transaction_service.proto

package sale_service

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

// StaffTransactionServiceClient is the client API for StaffTransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffTransactionServiceClient interface {
	Create(ctx context.Context, in *StaffTrCreateReq, opts ...grpc.CallOption) (*StaffTrCreateResp, error)
	GetList(ctx context.Context, in *StaffTrGetListReq, opts ...grpc.CallOption) (*StaffTrGetListResp, error)
	GetById(ctx context.Context, in *StaffTrIdReq, opts ...grpc.CallOption) (*StaffTransaction, error)
	Update(ctx context.Context, in *StaffTrUpdateReq, opts ...grpc.CallOption) (*StaffTrUpdateResp, error)
	Delete(ctx context.Context, in *StaffTrIdReq, opts ...grpc.CallOption) (*StaffTrDeleteResp, error)
}

type staffTransactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffTransactionServiceClient(cc grpc.ClientConnInterface) StaffTransactionServiceClient {
	return &staffTransactionServiceClient{cc}
}

func (c *staffTransactionServiceClient) Create(ctx context.Context, in *StaffTrCreateReq, opts ...grpc.CallOption) (*StaffTrCreateResp, error) {
	out := new(StaffTrCreateResp)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffTransactionServiceClient) GetList(ctx context.Context, in *StaffTrGetListReq, opts ...grpc.CallOption) (*StaffTrGetListResp, error) {
	out := new(StaffTrGetListResp)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffTransactionServiceClient) GetById(ctx context.Context, in *StaffTrIdReq, opts ...grpc.CallOption) (*StaffTransaction, error) {
	out := new(StaffTransaction)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffTransactionServiceClient) Update(ctx context.Context, in *StaffTrUpdateReq, opts ...grpc.CallOption) (*StaffTrUpdateResp, error) {
	out := new(StaffTrUpdateResp)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffTransactionServiceClient) Delete(ctx context.Context, in *StaffTrIdReq, opts ...grpc.CallOption) (*StaffTrDeleteResp, error) {
	out := new(StaffTrDeleteResp)
	err := c.cc.Invoke(ctx, "/sale_service.StaffTransactionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffTransactionServiceServer is the server API for StaffTransactionService service.
// All implementations must embed UnimplementedStaffTransactionServiceServer
// for forward compatibility
type StaffTransactionServiceServer interface {
	Create(context.Context, *StaffTrCreateReq) (*StaffTrCreateResp, error)
	GetList(context.Context, *StaffTrGetListReq) (*StaffTrGetListResp, error)
	GetById(context.Context, *StaffTrIdReq) (*StaffTransaction, error)
	Update(context.Context, *StaffTrUpdateReq) (*StaffTrUpdateResp, error)
	Delete(context.Context, *StaffTrIdReq) (*StaffTrDeleteResp, error)
	mustEmbedUnimplementedStaffTransactionServiceServer()
}

// UnimplementedStaffTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStaffTransactionServiceServer struct {
}

func (UnimplementedStaffTransactionServiceServer) Create(context.Context, *StaffTrCreateReq) (*StaffTrCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStaffTransactionServiceServer) GetList(context.Context, *StaffTrGetListReq) (*StaffTrGetListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedStaffTransactionServiceServer) GetById(context.Context, *StaffTrIdReq) (*StaffTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedStaffTransactionServiceServer) Update(context.Context, *StaffTrUpdateReq) (*StaffTrUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStaffTransactionServiceServer) Delete(context.Context, *StaffTrIdReq) (*StaffTrDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStaffTransactionServiceServer) mustEmbedUnimplementedStaffTransactionServiceServer() {
}

// UnsafeStaffTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffTransactionServiceServer will
// result in compilation errors.
type UnsafeStaffTransactionServiceServer interface {
	mustEmbedUnimplementedStaffTransactionServiceServer()
}

func RegisterStaffTransactionServiceServer(s grpc.ServiceRegistrar, srv StaffTransactionServiceServer) {
	s.RegisterService(&StaffTransactionService_ServiceDesc, srv)
}

func _StaffTransactionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffTrCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).Create(ctx, req.(*StaffTrCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffTransactionService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffTrGetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).GetList(ctx, req.(*StaffTrGetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffTransactionService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffTrIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).GetById(ctx, req.(*StaffTrIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffTransactionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffTrUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).Update(ctx, req.(*StaffTrUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffTransactionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffTrIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffTransactionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.StaffTransactionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffTransactionServiceServer).Delete(ctx, req.(*StaffTrIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffTransactionService_ServiceDesc is the grpc.ServiceDesc for StaffTransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffTransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sale_service.StaffTransactionService",
	HandlerType: (*StaffTransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StaffTransactionService_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _StaffTransactionService_GetList_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _StaffTransactionService_GetById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StaffTransactionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaffTransactionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff_transaction_service.proto",
}
