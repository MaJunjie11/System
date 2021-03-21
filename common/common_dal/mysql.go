package common_dal

import (
	"System/common"
	"System/pb_gen"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/skoo87/log4go"
)

func TeacherGetReqAddRoomStudentList(teacherId, roomId int64) ([]*pb_gen.UserInfo, error) {
	var (
		err          error
		userInfoList []*pb_gen.UserInfo = make([]*pb_gen.UserInfo, 0)
		userInfo     *pb_gen.UserInfo   = &pb_gen.UserInfo{}
	)
	common.InitDB()
	defer common.ClosDB()
	queryStr := fmt.Sprintf("select a.userInfo from user as a left join user_room as b on a.uid=b.userId where b.teacherId=%d and b.status=2 and b.roomId = %d", teacherId, roomId)
	rows, err := common.DB.Query(queryStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var userInfoStr string
		rows.Scan(&userInfoStr)
		jsoniter.UnmarshalFromString(userInfoStr, userInfo)
		userInfoList = append(userInfoList, userInfo)
	}
	return userInfoList, nil
}

func TeacherCheckStudentToRoom(roomId, userId int64) error {
	common.InitDB()
	defer common.ClosDB()
	// id userInfoJson uid phone  后续优化索引
	queryStr := fmt.Sprintf("update  user_room set status = 1 where roomId = %d and userId = %d)", roomId, userId)
	common.DB.Query(queryStr)
	return nil
}

func TeacherAddStudentToRoom(roomId, userId, teacherId int64) error {
	common.InitDB()
	defer common.ClosDB()
	// 通过状态
	queryStr := fmt.Sprintf("insert into user_room values('%d', '%d',1, '%d')", roomId, userId, teacherId)
	common.DB.Query(queryStr)
	return nil
}

func StudentApplyAddToRoom(roomId, userId, teacherId int64) error {
	common.InitDB()
	defer common.ClosDB()
	// id userInfoJson uid phone  后续优化索引
	queryStr := fmt.Sprintf("insert into user_room values('%d', '%d',2,'%d')", roomId, userId, teacherId)
	common.DB.Query(queryStr)
	return nil
}

// 通过phone拿到用户的Info 后续在userInfo中扩充字段
func GetUserInfoByPhone(phone string) (bool, *pb_gen.UserInfo, error) {
	var (
		err         error
		userInfoStr string
		userInfo    *pb_gen.UserInfo = &pb_gen.UserInfo{}
	)
	common.InitDB()
	defer common.ClosDB()

	queryStr := fmt.Sprintf("select userInfo from user where phone =%s", phone)
	row := common.DB.QueryRow("select userInfo from user where phone = ?", phone)
	if row == nil {
		log4go.Error("queryDB faild row is nil:queryStr:%s", queryStr)
		return false, nil, nil
	}
	if row.Err() != nil {
		log4go.Error("queryDB faild:queryStr:%s err:%v", queryStr, row.Err())
		return false, nil, err
	}
	if err = row.Scan(&userInfoStr); err != nil {
		log4go.Error("queryDB faild:queryStr:%s err:%v", queryStr, err)
		return false, nil, nil
	}
	// 没有当前用户
	if userInfoStr == "" {
		return false, nil, nil
	}
	if err = jsoniter.UnmarshalFromString(userInfoStr, userInfo); err != nil {
		log4go.Error("do not have this user or unmarshal faild:%s err:%v", queryStr, err)
		return false, nil, err
	}
	return true, userInfo, nil
}

func GetUsrInfoByUid(uid int64) (*pb_gen.UserInfo, error) {
	var (
		err         error
		userInfoStr string
		userInfo    *pb_gen.UserInfo = &pb_gen.UserInfo{}
	)
	common.InitDB()
	defer common.ClosDB()

	queryStr := fmt.Sprintf("select phone from user where uid=%d", uid)

	if err = common.DB.QueryRow(queryStr).Scan(userInfoStr); err != nil {
		log4go.Error("queryDB faild:queryStr:%s err:%v", queryStr, err)
		return nil, err
	}
	if err = jsoniter.UnmarshalFromString(userInfoStr, userInfo); err != nil {
		log4go.Error("queryDB faild:queryStr:%s err:%v", queryStr, err)
		return nil, err
	}
	return userInfo, nil

}
