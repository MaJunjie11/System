package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type ManagerResetPasswordHandler struct {
	Req  *pb_gen.ManagerResetPasswordRequest
	Resp *pb_gen.ManagerResetPasswordResponse
}

func NewManagerResetPasswordHandler(req *pb_gen.ManagerResetPasswordRequest, resp *pb_gen.ManagerResetPasswordResponse) *ManagerResetPasswordHandler {
	return &ManagerResetPasswordHandler{
		Req:  req,
		Resp: resp,
	}
}

func (m *ManagerResetPasswordHandler) Run() {

	// 管理员获得
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(m.Req.Token, util.StudentUserSecretKey); err != nil {
		m.Resp.Code = pb_gen.ErrNo_User_Token_Error
		m.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	fmt.Println(uid)
	course_dal.ManagerResetPassword(m.Req.Uuid)

	m.Resp.Code = pb_gen.ErrNo_Success
	m.Resp.Msg = "重置成功"
}
