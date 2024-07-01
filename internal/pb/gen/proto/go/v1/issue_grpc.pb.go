// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/issue.proto

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

// IssueServiceClient is the client API for IssueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IssueServiceClient interface {
	Issues(ctx context.Context, in *IssuesReq, opts ...grpc.CallOption) (*IssuesRes, error)
	IssuesUser(ctx context.Context, in *IssuesReq, opts ...grpc.CallOption) (*IssuesRes, error)
	CreateIssue(ctx context.Context, in *CreateIssueReq, opts ...grpc.CallOption) (*CreateIssueRes, error)
	DeleteIssue(ctx context.Context, in *DeleteIssueReq, opts ...grpc.CallOption) (*DeleteIssueRes, error)
	Issue(ctx context.Context, in *IssueReq, opts ...grpc.CallOption) (*IssueRes, error)
	IssueUpdateStatus(ctx context.Context, in *IssueUpdateStatusReq, opts ...grpc.CallOption) (*IssueUpdateStatusRes, error)
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentRes, error)
	DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentRes, error)
}

type issueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIssueServiceClient(cc grpc.ClientConnInterface) IssueServiceClient {
	return &issueServiceClient{cc}
}

func (c *issueServiceClient) Issues(ctx context.Context, in *IssuesReq, opts ...grpc.CallOption) (*IssuesRes, error) {
	out := new(IssuesRes)
	err := c.cc.Invoke(ctx, "/issue.IssueService/Issues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) IssuesUser(ctx context.Context, in *IssuesReq, opts ...grpc.CallOption) (*IssuesRes, error) {
	out := new(IssuesRes)
	err := c.cc.Invoke(ctx, "/issue.IssueService/IssuesUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) CreateIssue(ctx context.Context, in *CreateIssueReq, opts ...grpc.CallOption) (*CreateIssueRes, error) {
	out := new(CreateIssueRes)
	err := c.cc.Invoke(ctx, "/issue.IssueService/CreateIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) DeleteIssue(ctx context.Context, in *DeleteIssueReq, opts ...grpc.CallOption) (*DeleteIssueRes, error) {
	out := new(DeleteIssueRes)
	err := c.cc.Invoke(ctx, "/issue.IssueService/DeleteIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) Issue(ctx context.Context, in *IssueReq, opts ...grpc.CallOption) (*IssueRes, error) {
	out := new(IssueRes)
	err := c.cc.Invoke(ctx, "/issue.IssueService/Issue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) IssueUpdateStatus(ctx context.Context, in *IssueUpdateStatusReq, opts ...grpc.CallOption) (*IssueUpdateStatusRes, error) {
	out := new(IssueUpdateStatusRes)
	err := c.cc.Invoke(ctx, "/issue.IssueService/IssueUpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentRes, error) {
	out := new(AddCommentRes)
	err := c.cc.Invoke(ctx, "/issue.IssueService/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentRes, error) {
	out := new(DeleteCommentRes)
	err := c.cc.Invoke(ctx, "/issue.IssueService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssueServiceServer is the server API for IssueService service.
// All implementations must embed UnimplementedIssueServiceServer
// for forward compatibility
type IssueServiceServer interface {
	Issues(context.Context, *IssuesReq) (*IssuesRes, error)
	IssuesUser(context.Context, *IssuesReq) (*IssuesRes, error)
	CreateIssue(context.Context, *CreateIssueReq) (*CreateIssueRes, error)
	DeleteIssue(context.Context, *DeleteIssueReq) (*DeleteIssueRes, error)
	Issue(context.Context, *IssueReq) (*IssueRes, error)
	IssueUpdateStatus(context.Context, *IssueUpdateStatusReq) (*IssueUpdateStatusRes, error)
	AddComment(context.Context, *AddCommentReq) (*AddCommentRes, error)
	DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentRes, error)
	mustEmbedUnimplementedIssueServiceServer()
}

// UnimplementedIssueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIssueServiceServer struct {
}

func (UnimplementedIssueServiceServer) Issues(context.Context, *IssuesReq) (*IssuesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issues not implemented")
}
func (UnimplementedIssueServiceServer) IssuesUser(context.Context, *IssuesReq) (*IssuesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssuesUser not implemented")
}
func (UnimplementedIssueServiceServer) CreateIssue(context.Context, *CreateIssueReq) (*CreateIssueRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssue not implemented")
}
func (UnimplementedIssueServiceServer) DeleteIssue(context.Context, *DeleteIssueReq) (*DeleteIssueRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIssue not implemented")
}
func (UnimplementedIssueServiceServer) Issue(context.Context, *IssueReq) (*IssueRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (UnimplementedIssueServiceServer) IssueUpdateStatus(context.Context, *IssueUpdateStatusReq) (*IssueUpdateStatusRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueUpdateStatus not implemented")
}
func (UnimplementedIssueServiceServer) AddComment(context.Context, *AddCommentReq) (*AddCommentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedIssueServiceServer) DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedIssueServiceServer) mustEmbedUnimplementedIssueServiceServer() {}

// UnsafeIssueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IssueServiceServer will
// result in compilation errors.
type UnsafeIssueServiceServer interface {
	mustEmbedUnimplementedIssueServiceServer()
}

func RegisterIssueServiceServer(s grpc.ServiceRegistrar, srv IssueServiceServer) {
	s.RegisterService(&IssueService_ServiceDesc, srv)
}

func _IssueService_Issues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssuesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).Issues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/Issues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).Issues(ctx, req.(*IssuesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_IssuesUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssuesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).IssuesUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/IssuesUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).IssuesUser(ctx, req.(*IssuesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_CreateIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIssueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).CreateIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/CreateIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).CreateIssue(ctx, req.(*CreateIssueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_DeleteIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIssueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).DeleteIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/DeleteIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).DeleteIssue(ctx, req.(*DeleteIssueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/Issue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).Issue(ctx, req.(*IssueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_IssueUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueUpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).IssueUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/IssueUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).IssueUpdateStatus(ctx, req.(*IssueUpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issue.IssueService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).DeleteComment(ctx, req.(*DeleteCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IssueService_ServiceDesc is the grpc.ServiceDesc for IssueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IssueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "issue.IssueService",
	HandlerType: (*IssueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Issues",
			Handler:    _IssueService_Issues_Handler,
		},
		{
			MethodName: "IssuesUser",
			Handler:    _IssueService_IssuesUser_Handler,
		},
		{
			MethodName: "CreateIssue",
			Handler:    _IssueService_CreateIssue_Handler,
		},
		{
			MethodName: "DeleteIssue",
			Handler:    _IssueService_DeleteIssue_Handler,
		},
		{
			MethodName: "Issue",
			Handler:    _IssueService_Issue_Handler,
		},
		{
			MethodName: "IssueUpdateStatus",
			Handler:    _IssueService_IssueUpdateStatus_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _IssueService_AddComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _IssueService_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/issue.proto",
}