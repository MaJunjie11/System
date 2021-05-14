package room_method

import "System/pb_gen"

type StudentGetAlreadyRoomHandler struct {
	Req  *pb_gen.StudentGetAlreadyRoomRequest
	Resp *pb_gen.StudentGetAlreadyRoomResponse
}

func NewStudentGetAlreadyRoomHandler(req *pb_gen.StudentGetAlreadyRoomRequest, resp *pb_gen.StudentGetAlreadyRoomResponse) *StudentGetAlreadyRoomHandler {
	return &StudentGetAlreadyRoomHandler{
		Req:  req,
		Resp: resp,
	}
}

func (a *StudentGetAlreadyRoomHandler) Run() {
	// 学生获得可以加入的教室列表
	// 查询数据库 将所有的教室全部拿出来，在排除学生已经加入的教室和已经开课的教室
}
