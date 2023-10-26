// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: branch_product_transaction_service.proto

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

// BrPrTransactionServiceClient is the client API for BrPrTransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrPrTransactionServiceClient interface {
	Create(ctx context.Context, in *BrPrTrCreateReq, opts ...grpc.CallOption) (*BrPrTrCreateResp, error)
	GetList(ctx context.Context, in *BrPrTrGetListReq, opts ...grpc.CallOption) (*BrPrTrGetListResp, error)
	GetById(ctx context.Context, in *BrPrTrIdReq, opts ...grpc.CallOption) (*BrPrTransaction, error)
	Update(ctx context.Context, in *BrPrTrUpdateReq, opts ...grpc.CallOption) (*BrPrTrUpdateResp, error)
	Delete(ctx context.Context, in *BrPrTrIdReq, opts ...grpc.CallOption) (*BrPrTrDeleteResp, error)
}

type brPrTransactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrPrTransactionServiceClient(cc grpc.ClientConnInterface) BrPrTransactionServiceClient {
	return &brPrTransactionServiceClient{cc}
}

func (c *brPrTransactionServiceClient) Create(ctx context.Context, in *BrPrTrCreateReq, opts ...grpc.CallOption) (*BrPrTrCreateResp, error) {
	out := new(BrPrTrCreateResp)
	err := c.cc.Invoke(ctx, "/sale_service.BrPrTransactionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brPrTransactionServiceClient) GetList(ctx context.Context, in *BrPrTrGetListReq, opts ...grpc.CallOption) (*BrPrTrGetListResp, error) {
	out := new(BrPrTrGetListResp)
	err := c.cc.Invoke(ctx, "/sale_service.BrPrTransactionService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brPrTransactionServiceClient) GetById(ctx context.Context, in *BrPrTrIdReq, opts ...grpc.CallOption) (*BrPrTransaction, error) {
	out := new(BrPrTransaction)
	err := c.cc.Invoke(ctx, "/sale_service.BrPrTransactionService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brPrTransactionServiceClient) Update(ctx context.Context, in *BrPrTrUpdateReq, opts ...grpc.CallOption) (*BrPrTrUpdateResp, error) {
	out := new(BrPrTrUpdateResp)
	err := c.cc.Invoke(ctx, "/sale_service.BrPrTransactionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brPrTransactionServiceClient) Delete(ctx context.Context, in *BrPrTrIdReq, opts ...grpc.CallOption) (*BrPrTrDeleteResp, error) {
	out := new(BrPrTrDeleteResp)
	err := c.cc.Invoke(ctx, "/sale_service.BrPrTransactionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrPrTransactionServiceServer is the server API for BrPrTransactionService service.
// All implementations should embed UnimplementedBrPrTransactionServiceServer
// for forward compatibility
type BrPrTransactionServiceServer interface {
	Create(context.Context, *BrPrTrCreateReq) (*BrPrTrCreateResp, error)
	GetList(context.Context, *BrPrTrGetListReq) (*BrPrTrGetListResp, error)
	GetById(context.Context, *BrPrTrIdReq) (*BrPrTransaction, error)
	Update(context.Context, *BrPrTrUpdateReq) (*BrPrTrUpdateResp, error)
	Delete(context.Context, *BrPrTrIdReq) (*BrPrTrDeleteResp, error)
}

// UnimplementedBrPrTransactionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBrPrTransactionServiceServer struct {
}

func (UnimplementedBrPrTransactionServiceServer) Create(context.Context, *BrPrTrCreateReq) (*BrPrTrCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBrPrTransactionServiceServer) GetList(context.Context, *BrPrTrGetListReq) (*BrPrTrGetListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedBrPrTransactionServiceServer) GetById(context.Context, *BrPrTrIdReq) (*BrPrTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBrPrTransactionServiceServer) Update(context.Context, *BrPrTrUpdateReq) (*BrPrTrUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBrPrTransactionServiceServer) Delete(context.Context, *BrPrTrIdReq) (*BrPrTrDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeBrPrTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrPrTransactionServiceServer will
// result in compilation errors.
type UnsafeBrPrTransactionServiceServer interface {
	mustEmbedUnimplementedBrPrTransactionServiceServer()
}

func RegisterBrPrTransactionServiceServer(s grpc.ServiceRegistrar, srv BrPrTransactionServiceServer) {
	s.RegisterService(&BrPrTransactionService_ServiceDesc, srv)
}

func _BrPrTransactionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrPrTrCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrPrTransactionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BrPrTransactionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrPrTransactionServiceServer).Create(ctx, req.(*BrPrTrCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrPrTransactionService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrPrTrGetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrPrTransactionServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BrPrTransactionService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrPrTransactionServiceServer).GetList(ctx, req.(*BrPrTrGetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrPrTransactionService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrPrTrIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrPrTransactionServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BrPrTransactionService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrPrTransactionServiceServer).GetById(ctx, req.(*BrPrTrIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrPrTransactionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrPrTrUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrPrTransactionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BrPrTransactionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrPrTransactionServiceServer).Update(ctx, req.(*BrPrTrUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrPrTransactionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrPrTrIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrPrTransactionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.BrPrTransactionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrPrTransactionServiceServer).Delete(ctx, req.(*BrPrTrIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BrPrTransactionService_ServiceDesc is the grpc.ServiceDesc for BrPrTransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrPrTransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sale_service.BrPrTransactionService",
	HandlerType: (*BrPrTransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BrPrTransactionService_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _BrPrTransactionService_GetList_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _BrPrTransactionService_GetById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BrPrTransactionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BrPrTransactionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "branch_product_transaction_service.proto",
}
