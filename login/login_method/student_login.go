package login_method

import (
	"System/db"
	"System/pb_gen"
	"System/user/user_moudle"
	"System/util"
	"fmt"

	"github.com/skoo87/log4go"
	"golang.org/x/crypto/bcrypt"
)

type StudentLoginHandler struct {
	Req      *pb_gen.StudentLoginRequest
	Resp     *pb_gen.StudentLoginResponse
	userInfo *user_moudle.UserStudentLoginInfo
}

var basePassWord string = "123456"

func NewStudentLoginHandler(req *pb_gen.StudentLoginRequest, resp *pb_gen.StudentLoginResponse) *StudentLoginHandler {
	return &StudentLoginHandler{
		Req:  req,
		Resp: resp,
	}
}
func (u *StudentLoginHandler) CheckUserLogin() bool {
	// 初始密码 数据库存为nil 并且用户给的密码为basePassWord
	if u.userInfo.Password == "" && u.Req.Password == basePassWord {
		return true
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.userInfo.Password), []byte(u.Req.Password))
	if err != nil {
		return false
	}
	return true
}

func (s *StudentLoginHandler) Run() {

	// 登录主要就是校验密码
	var (
		count int
	)
	userInfo := &user_moudle.UserStudentLoginInfo{}
	db.Db.Where("uuid = ?", s.Req.Uuid).Find(userInfo).Count(&count)
	if count < 1 {
		log4go.Warn("当前用户不存在:email:%s", s.Req.Uuid)
		s.Resp.Code = pb_gen.ErrNo_User_Not_Exist
		s.Resp.Msg = "用户不存在"
		return
	}
	s.userInfo = userInfo
	// 检测用户密码是否通过非对称加密算法
	if !s.CheckUserLogin() {
		s.Resp.Code = pb_gen.ErrNo_Login_Password_Error
		s.Resp.Msg = "输入密码错误"
		return
	}

	s.Resp.Code = pb_gen.ErrNo_Success
	s.Resp.Msg = "登录成功"

	// 生成一个token
	userToken, err := util.GenToken(userInfo.Uid, userInfo.Limit, util.StudentUserExpireDuration, util.StudentUserSecretKey)
	if err != nil && userToken == "" {
		s.Resp.Code = pb_gen.ErrNo_Internal_Err
		s.Resp.Msg = "token生成失败"
		return
	}

	data := &pb_gen.StudentLoginData{
		Key:   "token",
		Token: userToken,
	}
	fmt.Printf("userToken:%s\n", userToken)
	s.Resp.Data = data
}
