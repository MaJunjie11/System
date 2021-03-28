package user_method

import (
	"System/db"
	"System/message/message_method"
	"System/pb_gen"
	"System/user/user_moudle"
	"System/util"
	"time"

	"github.com/skoo87/log4go"
	"golang.org/x/crypto/bcrypt"
)

type AddStudentHandler struct {
	Req   *pb_gen.AddStudentRequest
	Resp  *pb_gen.AddStudentResponse
	ECPwd string
}

func NewAddUserHandler(req *pb_gen.AddStudentRequest, resp *pb_gen.AddStudentResponse) *AddStudentHandler {
	return &AddStudentHandler{
		Req:  req,
		Resp: resp,
	}
}

func (u *AddStudentHandler) packStudent(uid int64) *user_moudle.UserStudent {
	return &user_moudle.UserStudent{
		Id:         1,
		Uid:        uid,
		Email:      u.Req.Email,
		Password:   u.ECPwd,
		Desc:       "",
		Status:     0,
		CreateTime: time.Now(),
	}
}

func (u *AddStudentHandler) enCodePwd() error {
	var (
		err     error
		hasCode []byte
	)
	if hasCode, err = bcrypt.GenerateFromPassword([]byte(u.Req.PassWord), bcrypt.DefaultCost); err != nil {
		log4go.Error("UserLogicMgr EnCodePwd err:%v", err)
		return err
	}
	// 保存在数据库中的密码 虽然每次不一样 但是只需要保存一份即可
	u.ECPwd = string(hasCode)
	return nil
}

func (u *AddStudentHandler) checkIfHasUser() bool {
	//判断是否有这个user
	var (
		hasUser bool
		count   int
	)

	userInfo := user_moudle.UserStudent{}
	db.Db.Where("email = ?", u.Req.Email).Find(&userInfo).Count(&count)
	if count < 1 {
		hasUser = false
	} else {
		hasUser = true
	}
	return hasUser
}

func (u *AddStudentHandler) Run() {
	var (
		err     error
		student *user_moudle.UserStudent
		hasUser bool
	)
	// 1.检测是否已注册
	hasUser = u.checkIfHasUser()
	// 如果已经有这个User
	if hasUser {
		u.Resp.Status = pb_gen.ErrNo_User_Already_Exist
		u.Resp.Msg = "已经有这个用户"
		return
	}

	// 2.对比验证码
	c := message_method.LocalCache
	authNum, ok := c.Get(u.Req.Email)
	if !ok {
		u.Resp.Status = pb_gen.ErrNo_Cache_No_Verification_code
		u.Resp.Msg = "没有次邮箱验证码"
		return
	}
	if authNum != u.Req.AutoNum {
		u.Resp.Status = pb_gen.ErrNo_Verification_Code_Error
		u.Resp.Msg = "验证码不正确"
		return
	}
	// 3.对比password
	if u.Req.GetPassWord() != u.Req.GetMutilPassWord() {
		u.Resp.Status = pb_gen.ErrNo_Password_Not_Same
		u.Resp.Msg = "两次密码输入不同"
		return
	}

	// 4.生成uid
	uid := util.Gen()
	// 生成密码加密密钥
	if err = u.enCodePwd(); err != nil {
		log4go.Error("user register u.enCodePwd err:%v", err)
		u.Resp.Status = pb_gen.ErrNo_Internal_Err
		u.Resp.Msg = "加密错误"
		return
	}
	student = u.packStudent(uid)
	// 5.user信息入库
	db.Db.Create(student)
	// 6.返回成功信息
	u.Resp.Status = pb_gen.ErrNo_Success
	u.Resp.Msg = "注册成功啦哈哈哈"

}
