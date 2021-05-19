package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type TeacherFeetStudentHandler struct {
	Req  *pb_gen.TeacherFeetStudentRequest
	Resp *pb_gen.TeacherFeetStudentResponse
}

func NewTeacherFeetStudentHandler(req *pb_gen.TeacherFeetStudentRequest, resp *pb_gen.TeacherFeetStudentResponse) *TeacherFeetStudentHandler {
	return &TeacherFeetStudentHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherFeetStudentHandler) Run() {
	// 1. 鉴权token
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
	fmt.Println(uid)

	course_dal.FeetStudentInDb(t.Req.Uid, t.Req.CourseId)
	uid = userToken.Uid
	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "删除成功"
}
