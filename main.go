package main

import (
	"System/server"

	"github.com/gin-gonic/gin"
	log "github.com/skoo87/log4go"
)

func main() {
	initLog()
	r := gin.Default()
	server.SetRoute(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func initLog() {
	if err := log.SetupLogWithConf("../conf/log.json"); err != nil {
		panic(err)
	}
	log.Info("启动成功")
}
