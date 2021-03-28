package user_moudle

import (
	"System/db"
	"time"
)

type UserStudent struct {
	Id         int
	Uid        int64
	Email      string
	Password   string
	Desc       string
	Status     int
	CreateTime time.Time
}

func (UserStudent) TableName() string {
	return "user_student"
}

func AddTableInDb() {
	db.Db.AutoMigrate(&UserStudent{})
}
