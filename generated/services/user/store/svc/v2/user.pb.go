// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: services/user/store/svc/v2/user.proto

package svcv2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email     *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	IsBanned  *wrapperspb.BoolValue   `protobuf:"bytes,4,opt,name=is_banned,json=isBanned,proto3" json:"is_banned,omitempty"`
	CreatedAt *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *User) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *User) GetEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *User) GetIsBanned() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsBanned
	}
	return nil
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *CreateUserRequest) GetEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UpdateUsernameRequest) Reset() {
	*x = UpdateUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUsernameRequest) ProtoMessage() {}

func (x *UpdateUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUsernameRequest.ProtoReflect.Descriptor instead.
func (*UpdateUsernameRequest) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUsernameRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdateUsernameRequest) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

type UpdateUsernameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUsernameResponse) Reset() {
	*x = UpdateUsernameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUsernameResponse) ProtoMessage() {}

func (x *UpdateUsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUsernameResponse.ProtoReflect.Descriptor instead.
func (*UpdateUsernameResponse) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUsernameResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByIDRequest) Reset() {
	*x = GetUserByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDRequest) ProtoMessage() {}

func (x *GetUserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIDRequest) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserByIDRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetUserByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByIDResponse) Reset() {
	*x = GetUserByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDResponse) ProtoMessage() {}

func (x *GetUserByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDResponse.ProtoReflect.Descriptor instead.
func (*GetUserByIDResponse) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserByIDResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetUserByEmailRequest) Reset() {
	*x = GetUserByEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByEmailRequest) ProtoMessage() {}

func (x *GetUserByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByEmailRequest.ProtoReflect.Descriptor instead.
func (*GetUserByEmailRequest) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserByEmailRequest) GetEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

type GetUserByEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByEmailResponse) Reset() {
	*x = GetUserByEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByEmailResponse) ProtoMessage() {}

func (x *GetUserByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByEmailResponse.ProtoReflect.Descriptor instead.
func (*GetUserByEmailResponse) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserByEmailResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserByUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetUserByUsernameRequest) Reset() {
	*x = GetUserByUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUsernameRequest) ProtoMessage() {}

func (x *GetUserByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUsernameRequest.ProtoReflect.Descriptor instead.
func (*GetUserByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserByUsernameRequest) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

type GetUserByUsernameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByUsernameResponse) Reset() {
	*x = GetUserByUsernameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_store_svc_v2_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByUsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUsernameResponse) ProtoMessage() {}

func (x *GetUserByUsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_store_svc_v2_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUsernameResponse.ProtoReflect.Descriptor instead.
func (*GetUserByUsernameResponse) Descriptor() ([]byte, []int) {
	return file_services_user_store_svc_v2_user_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserByUsernameResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_services_user_store_svc_v2_user_proto protoreflect.FileDescriptor

var file_services_user_store_svc_v2_user_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x37, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x73, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4a, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x7f, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x54, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x51, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x42, 0x84, 0x02,
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x32,
	0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x65, 0x72, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x76, 0x63,
	0x2f, 0x76, 0x32, 0x3b, 0x73, 0x76, 0x63, 0x76, 0x32, 0xa2, 0x02, 0x04, 0x53, 0x55, 0x53, 0x53,
	0xaa, 0x02, 0x1a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x76, 0x63, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x5c, 0x53, 0x76, 0x63, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x26, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x5c, 0x53, 0x76, 0x63, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a,
	0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x53, 0x76, 0x63,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_user_store_svc_v2_user_proto_rawDescOnce sync.Once
	file_services_user_store_svc_v2_user_proto_rawDescData = file_services_user_store_svc_v2_user_proto_rawDesc
)

func file_services_user_store_svc_v2_user_proto_rawDescGZIP() []byte {
	file_services_user_store_svc_v2_user_proto_rawDescOnce.Do(func() {
		file_services_user_store_svc_v2_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_user_store_svc_v2_user_proto_rawDescData)
	})
	return file_services_user_store_svc_v2_user_proto_rawDescData
}

