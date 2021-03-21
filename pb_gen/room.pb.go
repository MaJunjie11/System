// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: room.proto

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

type RoomCapcity int32

const (
	RoomCapcity_Small  RoomCapcity = 0
	RoomCapcity_Middle RoomCapcity = 1
	RoomCapcity_Big    RoomCapcity = 2
)

// Enum value maps for RoomCapcity.
var (
	RoomCapcity_name = map[int32]string{
		0: "Small",
		1: "Middle",
		2: "Big",
	}
	RoomCapcity_value = map[string]int32{
		"Small":  0,
		"Middle": 1,
		"Big":    2,
	}
)

func (x RoomCapcity) Enum() *RoomCapcity {
	p := new(RoomCapcity)
	*p = x
	return p
}

func (x RoomCapcity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomCapcity) Descriptor() protoreflect.EnumDescriptor {
	return file_room_proto_enumTypes[0].Descriptor()
}

func (RoomCapcity) Type() protoreflect.EnumType {
	return &file_room_proto_enumTypes[0]
}

func (x RoomCapcity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomCapcity.Descriptor instead.
func (RoomCapcity) EnumDescriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{0}
}

// 创建申请教室接口
type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherId   int64       `protobuf:"varint,1,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	RoomCapcity RoomCapcity `protobuf:"varint,2,opt,name=room_capcity,json=roomCapcity,proto3,enum=pb_gen.RoomCapcity" json:"room_capcity,omitempty"`
	RoomName    string      `protobuf:"bytes,3,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoomRequest) GetTeacherId() int64 {
	if x != nil {
		return x.TeacherId
	}
	return 0
}

func (x *CreateRoomRequest) GetRoomCapcity() RoomCapcity {
	if x != nil {
		return x.RoomCapcity
	}
	return RoomCapcity_Small
}

func (x *CreateRoomRequest) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

type CreateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo   int32  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips string `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomResponse) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *CreateRoomResponse) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

// 管理员拉取待审核创建教室列表
type ManagerGetReviewingRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerId int64 `protobuf:"varint,1,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
}

func (x *ManagerGetReviewingRoomRequest) Reset() {
	*x = ManagerGetReviewingRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerGetReviewingRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerGetReviewingRoomRequest) ProtoMessage() {}

func (x *ManagerGetReviewingRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerGetReviewingRoomRequest.ProtoReflect.Descriptor instead.
func (*ManagerGetReviewingRoomRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{2}
}

func (x *ManagerGetReviewingRoomRequest) GetManagerId() int64 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

type ManagerGetReviewingRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserType     int64       `protobuf:"varint,1,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	RoomInfoList []*RoomInfo `protobuf:"bytes,2,rep,name=room_info_list,json=roomInfoList,proto3" json:"room_info_list,omitempty"`
	ErrNo        int32       `protobuf:"varint,3,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips      string      `protobuf:"bytes,4,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
}

func (x *ManagerGetReviewingRoomResponse) Reset() {
	*x = ManagerGetReviewingRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerGetReviewingRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerGetReviewingRoomResponse) ProtoMessage() {}

func (x *ManagerGetReviewingRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerGetReviewingRoomResponse.ProtoReflect.Descriptor instead.
func (*ManagerGetReviewingRoomResponse) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{3}
}

func (x *ManagerGetReviewingRoomResponse) GetUserType() int64 {
	if x != nil {
		return x.UserType
	}
	return 0
}

func (x *ManagerGetReviewingRoomResponse) GetRoomInfoList() []*RoomInfo {
	if x != nil {
		return x.RoomInfoList
	}
	return nil
}

func (x *ManagerGetReviewingRoomResponse) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *ManagerGetReviewingRoomResponse) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

// 管理员审核教室创建接口
type ManagerCheckReviewingRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId    int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	TeacherId int64 `protobuf:"varint,2,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"` //申请的老师
}

func (x *ManagerCheckReviewingRoomRequest) Reset() {
	*x = ManagerCheckReviewingRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerCheckReviewingRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerCheckReviewingRoomRequest) ProtoMessage() {}

