package ginlog

import (
	"fmt"
	"go-gin/logger"
	"go-gin/util"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func GinLogger(fpath, fname string) gin.HandlerFunc {
	fileName := path.Join(fpath, fname)
	logger.Debug("GinLogger filename:", fileName)
	// 写入文件
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
	// 实例化
	logger := logrus.New()
	// 设置输出
	logger.Out = f
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新的日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最长保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割间隔时间
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfhook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: util.TimeLayout,
	})
	// 新增hook
	logger.AddHook(lfhook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
