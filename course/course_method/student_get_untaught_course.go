package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
	"fmt"
)

type StudentGetUntaughtCourseHandler struct {
	Req  *pb_gen.StudentGetUntaughtCourseRequest
	Resp *pb_gen.StudentGetUntaughtCourseResponse
}

func NewStudentGetAlreadyRoomHandler(req *pb_gen.StudentGetUntaughtCourseRequest, resp *pb_gen.StudentGetUntaughtCourseResponse) *StudentGetUntaughtCourseHandler {
	return &StudentGetUntaughtCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (a *StudentGetUntaughtCourseHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.StudentGetUntaughtCourseResponseData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(a.Req.Token, util.StudentUserSecretKey); err != nil {
		a.Resp.Code = pb_gen.ErrNo_User_Token_Error
		a.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	// 2.查数据库
	data = course_dal.GetStudentUntaughtCourseListFromDbByUid(uid)
	// 3. 组织数据返回
	if data == nil {
		fmt.Println("自己组织一个")
		data = &pb_gen.StudentGetUntaughtCourseResponseData{
			Total: 0,
			Items: make([]*pb_gen.StudentCourseData, 0),
		}
		data1 := &pb_gen.StudentCourseData{}
		data.Items = append(data.Items, data1)
	}
	a.Resp.Data = data
}
