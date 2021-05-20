package user_method

import (
	"System/pb_gen"
	"System/user/user_dal"
	"System/user/user_moudle"
	"System/util"
	"fmt"
)

type GetUserInfoHandler struct {
	Req  *pb_gen.GetUserInfoRequest
	Resp *pb_gen.GetUserInfoResponse
}

func NewGetUserInfoHandler(req *pb_gen.GetUserInfoRequest, resp *pb_gen.GetUserInfoResponse) *GetUserInfoHandler {
	return &GetUserInfoHandler{
		Req:  req,
		Resp: resp,
	}
}

func (g *GetUserInfoHandler) Run() {
	var (
		err          error
		userToken    *util.UserToken
		uid          int64
		userBaseInfo *user_moudle.UserStudentBaseInfo
	)

	// 1.鉴权
	if userToken, err = util.AuthToken(g.Req.Token, util.StudentUserSecretKey); err != nil {
		g.Resp.Code = pb_gen.ErrNo_User_Token_Error
		g.Resp.Msg = "token 认证失败"
		return
	}

	uid = userToken.Uid
	// 2.查数据库
	userBaseInfo, err = user_dal.GetUserBasicInfoFromDb(uid)
	if err != nil {
		g.Resp.Code = pb_gen.ErrNo_User_Not_Exist
		g.Resp.Msg = "没有用户"
		return
	}
	// 3.pack结构
	g.packRespData(userBaseInfo)
	// 4. 返回
	g.Resp.Code = pb_gen.ErrNo_Success
	g.Resp.Msg = "获取成功"
}

func (g *GetUserInfoHandler) packRespData(userBaseInfo *user_moudle.UserStudentBaseInfo) {
	data := &pb_gen.GetUserInfoData{
		Avatar:    userBaseInfo.Avatar,
		Signature: userBaseInfo.Signature,
		Name:      userBaseInfo.Name,
		Uuid:      userBaseInfo.Uuid,
	}
	//TODO: 后续需要优化
	if userBaseInfo.Limit == 2 {
		data.Limit = "/selectRoom,/myClass/startClass,/myClass/untaughtClass,/myClass/endClass,/nonjoinClass,/likeClass"
	} else if userBaseInfo.Limit == 1 {
		data.Limit = "/addRoom,/auditRoom,/teacherClass/startClass,/teacherClass/untaughtClass,/teacherClass/endClass,/refuseCourse"
	} else if userBaseInfo.Limit == 3 {
		data.Limit = "/managerAuditCourse,/resetPassword"
	}
	fmt.Println("Limit:", data.Limit)
	g.Resp.Data = data
}
