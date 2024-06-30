// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: coomunity.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{0}
}

type GroupId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *GroupId) Reset() {
	*x = GroupId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupId) ProtoMessage() {}

func (x *GroupId) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupId.ProtoReflect.Descriptor instead.
func (*GroupId) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{1}
}

func (x *GroupId) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId     *GroupId `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedBy   string   `protobuf:"bytes,4,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	CreatedAt   string   `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{2}
}

func (x *Group) GetGroupId() *GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Group) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Group) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type Groups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Groups) Reset() {
	*x = Groups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Groups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Groups) ProtoMessage() {}

func (x *Groups) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Groups.ProtoReflect.Descriptor instead.
func (*Groups) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{4}
}

func (x *Groups) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type JoinLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId  *GroupId `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	UserId   string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Role     string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	JoinedAt string   `protobuf:"bytes,4,opt,name=joinedAt,proto3" json:"joinedAt,omitempty"`
}

func (x *JoinLeave) Reset() {
	*x = JoinLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinLeave) ProtoMessage() {}

func (x *JoinLeave) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinLeave.ProtoReflect.Descriptor instead.
func (*JoinLeave) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{5}
}

func (x *JoinLeave) GetGroupId() *GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

func (x *JoinLeave) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *JoinLeave) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *JoinLeave) GetJoinedAt() string {
	if x != nil {
		return x.JoinedAt
	}
	return ""
}

type UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId *GroupId `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	UserId  string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Role    string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{6}
}

func (x *UserRole) GetGroupId() *GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

func (x *UserRole) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserRole) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type PostId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *PostId) Reset() {
	*x = PostId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostId) ProtoMessage() {}

func (x *PostId) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostId.ProtoReflect.Descriptor instead.
func (*PostId) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{7}
}

func (x *PostId) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId    *PostId  `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty"`
	GroupId   *GroupId `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	UserId    string   `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Content   string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string   `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{8}
}

func (x *Post) GetPostId() *PostId {
	if x != nil {
		return x.PostId
	}
	return nil
}

func (x *Post) GetGroupId() *GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

func (x *Post) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GroupPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId *GroupId `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	PostId  *PostId  `protobuf:"bytes,2,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *GroupPost) Reset() {
	*x = GroupPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPost) ProtoMessage() {}

func (x *GroupPost) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPost.ProtoReflect.Descriptor instead.
func (*GroupPost) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{9}
}

func (x *GroupPost) GetGroupId() *GroupId {
	if x != nil {
		return x.GroupId
	}
	return nil
}

func (x *GroupPost) GetPostId() *PostId {
	if x != nil {
		return x.PostId
	}
	return nil
}

type CommenId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId string `protobuf:"bytes,1,opt,name=commentId,proto3" json:"commentId,omitempty"`
}

func (x *CommenId) Reset() {
	*x = CommenId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommenId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommenId) ProtoMessage() {}

func (x *CommenId) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommenId.ProtoReflect.Descriptor instead.
func (*CommenId) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{10}
}

func (x *CommenId) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId *CommenId `protobuf:"bytes,1,opt,name=commentId,proto3" json:"commentId,omitempty"`
	PostId    *PostId   `protobuf:"bytes,2,opt,name=postId,proto3" json:"postId,omitempty"`
	UserId    string    `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Content   string    `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string    `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{11}
}

func (x *Comment) GetCommentId() *CommenId {
	if x != nil {
		return x.CommentId
	}
	return nil
}

func (x *Comment) GetPostId() *PostId {
	if x != nil {
		return x.PostId
	}
	return nil
}

func (x *Comment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type PostComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId    *PostId   `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty"`
	CommentId *CommenId `protobuf:"bytes,2,opt,name=commentId,proto3" json:"commentId,omitempty"`
}

func (x *PostComment) Reset() {
	*x = PostComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coomunity_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostComment) ProtoMessage() {}

func (x *PostComment) ProtoReflect() protoreflect.Message {
	mi := &file_coomunity_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostComment.ProtoReflect.Descriptor instead.
func (*PostComment) Descriptor() ([]byte, []int) {
	return file_coomunity_proto_rawDescGZIP(), []int{12}
}

func (x *PostComment) GetPostId() *PostId {
	if x != nil {
		return x.PostId
	}
	return nil
}

func (x *PostComment) GetCommentId() *CommenId {
	if x != nil {
		return x.CommentId
	}
	return nil
}

var File_coomunity_proto protoreflect.FileDescriptor

var file_coomunity_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x22, 0x05, 0x0a, 0x03,
	0x52, 0x65, 0x71, 0x22, 0x23, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x2c, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x32, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x28,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x09, 0x4a, 0x6f, 0x69,
	0x6e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x22, 0x64, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x07, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x22, 0xaf, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x64, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x08,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x49, 0x64, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x6b, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x49, 0x64, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x32, 0xde, 0x06,
	0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x11, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0d, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x62, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x1a,
	0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coomunity_proto_rawDescOnce sync.Once
	file_coomunity_proto_rawDescData = file_coomunity_proto_rawDesc
)

