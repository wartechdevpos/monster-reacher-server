// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: data_schema.proto

package data_schema

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WartechUserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"`
	User           string                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" bson:"user"`
	Email          string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" bson:"email"`
	Password       string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty" bson:"password"`
	EmailConfirmed bool                   `protobuf:"varint,5,opt,name=email_confirmed,json=emailConfirmed,proto3" json:"email_confirmed,omitempty" bson:"email_confirmed"`
	Birthday       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty" bson:"birthday"`
}

func (x *WartechUserData) Reset() {
	*x = WartechUserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WartechUserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WartechUserData) ProtoMessage() {}

func (x *WartechUserData) ProtoReflect() protoreflect.Message {
	mi := &file_data_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WartechUserData.ProtoReflect.Descriptor instead.
func (*WartechUserData) Descriptor() ([]byte, []int) {
	return file_data_schema_proto_rawDescGZIP(), []int{0}
}

func (x *WartechUserData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WartechUserData) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *WartechUserData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *WartechUserData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *WartechUserData) GetEmailConfirmed() bool {
	if x != nil {
		return x.EmailConfirmed
	}
	return false
}

func (x *WartechUserData) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

type ProfileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"`
	Name     string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Services map[string]string `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"services"`
}

func (x *ProfileData) Reset() {
	*x = ProfileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileData) ProtoMessage() {}

func (x *ProfileData) ProtoReflect() protoreflect.Message {
	mi := &file_data_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileData.ProtoReflect.Descriptor instead.
func (*ProfileData) Descriptor() ([]byte, []int) {
	return file_data_schema_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProfileData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfileData) GetServices() map[string]string {
	if x != nil {
		return x.Services
	}
	return nil
}

type AuthenticationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	CreateTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_timestamp,json=createTimestamp,proto3" json:"create_timestamp,omitempty" bson:"create_timestamp"`
	AccessToken     string                 `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty" bson:"access_token"`
	ExtendCount     int64                  `protobuf:"varint,4,opt,name=extend_count,json=extendCount,proto3" json:"extend_count,omitempty" bson:"extend_count"`
}

func (x *AuthenticationData) Reset() {
	*x = AuthenticationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationData) ProtoMessage() {}

func (x *AuthenticationData) ProtoReflect() protoreflect.Message {
	mi := &file_data_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationData.ProtoReflect.Descriptor instead.
func (*AuthenticationData) Descriptor() ([]byte, []int) {
	return file_data_schema_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticationData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthenticationData) GetCreateTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTimestamp
	}
	return nil
}

func (x *AuthenticationData) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthenticationData) GetExtendCount() int64 {
	if x != nil {
		return x.ExtendCount
	}
	return 0
}

var File_data_schema_proto protoreflect.FileDescriptor

var file_data_schema_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x0f, 0x57, 0x61, 0x72, 0x74, 0x65,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0x9a, 0x84, 0x9e, 0x03, 0x14, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x9a, 0x84, 0x9e, 0x03,
	0x0c, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0x9a, 0x84, 0x9e, 0x03, 0x0f, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x44, 0x0a, 0x0f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x16, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x22, 0x52, 0x0e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x4c, 0x0a, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0x9a, 0x84, 0x9e, 0x03,
	0x0f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22,
	0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0xf5, 0x01, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0x9a, 0x84, 0x9e, 0x03, 0x14, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x58, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x14, 0x9a, 0x84, 0x9e, 0x03, 0x0f, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x93, 0x02, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x63, 0x0a, 0x10, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x1c, 0x9a, 0x84, 0x9e, 0x03, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x52, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3b,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0x9a, 0x84, 0x9e, 0x03, 0x13, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x0c, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x18, 0x9a, 0x84, 0x9e, 0x03, 0x13, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x0b, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x4f, 0x5a, 0x4d, 0x77, 0x61, 0x72, 0x74,
	0x65, 0x63, 0x68, 0x2d, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3b, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_data_schema_proto_rawDescOnce sync.Once
	file_data_schema_proto_rawDescData = file_data_schema_proto_rawDesc
)

func file_data_schema_proto_rawDescGZIP() []byte {
	file_data_schema_proto_rawDescOnce.Do(func() {
		file_data_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_schema_proto_rawDescData)
	})
	return file_data_schema_proto_rawDescData
}

var file_data_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_data_schema_proto_goTypes = []interface{}{
	(*WartechUserData)(nil),       // 0: data_schema.WartechUserData
	(*ProfileData)(nil),           // 1: data_schema.ProfileData
	(*AuthenticationData)(nil),    // 2: data_schema.AuthenticationData
	nil,                           // 3: data_schema.ProfileData.ServicesEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_data_schema_proto_depIdxs = []int32{
	4, // 0: data_schema.WartechUserData.birthday:type_name -> google.protobuf.Timestamp
	3, // 1: data_schema.ProfileData.services:type_name -> data_schema.ProfileData.ServicesEntry
	4, // 2: data_schema.AuthenticationData.create_timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_data_schema_proto_init() }
func file_data_schema_proto_init() {
	if File_data_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WartechUserData); i {
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
		file_data_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileData); i {
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
		file_data_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationData); i {
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
			RawDescriptor: file_data_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_schema_proto_goTypes,
		DependencyIndexes: file_data_schema_proto_depIdxs,
		MessageInfos:      file_data_schema_proto_msgTypes,
	}.Build()
	File_data_schema_proto = out.File
	file_data_schema_proto_rawDesc = nil
	file_data_schema_proto_goTypes = nil
	file_data_schema_proto_depIdxs = nil
}
