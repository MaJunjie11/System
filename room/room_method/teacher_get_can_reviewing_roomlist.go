package room_method

import (
	"System/pb_gen"
	"System/room/room_dal"
)

type TeacherGetCanReviewingRoomListHandler struct {
	Req  *pb_gen.TeacherGetCanReviewingRoomListRequest
	Resp *pb_gen.TeacherGetCanReviewingRoomListResponse
}

func NewTeacherGetCanReviewingRoomListHandler(req *pb_gen.TeacherGetCanReviewingRoomListRequest, resp *pb_gen.TeacherGetCanReviewingRoomListResponse) *TeacherGetCanReviewingRoomListHandler {
	return &TeacherGetCanReviewingRoomListHandler{
		Req:  req,
		Resp: resp,
	}
}

func (c *TeacherGetCanReviewingRoomListHandler) Run() {
	roomInfoList, err := room_dal.GetCreatingRoomFromDB()
	if err != nil {
	}
	c.Resp.RoomInfoList = roomInfoList
}