func file_coomunity_proto_rawDescGZIP() []byte {
	file_coomunity_proto_rawDescOnce.Do(func() {
		file_coomunity_proto_rawDescData = protoimpl.X.CompressGZIP(file_coomunity_proto_rawDescData)
	})
	return file_coomunity_proto_rawDescData
}

var file_coomunity_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_coomunity_proto_goTypes = []any{
	(*Req)(nil),         // 0: community.Req
	(*GroupId)(nil),     // 1: community.GroupId
	(*Group)(nil),       // 2: community.Group
	(*Status)(nil),      // 3: community.Status
	(*Groups)(nil),      // 4: community.Groups
	(*JoinLeave)(nil),   // 5: community.JoinLeave
	(*UserRole)(nil),    // 6: community.UserRole
	(*PostId)(nil),      // 7: community.PostId
	(*Post)(nil),        // 8: community.Post
	(*GroupPost)(nil),   // 9: community.GroupPost
	(*CommenId)(nil),    // 10: community.CommenId
	(*Comment)(nil),     // 11: community.Comment
	(*PostComment)(nil), // 12: community.PostComment
}
var file_coomunity_proto_depIdxs = []int32{
	1,  // 0: community.Group.groupId:type_name -> community.GroupId
	2,  // 1: community.Groups.groups:type_name -> community.Group
	1,  // 2: community.JoinLeave.groupId:type_name -> community.GroupId
	1,  // 3: community.UserRole.GroupId:type_name -> community.GroupId
	7,  // 4: community.Post.postId:type_name -> community.PostId
	1,  // 5: community.Post.groupId:type_name -> community.GroupId
	1,  // 6: community.GroupPost.groupId:type_name -> community.GroupId
	7,  // 7: community.GroupPost.postId:type_name -> community.PostId
	10, // 8: community.Comment.commentId:type_name -> community.CommenId
	7,  // 9: community.Comment.postId:type_name -> community.PostId
	7,  // 10: community.PostComment.postId:type_name -> community.PostId
	10, // 11: community.PostComment.commentId:type_name -> community.CommenId
	2,  // 12: community.CommunityService.CreateGroup:input_type -> community.Group
	1,  // 13: community.CommunityService.GetGroup:input_type -> community.GroupId
	2,  // 14: community.CommunityService.UpdateGroup:input_type -> community.Group
	1,  // 15: community.CommunityService.DeleteGroup:input_type -> community.GroupId
	0,  // 16: community.CommunityService.GetAllGroups:input_type -> community.Req
	5,  // 17: community.CommunityService.JoinGroupUser:input_type -> community.JoinLeave
	5,  // 18: community.CommunityService.LeaveGroupUser:input_type -> community.JoinLeave
	6,  // 19: community.CommunityService.UpdateGroupMeber:input_type -> community.UserRole
	8,  // 20: community.CommunityService.CreatePost:input_type -> community.Post
	8,  // 21: community.CommunityService.UpdatePost:input_type -> community.Post
	7,  // 22: community.CommunityService.DeletePost:input_type -> community.PostId
	7,  // 23: community.CommunityService.GetPost:input_type -> community.PostId
	9,  // 24: community.CommunityService.GetGroupPost:input_type -> community.GroupPost
	11, // 25: community.CommunityService.CreatePostComments:input_type -> community.Comment
	12, // 26: community.CommunityService.GetPostComments:input_type -> community.PostComment
	3,  // 27: community.CommunityService.CreateGroup:output_type -> community.Status
	2,  // 28: community.CommunityService.GetGroup:output_type -> community.Group
	3,  // 29: community.CommunityService.UpdateGroup:output_type -> community.Status
	3,  // 30: community.CommunityService.DeleteGroup:output_type -> community.Status
	4,  // 31: community.CommunityService.GetAllGroups:output_type -> community.Groups
	3,  // 32: community.CommunityService.JoinGroupUser:output_type -> community.Status
	3,  // 33: community.CommunityService.LeaveGroupUser:output_type -> community.Status
	3,  // 34: community.CommunityService.UpdateGroupMeber:output_type -> community.Status
	3,  // 35: community.CommunityService.CreatePost:output_type -> community.Status
	3,  // 36: community.CommunityService.UpdatePost:output_type -> community.Status
	3,  // 37: community.CommunityService.DeletePost:output_type -> community.Status
	8,  // 38: community.CommunityService.GetPost:output_type -> community.Post
	8,  // 39: community.CommunityService.GetGroupPost:output_type -> community.Post
	3,  // 40: community.CommunityService.CreatePostComments:output_type -> community.Status
	11, // 41: community.CommunityService.GetPostComments:output_type -> community.Comment
	27, // [27:42] is the sub-list for method output_type
	12, // [12:27] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_coomunity_proto_init() }
func file_coomunity_proto_init() {
	if File_coomunity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coomunity_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Req); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GroupId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Group); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Groups); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*JoinLeave); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UserRole); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PostId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GroupPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*CommenId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coomunity_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*PostComment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coomunity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coomunity_proto_goTypes,
		DependencyIndexes: file_coomunity_proto_depIdxs,
		MessageInfos:      file_coomunity_proto_msgTypes,
	}.Build()
	File_coomunity_proto = out.File
	file_coomunity_proto_rawDesc = nil
	file_coomunity_proto_goTypes = nil
	file_coomunity_proto_depIdxs = nil
}
