package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type StudentGetStartCourseHandler struct {
	Req  *pb_gen.StudentGetStartCourseRequest
	Resp *pb_gen.StudentGetStartCourseResponse
}

func NewStudentGetStartCourseHandler(req *pb_gen.StudentGetStartCourseRequest, resp *pb_gen.StudentGetStartCourseResponse) *StudentGetStartCourseHandler {
	return &StudentGetStartCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (s *StudentGetStartCourseHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.StudentGetStartCourseResponseData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(s.Req.Token, util.StudentUserSecretKey); err != nil {
		s.Resp.Code = pb_gen.ErrNo_User_Token_Error
		s.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	// 2.查数据库
	data = course_dal.GetStudentStartCourseListFromDbByUid(uid)
	// 3. 组织数据返回
	if data == nil {
		fmt.Println("自己组织一个")
		data = &pb_gen.StudentGetStartCourseResponseData{
			Total: 0,
			Items: make([]*pb_gen.StudentCourseData, 0),
		}
		data1 := &pb_gen.StudentCourseData{}
		data.Items = append(data.Items, data1)
	}
	s.Resp.Data = data
}
