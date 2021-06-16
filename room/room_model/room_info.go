package room_model

import "System/db"

type RoomBaseInfo struct {
	Id           int `gorm:"primary_key;auto_increment"`
	RoomAddr     string
	RoomCapacity string
	Status       int // 表示教室状态 是否已经被占用 0 没有占用 1 已经占用 默认是0
}

func (RoomBaseInfo) TableName() string {
	return "room_base_info"
}

func AddTableInDb() {
	if db.Db.HasTable(&RoomBaseInfo{}) {
		db.Db.AutoMigrate(&RoomBaseInfo{})
	} else {
		db.Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&RoomBaseInfo{})
	}

}
