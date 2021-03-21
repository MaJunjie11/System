package room_method

import (
	"System/pb_gen"
)

type TeacherAddStudentToRoomHandler struct {
	Req  *pb_gen.TeacherAddStudentToRoomRequest
	Resp *pb_gen.TeacherAddStudentToRoomResponse
}

func NewTeacherAddStudentToRoomHandler(req *pb_gen.TeacherAddStudentToRoomRequest, resp *pb_gen.TeacherAddStudentToRoomResponse) *TeacherAddStudentToRoomHandler {
	return &TeacherAddStudentToRoomHandler{
		Req:  req,
		Resp: resp,
	}
}

func (a *TeacherAddStudentToRoomHandler) Run() {
}
