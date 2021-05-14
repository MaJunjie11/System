package user_moudle

import (
	"System/db"
)

type UserStudentLoginInfo struct {
	Id       int    `gorm:"primary_key;auto_increment"`
	Uid      int64  //系统给学生分配的uid也是token
	Uuid     string `gorm:"not null"` //学生的业务ID
	Password string //学生的密码 本系统没有注册，直接登陆，初始密码默认123456
	Limit    int    //鉴权信息
}

// 基础信息表
type UserStudentBaseInfo struct {
	Id        int `gorm:"primary_key;auto_increment"`
	Uid       int64
	Uuid      string //uuid 17060211111
	Name      string
	Signature string // 个性签名
	Avatar    string // 头像地址
	Limit     int    // 权限
}

type UserStudentDetailInfo struct {
	Id              int    `gorm:"primary_key;auto_increment"`
	Uid             int64  // 系统UID
	Avatar          string // 投降url
	Sex             string // 性别
	College         string // 所属学院
	Registration    string // 户籍
	SchoolingRecord string // 学历
}

// 学生选课表
type UserStudentClassInfo struct {
	Uid     int64
	ClassId int64
	//这两个需要做联合主键
}

func (UserStudentLoginInfo) TableName() string {
	return "user_student_login_info"
}

func (UserStudentBaseInfo) TableName() string {
	return "user_student_base_info"
}

func (UserStudentDetailInfo) TableName() string {
	return "user_student_detail_info"
}

func AddTableInDb() {
	db.Db.AutoMigrate(&UserStudentBaseInfo{})
	db.Db.AutoMigrate(&UserStudentLoginInfo{})
	db.Db.AutoMigrate(&UserStudentDetailInfo{})
}
