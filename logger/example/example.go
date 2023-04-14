package main

import (
	"go-gin/middleware/logger"
	"time"

	"github.com/fatih/color"
)

func main() {
	defer logger.Sync()
	a := logger.NewLog("", 0, false, int(10*time.Second))
	a.Info("foo", "aaaa", "bb")
	a.Info(color.New(color.BgYellow).Sprint("aaaa"), color.New(color.BgBlue).Sprint("bbbb"))
	logger.Level = logger.DEBUG
	test()
	logger.Info("bar")
	time.Sleep(10 * time.Second)

}

func test() {
	// 此方法的日志级别是DEBUG， 所以调试的时候必须将日志级别设置成DEBUG，不然不会显示
	logger.UpFunc(1, "who call me") // 2022-03-04 10:49:38 - [DEBUG] - DESKTOP-NENB5CA - C:/work/golog/example/example.go:16 - caller from C:/work/golog/example/example.go:11 -- who call me
}