func (x *ManagerCheckReviewingRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerCheckReviewingRoomRequest.ProtoReflect.Descriptor instead.
func (*ManagerCheckReviewingRoomRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{4}
}

func (x *ManagerCheckReviewingRoomRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *ManagerCheckReviewingRoomRequest) GetTeacherId() int64 {
	if x != nil {
		return x.TeacherId
	}
	return 0
}

type ManagerCheckReviewingRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo   int32  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips string `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
}

func (x *ManagerCheckReviewingRoomResponse) Reset() {
	*x = ManagerCheckReviewingRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerCheckReviewingRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerCheckReviewingRoomResponse) ProtoMessage() {}

func (x *ManagerCheckReviewingRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerCheckReviewingRoomResponse.ProtoReflect.Descriptor instead.
func (*ManagerCheckReviewingRoomResponse) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{5}
}

func (x *ManagerCheckReviewingRoomResponse) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *ManagerCheckReviewingRoomResponse) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

// 学生自主申请加入教室接口
type AddStudentToRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId int64 `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	RoomId    int64 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *AddStudentToRoomRequest) Reset() {
	*x = AddStudentToRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStudentToRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStudentToRoomRequest) ProtoMessage() {}

func (x *AddStudentToRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStudentToRoomRequest.ProtoReflect.Descriptor instead.
func (*AddStudentToRoomRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{6}
}

func (x *AddStudentToRoomRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *AddStudentToRoomRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type AddStudentToRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo   int32  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips string `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
}

func (x *AddStudentToRoomResponse) Reset() {
	*x = AddStudentToRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStudentToRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStudentToRoomResponse) ProtoMessage() {}

func (x *AddStudentToRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStudentToRoomResponse.ProtoReflect.Descriptor instead.
func (*AddStudentToRoomResponse) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{7}
}

func (x *AddStudentToRoomResponse) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *AddStudentToRoomResponse) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

// 老师添加学生加入教室接口
type TeacherAddStudentToRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId    int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	StudentId int64 `protobuf:"varint,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
}

func (x *TeacherAddStudentToRoomRequest) Reset() {
	*x = TeacherAddStudentToRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherAddStudentToRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherAddStudentToRoomRequest) ProtoMessage() {}

func (x *TeacherAddStudentToRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherAddStudentToRoomRequest.ProtoReflect.Descriptor instead.
func (*TeacherAddStudentToRoomRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{8}
}

func (x *TeacherAddStudentToRoomRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *TeacherAddStudentToRoomRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

type TeacherAddStudentToRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo   int32  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips string `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
}

func (x *TeacherAddStudentToRoomResponse) Reset() {
	*x = TeacherAddStudentToRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherAddStudentToRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherAddStudentToRoomResponse) ProtoMessage() {}

func (x *TeacherAddStudentToRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherAddStudentToRoomResponse.ProtoReflect.Descriptor instead.
func (*TeacherAddStudentToRoomResponse) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{9}
}

func (x *TeacherAddStudentToRoomResponse) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *TeacherAddStudentToRoomResponse) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

// 老师审核学生加入结果
type CheckStudentToRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId      int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	CheckStatus int64 `protobuf:"varint,2,opt,name=check_status,json=checkStatus,proto3" json:"check_status,omitempty"` // 0 通过  1不通过
}

func (x *CheckStudentToRoomRequest) Reset() {
	*x = CheckStudentToRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckStudentToRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStudentToRoomRequest) ProtoMessage() {}

func (x *CheckStudentToRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStudentToRoomRequest.ProtoReflect.Descriptor instead.
func (*CheckStudentToRoomRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{10}
}

func (x *CheckStudentToRoomRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *CheckStudentToRoomRequest) GetCheckStatus() int64 {
	if x != nil {
		return x.CheckStatus
	}
	return 0
}

type CheckStudentToRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo   int32  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips string `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
}

func (x *CheckStudentToRoomResponse) Reset() {
	*x = CheckStudentToRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckStudentToRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStudentToRoomResponse) ProtoMessage() {}

func (x *CheckStudentToRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStudentToRoomResponse.ProtoReflect.Descriptor instead.
func (*CheckStudentToRoomResponse) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{11}
}

func (x *CheckStudentToRoomResponse) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *CheckStudentToRoomResponse) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

var File_room_proto protoreflect.FileDescriptor

var file_room_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x62,
	0x5f, 0x67, 0x65, 0x6e, 0x1a, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0c, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x63, 0x61, 0x70, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x61,
	0x70, 0x63, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x61, 0x70, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x46, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x72, 0x72, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x22, 0x3f, 0x0a, 0x1e, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x1f, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x0e, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f,
	0x74, 0x69, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x54,
	0x69, 0x70, 0x73, 0x22, 0x5a, 0x0a, 0x20, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x55, 0x0a, 0x21, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x72, 0x72, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x22, 0x51, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x18, 0x41, 0x64, 0x64,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x72, 0x72, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x22, 0x58, 0x0a, 0x1e, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x53, 0x0a, 0x1f, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x41, 0x64, 0x64, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x72, 0x72, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x22, 0x57, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x4e, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x54,
	0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x74, 0x69, 0x70, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x2a,
	0x2d, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x61, 0x70, 0x63, 0x69, 0x74, 0x79, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x69, 0x67, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_proto_rawDescOnce sync.Once
	file_room_proto_rawDescData = file_room_proto_rawDesc
)

