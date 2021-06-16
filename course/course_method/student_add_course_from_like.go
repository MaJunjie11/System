package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type StudentAddCourseFromLikeHandler struct {
	Req  *pb_gen.StudentAddCourseFromLikeRequest
	Resp *pb_gen.StudentAddCourseFromLikeResponse
}

func NewStudentAddCouserFromLikeHandler(req *pb_gen.StudentAddCourseFromLikeRequest, resp *pb_gen.StudentAddCourseFromLikeResponse) *StudentAddCourseFromLikeHandler {
	return &StudentAddCourseFromLikeHandler{
		Req:  req,
		Resp: resp,
	}
}

func (s *StudentAddCourseFromLikeHandler) Run() {
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
	if err = course_dal.StudentAddCourseFromLikeToDb(uid, s.Req.CourseId); err != nil {
		s.Resp.Code = pb_gen.ErrNo_Internal_Err
		s.Resp.Msg = "人数已满或已经开课"
		return
	}

	s.Resp.Code = pb_gen.ErrNo_Success
	s.Resp.Msg = "添加成功"
}
