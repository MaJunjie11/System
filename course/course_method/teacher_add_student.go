package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type TeacherAddStudentHandler struct {
	Req  *pb_gen.TeacherAddStudentRequest
	Resp *pb_gen.TeacherAddStudentResponse
}

func NewTeacherAddStudentHandler(req *pb_gen.TeacherAddStudentRequest, resp *pb_gen.TeacherAddStudentResponse) *TeacherAddStudentHandler {
	return &TeacherAddStudentHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherAddStudentHandler) Run() {
	var (
		err       error
		userToken *util.UserToken
		uid       int64
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(t.Req.Token, util.StudentUserSecretKey); err != nil {
		t.Resp.Code = pb_gen.ErrNo_User_Token_Error
		t.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	// TODO:
	fmt.Println("uid:", uid)

	if err = course_dal.TeacherAddStudent(t.Req.Uuid, t.Req.CourseId); err != nil {
		t.Resp.Code = pb_gen.ErrNo_User_Not_Exist
		t.Resp.Msg = "学生不存在"
		return
	}

	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "申请成功"

}
