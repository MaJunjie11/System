package room_method

import (
	"System/db"
	"System/pb_gen"
	"System/room/room_model"
	"System/util"
	"fmt"
)

type ManagerGetUseRoomHandler struct {
	Req  *pb_gen.ManagerGetUseRoomRequest
	Resp *pb_gen.ManagerGetUseRoomResponse
}

func NewManagerGetUseRoomHandler(req *pb_gen.ManagerGetUseRoomRequest, resp *pb_gen.ManagerGetUseRoomResponse) *ManagerGetUseRoomHandler {
	return &ManagerGetUseRoomHandler{
		Req:  req,
		Resp: resp,
	}
}

func (m *ManagerGetUseRoomHandler) Run() {
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
	m.getUseRoomFromDb()
	m.Resp.Code = pb_gen.ErrNo_Success
	m.Resp.Msg = "查询成功"
}

func (m *ManagerGetUseRoomHandler) getUseRoomFromDb() {

	// 获取可以使用的room

	var (
		result  []room_model.RoomBaseInfo = make([]room_model.RoomBaseInfo, 0)
		count   int
		data    *pb_gen.ManagerGetUseRoomData       = &pb_gen.ManagerGetUseRoomData{}
		options []*pb_gen.ManagerGetUseRoomDataItem = make([]*pb_gen.ManagerGetUseRoomDataItem, 0)
	)
	db.Db.Table("room_base_info").Where("status = 0").Scan(&result).Count(&count)

	for _, val := range result {
		item := &pb_gen.ManagerGetUseRoomDataItem{
			Label: fmt.Sprintf("%s(%s)", val.RoomAddr, val.RoomCapacity),
			Value: fmt.Sprintf("%d", val.Id),
		}
		options = append(options, item)
	}

	data.Options = options
	m.Resp.Data = data
}
