package room_logic

import (
	"System/pb_gen"
	"System/room/room_dal"
	"System/util"
)

type CreateRoomManager struct {
	ManagerId   int64
	RoomCapcity pb_gen.RoomCapcity
	RoomName    string
}

func NewCreateRoomManager(teacherId int64, roomCapcity pb_gen.RoomCapcity, roomName string) *CreateRoomManager {
	return &CreateRoomManager{
		ManagerId:   teacherId,
		RoomCapcity: roomCapcity,
		RoomName:    roomName,
	}
}

// 大中小房间对应数据
func (c *CreateRoomManager) getRoomCap() int64 {
	var roomCap int64
	switch c.RoomCapcity {
	case pb_gen.RoomCapcity_Big:
		roomCap = 150
	case pb_gen.RoomCapcity_Middle:
		roomCap = 100
	case pb_gen.RoomCapcity_Small:
		roomCap = 50
	default:
		roomCap = 50
	}
	return roomCap
}

func (c *CreateRoomManager) Process() error {
	var (
		err      error
		roomId   int64
		roomInfo *pb_gen.RoomInfo
		roomCap  int64
	)
	// 第一步 激活房间
	roomId = util.Gen()

	roomCap = c.getRoomCap()

	// pack roomInfo
	roomInfo = c.packRoomInfo(roomId, roomCap)

	// 存入数据库中
	if err = room_dal.CreateRoomToDB(roomInfo); err != nil {
		return err
	}
	return nil
}

func (c *CreateRoomManager) packRoomInfo(roomId, roomCap int64) *pb_gen.RoomInfo {
	return &pb_gen.RoomInfo{
		RoomId:         roomId,
		RoomReviewerId: c.ManagerId,
		RoomName:       c.RoomName,
		RoomStatus:     pb_gen.RoomStatus_Creating, //房间处于创建状态
		RoomCapcity:    roomCap,
	}
}
