package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type StudentAddLikeCourseHandler struct {
	Req  *pb_gen.StudentAddLikeCourseRequest
	Resp *pb_gen.StudentAddLikeCourseResponse
}

func NewStudentAddLikeCouserHandler(req *pb_gen.StudentAddLikeCourseRequest, resp *pb_gen.StudentAddLikeCourseResponse) *StudentAddLikeCourseHandler {
	return &StudentAddLikeCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (s *StudentAddLikeCourseHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(s.Req.Token, util.StudentUserSecretKey); err != nil {
		s.Resp.Code = pb_gen.ErrNo_User_Token_Error
		s.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	if err = course_dal.StudentAddLikeCourse(uid, s.Req.CourseId); err != nil {
		s.Resp.Code = pb_gen.ErrNo_Internal_Err
		s.Resp.Msg = "加入表失败"
		return
	}

	s.Resp.Code = pb_gen.ErrNo_Success
	s.Resp.Msg = "添加成功"
}
