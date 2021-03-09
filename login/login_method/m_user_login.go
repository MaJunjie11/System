package login_method

import (
	"OpenExperimentalManagementSystem/pb_gen"
	"fmt"
	"log"
)

type UserLoginHander struct {
	Req  *pb_gen.UserLoginRequest
	Resp *pb_gen.UserLoginResponse
}

func (u *UserLoginHander) Run() {
	// checkParam
	var err error
	if err = u.checkParam(); err != nil {
		log.Fatal(fmt.Sprintf("err in UserLoginHander checkParam:%v", err))
		u.Resp.ErrNo = 0
		u.Resp.ErrTips = "param error"
		u.Resp.LoginStatus = 0
		return
	}

	// TODO: 处理登陆业务逻辑
}

func (u *UserLoginHander) checkParam() error {
	return nil
}
