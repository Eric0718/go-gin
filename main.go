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
	logger.InitLogger(config.LogConf)

	gin.SetMode(config.HttpConf.RunMode)
	r := gin.Default()
	r.Use(ginlog.GinLogger(config.LogConf.GinPath, config.LogConf.GinName))
	routers.InitRouters(r)
	logger.Info("start server: localhost", config.HttpConf.HttpPort)
	r.Run(config.HttpConf.HttpPort)
}
