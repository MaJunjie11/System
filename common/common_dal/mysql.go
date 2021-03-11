package common_dal

import (
	"System/common"
	"System/user/user_moudle"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/skoo87/log4go"
)

// 通过phone拿到用户的Info 后续在userInfo中扩充字段
func GetUserInfoByPhone(phone string) (bool, *user_moudle.UserInfo, error) {
	var (
		err         error
		userInfoStr string
		userInfo    *user_moudle.UserInfo = &user_moudle.UserInfo{}
	)
	common.InitDB()
	defer common.ClosDB()

	queryStr := fmt.Sprintf("select phone from user where phone=%s", phone)
	row := common.DB.QueryRow(queryStr)
	if row == nil {
		log4go.Error("queryDB faild row is nil:queryStr:%s", queryStr)
		return false, nil, nil
	}
	if row.Err() != nil {
		log4go.Error("queryDB faild:queryStr:%s err:%v", queryStr, row.Err())
		return false, nil, err
	}
	if err = row.Scan(userInfoStr); err != nil {
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

func GetUsrInfoByUid(uid int64) (*user_moudle.UserInfo, error) {
	var (
		err         error
		userInfoStr string
		userInfo    *user_moudle.UserInfo = &user_moudle.UserInfo{}
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
