package user_dal

import (
	"System/db"
	"System/user/user_moudle"
	"fmt"
)

func GetUserBasicInfoFromDb(uid int64) (*user_moudle.UserStudentBaseInfo, error) {
	var (
		userBaseInfo *user_moudle.UserStudentBaseInfo = &user_moudle.UserStudentBaseInfo{}
		err          error
		count        int
	)
	//
	db.Db.Where("uid = ?", uid).Find(userBaseInfo).Count(&count)
	if count < 1 {
		return nil, fmt.Errorf("no baseInfo")
	}
	return userBaseInfo, err

}

func ChangeUserPasswordToDbByUid(uid int64, password string) error {
	db.Db.Model(&user_moudle.UserStudentLoginInfo{}).Where("uid = ?", uid).Update("password", password)
	return nil
}
