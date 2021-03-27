package user_logic

import (
	"System/pb_gen"
	"System/user/user_dal"
	"System/util"
	"fmt"

	"github.com/skoo87/log4go"
	"golang.org/x/crypto/bcrypt"
)

type UserLogicMgr struct {
	UserName     string
	Passworld    string
	TelephoneNum string
	UserType     pb_gen.UserType
	Sex          int32
	Uid          int64
	Age          int32
	ECPwd        string
}

func NewUserLogicMgr(userName, passWorld, telephoneNum string, sex, age int32, userType pb_gen.UserType) *UserLogicMgr {
	return &UserLogicMgr{
		UserName:     userName,
		Passworld:    passWorld,
		TelephoneNum: telephoneNum,
		Sex:          sex,
		Age:          age,
		UserType:     userType,
	}
}

// 对于用户登陆的密码进行加密算法处理 并且可以防止字典攻击
func (u *UserLogicMgr) EnCodePwd() error {
	var (
		err     error
		hasCode []byte
	)
	if hasCode, err = bcrypt.GenerateFromPassword([]byte(u.Passworld), bcrypt.DefaultCost); err != nil {
		log4go.Error("UserLogicMgr EnCodePwd err:%v", err)
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
		userInfo *pb_gen.UserInfo
	)
	// 1.幂等校验
	if hasUser, err = u.checkIfHasUser(); err != nil {
		log4go.Error("checkIfHasUser err:%v", err)
	} else if hasUser {
		// 已经拥有这个User 那么就 返回增加用户的错误信息
		return fmt.Errorf("addUserHadler hasUser")
	}
	// 2.生成用户UID
	u.Uid = util.Gen()

	// 3. 加密用户的passworld
	if err = u.EnCodePwd(); err != nil {
		log4go.Error("EnCodePwd err:%v", err)
		return err
	}

	// 4.pack 一下userInfo
	userInfo = u.packUserInfo()

	// 5.将数据存入数据库当中
	if err = u.addUserToDB(userInfo); err != nil {
		log4go.Error("save user_info into DB err:%v", err)
		return err
	}

	// 用户添加成功
	return nil
}

// 规定电话号码为唯一的不可重复键
func (u *UserLogicMgr) checkIfHasUser() (bool, error) {
	return user_dal.CheckIfHasUserFromDB(u.TelephoneNum)
}

func (u *UserLogicMgr) packUserInfo() *pb_gen.UserInfo {
	return &pb_gen.UserInfo{
		Uid:      u.Uid,
		Name:     u.UserName,
		Sex:      pb_gen.Sex(u.Sex),
		Email:    u.TelephoneNum,
		PassWord: u.ECPwd,
	}
}

func (u *UserLogicMgr) addUserToDB(userInfo *pb_gen.UserInfo) error {
	return user_dal.AddUserInfoToDB(userInfo)
}
