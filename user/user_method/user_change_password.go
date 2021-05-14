package user_method

import (
	"System/db"
	"System/pb_gen"
	"System/user/user_dal"
	"System/user/user_moudle"
	"System/util"
	"fmt"

	"github.com/skoo87/log4go"
	"golang.org/x/crypto/bcrypt"
)

type UserChangePasswordHandler struct {
	Req      *pb_gen.UserChangePasswordRequest
	Resp     *pb_gen.UserChangePasswordResponse
	userInfo *user_moudle.UserStudentLoginInfo
	ECPwd    string
}

var (
	basePassWord string = "123456"
)

func NewUserChangePasswordHandler(req *pb_gen.UserChangePasswordRequest, resp *pb_gen.UserChangePasswordResponse) *UserChangePasswordHandler {
	return &UserChangePasswordHandler{
		Req:  req,
		Resp: resp,
	}
}

func (u *UserChangePasswordHandler) Run() {
	var (
		err           error
		userToken     *util.UserToken
		uid           int64
		userLoginInfo *user_moudle.UserStudentLoginInfo = &user_moudle.UserStudentLoginInfo{}
		count         int
	)
	// 1.鉴权
	// 1.鉴权
	if userToken, err = util.AuthToken(u.Req.Token, util.StudentUserSecretKey); err != nil {
		u.Resp.Code = pb_gen.ErrNo_User_Token_Error
		u.Resp.Msg = "token 认证失败"
		return
	}

	uid = userToken.Uid
	// 2.查数据
	db.Db.Where("uid = ?", uid).Find(userLoginInfo).Count(&count)
	if count < 1 {
		log4go.Warn("当前用户不存在,uid:%s", uid)
		u.Resp.Code = pb_gen.ErrNo_User_Not_Exist
		u.Resp.Msg = "用户不存在"
		return
	}
	u.userInfo = userLoginInfo
	// 3.鉴定old密码
	if !u.checkUserLogin() {
		log4go.Warn("原始密码错误,uid:%s", uid)
		u.Resp.Code = pb_gen.ErrNo_Login_Password_Error
		u.Resp.Msg = "原始密码错误"
		return
	}

	// 4. 修改密码
	if u.Req.Password != u.Req.ConfirmPassword {
		log4go.Warn("两次密码不同,uid:%s", uid)
		u.Resp.Code = pb_gen.ErrNo_Password_Not_Same
		u.Resp.Msg = "两次密码不同"
		return
	}

	// 5.修改密码
	if err = u.enCodePwd(); err != nil {
		log4go.Error("user u.enCodePwd err:%v", err)
		u.Resp.Code = pb_gen.ErrNo_Internal_Err
		u.Resp.Msg = "加密错误"
		return
	}

	if err = user_dal.ChangeUserPasswordToDbByUid(userLoginInfo.Uid, u.ECPwd); err != nil {
		u.Resp.Code = pb_gen.ErrNo_Internal_Err
		u.Resp.Msg = "修改密码错误"
		return
	}

	u.Resp.Code = pb_gen.ErrNo_Success
	u.Resp.Msg = "修改成功"
}

func (u *UserChangePasswordHandler) checkUserLogin() bool {
	// 初始密码 数据库存为nil 并且用户给的密码为basePassWord
	if u.userInfo.Password == "" && u.Req.OldPassword == basePassWord {
		fmt.Println("修改了初始密码")
		return true
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.userInfo.Password), []byte(u.Req.Password))
	if err != nil {
		return false
	}
	return true
}

func (u *UserChangePasswordHandler) enCodePwd() error {
	var (
		err     error
		hasCode []byte
	)
	if hasCode, err = bcrypt.GenerateFromPassword([]byte(u.Req.Password), bcrypt.DefaultCost); err != nil {
		log4go.Error("UserLogicMgr EnCodePwd err:%v", err)
		return err
	}
	// 保存在数据库中的密码 虽然每次不一样 但是只需要保存一份即可
	u.ECPwd = string(hasCode)
	return nil
}
