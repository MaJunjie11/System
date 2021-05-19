package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type GetCourseStudentHandler struct {
	Req  *pb_gen.GetCourseStudentRequest
	Resp *pb_gen.GetCourseStudentResponse
}

func NewGetCourseStudentHandler(req *pb_gen.GetCourseStudentRequest, resp *pb_gen.GetCourseStudentResponse) *GetCourseStudentHandler {
	return &GetCourseStudentHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *GetCourseStudentHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.GetCourseStudentResponseData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(t.Req.Token, util.StudentUserSecretKey); err != nil {
		t.Resp.Code = pb_gen.ErrNo_User_Token_Error
		t.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	data = course_dal.GetCourseStudentFromDbByCourseId(t.Req.CourseId, uid)
	if data == nil {
		data = &pb_gen.GetCourseStudentResponseData{
			Total: 0,
			Items: make([]*pb_gen.CourseStudentData, 0),
		}
		data1 := &pb_gen.CourseStudentData{}
		data.Items = append(data.Items, data1)
	}
	// 获取数据
	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "获取成功"
	t.Resp.Data = data
}
