package login_method

import (
	"System/db"
	"System/pb_gen"
	"System/user/user_moudle"

	"github.com/skoo87/log4go"
	"golang.org/x/crypto/bcrypt"
)

type StudentLoginHandler struct {
	Req      *pb_gen.StudentLoginRequest
	Resp     *pb_gen.StudentLoginResponse
	userInfo *user_moudle.UserStudent
}

func NewStudentLoginHandler(req *pb_gen.StudentLoginRequest, resp *pb_gen.StudentLoginResponse) *StudentLoginHandler {
	return &StudentLoginHandler{
		Req:  req,
		Resp: resp,
	}
}
func (u *StudentLoginHandler) CheckUserLogin() bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Req.Password), []byte(u.userInfo.Password))
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
	userInfo := &user_moudle.UserStudent{}
	db.Db.Where("email = ?", s.Req.Email).Find(userInfo).Count(&count)
	if count < 1 {
		log4go.Warn("当前用户不存在:email:%s", s.Req.Email)
		s.Resp.Status = pb_gen.ErrNo_User_Not_Exist
		s.Resp.Msg = "用户不存在"
		return
	}
	s.userInfo = userInfo

	// 检测用户密码是否通过非对称加密算法
	if s.CheckUserLogin() {
		s.Resp.Status = pb_gen.ErrNo_Login_Password_Error
		s.Resp.Msg = "输入密码错误"
		return
	}

	s.Resp.Status = pb_gen.ErrNo_Success
	s.Resp.Msg = "登录成功"
	s.Resp.Uid = userInfo.Uid
	s.Resp.Email = userInfo.Email
	s.Resp.Token = "虚假token"
}
