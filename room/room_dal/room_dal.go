package room_dal

import (
	"System/common"
	"System/pb_gen"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/skoo87/log4go"
)

func GetCreatingRoomFromDB() ([]*pb_gen.RoomInfo, error) {
	var (
		err          error
		roomInfoList []*pb_gen.RoomInfo = make([]*pb_gen.RoomInfo, 0)
		roomInfo     *pb_gen.RoomInfo   = &pb_gen.RoomInfo{}
	)
	common.InitDB()
	defer common.ClosDB()
	queryStr := fmt.Sprintf("select roomInfo from room where roomStatus = %d", int32(pb_gen.RoomStatus_Creating))
	rows, err := common.DB.Query(queryStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var roomInfoStr string
		rows.Scan(&roomInfoStr)
		jsoniter.UnmarshalFromString(roomInfoStr, roomInfo)
		roomInfoList = append(roomInfoList, roomInfo)
	}
	return roomInfoList, nil
}

func CreateRoomToDB(roomInfo *pb_gen.RoomInfo) error {
	var roomInfoStr string
	roomInfoStr, _ = jsoniter.MarshalToString(roomInfo)
	common.InitDB()
	defer common.ClosDB()
	fmt.Print(roomInfoStr)
	queryStr := fmt.Sprintf("insert into room values(1,'%s','%d','%d', '%d','%d')", roomInfoStr, roomInfo.RoomId, int32(roomInfo.RoomStatus), roomInfo.RoomReviewerId, roomInfo.TeacherId)
	common.DB.Query(queryStr)
	return nil
}

// 目前看来创建以后只有status会改变
func ModifyRoomStatus(roomStatus pb_gen.RoomStatus, roomId int64) error {
	common.InitDB()
	defer common.ClosDB()
	queryStr := fmt.Sprintf("update room set roomStatus = %d where roomId = %d", int32(roomStatus), roomId)
	common.DB.Query(queryStr)
	return nil
}

func ModifyRoomTeacher(roomId, teacherId int64) error {
	common.InitDB()
	defer common.ClosDB()
	queryStr := fmt.Sprintf("update room set teacherId = %d where roomId = %d", teacherId, roomId)
	common.DB.Query(queryStr)
	return nil

}

// 管理员获取审核中教室列表
func MangerGetReviewingRoom() ([]*pb_gen.RoomInfo, error) {
	var (
		roomInfoList []*pb_gen.RoomInfo = make([]*pb_gen.RoomInfo, 0)
		roomInfo     *pb_gen.RoomInfo   = &pb_gen.RoomInfo{}
		err          error
	)
	common.InitDB()
	defer common.ClosDB()
	queryStr := fmt.Sprintf("select roomInfo from room where roomStatus = %d", int32(pb_gen.RoomStatus_Reviewing))
	rows, err := common.DB.Query(queryStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var roomInfoStr string
		rows.Scan(&roomInfoStr)
		jsoniter.UnmarshalFromString(roomInfoStr, roomInfo)
		roomInfoList = append(roomInfoList, roomInfo)
	}
	return roomInfoList, nil
}

// 老师获取审核中的教室
func TeacherGetReviewingRoom(teacherId int64) []*pb_gen.RoomInfo {
	return nil
}

// 老师获取可使用的教室
func TeacherGetUsingRoom(teacherId int64) []*pb_gen.RoomInfo {
	return nil
}

func GetRoomInfoByRoomId(roomId int64) (bool, *pb_gen.RoomInfo, error) {
	var (
		err         error
		roomInfoStr string
		roomInfo    *pb_gen.RoomInfo = &pb_gen.RoomInfo{}
	)
	common.InitDB()
	defer common.ClosDB()

	queryStr := fmt.Sprintf("select roomInfo from user where roomId =%d", roomId)
	row := common.DB.QueryRow("select roomInfo from user where roomId = ?", roomId)
	if row == nil {
		log4go.Error("queryDB faild row is nil:queryStr:%s", queryStr)
		return false, nil, nil
	}
	if row.Err() != nil {
		log4go.Error("queryDB faild:queryStr:%s err:%v", queryStr, row.Err())
		return false, nil, err
	}
	if err = row.Scan(&roomInfoStr); err != nil {
		log4go.Error("queryDB faild:queryStr:%s err:%v", queryStr, err)
		return false, nil, nil
	}
	// 没有当前用户
	if roomInfoStr == "" {
		return false, nil, nil
	}
	if err = jsoniter.UnmarshalFromString(roomInfoStr, roomInfo); err != nil {
		log4go.Error("do not have this user or unmarshal faild:%s err:%v", queryStr, err)
		return false, nil, err
	}
	return true, roomInfo, nil
}

// 修改数据库中表信息
func ModifyRoomInfo(roomInfo *pb_gen.RoomInfo) error {
	common.InitDB()
	defer common.ClosDB()
	roomInfoStr, err := jsoniter.MarshalToString(roomInfo)
	if err != nil {
		log4go.Error("err:%v", err)
		return err
	}
	queryStr := fmt.Sprintf("update room set roomInfo = %s", roomInfoStr)
	common.DB.Query(queryStr)
	return nil
}
