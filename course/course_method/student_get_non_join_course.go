package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type StudentGetNonJoinCourseHandler struct {
	Req  *pb_gen.StudentGetNonJoinCourseRequset
	Resp *pb_gen.StudentGetNonJoinCourseResponse
}

func NewStudentGetNonJoinCourseHandler(req *pb_gen.StudentGetNonJoinCourseRequset, resp *pb_gen.StudentGetNonJoinCourseResponse) *StudentGetNonJoinCourseHandler {
	return &StudentGetNonJoinCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (s *StudentGetNonJoinCourseHandler) Run() {
	// 获取学生不能加入的课程列表
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.StudentGetNonJoinCourseResponseData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(s.Req.Token, util.StudentUserSecretKey); err != nil {
		s.Resp.Code = pb_gen.ErrNo_User_Token_Error
		s.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	// 2.查数据库
	data = course_dal.GetStudentNonJoinCourseByUid(uid)
	// 3. 组织数据返回
	if data == nil {
		fmt.Println("自己组织一个")
		data = &pb_gen.StudentGetNonJoinCourseResponseData{
			Total: 0,
			Items: make([]*pb_gen.StudentNonJoinCourseData, 0),
		}
		data1 := &pb_gen.StudentNonJoinCourseData{}
		data.Items = append(data.Items, data1)
	}
	s.Resp.Data = data
}
