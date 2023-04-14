package userinfo

import (
	"fmt"
	"go-gin/app/errmsg"
	ginresponse "go-gin/app/response"
	"go-gin/dbutil"
	"go-gin/logger"
	"go-gin/models"
	"net/http"

	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

func DealBlacklist(c *gin.Context) {
	logger.PrintHttpRequest(c)
	resG := ginresponse.Gin{C: c}
	email := c.Query("email")
	flag := c.Query("flag")
	state := com.StrTo(flag).MustInt()
	if !(state == 0 || state == 1 || state == 3) {
		c.Error(fmt.Errorf("state[%v] not wright!", state))
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	}

	logger.Info("email:", email, "flag:", flag, "state:", state)
	var user models.UserBase
	ok, err := dbutil.SqlEngine.Where("email = ?", email).Get(&user)
	if !ok || err != nil {
		logger.Info("err1:", err)
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	}

	_, err = dbutil.SqlEngine.Where("email = ?", email).Cols("del_flag").Update(models.UserBase{DelFlag: state})
	if err != nil {
		logger.Info("err2:", err)
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	}
	resG.Response(http.StatusOK, errmsg.SUCCESS, nil)
	return
}
