package task

import (
	"go-gin/config"
	"go-gin/logger"
	"time"
)

func init() {
	go func() {
		for {
			time.Sleep(time.Second)
			logger.Infof("runmode=%v", config.HttpConf.RunMode)
		}
	}()
	//go task()
}

func task() {
}
