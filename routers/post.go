package routers

import (
	"go-gin/logger"
	"go-gin/util/resmsg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func apiPost(c *gin.Context) {
	logger.PrintHttpRequest(c)
	resG := Gin{C: c}
	id := -1
	if Id := c.Query("id"); Id != "" {
		id = com.StrTo(Id).MustInt()
	} else {
		logger.Errorf("request id[%v] empty!\n", Id)
		resG.Response(http.StatusBadRequest, resmsg.INVALID_PARAMS, nil)
		return
	}
	var method string
	if method = c.Query("method"); method == "" {
		logger.Errorf("request method[%v] empty!\n", method)
		resG.Response(http.StatusBadRequest, resmsg.INVALID_PARAMS, nil)
		return
	}

	var param string
	if param = c.Query("param"); param == "" {
		logger.Errorf("request param [%v] empty!\n", param)
		resG.Response(http.StatusBadRequest, resmsg.INVALID_PARAMS, nil)
		return
	}

	handleFunc, ok := handlers[id]
	if !ok {
		logger.Errorf("request id[%v] not exist!\n", id)
		resG.Response(http.StatusBadRequest, resmsg.INVALID_PARAMS, nil)
		return
	}

	if idToFunc[id] != method {
		logger.Errorf("request method[%v] not registed!\n", method)
		resG.Response(http.StatusBadRequest, resmsg.INVALID_PARAMS, nil)
		return
	}
	resCode, resData := handleFunc(param)
	if resCode != resmsg.SUCCESS {
		logger.Error("resCode:", resCode)
		resG.Response(http.StatusBadRequest, resCode, nil)
		return
	}

	resG.Response(http.StatusOK, resCode, resData)
	return
}
