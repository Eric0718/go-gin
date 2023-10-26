package task

import (
	"go-gin/logger"
	"go-gin/util"
	"time"
)

func init() {
	go testlog()
}

func testlog() {
	for {
		logger.Info("test log time:", time.Now().Format(util.TimeLayout))
		time.Sleep(time.Second)
	}
}
