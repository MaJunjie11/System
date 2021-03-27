package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB
var err error

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogoMode bool   `json:"logo_mode"`
}

func LoadMysqlConf() *MysqlConf {

	mysql_conf := MysqlConf{}

	// 可执行文件在 output 这里注意获取conf地址
	file, err := os.Open("../conf/mysql_conf.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	byte_data, err2 := ioutil.ReadAll(file)

	if err2 != nil {
		panic(err2)
	}

	err3 := json.Unmarshal(byte_data, &mysql_conf)

	if err3 != nil {
		panic(err3)
	}

	return &mysql_conf

}

func init() {

	mysql_conf := LoadMysqlConf()

	logo_mode := mysql_conf.LogoMode

	data_source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		mysql_conf.UserName,
		mysql_conf.Password,
		mysql_conf.Host,
		mysql_conf.Port,
		mysql_conf.DataBase,
	)

	Db, err = gorm.Open("mysql", data_source)

	if err != nil {
		panic(err)
	}

	Db.LogMode(logo_mode)

	Db.DB().SetMaxOpenConns(100) // 最大连接数
	Db.DB().SetMaxIdleConns(50)  // 最大空闲数

	// 这里设置db的表暂时不使用
	//Db.AutoMigrate(&models.FrontUser{}, &models.AdminUser{})
}
