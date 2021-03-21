package room_method

import (
	"System/common/common_dal"
	"System/pb_gen"
	"System/room/room_dal"
)

type AddStudentToRoomHandler struct {
	Req  *pb_gen.AddStudentToRoomRequest
	Resp *pb_gen.AddStudentToRoomResponse
}

func NewAddStudentToRoomHandler(req *pb_gen.AddStudentToRoomRequest, resp *pb_gen.AddStudentToRoomResponse) *AddStudentToRoomHandler {
	return &AddStudentToRoomHandler{
		Req:  req,
		Resp: resp,
	}
}

func (a *AddStudentToRoomHandler) Run() {
	hasRoom, roomInfo, err := room_dal.GetRoomInfoByRoomId(a.Req.GetRoomId())
	if !hasRoom || err != nil {
		return
	}
	// 教室不在使用中 不能添加
	if roomInfo.RoomStatus != pb_gen.RoomStatus_Usering {
		return
	}
	common_dal.StudentApplyAddToRoom(a.Req.GetRoomId(), a.Req.GetStudentId(), roomInfo.TeacherId)
}
