package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type TeacherDeleteRefuseCourseHandler struct {
	Req  *pb_gen.TeacherDeleteRefuseCourseRequest
	Resp *pb_gen.TeacherDeleteRefuseCourseResponse
}

func NewTeacherDeleteRefuseCourseHandler(req *pb_gen.TeacherDeleteRefuseCourseRequest, resp *pb_gen.TeacherDeleteRefuseCourseResponse) *TeacherDeleteRefuseCourseHandler {
	return &TeacherDeleteRefuseCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherDeleteRefuseCourseHandler) Run() {
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
	course_dal.TeacherDeleteRefuseCourse(t.Req.CourseId)
	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "删除成功"

}
