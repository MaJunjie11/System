package room_method

import "System/pb_gen"

type CreateRoomHandler struct {
	Req  *pb_gen.CreateRoomRequest
	Resp *pb_gen.CreateRoomResponse
}

func NewCreateRoomHandler(req *pb_gen.CreateRoomRequest, resp *pb_gen.CreateRoomResponse) *CreateRoomHandler {
	return &CreateRoomHandler{
		Req:  req,
		Resp: resp,
	}
}

func (c *CreateRoomHandler) Run() {
	// 在数据库中创建一个待审核房间
}