func file_room_proto_rawDescGZIP() []byte {
	file_room_proto_rawDescOnce.Do(func() {
		file_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_proto_rawDescData)
	})
	return file_room_proto_rawDescData
}

var file_room_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_room_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_room_proto_goTypes = []interface{}{
	(RoomCapcity)(0),                          // 0: pb_gen.RoomCapcity
	(*CreateRoomRequest)(nil),                 // 1: pb_gen.CreateRoomRequest
	(*CreateRoomResponse)(nil),                // 2: pb_gen.CreateRoomResponse
	(*ManagerGetReviewingRoomRequest)(nil),    // 3: pb_gen.ManagerGetReviewingRoomRequest
	(*ManagerGetReviewingRoomResponse)(nil),   // 4: pb_gen.ManagerGetReviewingRoomResponse
	(*ManagerCheckReviewingRoomRequest)(nil),  // 5: pb_gen.ManagerCheckReviewingRoomRequest
	(*ManagerCheckReviewingRoomResponse)(nil), // 6: pb_gen.ManagerCheckReviewingRoomResponse
	(*AddStudentToRoomRequest)(nil),           // 7: pb_gen.AddStudentToRoomRequest
	(*AddStudentToRoomResponse)(nil),          // 8: pb_gen.AddStudentToRoomResponse
	(*TeacherAddStudentToRoomRequest)(nil),    // 9: pb_gen.TeacherAddStudentToRoomRequest
	(*TeacherAddStudentToRoomResponse)(nil),   // 10: pb_gen.TeacherAddStudentToRoomResponse
	(*CheckStudentToRoomRequest)(nil),         // 11: pb_gen.CheckStudentToRoomRequest
	(*CheckStudentToRoomResponse)(nil),        // 12: pb_gen.CheckStudentToRoomResponse
	(*RoomInfo)(nil),                          // 13: pb_gen.RoomInfo
}
var file_room_proto_depIdxs = []int32{
	0,  // 0: pb_gen.CreateRoomRequest.room_capcity:type_name -> pb_gen.RoomCapcity
	13, // 1: pb_gen.ManagerGetReviewingRoomResponse.room_info_list:type_name -> pb_gen.RoomInfo
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_room_proto_init() }
func file_room_proto_init() {
	if File_room_proto != nil {
		return
	}
	file_room_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRequest); i {
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
		file_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomResponse); i {
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
		file_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerGetReviewingRoomRequest); i {
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
		file_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerGetReviewingRoomResponse); i {
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
		file_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerCheckReviewingRoomRequest); i {
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
		file_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerCheckReviewingRoomResponse); i {
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
		file_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStudentToRoomRequest); i {
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
		file_room_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStudentToRoomResponse); i {
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
		file_room_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherAddStudentToRoomRequest); i {
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
		file_room_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherAddStudentToRoomResponse); i {
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
		file_room_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckStudentToRoomRequest); i {
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
		file_room_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckStudentToRoomResponse); i {
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
			RawDescriptor: file_room_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_room_proto_goTypes,
		DependencyIndexes: file_room_proto_depIdxs,
		EnumInfos:         file_room_proto_enumTypes,
		MessageInfos:      file_room_proto_msgTypes,
	}.Build()
	File_room_proto = out.File
	file_room_proto_rawDesc = nil
	file_room_proto_goTypes = nil
	file_room_proto_depIdxs = nil
}
