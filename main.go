package main

import (
	"go-gin/config"
	"go-gin/logger"
	"go-gin/routers"
	_ "go-gin/task"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	logger.InitLogger(config.LogConf)
	defer logger.Sync()

	gin.SetMode(config.HttpConf.RunMode)
	r := gin.Default()
	//r.Use(ginlog.GinLogger(config.LogConf.GinPath, config.LogConf.GinName))
	routers.InitRouters(r)

	logger.Info("start server: localhost", config.HttpConf.HttpPort)
	r.Run(config.HttpConf.HttpPort)
}
