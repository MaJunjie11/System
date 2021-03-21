package room_method

import (
	"System/common/common_dal"
	"System/pb_gen"
	"System/room/room_dal"
)

type TeacherCheckStudentToRoomHandler struct {
	Req  *pb_gen.TeacherCheckStudentToRoomRequest
	Resp *pb_gen.TeacherCheckStudentToRoomResponse
}

func NewTeacherCheckStudentToRoomHandler(req *pb_gen.TeacherCheckStudentToRoomRequest, resp *pb_gen.TeacherCheckStudentToRoomResponse) *TeacherCheckStudentToRoomHandler {
	return &TeacherCheckStudentToRoomHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherCheckStudentToRoomHandler) Run() {
	if t.Req.CheckStatus == 1 {
		return
	}
	hasRoom, roomInfo, err := room_dal.GetRoomInfoByRoomId(t.Req.RoomId)
	if err != nil || !hasRoom {
		return
	}
	//做个鉴权
	if roomInfo.TeacherId != t.Req.TeacherId {
		return
	}
	common_dal.TeacherCheckStudentToRoom(t.Req.RoomId, t.Req.StudentId)
}
