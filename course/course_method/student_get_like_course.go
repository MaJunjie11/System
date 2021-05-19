package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type StudentGetLikeCourseHandler struct {
	Req  *pb_gen.StudentGetLikeCourseRequest
	Resp *pb_gen.StudentGetLikeCourseResponse
}

func NewStudentGetLikeCourseHandler(req *pb_gen.StudentGetLikeCourseRequest, resp *pb_gen.StudentGetLikeCourseResponse) *StudentGetLikeCourseHandler {
	return &StudentGetLikeCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (s *StudentGetLikeCourseHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.StudentGetLikeCourseResponseData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(s.Req.Token, util.StudentUserSecretKey); err != nil {
		s.Resp.Code = pb_gen.ErrNo_User_Token_Error
		s.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	// 2.查数据库
	data = course_dal.GetStudentLikeCourseListFromDbByUid(uid)
	// 3. 组织数据返回
	if data == nil {
		fmt.Println("自己组织一个")
		data = &pb_gen.StudentGetLikeCourseResponseData{
			Total: 0,
			Items: make([]*pb_gen.StudentGetLikeCourseData, 0),
		}
		data1 := &pb_gen.StudentGetLikeCourseData{}
		data.Items = append(data.Items, data1)
	}
	s.Resp.Data = data
}
