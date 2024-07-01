package service

import (
 pb "community/generated"
 "community/storage/postgres"
 "context"
)

type Server struct {
 pb.UnimplementedCommunityServiceServer
 C postgres.NewCommunity
}

func NewCommunity(c postgres.NewCommunity) *Server {
 return &Server{C: c}
}

func (S *Server) CreateGroup(ctx context.Context, group *pb.Group) (*pb.Status, error) {
 status, err := S.C.CreateGroup(group)
 if err != nil {
  return nil, err
 }
 return status, err
}

func (S *Server) GetGroup(ctx context.Context, groupId *pb.GroupId) (*pb.Group, error) {
 group, err := S.C.GetGroup(groupId)
 if err != nil {
  return nil, err
 }
 return group, nil
}

func (S *Server) UpdateGroup(ctx context.Context, group *pb.Group) (*pb.Status, error) {
 status, err := S.C.UpdateGroup(group)
 if err != nil {
  return nil, err
 }
 return status, nil
}

func (S *Server) DeleteGroup(ctx context.Context, groupId *pb.GroupId) (*pb.Status, error) {
 status, err := S.C.DeleteGroup(groupId)
 if err != nil {
  return nil, err
 }
 return status, nil
}

func (S *Server) GetAllGroups(ctx context.Context, req *pb.Req) (*pb.Groups, error) {
 grops, err := S.C.GetAllGroups(req)
 if err != nil {
  return nil, err
 }
 return grops, nil
}

func (S *Server) JoinGroupUser(ctx context.Context, joinLeave *pb.JoinLeave) (*pb.Status, error) {
 status, err := S.C.JoinGroupUser(joinLeave)
 if err != nil {
  return nil, err
 }
 return status, nil
}

func (S *Server) LeaveGroupUser(ctx context.Context, joinLeave *pb.JoinLeave) (*pb.Status, error) {
 status, err := S.C.JoinGroupUser(joinLeave)
 if err != nil {
  return nil, err
 }
 return status, nil
}

func (S *Server) UpdateGroupMeber(ctx context.Context, userRole *pb.UserRole) (*pb.Status, error) {
 status, err := S.C.UpdateGroupMeber(userRole)
 if err != nil {
  return nil, err
 }
 return status, nil
}

func (S *Server) CreatePost(ctx context.Context, post *pb.Post) (*pb.Status, error) {
 status, err := S.C.CreatePost(post)
 if err != nil {
  return nil, err
 }
 return status, nil
}

func (S *Server) UpdatePost(ctx context.Context, post *pb.Post) (*pb.Status, error) {
 status, err := S.C.UpdatePost(post)
 if err != nil {
  return nil, err
 }
 return status, nil
}

func (S *Server) DeletePost(ctx context.Context, postId *pb.PostId) (*pb.Status, error) {
 status, err := S.C.DeletePost(postId)
 if err != nil {
  return nil, err
 }
 return status, nil
}

func (S *Server) GetPost(ctx context.Context, postId *pb.PostId) (*pb.Post, error) {
 post, err := S.C.GetPost(postId)
 if err != nil {
  return nil, err
 }
 return post, nil
}

func (S *Server) GetGroupPost(ctx context.Context, gropPost *pb.GroupPost) (*pb.Post, error) {
 post, err := S.C.GetGroupPost(gropPost)
 if err != nil {
  return nil, err
 }
 return post, nil
}

func (S *Server) CreatePostComments(ctx context.Context, comment *pb.Comment) (*pb.Status, error) {
 status, err := S.C.CreatePostComments(comment)
 if err != nil {
  return nil, err
 }
 return status, nil
}

func (S *Server) GetPostComments(ctx context.Context, postComment *pb.PostComment) (*pb.Comment, error) {
 comment, err := S.C.GetPostComments(postComment)
 if err != nil {
  return nil, err
 }
 return comment, nil
}
