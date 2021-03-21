package room_method

import (
	"System/pb_gen"
	"System/room/room_dal"
)

type ManagerCheckReviewingRoomHandler struct {
	Req  *pb_gen.ManagerCheckReviewingRoomRequest
	Resp *pb_gen.ManagerCheckReviewingRoomResponse
}

func NewManagerCheckjReviewindRoomHandler(req *pb_gen.ManagerCheckReviewingRoomRequest, resp *pb_gen.ManagerCheckReviewingRoomResponse) *ManagerCheckReviewingRoomHandler {
	return &ManagerCheckReviewingRoomHandler{
		Req:  req,
		Resp: resp,
	}
}

func (m *ManagerCheckReviewingRoomHandler) Run() {
	// 获取roomInfo
	hasRoom, roomInfo, err := room_dal.GetRoomInfoByRoomId(m.Req.RoomId)
	if err != nil {
		return
	}
	if !hasRoom {
		return
	}
	if roomInfo.RoomStatus != pb_gen.RoomStatus_Reviewing {
		return
	}
	roomInfo.RoomStatus = pb_gen.RoomStatus_Usering
	err = room_dal.ModifyRoomInfo(roomInfo)
	if err != nil {

	}
	room_dal.ModifyRoomStatus(pb_gen.RoomStatus_Usering, m.Req.RoomId)
}
