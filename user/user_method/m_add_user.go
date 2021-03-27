package user_method

import (
	"System/pb_gen"
)

type AddStudentHandler struct {
	Req  *pb_gen.AddStudentRequest
	Resp *pb_gen.AddStudentResponse
}

func NewAddUserHandler(req *pb_gen.AddStudentRequest, resp *pb_gen.AddStudentResponse) *AddStudentHandler {
	return &AddStudentHandler{
		Req:  req,
		Resp: resp,
	}
}

func (u *AddStudentHandler) Run() {
	// 1.对比验证码
	// 2.检测是否已注册
	// 3.对比password
	// 4.生成uid
	// 5.user信息入库
	// 6.返回成功信息

}
