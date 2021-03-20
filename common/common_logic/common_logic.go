package common_logic

import (
	"System/user/user_moudle"

	"github.com/skoo87/log4go"
)

// 给User鉴权的函数 返回user的权限 0: 没权限   1：学生 2：老师 3：管理员
func AuthenticationUserByUserId(userInfo *user_moudle.UserInfo) user_moudle.UserType {
	if userInfo == nil {
		log4go.Warn("user no authentication")
		return user_moudle.UserType_Unknow
	}
	return userInfo.UserType
}
