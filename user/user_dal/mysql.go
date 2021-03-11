package user_dal

import (
	"System/common"
	"System/user/user_moudle"
	"fmt"

	"github.com/skoo87/log4go"
)

func AddUserInfoToDB(userInfo *user_moudle.UserInfo) error {
	// 先做简单一些 后期可以加入连接池等操作 目前使用链接断开的简单方式
	common.InitDB()
	defer common.ClosDB()
	// 在这里将用户数据传入DB
	// id name age sex phone passworld Uid
	queryStr := fmt.Sprintf("insert into user values(1,'%s', %d, %d, '%s', '%s', %d)", userInfo.Name, userInfo.Age, userInfo.Sex, userInfo.Phone, userInfo.Passworld, userInfo.Uid)
	fmt.Println(queryStr)
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
