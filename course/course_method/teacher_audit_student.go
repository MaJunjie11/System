package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type TeacherAuditStudentHandler struct {
	Req  *pb_gen.TeacherAuditStudentRequest
	Resp *pb_gen.TeacherAuditStudentResponse
}

func NewTeacherAuditStudentHandler(req *pb_gen.TeacherAuditStudentRequest, resp *pb_gen.TeacherAuditStudentResponse) *TeacherAuditStudentHandler {
	return &TeacherAuditStudentHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherAuditStudentHandler) Run() {

	// 1. 鉴权token
	var (
		err         error
		userToken   *util.UserToken
		teacher_uid int64
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(t.Req.Token, util.StudentUserSecretKey); err != nil {
		t.Resp.Code = pb_gen.ErrNo_User_Token_Error
		t.Resp.Msg = "token 认证失败"
		return
	}
	teacher_uid = userToken.Uid
	fmt.Println("courseId:", t.Req.CourseId)
	fmt.Println("teacher_uid:", teacher_uid)
	fmt.Println("uid:", t.Req.Uid)

	course_dal.TeacherAuditStudentToDb(t.Req.CourseId, t.Req.Uid)
	// 修改数据库数据
	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "添加成功"
}
