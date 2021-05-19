package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type TeacherGetEndCourseHandler struct {
	Req  *pb_gen.TeacherGetEndCourseRequest
	Resp *pb_gen.TeacherGetEndCourseResponse
}

func NewTeacherGetEndCourseHandler(req *pb_gen.TeacherGetEndCourseRequest, resp *pb_gen.TeacherGetEndCourseResponse) *TeacherGetEndCourseHandler {
	return &TeacherGetEndCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherGetEndCourseHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.TeacherGetEndCourseResponseData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(t.Req.Token, util.StudentUserSecretKey); err != nil {
		t.Resp.Code = pb_gen.ErrNo_User_Token_Error
		t.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	if data = course_dal.TeacherGetEndCouserFromDb(uid); data == nil {
		data = &pb_gen.TeacherGetEndCourseResponseData{
			Total: 0,
			Items: make([]*pb_gen.TeacherCourseDetailData, 0),
		}
		data1 := &pb_gen.TeacherCourseDetailData{}
		data.Items = append(data.Items, data1)
	}
	// 获取数据
	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "申请成功"
	t.Resp.Data = data
}
