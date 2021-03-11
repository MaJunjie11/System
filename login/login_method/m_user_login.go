package login_method

import (
	"System/login/login_logic"
	"System/pb_gen"
	"fmt"

	"github.com/skoo87/log4go"
)

type UserLoginHander struct {
	Req  *pb_gen.UserLoginRequest
	Resp *pb_gen.UserLoginResponse
}

func (u *UserLoginHander) Run() {
	// checkParam
	var (
		err          error
		userLoginMgr *login_logic.UserLoginLogicMgr
	)
	if err = u.checkParam(); err != nil {
		log4go.Error(fmt.Sprintf("err in UserLoginHander checkParam:%v", err))
		u.Resp.ErrNo = 1 // 错误
		u.Resp.ErrTips = "param error"
		u.Resp.LoginStatus = 0
		return
	}
	userLoginMgr = login_logic.NewUserLogicMgr(u.Req.Phone, u.Req.Phone)
	if err = userLoginMgr.Process(); err != nil {
		log4go.Error(fmt.Sprintf("err in UserLogin process:%v", err))
		u.Resp.ErrNo = 1
		u.Resp.ErrTips = err.Error()
		u.Resp.LoginStatus = 3 //  未知登陆
		return
	}
	// 汇报登陆状态
	u.Resp.LoginStatus = userLoginMgr.LoginStatus
}

func (u *UserLoginHander) checkParam() error {
	return nil
}
