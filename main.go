package main

import (
	"System/course/course_model"
	"System/server"
	"System/user/user_moudle"

	_ "System/db"

	"github.com/gin-gonic/gin"
	log "github.com/skoo87/log4go"
)

func main() {
	initLog()
	//initDbTable()
	r := gin.Default()
	server.SetRoute(r)
	r.Run(":9090") // 监听并在 0.0.0.0:9090 上启动服务
}

func initLog() {
	if err := log.SetupLogWithConf("../conf/log.json"); err != nil {
		panic(err)
	}
	log.Info("启动成功")
}

func initDbTable() {
	user_moudle.AddTableInDb()
	course_model.AddTableInDb()
}