var file_services_user_store_svc_v2_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_services_user_store_svc_v2_user_proto_goTypes = []interface{}{
	(*User)(nil),                      // 0: services.user.store.svc.v2.User
	(*CreateUserRequest)(nil),         // 1: services.user.store.svc.v2.CreateUserRequest
	(*CreateUserResponse)(nil),        // 2: services.user.store.svc.v2.CreateUserResponse
	(*UpdateUsernameRequest)(nil),     // 3: services.user.store.svc.v2.UpdateUsernameRequest
	(*UpdateUsernameResponse)(nil),    // 4: services.user.store.svc.v2.UpdateUsernameResponse
	(*GetUserByIDRequest)(nil),        // 5: services.user.store.svc.v2.GetUserByIDRequest
	(*GetUserByIDResponse)(nil),       // 6: services.user.store.svc.v2.GetUserByIDResponse
	(*GetUserByEmailRequest)(nil),     // 7: services.user.store.svc.v2.GetUserByEmailRequest
	(*GetUserByEmailResponse)(nil),    // 8: services.user.store.svc.v2.GetUserByEmailResponse
	(*GetUserByUsernameRequest)(nil),  // 9: services.user.store.svc.v2.GetUserByUsernameRequest
	(*GetUserByUsernameResponse)(nil), // 10: services.user.store.svc.v2.GetUserByUsernameResponse
	(*wrapperspb.StringValue)(nil),    // 11: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),      // 12: google.protobuf.BoolValue
	(*timestamppb.Timestamp)(nil),     // 13: google.protobuf.Timestamp
}
var file_services_user_store_svc_v2_user_proto_depIdxs = []int32{
	11, // 0: services.user.store.svc.v2.User.id:type_name -> google.protobuf.StringValue
	11, // 1: services.user.store.svc.v2.User.username:type_name -> google.protobuf.StringValue
	11, // 2: services.user.store.svc.v2.User.email:type_name -> google.protobuf.StringValue
	12, // 3: services.user.store.svc.v2.User.is_banned:type_name -> google.protobuf.BoolValue
	13, // 4: services.user.store.svc.v2.User.created_at:type_name -> google.protobuf.Timestamp
	13, // 5: services.user.store.svc.v2.User.updated_at:type_name -> google.protobuf.Timestamp
	11, // 6: services.user.store.svc.v2.CreateUserRequest.username:type_name -> google.protobuf.StringValue
	11, // 7: services.user.store.svc.v2.CreateUserRequest.email:type_name -> google.protobuf.StringValue
	0,  // 8: services.user.store.svc.v2.CreateUserResponse.user:type_name -> services.user.store.svc.v2.User
	11, // 9: services.user.store.svc.v2.UpdateUsernameRequest.id:type_name -> google.protobuf.StringValue
	11, // 10: services.user.store.svc.v2.UpdateUsernameRequest.username:type_name -> google.protobuf.StringValue
	0,  // 11: services.user.store.svc.v2.UpdateUsernameResponse.user:type_name -> services.user.store.svc.v2.User
	11, // 12: services.user.store.svc.v2.GetUserByIDRequest.id:type_name -> google.protobuf.StringValue
	0,  // 13: services.user.store.svc.v2.GetUserByIDResponse.user:type_name -> services.user.store.svc.v2.User
	11, // 14: services.user.store.svc.v2.GetUserByEmailRequest.email:type_name -> google.protobuf.StringValue
	0,  // 15: services.user.store.svc.v2.GetUserByEmailResponse.user:type_name -> services.user.store.svc.v2.User
	11, // 16: services.user.store.svc.v2.GetUserByUsernameRequest.username:type_name -> google.protobuf.StringValue
	0,  // 17: services.user.store.svc.v2.GetUserByUsernameResponse.user:type_name -> services.user.store.svc.v2.User
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_services_user_store_svc_v2_user_proto_init() }
func file_services_user_store_svc_v2_user_proto_init() {
	if File_services_user_store_svc_v2_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_user_store_svc_v2_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUsernameRequest); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUsernameResponse); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDRequest); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDResponse); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByEmailRequest); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByEmailResponse); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByUsernameRequest); i {
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
		file_services_user_store_svc_v2_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByUsernameResponse); i {
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
			RawDescriptor: file_services_user_store_svc_v2_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_user_store_svc_v2_user_proto_goTypes,
		DependencyIndexes: file_services_user_store_svc_v2_user_proto_depIdxs,
		MessageInfos:      file_services_user_store_svc_v2_user_proto_msgTypes,
	}.Build()
	File_services_user_store_svc_v2_user_proto = out.File
	file_services_user_store_svc_v2_user_proto_rawDesc = nil
	file_services_user_store_svc_v2_user_proto_goTypes = nil
	file_services_user_store_svc_v2_user_proto_depIdxs = nil
}
