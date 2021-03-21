// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: user.proto

package pb_gen

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AddUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName     string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Passworld    string   `protobuf:"bytes,2,opt,name=passworld,proto3" json:"passworld,omitempty"`
	TelephoneNum string   `protobuf:"bytes,3,opt,name=telephone_num,json=telephoneNum,proto3" json:"telephone_num,omitempty"`
	Sex          Sex      `protobuf:"varint,4,opt,name=sex,proto3,enum=pb_gen.Sex" json:"sex,omitempty"`
	Age          int32    `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	UserType     UserType `protobuf:"varint,6,opt,name=user_type,json=userType,proto3,enum=pb_gen.UserType" json:"user_type,omitempty"`
}

func (x *AddUserRequest) Reset() {
	*x = AddUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRequest) ProtoMessage() {}

func (x *AddUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRequest.ProtoReflect.Descriptor instead.
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AddUserRequest) GetPassworld() string {
	if x != nil {
		return x.Passworld
	}
	return ""
}

func (x *AddUserRequest) GetTelephoneNum() string {
	if x != nil {
		return x.TelephoneNum
	}
	return ""
}

func (x *AddUserRequest) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_SexUnknow
}

func (x *AddUserRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *AddUserRequest) GetUserType() UserType {
	if x != nil {
		return x.UserType
	}
	return UserType_UserTypeUnknow
}

type AddUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo    int32    `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips  string   `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
	Uid      int64    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	UserType UserType `protobuf:"varint,4,opt,name=user_type,json=userType,proto3,enum=pb_gen.UserType" json:"user_type,omitempty"`
}

func (x *AddUserResponse) Reset() {
	*x = AddUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserResponse) ProtoMessage() {}

func (x *AddUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserResponse.ProtoReflect.Descriptor instead.
func (*AddUserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserResponse) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *AddUserResponse) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

func (x *AddUserResponse) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AddUserResponse) GetUserType() UserType {
	if x != nil {
		return x.UserType
	}
	return UserType_UserTypeUnknow
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x62,
	0x5f, 0x67, 0x65, 0x6e, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x53,
	0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70,
	0x62, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x2d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x32,
	0x4c, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x5f,
	0x67, 0x65, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_proto_goTypes = []interface{}{
	(*AddUserRequest)(nil),  // 0: pb_gen.AddUserRequest
	(*AddUserResponse)(nil), // 1: pb_gen.AddUserResponse
	(Sex)(0),                // 2: pb_gen.Sex
	(UserType)(0),           // 3: pb_gen.UserType
}
var file_user_proto_depIdxs = []int32{
	2, // 0: pb_gen.AddUserRequest.sex:type_name -> pb_gen.Sex
	3, // 1: pb_gen.AddUserRequest.user_type:type_name -> pb_gen.UserType
	3, // 2: pb_gen.AddUserResponse.user_type:type_name -> pb_gen.UserType
	0, // 3: pb_gen.LoginService.AddUser:input_type -> pb_gen.AddUserRequest
	1, // 4: pb_gen.LoginService.AddUser:output_type -> pb_gen.AddUserResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_user_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRequest); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserResponse); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
