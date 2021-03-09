package user_method

import (
	"OpenExperimentalManagementSystem/pb_gen"
	"OpenExperimentalManagementSystem/user/user_logic"
	"fmt"
	"log"
)

type AddUserHander struct {
	Req  *pb_gen.AddUserRequest
	Resp *pb_gen.AddUserResponse
}

func (u *AddUserHander) Run() {
	// checkParam
	var err error
	if err = u.checkParam(); err != nil {
		log.Fatal(fmt.Sprintf("err in UserLoginHander checkParam:%v", err))
		return
	}

	userName := u.Req.GetUserName()
	passworld := u.Req.GetPassworld()
	sex := u.Req.GetSex()
	phoneNum := u.Req.GetTelephoneNum()
	age := u.Req.GetAge()
	// 1. 到数据库查询是否已经有这个用户
	// 2. 如果不存在 那么将这个用户添加
	userLogicMgr := user_logic.NewUserLogicMgr(userName, passworld, phoneNum, int32(sex), age)
	if err = userLogicMgr.Process(); err != nil {
		return
	}
	u.Resp.Uid = userLogicMgr.Uid
}

func (u *AddUserHander) checkParam() error {
	return nil
}
