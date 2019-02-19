package Controllers

import "github.com/kataras/iris/context"

type Json context.Map

func SendJson(code int, data interface{}) (result *Json) {
	result = &Json{"code": code, "data": data}
	return
}

