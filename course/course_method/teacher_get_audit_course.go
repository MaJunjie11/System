package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type TeacherGetAuditCourseHandler struct {
	Req  *pb_gen.TeacherGetAuditCourseRequest
	Resp *pb_gen.TeacherGetAuditCourseResponse
}

func NewTeacherGetAuditCourseHandler(req *pb_gen.TeacherGetAuditCourseRequest, resp *pb_gen.TeacherGetAuditCourseResponse) *TeacherGetAuditCourseHandler {
	return &TeacherGetAuditCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherGetAuditCourseHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.TeacherGetAuditCourseResponseData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(t.Req.Token, util.StudentUserSecretKey); err != nil {
		t.Resp.Code = pb_gen.ErrNo_User_Token_Error
		t.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	if data = course_dal.TeacherGetAuditInfoFromDb(uid); data == nil {
		data = &pb_gen.TeacherGetAuditCourseResponseData{
			Total: 0,
			Items: make([]*pb_gen.AuditCourseData, 0),
		}
		data1 := &pb_gen.AuditCourseData{}
		data.Items = append(data.Items, data1)
	}

	// 获取数据

	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "申请成功"
	t.Resp.Data = data
}
