package room_method

import (
	"System/pb_gen"
	"System/room/room_logic"
)

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
	var (
		err               error
		createRoomManager *room_logic.CreateRoomManager
	)
	// 参数检测
	if err = c.checkParam(); err != nil {
		c.Resp.ErrNo = 1
		c.Resp.ErrTips = "参数错误"
		return
	}
	createRoomManager = room_logic.NewCreateRoomManager(c.Req.ManagerId, c.Req.GetRoomCapcity(), c.Req.GetRoomName())
	if err = createRoomManager.Process(); err != nil {
		return
	}
}

func (c *CreateRoomHandler) checkParam() error {
	return nil
}
