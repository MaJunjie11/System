package course_model

import (
	"System/db"
)

//
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

type CourseOpenInfo struct {
	Id     int `gorm:"primary_key;auto_increment"`
	Status int // 就一个标记位 0:关闭 1：开启
}

func (CourseOpenInfo) TableName() string {
	return "course_open_info"
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

	if db.Db.HasTable(&CourseBaseInfo{}) {
		db.Db.AutoMigrate(&CourseBaseInfo{})
	} else {
		db.Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&CourseBaseInfo{})
	}

	if db.Db.HasTable(&CourseDetailInfo{}) {
		db.Db.AutoMigrate(&CourseDetailInfo{})
	} else {
		db.Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&CourseDetailInfo{})
	}

	if db.Db.HasTable(&StudentSelectCourseInfo{}) {
		db.Db.AutoMigrate(&StudentSelectCourseInfo{})
	} else {
		db.Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&StudentSelectCourseInfo{})
	}

	if db.Db.HasTable(&StudentLikeCourseInfo{}) {
		db.Db.AutoMigrate(&StudentLikeCourseInfo{})
	} else {
		db.Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&StudentLikeCourseInfo{})
	}

	if db.Db.HasTable(&CourseAuditInfo{}) {
		db.Db.AutoMigrate(&CourseAuditInfo{})
	} else {
		db.Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&CourseAuditInfo{})
	}

	if db.Db.HasTable(&CourseRefuseInfo{}) {
		db.Db.AutoMigrate(&CourseRefuseInfo{})
	} else {
		db.Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&CourseRefuseInfo{})
	}

	if db.Db.HasTable(&CourseOpenInfo{}) {
		db.Db.AutoMigrate(&CourseOpenInfo{})
	} else {
		db.Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&CourseOpenInfo{})
	}

}
