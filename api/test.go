package api

import (
	"encoding/json"
	"go-gin/logger"

	"github.com/gin-gonic/gin"
)

// func TestApi(c *gin.Context) (int, interface{}) {
// 	pm, err := hex.DecodeString(param)
// 	if err != nil {
// 		logger.Error("err:", err)
// 		return INVALID_PARAMS, nil
// 	}

// 	var ureq models.UserRequest
// 	err = json.Unmarshal(pm, &ureq)
// 	if err != nil {
// 		logger.Error("err:", err)
// 		return INVALID_PARAMS, nil
// 	}

// 	email := ureq.Email
// 	state := ureq.Flag

// 	if !(state == 0 || state == 1 || state == 3) {
// 		logger.Error("err:", fmt.Errorf("state[%v] not wright!", state))
// 		return INVALID_PARAMS, nil
// 	}
// 	logger.Debug("email:", email, "state:", state)
// 	var user models.UserBase
// 	ok, err := dbutil.SqlEngine.Where("email = ?", email).Get(&user)
// 	if !ok || err != nil {
// 		logger.Error("err:", fmt.Errorf("email[%v] not exist or sql error!", email))
// 		return INVALID_PARAMS, nil
// 	}

// 	_, err = dbutil.SqlEngine.Where("email = ?", email).Cols("del_flag").Update(models.UserBase{DelFlag: state})
// 	if err != nil {
// 		logger.Error("err:", err)
// 		return INVALID_PARAMS, nil
// 	}
// 	return SUCCESS, nil
// }

func TestApi(c *gin.Context) {
	logger.Debug("Into TestApi=======")
	var method string
	if method = c.Query("method"); method == "" {
		logger.Errorf("request method[%v] empty!", method)
		Response(c, INVALID_PARAMS, nil)
		return
	}

	var param string
	if param = c.Query("param"); param == "" {
		logger.Errorf("request param [%v] empty!", param)
		Response(c, INVALID_PARAMS, nil)
		return
	}

	data, _ := c.GetRawData()
	var mp map[string]interface{}
	json.Unmarshal(data, &mp)

	Response(c, SUCCESS, mp)
	return
}

func TestApi2(c *gin.Context) {
	logger.Debug("Into TestApi 2 =======")
	data, _ := c.GetRawData()
	var mp map[string]interface{}
	json.Unmarshal(data, &mp)

	Response(c, SUCCESS, mp)
	return
}
