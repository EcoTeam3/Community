// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: coomunity.proto

package generated

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CommunityService_CreateGroup_FullMethodName        = "/community.CommunityService/CreateGroup"
	CommunityService_GetGroup_FullMethodName           = "/community.CommunityService/GetGroup"
	CommunityService_UpdateGroup_FullMethodName        = "/community.CommunityService/UpdateGroup"
	CommunityService_DeleteGroup_FullMethodName        = "/community.CommunityService/DeleteGroup"
	CommunityService_GetAllGroups_FullMethodName       = "/community.CommunityService/GetAllGroups"
	CommunityService_JoinGroupUser_FullMethodName      = "/community.CommunityService/JoinGroupUser"
	CommunityService_LeaveGroupUser_FullMethodName     = "/community.CommunityService/LeaveGroupUser"
	CommunityService_UpdateGroupMeber_FullMethodName   = "/community.CommunityService/UpdateGroupMeber"
	CommunityService_CreatePost_FullMethodName         = "/community.CommunityService/CreatePost"
	CommunityService_UpdatePost_FullMethodName         = "/community.CommunityService/UpdatePost"
	CommunityService_DeletePost_FullMethodName         = "/community.CommunityService/DeletePost"
	CommunityService_GetPost_FullMethodName            = "/community.CommunityService/GetPost"
	CommunityService_GetGroupPost_FullMethodName       = "/community.CommunityService/GetGroupPost"
	CommunityService_CreatePostComments_FullMethodName = "/community.CommunityService/CreatePostComments"
	CommunityService_GetPostComments_FullMethodName    = "/community.CommunityService/GetPostComments"
)

// CommunityServiceClient is the client API for CommunityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunityServiceClient interface {
	CreateGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Status, error)
	GetGroup(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*Group, error)
	UpdateGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Status, error)
	DeleteGroup(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*Status, error)
	GetAllGroups(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Groups, error)
	JoinGroupUser(ctx context.Context, in *JoinLeave, opts ...grpc.CallOption) (*Status, error)
	LeaveGroupUser(ctx context.Context, in *JoinLeave, opts ...grpc.CallOption) (*Status, error)
	UpdateGroupMeber(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*Status, error)
	CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error)
	UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error)
	DeletePost(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Status, error)
	GetPost(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Post, error)
	GetGroupPost(ctx context.Context, in *GroupPost, opts ...grpc.CallOption) (*Post, error)
	CreatePostComments(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Status, error)
	GetPostComments(ctx context.Context, in *PostComment, opts ...grpc.CallOption) (*Comment, error)
}

type communityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunityServiceClient(cc grpc.ClientConnInterface) CommunityServiceClient {
	return &communityServiceClient{cc}
}

func (c *communityServiceClient) CreateGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_CreateGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetGroup(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*Group, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Group)
	err := c.cc.Invoke(ctx, CommunityService_GetGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) UpdateGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_UpdateGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) DeleteGroup(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_DeleteGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetAllGroups(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Groups, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Groups)
	err := c.cc.Invoke(ctx, CommunityService_GetAllGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) JoinGroupUser(ctx context.Context, in *JoinLeave, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_JoinGroupUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) LeaveGroupUser(ctx context.Context, in *JoinLeave, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_LeaveGroupUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) UpdateGroupMeber(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_UpdateGroupMeber_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_UpdatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) DeletePost(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_DeletePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetPost(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, CommunityService_GetPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetGroupPost(ctx context.Context, in *GroupPost, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, CommunityService_GetGroupPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) CreatePostComments(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, CommunityService_CreatePostComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetPostComments(ctx context.Context, in *PostComment, opts ...grpc.CallOption) (*Comment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Comment)
	err := c.cc.Invoke(ctx, CommunityService_GetPostComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunityServiceServer is the server API for CommunityService service.
// All implementations must embed UnimplementedCommunityServiceServer
// for forward compatibility
type CommunityServiceServer interface {
	CreateGroup(context.Context, *Group) (*Status, error)
	GetGroup(context.Context, *GroupId) (*Group, error)
	UpdateGroup(context.Context, *Group) (*Status, error)
	DeleteGroup(context.Context, *GroupId) (*Status, error)
	GetAllGroups(context.Context, *Req) (*Groups, error)
	JoinGroupUser(context.Context, *JoinLeave) (*Status, error)
	LeaveGroupUser(context.Context, *JoinLeave) (*Status, error)
	UpdateGroupMeber(context.Context, *UserRole) (*Status, error)
	CreatePost(context.Context, *Post) (*Status, error)
	UpdatePost(context.Context, *Post) (*Status, error)
	DeletePost(context.Context, *PostId) (*Status, error)
	GetPost(context.Context, *PostId) (*Post, error)
	GetGroupPost(context.Context, *GroupPost) (*Post, error)
	CreatePostComments(context.Context, *Comment) (*Status, error)
	GetPostComments(context.Context, *PostComment) (*Comment, error)
	mustEmbedUnimplementedCommunityServiceServer()
}

// UnimplementedCommunityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommunityServiceServer struct {
}

func (UnimplementedCommunityServiceServer) CreateGroup(context.Context, *Group) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedCommunityServiceServer) GetGroup(context.Context, *GroupId) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedCommunityServiceServer) UpdateGroup(context.Context, *Group) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedCommunityServiceServer) DeleteGroup(context.Context, *GroupId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedCommunityServiceServer) GetAllGroups(context.Context, *Req) (*Groups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGroups not implemented")
}
func (UnimplementedCommunityServiceServer) JoinGroupUser(context.Context, *JoinLeave) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroupUser not implemented")
}
func (UnimplementedCommunityServiceServer) LeaveGroupUser(context.Context, *JoinLeave) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGroupUser not implemented")
}
func (UnimplementedCommunityServiceServer) UpdateGroupMeber(context.Context, *UserRole) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupMeber not implemented")
}
func (UnimplementedCommunityServiceServer) CreatePost(context.Context, *Post) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedCommunityServiceServer) UpdatePost(context.Context, *Post) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedCommunityServiceServer) DeletePost(context.Context, *PostId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedCommunityServiceServer) GetPost(context.Context, *PostId) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedCommunityServiceServer) GetGroupPost(context.Context, *GroupPost) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupPost not implemented")
}
func (UnimplementedCommunityServiceServer) CreatePostComments(context.Context, *Comment) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePostComments not implemented")
}
func (UnimplementedCommunityServiceServer) GetPostComments(context.Context, *PostComment) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostComments not implemented")
}
func (UnimplementedCommunityServiceServer) mustEmbedUnimplementedCommunityServiceServer() {}

// UnsafeCommunityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunityServiceServer will
// result in compilation errors.
type UnsafeCommunityServiceServer interface {
	mustEmbedUnimplementedCommunityServiceServer()
}

func RegisterCommunityServiceServer(s grpc.ServiceRegistrar, srv CommunityServiceServer) {
	s.RegisterService(&CommunityService_ServiceDesc, srv)
}

func _CommunityService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).CreateGroup(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetGroup(ctx, req.(*GroupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_UpdateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).UpdateGroup(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_DeleteGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).DeleteGroup(ctx, req.(*GroupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetAllGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetAllGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_GetAllGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetAllGroups(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_JoinGroupUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinLeave)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).JoinGroupUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_JoinGroupUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).JoinGroupUser(ctx, req.(*JoinLeave))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_LeaveGroupUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinLeave)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).LeaveGroupUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_LeaveGroupUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).LeaveGroupUser(ctx, req.(*JoinLeave))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_UpdateGroupMeber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).UpdateGroupMeber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_UpdateGroupMeber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).UpdateGroupMeber(ctx, req.(*UserRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).CreatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).UpdatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).DeletePost(ctx, req.(*PostId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetPost(ctx, req.(*PostId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetGroupPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetGroupPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_GetGroupPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetGroupPost(ctx, req.(*GroupPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_CreatePostComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).CreatePostComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_CreatePostComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).CreatePostComments(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetPostComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetPostComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunityService_GetPostComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetPostComments(ctx, req.(*PostComment))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunityService_ServiceDesc is the grpc.ServiceDesc for CommunityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "community.CommunityService",
	HandlerType: (*CommunityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _CommunityService_CreateGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _CommunityService_GetGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _CommunityService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _CommunityService_DeleteGroup_Handler,
		},
		{
			MethodName: "GetAllGroups",
			Handler:    _CommunityService_GetAllGroups_Handler,
		},
		{
			MethodName: "JoinGroupUser",
			Handler:    _CommunityService_JoinGroupUser_Handler,
		},
		{
			MethodName: "LeaveGroupUser",
			Handler:    _CommunityService_LeaveGroupUser_Handler,
		},
		{
			MethodName: "UpdateGroupMeber",
			Handler:    _CommunityService_UpdateGroupMeber_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _CommunityService_CreatePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _CommunityService_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _CommunityService_DeletePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _CommunityService_GetPost_Handler,
		},
		{
			MethodName: "GetGroupPost",
			Handler:    _CommunityService_GetGroupPost_Handler,
		},
		{
			MethodName: "CreatePostComments",
			Handler:    _CommunityService_CreatePostComments_Handler,
		},
		{
			MethodName: "GetPostComments",
			Handler:    _CommunityService_GetPostComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coomunity.proto",
}
