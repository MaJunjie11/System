package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type TeacherStartCourseHandler struct {
	Req  *pb_gen.TeacherStartCourseRequest
	Resp *pb_gen.TeacherStartCourseResponse
}

func NewTeacherStartCourseHandler(req *pb_gen.TeacherStartCourseRequest, resp *pb_gen.TeacherStartCourseResponse) *TeacherStartCourseHandler {
	return &TeacherStartCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherStartCourseHandler) Run() {
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
	course_dal.TeacherStartCourse(t.Req.CourseId)
	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "开课成功"

}
