package routers

import (
	"go-gin/util/resmsg"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  getMsg(errCode),
		Data: data,
	})
	return
}

func getMsg(code int) string {
	msg, ok := resmsg.MsgFlags[code]
	if ok {
		return msg
	}

	return resmsg.MsgFlags[resmsg.ERROR]
}
