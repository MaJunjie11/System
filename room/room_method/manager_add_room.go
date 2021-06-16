package room_method

import (
	"System/db"
	"System/pb_gen"
	"System/room/room_model"
	"System/util"
	"fmt"
)

type ManagerAddRoomHandler struct {
	Req  *pb_gen.ManagerAddRoomRequest
	Resp *pb_gen.ManagerAddRoomResponse
}

func NewManagerAddRoomHandelr(req *pb_gen.ManagerAddRoomRequest, resp *pb_gen.ManagerAddRoomResponse) *ManagerAddRoomHandler {
	return &ManagerAddRoomHandler{
		Req:  req,
		Resp: resp,
	}
}

func (m *ManagerAddRoomHandler) Run() {
	var (
		err       error
		userToken *util.UserToken
		uid       int64
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(m.Req.Token, util.StudentUserSecretKey); err != nil {
		m.Resp.Code = pb_gen.ErrNo_User_Token_Error
		m.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	fmt.Println("uid:", uid)
	m.addRoomToDb()
	m.Resp.Code = pb_gen.ErrNo_Success
	m.Resp.Msg = "添加成功"
}

func (m *ManagerAddRoomHandler) addRoomToDb() {
	roomInfo := &room_model.RoomBaseInfo{
		Id:           0,
		RoomAddr:     m.Req.RoomAddr,
		RoomCapacity: m.Req.RoomCapacity,
	}
	db.Db.Create(roomInfo)
}
