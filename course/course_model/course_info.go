package course_model

import (
	"System/db"
)

// 课程的基础信息 展示用
type CourseBaseInfo struct {
	CourseId        string
	TeacherName     string //老师名
	ClassName       string //课程名称
	ClassStudentNum int    //目前报名学生人数
	ClassCapacity   int    //课程最多报名人数
	CourseStatus    int    //课程状态 未开课（可报名），开课，结课
}

type CourseDetailInfo struct {
	CourseId        string
	TeacherUid      int64  // 老师uid
	TeacherUuid     string // 老师工号
	CourseAddr      string // 课程上课地址
	CourseDate      string // 开课日期
	CourseIntroduce string // 课程简介
	CourseMajor     string // 所属专业
	CourseStatus    int    // 课程状态
	ManagerUid      int64  // 审核课程的管理员UID
}

// 课程审核表
type CourseAuditInfo struct {
	CourseId        string
	TeacherUid      int64  // 老师UID
	ClassName       string // 课程名
	CourseDate      string // 开课日期
	ClassCapacity   int    // 课程容量
	CourseMajor     string // 课程专业
	CourseIntroduce string // 课程介绍
	ManagerUid      int64  // 管理员UID
}

type CourseRefuseInfo struct {
	CourseId        string
	TeacherUid      int64  // 老师UID
	ClassName       string // 课程名
	CourseDate      string // 开课日期
	ClassCapacity   int    // 课程容量
	CourseMajor     string // 课程专业
	CourseIntroduce string // 课程介绍
	ManagerUid      int64  // 管理员UID
	RefuseReason    string // 被拒绝的理由
}

func (CourseRefuseInfo) TableName() string {
	return "course_refuse_info"
}

func (CourseAuditInfo) TableName() string {
	return "course_audit_infos"
}

// 学生选课信息表
type StudentSelectCourseInfo struct {
	CourseId string `gorm:"primaryKey;autoIncrement:false"`
	Uid      int64  `gorm:"primaryKey;autoIncrement:false"` //学生UID
	Status   int    //选课的状态 审核中 已通过 (没有未通过，未通过直接从这个表中删除数据)
}

type StudentLikeCourseInfo struct {
	CourseId string `gorm:"primaryKey;autoIncrement:false"`
	Uid      int64  `gorm:"primaryKey;autoIncrement:false"` //学生UID
}

func (CourseBaseInfo) TableName() string {
	return "course_base_info"
}
func (CourseDetailInfo) TableName() string {
	return "course_detail_info"
}

func (StudentSelectCourseInfo) TableName() string {
	return "student_select_course_infos"
}

func (StudentLikeCourseInfo) TableName() string {
	return "student_like_course_infos"
}

func AddTableInDb() {
	db.Db.AutoMigrate(&CourseBaseInfo{})
	db.Db.AutoMigrate(&CourseDetailInfo{})
	db.Db.AutoMigrate(&StudentSelectCourseInfo{})
	db.Db.AutoMigrate(&StudentLikeCourseInfo{})
	db.Db.AutoMigrate(&CourseAuditInfo{})
	db.Db.AutoMigrate(&CourseRefuseInfo{})
}
