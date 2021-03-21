package room_method

import (
	"System/pb_gen"
	"System/room/room_dal"
)

type ManagerGetReviewingRoomHandler struct {
	Req  *pb_gen.ManagerGetReviewingRoomRequest
	Resp *pb_gen.ManagerGetReviewingRoomResponse
}

func NewManagerGetReviewindRoomHandler(req *pb_gen.ManagerGetReviewingRoomRequest, resp *pb_gen.ManagerGetReviewingRoomResponse) *ManagerGetReviewingRoomHandler {
	return &ManagerGetReviewingRoomHandler{
		Req:  req,
		Resp: resp,
	}
}

func (m *ManagerGetReviewingRoomHandler) Run() {
	roomInfoList, err := room_dal.MangerGetReviewingRoom()
	if err != nil {
	}
	m.Resp.RoomInfoList = roomInfoList

}
