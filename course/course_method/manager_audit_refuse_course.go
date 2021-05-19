package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type ManagerAuditRefuseCourseHandler struct {
	Req  *pb_gen.ManagerAuditRefuseCourseRequest
	Resp *pb_gen.ManagerAuditRefuseCourseResponse
}

func NewManagerAuditRefuseCourseHandler(req *pb_gen.ManagerAuditRefuseCourseRequest, resp *pb_gen.ManagerAuditRefuseCourseResponse) *ManagerAuditRefuseCourseHandler {
	return &ManagerAuditRefuseCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (m *ManagerAuditRefuseCourseHandler) Run() {
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
	course_dal.ManagerAddRefuseCourseToDb(uid, m.Req.CourseId, m.Req.RefuseReason)

	m.Resp.Code = pb_gen.ErrNo_Success
	m.Resp.Msg = "添加成功"
}
