package server

import (
	"System/login/login_method"
	"System/pb_gen"
	"System/user/user_method"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine) {
	r.GET("/user_login", warpUserLogin)
	r.GET("/user_add", warpAddUser)
}

func warpAddUser(c *gin.Context) {
	req := pb_gen.AddUserRequest{}
	resp := pb_gen.AddUserResponse{}

	sexStr := c.Query("sex")
	userName := c.Query("userName")
	passWord := c.Query("passWord")
	telephoneNum := c.Query("telephoneNum")
	ageStr := c.Query("age")

	sex, err := strconv.ParseInt(sexStr, 10, 64)
	age, err := strconv.ParseInt(ageStr, 10, 64)

	if err != nil {

	}
	req.UserName = userName
	req.Passworld = passWord
	req.Sex = pb_gen.Sex(sex)
	req.TelephoneNum = telephoneNum
	req.Age = int32(age)
	addUserHandler := user_method.AddUserHander{
		Req:  &req,
		Resp: &resp,
	}
	addUserHandler.Run()
	c.JSON(200, gin.H{
		"errNo":   resp.ErrNo,
		"errTips": resp.ErrTips,
		"uid":     resp.Uid,
	})
}

func warpUserLogin(c *gin.Context) {
	req := pb_gen.UserLoginRequest{}
	resp := pb_gen.UserLoginResponse{}

	uidStr := c.Query("uid")
	passward := c.Query("passward")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
	}

	req.Uid = uid
	req.Passward = passward
	UserLoginHandler := login_method.UserLoginHander{
		Req:  &req,
		Resp: &resp,
	}
	UserLoginHandler.Run()
	c.JSON(200, gin.H{
		"loginStatus": resp.LoginStatus,
		"errNo":       resp.ErrNo,
		"errTips":     resp.ErrTips,
	})
}
