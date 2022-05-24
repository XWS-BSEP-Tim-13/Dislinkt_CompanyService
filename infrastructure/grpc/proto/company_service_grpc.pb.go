// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: company_service.proto

package proto

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

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	CreateCompany(ctx context.Context, in *NewCompany, opts ...grpc.CallOption) (*NewCompany, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/company.CompanyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/company.CompanyService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) CreateCompany(ctx context.Context, in *NewCompany, opts ...grpc.CallOption) (*NewCompany, error) {
	out := new(NewCompany)
	err := c.cc.Invoke(ctx, "/company.CompanyService/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
// All implementations must embed UnimplementedCompanyServiceServer
// for forward compatibility
type CompanyServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	CreateCompany(context.Context, *NewCompany) (*NewCompany, error)
	mustEmbedUnimplementedCompanyServiceServer()
}

// UnimplementedCompanyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyServiceServer struct {
}

func (UnimplementedCompanyServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCompanyServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCompanyServiceServer) CreateCompany(context.Context, *NewCompany) (*NewCompany, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedCompanyServiceServer) mustEmbedUnimplementedCompanyServiceServer() {}

// UnsafeCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServiceServer will
// result in compilation errors.
type UnsafeCompanyServiceServer interface {
	mustEmbedUnimplementedCompanyServiceServer()
}

func RegisterCompanyServiceServer(s grpc.ServiceRegistrar, srv CompanyServiceServer) {
	s.RegisterService(&CompanyService_ServiceDesc, srv)
}

func _CompanyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateCompany(ctx, req.(*NewCompany))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyService_ServiceDesc is the grpc.ServiceDesc for CompanyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "company.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CompanyService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CompanyService_GetAll_Handler,
		},
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyService_CreateCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "company_service.proto",
}
