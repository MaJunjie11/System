package server

import (
	"System/login/login_method"
	"System/pb_gen"
	"System/room/room_method"
	"System/user/user_method"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin") //请求头部
		if len(origin) == 0 {
			origin = c.Request.Header.Get("Origin")
		}
		//接收客户端发送的origin （重要！）
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//允许客户端传递校验信息比如 cookie (重要)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//服务器支持的所有跨域请求的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 设置预验请求有效期为 86400 秒
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
func SetRoute(r *gin.Engine) {
	r.Use(Cors())
	r.GET("/user_login", warpUserLogin)
	r.GET("/user_add", warpAddUser)
	r.GET("/create_room", warpCreateRoom)
	r.GET("/manager_get_reviewing_room", warpManagerGetReeviewingRoom)
	r.GET("/manager_check_reviewing_room", warpMangerCheckReviewingRoom)
}

func warpAddStudentRoRoom(c *gin.Context) {
	req := &pb_gen.AddStudentToRoomRequest{}
	resp := &pb_gen.AddStudentToRoomResponse{}
	c.ShouldBindQuery(&req) // 按照json格式解析出来

	c.JSON(200, resp)
}

func warpMangerCheckReviewingRoom(c *gin.Context) {
	req := &pb_gen.CreateRoomRequest{}
	resp := &pb_gen.CreateRoomResponse{}
	c.ShouldBindQuery(&req) // 按照json格式解析出来
	createRoomHandler := &room_method.CreateRoomHandler{
		Req:  req,
		Resp: resp,
	}
	createRoomHandler.Run()
	c.JSON(200, resp)

}

func warpCreateRoom(c *gin.Context) {
	req := &pb_gen.CreateRoomRequest{}
	resp := &pb_gen.CreateRoomResponse{}
	c.ShouldBindQuery(&req) // 按照json格式解析出来
	createRoomHandler := &room_method.CreateRoomHandler{
		Req:  req,
		Resp: resp,
	}
	createRoomHandler.Run()
	c.JSON(200, resp)
}

func warpManagerGetReeviewingRoom(c *gin.Context) {
	req := &pb_gen.ManagerGetReviewingRoomRequest{}
	resp := &pb_gen.ManagerGetReviewingRoomResponse{}
	c.ShouldBindUri(&req)
	managerGetReviewingRoomHandler := room_method.ManagerGetReviewingRoomHandler{
		Req:  req,
		Resp: resp,
	}
	managerGetReviewingRoomHandler.Run()
	c.JSON(200, resp)

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

	phone := c.Query("phone")
	passward := c.Query("passward")
	req.Phone = phone
	req.Passward = passward
	UserLoginHandler := login_method.UserLoginHander{
		Req:  &req,
		Resp: &resp,
	}
	UserLoginHandler.Run()
	c.Handler()
	c.JSON(200, gin.H{
		"loginStatus": resp.LoginStatus,
		"errNo":       resp.ErrNo,
		"errTips":     resp.ErrTips,
	})
}
