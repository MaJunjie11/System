package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type ManagerAuditDoneCourseHandler struct {
	Req  *pb_gen.ManagerAuditDoneCourseRequest
	Resp *pb_gen.ManagerAuditDoneCourseResponse
}

func NewManagerAuditDoneCourseHandler(req *pb_gen.ManagerAuditDoneCourseRequest, resp *pb_gen.ManagerAuditDoneCourseResponse) *ManagerAuditDoneCourseHandler {
	return &ManagerAuditDoneCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (m *ManagerAuditDoneCourseHandler) Run() {
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

	course_dal.ManagerAuditDownCourseToDb(uid, m.Req.CourseId, m.Req.CourseAddr)

	m.Resp.Code = pb_gen.ErrNo_Success
	m.Resp.Msg = "添加成功"
}
