package server

import (
	"System/login/login_method"
	"System/message/message_method"
	"System/pb_gen"
	"System/room/room_method"
	"System/user/user_method"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skoo87/log4go"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin") //请求头部
		if len(origin) == 0 {
			origin = c.Request.Header.Get("Origin")
		}
		//接收客户端发送的origin （重要！）
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		//允许客户端传递校验信息比如 cookie (重要)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Action-Addr,x-token,*")
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
	r.POST("/amisdemo", warpAmisDemo)
	r.GET("/user_login", warpUserLogin)

	//r.POST("/user_add", warpAddUser)
	r.GET("/create_room", warpCreateRoom)
	r.GET("/manager_get_reviewing_room", warpManagerGetReeviewingRoom)
	r.GET("/manager_check_reviewing_room", warpMangerCheckReviewingRoom)
	r.GET("/add_student_to_room", warpAddStudentRoRoom)
	r.GET("/teacher_get_can_reviewingy_roomlist", warpTeacherGetCanReviewingRoomList)
	r.GET("/teacher_check_student_to_room", warpTeacherCheckStudentToRoom)
	r.GET("/teacher_add_student_to_room", warpTeacherAddStudentToRoom)
	r.GET("/teacher_get_req_add_room_studentlist", warpTeacherGetReqAddRoomStudentList)
	r.POST("/send_email", warpSendEmail)

	// 1.用户登陆
	r.POST("/student_login", warpStudentLogin)
	// 2.学生拉教室列表
	r.GET("/student_get_already_room", warpStudentGetAlreadyRoom)
	// 3.学生删除课程
	r.POST("/student_delete_room", warpStudentDeleteRoom)
	// 4.获取用户信息
	r.GET("/get_user_info", warpGetUserInfo)
	// 5.用户修改密码
	r.PUT("/user_change_password", warpUserChangePassword)
}

