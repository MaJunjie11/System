package room_method

import (
	"System/common/common_dal"
	"System/pb_gen"
)

type TeacherGetReqAddRoomStudentListHandler struct {
	Req  *pb_gen.TeacherGetReqAddRoomStudentListRequest
	Resp *pb_gen.TeacherGetReqAddRoomStudentListResponse
}

func NewTeacherGetReqAddRoomStudentList(req *pb_gen.TeacherGetReqAddRoomStudentListRequest, resp *pb_gen.TeacherGetReqAddRoomStudentListResponse) *TeacherGetReqAddRoomStudentListHandler {
	return &TeacherGetReqAddRoomStudentListHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherGetReqAddRoomStudentListHandler) Run() {

	//TODO: 后续在做鉴权等优化
	common_dal.TeacherGetReqAddRoomStudentList(t.Req.TeacherId, t.Req.RoomId)
}
