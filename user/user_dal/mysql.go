package user_dal

import (
	"System/common"
	"System/user/user_moudle"
	"fmt"

	jsoniter "github.com/json-iterator/go"

	"github.com/skoo87/log4go"
)

func AddUserInfoToDB(userInfo *user_moudle.UserInfo) error {
	var userInfoStr string
	userInfoStr, _ = jsoniter.MarshalToString(userInfo)
	common.InitDB()
	defer common.ClosDB()
	fmt.Print(userInfoStr)
	// id userInfoJson uid phone  后续优化索引
	queryStr := fmt.Sprintf("insert into user values(1,'%s','%d', '%s')", userInfoStr, userInfo.Uid, userInfo.Phone)
	common.DB.Query(queryStr)
	return nil
}

func CheckIfHasUserFromDB(telephoneNum string) (bool, error) {
	var (
		err     error
		hasUser bool = false
		phone   string
	)
	common.InitDB()
	defer common.ClosDB()

	queryStr := fmt.Sprintf("select phone from user where phone=%s", telephoneNum)

	if err = common.DB.QueryRow(queryStr).Scan(&phone); err != nil {
		log4go.Error("queryDB faild:queryStr:%s err:%v", queryStr, err)
		return false, err
	}
	if phone != "" {
		hasUser = true
	}
	return hasUser, nil
}