func warpUserChangePassword(c *gin.Context) {
	req := &pb_gen.UserChangePasswordRequest{}
	resp := &pb_gen.UserChangePasswordResponse{}
	c.ShouldBindJSON(req)
	fmt.Println("passward:", req.OldPassword)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := user_method.NewUserChangePasswordHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		fmt.Println("resp.Code:", resp.Code)
		c.JSON(404, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpGetUserInfo(c *gin.Context) {
	req := &pb_gen.GetUserInfoRequest{}
	resp := &pb_gen.GetUserInfoResponse{}

	token := c.Request.Header.Get("token")
	req.Token = token
	handler := user_method.NewGetUserInfoHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		// 如果token是空的话
		fmt.Println("resp.Code:", resp.Code)
		c.JSON(404, resp)
	} else {
		c.JSON(http.StatusOK, resp)

	}
}

func warpStudentDeleteRoom(c *gin.Context) {
	type Req struct {
		RoomId int `json:"room_id"`
	}
	req := &Req{}

	c.ShouldBindJSON(req)
	c.JSON(http.StatusOK, "haha")
}

func warpStudentGetAlreadyRoom(c *gin.Context) {
	req := &pb_gen.StudentGetAlreadyRoomRequest{}
	resp := &pb_gen.StudentGetAlreadyRoomResponse{}
	c.ShouldBindJSON(req)
	handler := room_method.NewStudentGetAlreadyRoomHandler(req, resp)
	handler.Run()

	var data pb_gen.StudentGetAlreadyRoomResponseData
	data.Count = 3

	data1 := &pb_gen.RoomInfo{
		RoomId:      213890,
		TeacherId:   12390,
		StudentCnt:  30,
		RoomCapcity: 130,
		RoomName:    "MJJ",
	}
	data2 := &pb_gen.RoomInfo{
		RoomId:      213890,
		TeacherId:   12390,
		StudentCnt:  30,
		RoomCapcity: 130,
		RoomName:    "PJJ",
	}
	data3 := &pb_gen.RoomInfo{
		RoomId:      213890,
		TeacherId:   12390,
		StudentCnt:  30,
		RoomCapcity: 130,
		RoomName:    "PJJ",
	}

	var rows []*pb_gen.RoomInfo = make([]*pb_gen.RoomInfo, 0)
	rows = append(rows, data1)
	rows = append(rows, data2)
	rows = append(rows, data3)
	data.Rows = rows
	resp.Data = &data
	resp.Msg = "ok"
	resp.Status = pb_gen.ErrNo_Success
	if resp.Status != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpStudentLogin(c *gin.Context) {
	req := &pb_gen.StudentLoginRequest{}
	resp := &pb_gen.StudentLoginResponse{}
	c.ShouldBindJSON(req)
	handler := login_method.NewStudentLoginHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		fmt.Println("code:", resp.Code)
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpSendEmail(c *gin.Context) {
	req := &pb_gen.SendEmailRequest{}
	resp := &pb_gen.SendEmailResponse{}
	c.ShouldBindJSON(req)
	handler := message_method.NewSendEmailHandler(req, resp)
	handler.Run()
	c.JSON(http.StatusOK, resp)
}

func warpTeacherGetReqAddRoomStudentList(c *gin.Context) {
	req := &pb_gen.TeacherGetReqAddRoomStudentListRequest{}
	resp := &pb_gen.TeacherGetReqAddRoomStudentListResponse{}
	c.ShouldBindQuery(req)
	handler := room_method.TeacherGetReqAddRoomStudentListHandler{
		Req:  req,
		Resp: resp,
	}
	handler.Run()
	c.JSON(200, resp)

}

func warpTeacherAddStudentToRoom(c *gin.Context) {
	req := &pb_gen.AddStudentToRoomRequest{}
	resp := &pb_gen.AddStudentToRoomResponse{}
	c.ShouldBindQuery(req)
	handler := room_method.AddStudentToRoomHandler{
		Req:  req,
		Resp: resp,
	}
	handler.Run()
	c.JSON(200, resp)

}

func warpTeacherCheckStudentToRoom(c *gin.Context) {
	req := &pb_gen.TeacherCheckStudentToRoomRequest{}
	resp := &pb_gen.TeacherCheckStudentToRoomResponse{}
	c.ShouldBindQuery(req)
	handler := room_method.TeacherCheckStudentToRoomHandler{
		Req:  req,
		Resp: resp,
	}
	handler.Run()
	c.JSON(200, resp)
}

func warpTeacherGetCanReviewingRoomList(c *gin.Context) {
	req := &pb_gen.TeacherGetCanReviewingRoomListRequest{}
	resp := &pb_gen.TeacherGetCanReviewingRoomListResponse{}
	c.ShouldBindQuery(req)
	handler := room_method.TeacherGetCanReviewingRoomListHandler{
		Req:  req,
		Resp: resp,
	}
	handler.Run()
	c.JSON(200, resp)
}

func warpAddStudentRoRoom(c *gin.Context) {
	req := &pb_gen.AddStudentToRoomRequest{}
	resp := &pb_gen.AddStudentToRoomResponse{}
	c.ShouldBindQuery(&req) // 按照json格式解析出来
	handler := room_method.AddStudentToRoomHandler{
		Req:  req,
		Resp: resp,
	}
	handler.Run()
	c.JSON(200, resp)
}

func warpMangerCheckReviewingRoom(c *gin.Context) {
	req := &pb_gen.ManagerCheckReviewingRoomRequest{}
	resp := &pb_gen.ManagerCheckReviewingRoomResponse{}
	c.ShouldBindQuery(&req) // 按照json格式解析出来
	handler := room_method.ManagerCheckReviewingRoomHandler{
		Req:  req,
		Resp: resp,
	}
	handler.Run()
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

// 屏蔽注册逻辑

//func warpAddUser(c *gin.Context) {
//req := &pb_gen.AddStudentRequest{}
//resp := &pb_gen.AddStudentResponse{}
//c.ShouldBindJSON(req)
//handler := user_method.NewAddUserHandler(req, resp)
//handler.Run()
//if resp.GetStatus() != 1 {
//c.JSON(500, resp)
//} else {
//c.JSON(http.StatusOK, resp)
//}
//}

func warpAmisDemo(c *gin.Context) {
	name := c.PostForm("name")
	log4go.Info("name:%s", name)
	resp := pb_gen.UserLoginResponse{}
	resp.ErrNo = 1
	resp.ErrTips = "2"
	resp.LoginStatus = 3
	c.JSON(200, &resp)
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
