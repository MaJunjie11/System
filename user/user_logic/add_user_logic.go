package user_logic

import (
	"OpenExperimentalManagementSystem/user/user_dal"
	"OpenExperimentalManagementSystem/user/user_moudle"
	"OpenExperimentalManagementSystem/util"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserLogicMgr struct {
	UserName     string
	Passworld    string
	TelephoneNum string
	Sex          int32
	Uid          int64
	Age          int32
	ECPwd        string
}

func NewUserLogicMgr(userName, passWorld, telephoneNum string, sex, age int32) *UserLogicMgr {
	return &UserLogicMgr{
		UserName:     userName,
		Passworld:    passWorld,
		TelephoneNum: telephoneNum,
		Sex:          sex,
		Age:          age,
	}
}

// 对于用户登陆的密码进行加密算法处理 并且可以防止字典攻击
func (u *UserLogicMgr) EnCodePwd() error {
	var (
		err     error
		hasCode []byte
	)
	if hasCode, err = bcrypt.GenerateFromPassword([]byte(u.Passworld), bcrypt.DefaultCost); err != nil {
		// TODO: 后续增加log模块
		return err
	}
	// 保存在数据库中的密码 虽然每次不一样 但是只需要保存一份即可
	u.ECPwd = string(hasCode)
	return nil
}

func (u *UserLogicMgr) Process() error {
	// 1. 首先进行校验逻辑 可以根据用户的phone来校验
	var (
		hasUser  bool
		err      error
		userInfo *user_moudle.UserInfo
	)
	// 1.幂等校验
	if hasUser, err = u.checkIfHasUser(); err != nil {
		//TODO: 增加log
	} else if hasUser {
		// 已经拥有这个User 那么就 返回增加用户的错误信息
		return fmt.Errorf("addUserHadler hasUser")
	}
	// 2.生成用户UID
	u.Uid = util.Gen()

	// 3. 加密用户的passworld
	if err = u.EnCodePwd(); err != nil {
		return err
	}

	// 4.pack 一下userInfo
	userInfo = u.packUserInfo()

	// 5.将数据存入数据库当中
	if err = u.addUserToDB(userInfo); err != nil {
		//
		return err
	}

	// 用户添加成功
	return nil
}

// 规定电话号码为唯一的不可重复键
func (u *UserLogicMgr) checkIfHasUser() (bool, error) {
	return user_dal.CheckIfHasUserFromDB(u.TelephoneNum)
}

func (u *UserLogicMgr) packUserInfo() *user_moudle.UserInfo {
	return &user_moudle.UserInfo{
		Uid:       u.Uid,
		Name:      u.UserName,
		Sex:       u.Sex,
		Phone:     u.TelephoneNum,
		Passworld: u.ECPwd,
	}
}

func (u *UserLogicMgr) addUserToDB(userInfo *user_moudle.UserInfo) error {
	return user_dal.AddUserInfoToDB(userInfo)
}
