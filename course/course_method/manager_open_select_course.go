package course_method

import (
	"System/course/course_model"
	"System/db"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type ManagerOpenSelectCourseHandler struct {
	Req  *pb_gen.ManagerOpenSelectCourseRequest
	Resp *pb_gen.ManagerOpenSelectCourseResponse
}

func NewManagerOpenSelectCourseHandler(req *pb_gen.ManagerOpenSelectCourseRequest, resp *pb_gen.ManagerOpenSelectCourseResponse) *ManagerOpenSelectCourseHandler {
	return &ManagerOpenSelectCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (m *ManagerOpenSelectCourseHandler) Run() {
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

	db.Db.Model(&course_model.CourseOpenInfo{}).Where("id=?", 1).Update("status", 1)
	m.Resp.Code = pb_gen.ErrNo_Success
	m.Resp.Msg = "开启成功"

}
