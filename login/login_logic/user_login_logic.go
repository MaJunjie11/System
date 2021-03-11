package login_logic

import (
	"System/common/common_dal"
	"System/user/user_moudle"

	"golang.org/x/crypto/bcrypt"
)

type UserLoginLogicMgr struct {
	phone       string
	passWord    string
	userInfo    *user_moudle.UserInfo
	LoginStatus int32
}

func NewUserLogicMgr(phone, passWorld string) *UserLoginLogicMgr {
	return &UserLoginLogicMgr{
		phone:    phone,
		passWord: passWorld,
	}
}

func (u *UserLoginLogicMgr) Process() error {

	// 然后继续考虑加密算法
	var (
		err          error
		loginSuccess bool
	)
	if err = u.loadUserInfo(); err != nil {
		return err
	}
	if u.userInfo == nil {
		u.LoginStatus = 2 // 当前user没有注册
		return nil
	}
	// 检测密码
	loginSuccess = u.CheckUserLogin()
	if loginSuccess {
		u.LoginStatus = 0 //登陆成功
	} else {
		u.LoginStatus = 1 //密码错误
	}
	return nil
}

func (u *UserLoginLogicMgr) CheckUserLogin() bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.userInfo.Passworld), []byte(u.passWord))
	if err != nil {
		return false
	}
	return true
}

func (u *UserLoginLogicMgr) loadUserInfo() error {
	hasUser, userInfo, err := common_dal.GetUserInfoByPhone(u.phone)
	if err != nil {
		return err
	}
	if hasUser {
		u.userInfo = userInfo
	}
	return nil
}
