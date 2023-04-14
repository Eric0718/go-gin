package routers

import "fmt"

type Handler func(param string) (error, interface{})

var handlers map[int]Handler
var idToFunc map[int]string

func init() {
	handlers = make(map[int]Handler)
	idToFunc = make(map[int]string)
	//前端接口注册
	registerFuncs()
}

//id start from 1 ,type int
func register(id int, method string, h Handler) error {
	if _, ok := handlers[id]; ok {
		return fmt.Errorf("handlers id[%v],method[%v] already registed!", id, method)
	}

	if _, ok := idToFunc[id]; ok {
		return fmt.Errorf("idToFunc id[%v],method[%v] already registed!", id, method)
	}

	handlers[id] = h
	idToFunc[id] = method
	return nil
}
