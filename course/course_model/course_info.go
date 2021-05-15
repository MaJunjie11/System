package course_model

import (
	"System/db"
)

// 课程的基础信息 展示用
type CourseBaseInfo struct {
	CourseId        int    `gorm:"primary_key;auto_increment"`
	TeacherName     string //老师名
	ClassName       string //课程名称
	ClassStudentNum int    //目前报名学生人数
	ClassCapacity   int    //课程最多报名人数
}

type CourseDetailInfo struct {
	CourseId     int    `gorm:"primary_key;auto_increment"`
	TeacherUid   int64  //老师uid
	TeacherUUid  string //老师工号
	CourseStatus int    //课程状态 未开课（可报名），开课，结课
	CourseAddr   string //课程上课地址
}

// 学生选课信息表
type StudentSelectCourseInfo struct {
	CourseId int   //课程ID
	Uid      int64 //学生UID
}

func (CourseBaseInfo) TableName() string {
	return "course_base_info"
}
func (CourseDetailInfo) TableName() string {
	return "course_detail_info"
}

func AddTableInDb() {
	db.Db.AutoMigrate(&CourseBaseInfo{})
	db.Db.AutoMigrate(&CourseDetailInfo{})
}
