package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type TeacherEndCourseHandler struct {
	Req  *pb_gen.TeacherEndCourseRequest
	Resp *pb_gen.TeacherEndCourseResponse
}

func NewTeacherEndCourseHandler(req *pb_gen.TeacherEndCourseRequest, resp *pb_gen.TeacherEndCourseResponse) *TeacherEndCourseHandler {
	return &TeacherEndCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherEndCourseHandler) Run() {
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
	course_dal.TeacherEndCourse(t.Req.CourseId)
	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "结课成功"

}
