// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: battle/v1/monster.proto

package v1

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

type CreateMonsterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateMonsterRequest) Reset() {
	*x = CreateMonsterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMonsterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMonsterRequest) ProtoMessage() {}

func (x *CreateMonsterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMonsterRequest.ProtoReflect.Descriptor instead.
func (*CreateMonsterRequest) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMonsterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateMonsterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateMonsterReply) Reset() {
	*x = CreateMonsterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMonsterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMonsterReply) ProtoMessage() {}

func (x *CreateMonsterReply) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMonsterReply.ProtoReflect.Descriptor instead.
func (*CreateMonsterReply) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMonsterReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateMonsterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateMonsterRequest) Reset() {
	*x = UpdateMonsterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMonsterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMonsterRequest) ProtoMessage() {}

func (x *UpdateMonsterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMonsterRequest.ProtoReflect.Descriptor instead.
func (*UpdateMonsterRequest) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMonsterRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateMonsterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateMonsterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMonsterReply) Reset() {
	*x = UpdateMonsterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMonsterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMonsterReply) ProtoMessage() {}

func (x *UpdateMonsterReply) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMonsterReply.ProtoReflect.Descriptor instead.
func (*UpdateMonsterReply) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{3}
}

type DeleteMonsterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMonsterRequest) Reset() {
	*x = DeleteMonsterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMonsterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMonsterRequest) ProtoMessage() {}

func (x *DeleteMonsterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMonsterRequest.ProtoReflect.Descriptor instead.
func (*DeleteMonsterRequest) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMonsterRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteMonsterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMonsterReply) Reset() {
	*x = DeleteMonsterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMonsterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMonsterReply) ProtoMessage() {}

func (x *DeleteMonsterReply) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMonsterReply.ProtoReflect.Descriptor instead.
func (*DeleteMonsterReply) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{5}
}

type GetMonsterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMonsterRequest) Reset() {
	*x = GetMonsterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMonsterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMonsterRequest) ProtoMessage() {}

func (x *GetMonsterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMonsterRequest.ProtoReflect.Descriptor instead.
func (*GetMonsterRequest) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{6}
}

func (x *GetMonsterRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMonsterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetMonsterReply) Reset() {
	*x = GetMonsterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMonsterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMonsterReply) ProtoMessage() {}

func (x *GetMonsterReply) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMonsterReply.ProtoReflect.Descriptor instead.
func (*GetMonsterReply) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{7}
}

func (x *GetMonsterReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListMonsterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListMonsterRequest) Reset() {
	*x = ListMonsterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMonsterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMonsterRequest) ProtoMessage() {}

func (x *ListMonsterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMonsterRequest.ProtoReflect.Descriptor instead.
func (*ListMonsterRequest) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{8}
}

type ListMonsterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ListMonsterReply_Data `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListMonsterReply) Reset() {
	*x = ListMonsterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMonsterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMonsterReply) ProtoMessage() {}

func (x *ListMonsterReply) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMonsterReply.ProtoReflect.Descriptor instead.
func (*ListMonsterReply) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{9}
}

func (x *ListMonsterReply) GetData() []*ListMonsterReply_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListMonsterReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListMonsterReply_Data) Reset() {
	*x = ListMonsterReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_v1_monster_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMonsterReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMonsterReply_Data) ProtoMessage() {}

func (x *ListMonsterReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_battle_v1_monster_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMonsterReply_Data.ProtoReflect.Descriptor instead.
func (*ListMonsterReply_Data) Descriptor() ([]byte, []int) {
	return file_battle_v1_monster_proto_rawDescGZIP(), []int{9, 0}
}

func (x *ListMonsterReply_Data) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListMonsterReply_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_battle_v1_monster_proto protoreflect.FileDescriptor

var file_battle_v1_monster_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x79, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x2a, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xc1, 0x03, 0x0a, 0x07, 0x4d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x59, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x1c, 0x5a, 0x1a,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_battle_v1_monster_proto_rawDescOnce sync.Once
	file_battle_v1_monster_proto_rawDescData = file_battle_v1_monster_proto_rawDesc
)

func file_battle_v1_monster_proto_rawDescGZIP() []byte {
	file_battle_v1_monster_proto_rawDescOnce.Do(func() {
		file_battle_v1_monster_proto_rawDescData = protoimpl.X.CompressGZIP(file_battle_v1_monster_proto_rawDescData)
	})
	return file_battle_v1_monster_proto_rawDescData
}

var file_battle_v1_monster_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_battle_v1_monster_proto_goTypes = []interface{}{
	(*CreateMonsterRequest)(nil),  // 0: api.service.v1.CreateMonsterRequest
	(*CreateMonsterReply)(nil),    // 1: api.service.v1.CreateMonsterReply
	(*UpdateMonsterRequest)(nil),  // 2: api.service.v1.UpdateMonsterRequest
	(*UpdateMonsterReply)(nil),    // 3: api.service.v1.UpdateMonsterReply
	(*DeleteMonsterRequest)(nil),  // 4: api.service.v1.DeleteMonsterRequest
	(*DeleteMonsterReply)(nil),    // 5: api.service.v1.DeleteMonsterReply
	(*GetMonsterRequest)(nil),     // 6: api.service.v1.GetMonsterRequest
	(*GetMonsterReply)(nil),       // 7: api.service.v1.GetMonsterReply
	(*ListMonsterRequest)(nil),    // 8: api.service.v1.ListMonsterRequest
	(*ListMonsterReply)(nil),      // 9: api.service.v1.ListMonsterReply
	(*ListMonsterReply_Data)(nil), // 10: api.service.v1.ListMonsterReply.Data
}
var file_battle_v1_monster_proto_depIdxs = []int32{
	10, // 0: api.service.v1.ListMonsterReply.data:type_name -> api.service.v1.ListMonsterReply.Data
	0,  // 1: api.service.v1.Monster.CreateMonster:input_type -> api.service.v1.CreateMonsterRequest
	2,  // 2: api.service.v1.Monster.UpdateMonster:input_type -> api.service.v1.UpdateMonsterRequest
	4,  // 3: api.service.v1.Monster.DeleteMonster:input_type -> api.service.v1.DeleteMonsterRequest
	6,  // 4: api.service.v1.Monster.GetMonster:input_type -> api.service.v1.GetMonsterRequest
	8,  // 5: api.service.v1.Monster.ListMonster:input_type -> api.service.v1.ListMonsterRequest
	1,  // 6: api.service.v1.Monster.CreateMonster:output_type -> api.service.v1.CreateMonsterReply
	3,  // 7: api.service.v1.Monster.UpdateMonster:output_type -> api.service.v1.UpdateMonsterReply
	5,  // 8: api.service.v1.Monster.DeleteMonster:output_type -> api.service.v1.DeleteMonsterReply
	7,  // 9: api.service.v1.Monster.GetMonster:output_type -> api.service.v1.GetMonsterReply
	9,  // 10: api.service.v1.Monster.ListMonster:output_type -> api.service.v1.ListMonsterReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_battle_v1_monster_proto_init() }
func file_battle_v1_monster_proto_init() {
	if File_battle_v1_monster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_battle_v1_monster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMonsterRequest); i {
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
		file_battle_v1_monster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMonsterReply); i {
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
		file_battle_v1_monster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMonsterRequest); i {
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
		file_battle_v1_monster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMonsterReply); i {
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
		file_battle_v1_monster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMonsterRequest); i {
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
		file_battle_v1_monster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMonsterReply); i {
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
		file_battle_v1_monster_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMonsterRequest); i {
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
		file_battle_v1_monster_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMonsterReply); i {
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
		file_battle_v1_monster_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMonsterRequest); i {
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
		file_battle_v1_monster_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMonsterReply); i {
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
		file_battle_v1_monster_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMonsterReply_Data); i {
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
			RawDescriptor: file_battle_v1_monster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_battle_v1_monster_proto_goTypes,
		DependencyIndexes: file_battle_v1_monster_proto_depIdxs,
		MessageInfos:      file_battle_v1_monster_proto_msgTypes,
	}.Build()
	File_battle_v1_monster_proto = out.File
	file_battle_v1_monster_proto_rawDesc = nil
	file_battle_v1_monster_proto_goTypes = nil
	file_battle_v1_monster_proto_depIdxs = nil
}