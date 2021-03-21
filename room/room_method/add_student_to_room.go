package room_method

import (
	"System/common/common_dal"
	"System/pb_gen"
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
	common_dal.AddStudentToRoom(a.Req.GetRoomId(), a.Req.GetStudentId())
}
