package routers

import (
	"go-gin/app/user"
	"log"
)

//所有模块前端调用接口这里添加注册,id 从1开始递增，且不能重复，method 参数格式: req_+函数名(第一个字母小写)
func registerFuncs() {
	var err error
	err = register(1, "req_blacklist", user.Blacklist)
	if err != nil {
		log.Fatal(err)
	}
}
