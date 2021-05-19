package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type ManagerGetAuditCourseHandler struct {
	Req  *pb_gen.ManagerGetAuditCourseRequest
	Resp *pb_gen.ManagerGetAuditCourseResponse
}

func NewManagerGetAuditCourseHandler(req *pb_gen.ManagerGetAuditCourseRequest, resp *pb_gen.ManagerGetAuditCourseResponse) *ManagerGetAuditCourseHandler {
	return &ManagerGetAuditCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (m *ManagerGetAuditCourseHandler) Run() {
	// 管理员获得
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.ManagerGetAuditCourseResponseData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(m.Req.Token, util.StudentUserSecretKey); err != nil {
		m.Resp.Code = pb_gen.ErrNo_User_Token_Error
		m.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	if data = course_dal.ManagerGetAuditCourseFromDb(uid); data == nil {
		data = &pb_gen.ManagerGetAuditCourseResponseData{
			Total: 0,
			Items: make([]*pb_gen.AuditCourseData, 0),
		}
		data1 := &pb_gen.AuditCourseData{}
		data.Items = append(data.Items, data1)
	}

	// 获取数据

	m.Resp.Code = pb_gen.ErrNo_Success
	m.Resp.Msg = "申请成功"
	m.Resp.Data = data
}
