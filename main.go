package main

import (
	_ "go-gin/app/task"
	"go-gin/config"
	_ "go-gin/config"
	_ "go-gin/dbutil"
	"go-gin/logger"
	"go-gin/middleware/ginlog"
	"go-gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	defer logger.Sync()
	logger.Level = logger.DEBUG
	logger.InitLogger("./log/debug/debug.log", 0 /*size*/, true /*day cut*/, 7 /*day age*/)

	gin.SetMode(config.HttpConf.RunMode)
	r := gin.Default()
	r.Use(ginlog.GinLogger("./log/ginlog/", "gin.log"))
	routers.InitRouters(r)
	logger.Info("start server: localhost", config.HttpConf.HttpPort)
	r.Run(config.HttpConf.HttpPort) // 监听,并启动服务
}
