package server

import (
	"System/course/course_method"
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
	r.GET("/student_get_untaught_course", warpStudentGetUntaughtCourse)
	// 3.学生删除课程
	r.POST("/student_delete_room", warpStudentDeleteRoom)
	// 4.获取用户信息
	r.GET("/get_user_info", warpGetUserInfo)
	// 5.用户修改密码
	r.PUT("/user_change_password", warpUserChangePassword)
	// 6.获取用户主页展示的可以添加的课程
	r.GET("/student_get_add_course_list", warpStudentGetAddCourseList)
	// 7. 学生加入课程
	r.POST("/student_add_course", warpStudentAddCourse)
	// 8.学生不可以加入的课程列表
	r.GET("/student_get_non_join_course_list", warpStudentGetNonJoinCourse)
	// 9.学生添加心愿单
	r.POST("/student_add_like_course", warpStudentAddLikeCourse)
	// 10.学生心愿课程展示
	r.GET("/student_get_like_course", warpStudentGetLikeCourse)
	// 11.学生获取已经开课课程
	r.GET("/student_get_start_course", warpStudentGetStartCourse)
	// 12.学生获取结束的课程
	r.GET("/student_get_end_course", warpStudentGetEndCourse)
	// 13.学生退出一个未开课课程
	r.POST("/student_feet_course", warpStudentFeetCourse)

	/*
		以下是老师端接口
	*/

	// 1.老师添加课程
	r.POST("/teacher_add_course", warpTeacherAddCourse)
	// 2 老师获取审核中课程
	r.GET("/teacher_get_audit_course", warpTeacherGetAuditCourse)
	// 3.获取审核中课程详细信息
	r.GET("/audit_course_info", warpTeacherGetAuditCourseDetailInfo)
	// 4.老师获取自己的所有课程
	r.GET("/teacher_get_untaught_course", warpTeacherGetUntaughtCourse)
	// 5.老师获取某个课程的所有学生信息
	r.POST("/get_course_student", warpGetCourseStudent)
	// 6. 获取当前课程等待审核的学生列表
	r.POST("/get_audit_course_student", warpGetAuditCourseStudent)
	// 7.老师同意学生申请加入课程
	r.POST("/teacher_audit_student", warpTeacherAuditStudent)
	// 8.老师开课
	r.POST("/teacher_start_course", warpTeacherStartCourse)
	// 9.老师结课
	r.POST("/teacher_end_course", warpTeacherEndCourse)
	// 10.老师添加学生
	r.POST("/teacher_add_student", warpTeacherAddStudent)
	// 11.老师获取已经开课课程
	r.GET("/teacher_get_start_course", warpTeacherGetStartCourse)
	// 12.老师获取已经结课课
	r.GET("/teacher_get_end_course", warpTeacherGetEndCourse)
	// 13.老师踢出学生 或拒绝学生加入
	r.POST("teacher_feet_student", warpTeacherFeetStudent)
	// 14.老师获取管理员审核不同过的课程
	r.GET("/teacher_get_refuse_course", warpTeacherGetRefuseCourse)
	// 15.老师删除没通过的课程
	r.POST("/teacher_delete_refuse_course", warpTeacherDeleteRefuseCourse)

	/*
		管理员功能
	*/
	// 1.管理员拉取等待审核课程列表
	r.GET("/manager_get_audit_course", warpManagerGetAuditCourse)
	// 2.
	r.POST("/manager_audit_down_course", warpManagerAuditDownCourse)
	// 3.
	r.POST("/manager_audit_refuse_course", warpManagerAuditRefuseCourse)
	// 4.管理员重置学生老师的密码
	r.POST("manager_reset_password", warpManagerResetPassword)
	// 5.管理员开启选课通道
	r.GET("/manager_open_select_course", warpManagerOpenSelectCourse)
	r.GET("/manager_close_select_course", warpManagerCloseSelectCourse)

	//TODO: 教室功能
	r.POST("/manager_add_room", warpManagerAddRoom)

	r.GET("/manager_get_use_room", warpManagerGetUseRoom)

}

