package main

import (
	"System/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	server.SetRoute(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
