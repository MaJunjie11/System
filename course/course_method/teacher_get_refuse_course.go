package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type TeacherGetRefuseCourseHandler struct {
	Req  *pb_gen.TeacherGetRefuseCourseRequest
	Resp *pb_gen.TeacherGetRefuseCourseResponse
}

func NewTeacherGetRefuseCourseHandler(req *pb_gen.TeacherGetRefuseCourseRequest, resp *pb_gen.TeacherGetRefuseCourseResponse) *TeacherGetRefuseCourseHandler {
	return &TeacherGetRefuseCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherGetRefuseCourseHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.TeacherGetRefuseCourseResponseData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(t.Req.Token, util.StudentUserSecretKey); err != nil {
		t.Resp.Code = pb_gen.ErrNo_User_Token_Error
		t.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	if data = course_dal.TeacherGetRefuseInfoFromDb(uid); data == nil {
		data = &pb_gen.TeacherGetRefuseCourseResponseData{
			Total: 0,
			Items: make([]*pb_gen.RefuseCourseData, 0),
		}
		data1 := &pb_gen.RefuseCourseData{}
		data.Items = append(data.Items, data1)
	}

	// 获取数据

	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "申请成功"
	t.Resp.Data = data
}
