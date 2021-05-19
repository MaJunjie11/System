package course_method

import (
	"System/course/course_dal"
	"System/course/course_model"
	"System/pb_gen"
	"System/util"
	"fmt"
	"strconv"
	"time"
)

type TeacherAddCourseHandler struct {
	Req             *pb_gen.TeacherAddCourseRequest
	Resp            *pb_gen.TeacherAddCourseResponse
	courseAuditInfo *course_model.CourseAuditInfo
}

func NewTeacherAddCourseHandler(req *pb_gen.TeacherAddCourseRequest, resp *pb_gen.TeacherAddCourseResponse) *TeacherAddCourseHandler {
	return &TeacherAddCourseHandler{
		Req:  req,
		Resp: resp,
	}
}

func (t *TeacherAddCourseHandler) Run() {
	// 1. 鉴权token
	var (
		err       error
		userToken *util.UserToken
		uid       int64
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(t.Req.Token, util.StudentUserSecretKey); err != nil {
		t.Resp.Code = pb_gen.ErrNo_User_Token_Error
		t.Resp.Msg = "token 认证失败"
		return
	}
	uid = userToken.Uid

	managerUid := course_dal.GetManagerUidFromUuid(t.Req.ManagerUuid)
	if managerUid == -1 {
		t.Resp.Code = pb_gen.ErrNo_User_Not_Exist
		t.Resp.Msg = "没有当前管理员"
		return
	}

	// 打包数据
	t.packCourseAuditInfo(uid, managerUid)

	// 存入数据库
	course_dal.AddCourseAuditInfo(t.courseAuditInfo)

	t.Resp.Code = pb_gen.ErrNo_Success
	t.Resp.Msg = "申请成功"
}

func (t *TeacherAddCourseHandler) packCourseAuditInfo(uid, managerUid int64) {
	cap, _ := strconv.ParseInt(t.Req.ClassCapacity, 10, 32)
	data := &course_model.CourseAuditInfo{
		CourseId:        fmt.Sprintf("%d", time.Now().UnixNano()),
		TeacherUid:      uid,
		ClassName:       t.Req.ClassName,
		CourseDate:      t.Req.CourseDate,
		CourseMajor:     t.Req.CourseMajor,
		CourseIntroduce: t.Req.CourseIntroduce,
		ClassCapacity:   int(cap),
		ManagerUid:      managerUid,
	}
	t.courseAuditInfo = data
}
