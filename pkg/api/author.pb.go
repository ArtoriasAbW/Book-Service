// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: author.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AuthorGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorGetRequest) Reset() {
	*x = AuthorGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorGetRequest) ProtoMessage() {}

func (x *AuthorGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorGetRequest.ProtoReflect.Descriptor instead.
func (*AuthorGetRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorGetRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AuthorGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthorGetResponse) Reset() {
	*x = AuthorGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorGetResponse) ProtoMessage() {}

func (x *AuthorGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorGetResponse.ProtoReflect.Descriptor instead.
func (*AuthorGetResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorGetResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthorGetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AuthorCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthorCreateRequest) Reset() {
	*x = AuthorCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorCreateRequest) ProtoMessage() {}

func (x *AuthorCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorCreateRequest.ProtoReflect.Descriptor instead.
func (*AuthorCreateRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AuthorCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorCreateResponse) Reset() {
	*x = AuthorCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorCreateResponse) ProtoMessage() {}

func (x *AuthorCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorCreateResponse.ProtoReflect.Descriptor instead.
func (*AuthorCreateResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{3}
}

type AuthorListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorListRequest) Reset() {
	*x = AuthorListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorListRequest) ProtoMessage() {}

func (x *AuthorListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorListRequest.ProtoReflect.Descriptor instead.
func (*AuthorListRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{4}
}

type AuthorListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*AuthorListResponse_Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *AuthorListResponse) Reset() {
	*x = AuthorListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorListResponse) ProtoMessage() {}

func (x *AuthorListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorListResponse.ProtoReflect.Descriptor instead.
func (*AuthorListResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{5}
}

func (x *AuthorListResponse) GetAuthors() []*AuthorListResponse_Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type AuthorUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthorUpdateRequest) Reset() {
	*x = AuthorUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorUpdateRequest) ProtoMessage() {}

func (x *AuthorUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorUpdateRequest.ProtoReflect.Descriptor instead.
func (*AuthorUpdateRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{6}
}

func (x *AuthorUpdateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthorUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AuthorUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorUpdateResponse) Reset() {
	*x = AuthorUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorUpdateResponse) ProtoMessage() {}

func (x *AuthorUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorUpdateResponse.ProtoReflect.Descriptor instead.
func (*AuthorUpdateResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{7}
}

type AuthorDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorDeleteRequest) Reset() {
	*x = AuthorDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorDeleteRequest) ProtoMessage() {}

func (x *AuthorDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorDeleteRequest.ProtoReflect.Descriptor instead.
func (*AuthorDeleteRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{8}
}

func (x *AuthorDeleteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AuthorDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorDeleteResponse) Reset() {
	*x = AuthorDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorDeleteResponse) ProtoMessage() {}

func (x *AuthorDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorDeleteResponse.ProtoReflect.Descriptor instead.
func (*AuthorDeleteResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{9}
}

type AuthorListResponse_Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthorListResponse_Author) Reset() {
	*x = AuthorListResponse_Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorListResponse_Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorListResponse_Author) ProtoMessage() {}

func (x *AuthorListResponse_Author) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorListResponse_Author.ProtoReflect.Descriptor instead.
func (*AuthorListResponse_Author) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{5, 0}
}

func (x *AuthorListResponse_Author) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthorListResponse_Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_author_proto protoreflect.FileDescriptor

var file_author_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x22, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29,
	0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7c, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x1a, 0x2c, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x16, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16,
	0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc5, 0x03, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x53, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x47, 0x65, 0x74, 0x12, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5a, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3a,
	0x01, 0x2a, 0x12, 0x52, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x5a, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x1a, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3a,
	0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a,
	0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2f, 0x41, 0x72, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x73, 0x41, 0x62, 0x57, 0x2f, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x30, 0x31, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_author_proto_rawDescOnce sync.Once
	file_author_proto_rawDescData = file_author_proto_rawDesc
)

func file_author_proto_rawDescGZIP() []byte {
	file_author_proto_rawDescOnce.Do(func() {
		file_author_proto_rawDescData = protoimpl.X.CompressGZIP(file_author_proto_rawDescData)
	})
	return file_author_proto_rawDescData
}

var file_author_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_author_proto_goTypes = []interface{}{
	(*AuthorGetRequest)(nil),          // 0: api.AuthorGetRequest
	(*AuthorGetResponse)(nil),         // 1: api.AuthorGetResponse
	(*AuthorCreateRequest)(nil),       // 2: api.AuthorCreateRequest
	(*AuthorCreateResponse)(nil),      // 3: api.AuthorCreateResponse
	(*AuthorListRequest)(nil),         // 4: api.AuthorListRequest
	(*AuthorListResponse)(nil),        // 5: api.AuthorListResponse
	(*AuthorUpdateRequest)(nil),       // 6: api.AuthorUpdateRequest
	(*AuthorUpdateResponse)(nil),      // 7: api.AuthorUpdateResponse
	(*AuthorDeleteRequest)(nil),       // 8: api.AuthorDeleteRequest
	(*AuthorDeleteResponse)(nil),      // 9: api.AuthorDeleteResponse
	(*AuthorListResponse_Author)(nil), // 10: api.AuthorListResponse.Author
}
var file_author_proto_depIdxs = []int32{
	10, // 0: api.AuthorListResponse.authors:type_name -> api.AuthorListResponse.Author
	0,  // 1: api.Author.AuthorGet:input_type -> api.AuthorGetRequest
	2,  // 2: api.Author.AuthorCreate:input_type -> api.AuthorCreateRequest
	4,  // 3: api.Author.AuthorList:input_type -> api.AuthorListRequest
	6,  // 4: api.Author.AuthorUpdate:input_type -> api.AuthorUpdateRequest
	8,  // 5: api.Author.AuthorDelete:input_type -> api.AuthorDeleteRequest
	1,  // 6: api.Author.AuthorGet:output_type -> api.AuthorGetResponse
	3,  // 7: api.Author.AuthorCreate:output_type -> api.AuthorCreateResponse
	5,  // 8: api.Author.AuthorList:output_type -> api.AuthorListResponse
	7,  // 9: api.Author.AuthorUpdate:output_type -> api.AuthorUpdateResponse
	9,  // 10: api.Author.AuthorDelete:output_type -> api.AuthorDeleteResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_author_proto_init() }
func file_author_proto_init() {
	if File_author_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_author_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorGetRequest); i {
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
		file_author_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorGetResponse); i {
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
		file_author_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorCreateRequest); i {
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
		file_author_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorCreateResponse); i {
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
		file_author_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorListRequest); i {
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
		file_author_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorListResponse); i {
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
		file_author_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorUpdateRequest); i {
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
		file_author_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorUpdateResponse); i {
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
		file_author_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorDeleteRequest); i {
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
		file_author_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorDeleteResponse); i {
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
		file_author_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorListResponse_Author); i {
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
			RawDescriptor: file_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_author_proto_goTypes,
		DependencyIndexes: file_author_proto_depIdxs,
		MessageInfos:      file_author_proto_msgTypes,
	}.Build()
	File_author_proto = out.File
	file_author_proto_rawDesc = nil
	file_author_proto_goTypes = nil
	file_author_proto_depIdxs = nil
}
