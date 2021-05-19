package course_method

import (
	"System/course/course_dal"
	"System/pb_gen"
	"System/util"
)

type GetStudentAddCourseListHandler struct {
	Req  *pb_gen.StudentGetAddCourseListRequest
	Resp *pb_gen.StudentGetAddCourseListResponse
}

func NewGetStudentAddCourseListHandler(req *pb_gen.StudentGetAddCourseListRequest, resp *pb_gen.StudentGetAddCourseListResponse) *GetStudentAddCourseListHandler {
	return &GetStudentAddCourseListHandler{
		Req:  req,
		Resp: resp,
	}
}
func (g *GetStudentAddCourseListHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
		data      *pb_gen.StudentGetAddCourseListData
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(g.Req.Token, util.StudentUserSecretKey); err != nil {
		g.Resp.Code = pb_gen.ErrNo_User_Token_Error
		g.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid
	// 2.查数据库
	data = course_dal.GetStudentAddCourseListFromDbByUid(uid)
	// 3. 组织数据返回
	if data == nil {
		data = &pb_gen.StudentGetAddCourseListData{
			Count: 1,
			Rows:  make([]*pb_gen.AddCourseData, 0),
		}
		data1 := &pb_gen.AddCourseData{}
		data.Rows = append(data.Rows, data1)
	}
	g.Resp.Data = data
}