func warpManagerCloseSelectCourse(c *gin.Context) {
	req := &pb_gen.ManagerCloseSelectCourseRequest{}
	resp := &pb_gen.ManagerCloseSelectCourseResponse{}
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewManagerCloseSelectCourseHandler(req, resp)
	handler.Run()

	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpManagerOpenSelectCourse(c *gin.Context) {
	req := &pb_gen.ManagerOpenSelectCourseRequest{}
	resp := &pb_gen.ManagerOpenSelectCourseResponse{}
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewManagerOpenSelectCourseHandler(req, resp)
	handler.Run()

	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpManagerGetUseRoom(c *gin.Context) {
	req := &pb_gen.ManagerGetUseRoomRequest{}
	resp := &pb_gen.ManagerGetUseRoomResponse{}

	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := room_method.NewManagerGetUseRoomHandler(req, resp)
	handler.Run()

	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpManagerAddRoom(c *gin.Context) {
	req := &pb_gen.ManagerAddRoomRequest{}
	resp := &pb_gen.ManagerAddRoomResponse{}

	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := room_method.NewManagerAddRoomHandelr(req, resp)
	handler.Run()

	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpManagerResetPassword(c *gin.Context) {
	req := &pb_gen.ManagerResetPasswordRequest{}
	resp := &pb_gen.ManagerResetPasswordResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewManagerResetPasswordHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpManagerAuditRefuseCourse(c *gin.Context) {
	req := &pb_gen.ManagerAuditRefuseCourseRequest{}
	resp := &pb_gen.ManagerAuditRefuseCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewManagerAuditRefuseCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpStudentFeetCourse(c *gin.Context) {
	req := &pb_gen.StudentFeetCourseRequest{}
	resp := &pb_gen.StudentFeetCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewStudentFeetCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherFeetStudent(c *gin.Context) {
	req := &pb_gen.TeacherFeetStudentRequest{}
	resp := &pb_gen.TeacherFeetStudentResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherFeetStudentHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpManagerAuditDownCourse(c *gin.Context) {
	req := &pb_gen.ManagerAuditDoneCourseRequest{}
	resp := &pb_gen.ManagerAuditDoneCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewManagerAuditDoneCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpManagerGetAuditCourse(c *gin.Context) {
	req := &pb_gen.ManagerGetAuditCourseRequest{}
	resp := &pb_gen.ManagerGetAuditCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewManagerGetAuditCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpTeacherAddStudent(c *gin.Context) {
	req := &pb_gen.TeacherAddStudentRequest{}
	resp := &pb_gen.TeacherAddStudentResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherAddStudentHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherEndCourse(c *gin.Context) {
	req := &pb_gen.TeacherEndCourseRequest{}
	resp := &pb_gen.TeacherEndCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherEndCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherStartCourse(c *gin.Context) {
	req := &pb_gen.TeacherStartCourseRequest{}
	resp := &pb_gen.TeacherStartCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherStartCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherAuditStudent(c *gin.Context) {
	req := &pb_gen.TeacherAuditStudentRequest{}
	resp := &pb_gen.TeacherAuditStudentResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherAuditStudentHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpGetAuditCourseStudent(c *gin.Context) {
	req := &pb_gen.GetAuditCourseStudentRequest{}
	resp := &pb_gen.GetAuditCourseStudentResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewGetAuditCourseStudentHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpGetCourseStudent(c *gin.Context) {
	req := &pb_gen.GetCourseStudentRequest{}
	resp := &pb_gen.GetCourseStudentResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewGetCourseStudentHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherGetUntaughtCourse(c *gin.Context) {
	req := &pb_gen.TeacherGetUntaughtCourseRequest{}
	resp := &pb_gen.TeacherGetUntaughtCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherGetUntaughtCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherGetStartCourse(c *gin.Context) {
	req := &pb_gen.TeacherGetStartCourseRequest{}
	resp := &pb_gen.TeacherGetStartCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherGetStartCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherGetEndCourse(c *gin.Context) {
	req := &pb_gen.TeacherGetEndCourseRequest{}
	resp := &pb_gen.TeacherGetEndCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherGetEndCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherGetAuditCourseDetailInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"className":   "haha",
		"courseMajro": "物理",
	})
}

func warpTeacherDeleteRefuseCourse(c *gin.Context) {
	req := &pb_gen.TeacherDeleteRefuseCourseRequest{}
	resp := &pb_gen.TeacherDeleteRefuseCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherDeleteRefuseCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherGetRefuseCourse(c *gin.Context) {
	req := &pb_gen.TeacherGetRefuseCourseRequest{}
	resp := &pb_gen.TeacherGetRefuseCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherGetRefuseCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherGetAuditCourse(c *gin.Context) {
	req := &pb_gen.TeacherGetAuditCourseRequest{}
	resp := &pb_gen.TeacherGetAuditCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherGetAuditCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpTeacherAddCourse(c *gin.Context) {
	req := &pb_gen.TeacherAddCourseRequest{}
	resp := &pb_gen.TeacherAddCourseResponse{}
	// 接受请求
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewTeacherAddCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpStudentGetLikeCourse(c *gin.Context) {
	req := &pb_gen.StudentGetLikeCourseRequest{}
	resp := &pb_gen.StudentGetLikeCourseResponse{}
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewStudentGetLikeCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpStudentAddLikeCourse(c *gin.Context) {
	req := &pb_gen.StudentAddLikeCourseRequest{}
	resp := &pb_gen.StudentAddLikeCourseResponse{}
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token

	handler := course_method.NewStudentAddLikeCouserHandler(req, resp)
	handler.Run()

	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpStudentGetNonJoinCourse(c *gin.Context) {
	req := &pb_gen.StudentGetNonJoinCourseRequset{}
	resp := &pb_gen.StudentGetNonJoinCourseResponse{}
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewStudentGetNonJoinCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpStudentAddCourse(c *gin.Context) {
	req := &pb_gen.StudentAddCourseRequest{}
	resp := &pb_gen.StudentAddCourseResponse{}
	token := c.Request.Header.Get("token")
	req.Token = token
	c.ShouldBindJSON(req)

	handler := course_method.NewStudentAddCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
func warpStudentGetEndCourse(c *gin.Context) {
	req := &pb_gen.StudentGetEndCourseRequest{}
	resp := &pb_gen.StudentGetEndCourseResponse{}
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewStudentGetEndCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpStudentGetStartCourse(c *gin.Context) {
	req := &pb_gen.StudentGetStartCourseRequest{}
	resp := &pb_gen.StudentGetStartCourseResponse{}
	c.ShouldBindJSON(req)
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewStudentGetStartCourseHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

func warpStudentGetUntaughtCourse(c *gin.Context) {
	req := &pb_gen.StudentGetUntaughtCourseRequest{}
	resp := &pb_gen.StudentGetUntaughtCourseResponse{}
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewStudentGetAlreadyRoomHandler(req, resp)
	handler.Run()
	if resp.Code != pb_gen.ErrNo_Success {
		c.JSON(500, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpStudentGetAddCourseList(c *gin.Context) {
	req := &pb_gen.StudentGetAddCourseListRequest{}
	resp := &pb_gen.StudentGetAddCourseListResponse{}
	token := c.Request.Header.Get("token")
	req.Token = token
	handler := course_method.NewGetStudentAddCourseListHandler(req, resp)
	handler.Run()

	if resp.Code != pb_gen.ErrNo_Success {
		fmt.Println("resp.Code:", resp.Code)
		c.JSON(404, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func warpUserChangePassword(c *gin.Context) {
	req := &pb_gen.UserChangePasswordRequest{}
	resp := &pb_gen.UserChangePasswordResponse{}
	c.ShouldBindJSON(req)
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
