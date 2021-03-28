package main

import (
	"System/server"
	"System/user/user_moudle"

	_ "System/db"

	"github.com/gin-gonic/gin"
	log "github.com/skoo87/log4go"
)

func main() {
	initLog()
	r := gin.Default()
	server.SetRoute(r)
	user_moudle.AddTableInDb()
	r.Run(":9090") // 监听并在 0.0.0.0:9090 上启动服务
}

func initLog() {
	if err := log.SetupLogWithConf("../conf/log.json"); err != nil {
		panic(err)
	}
	log.Info("启动成功")
}
