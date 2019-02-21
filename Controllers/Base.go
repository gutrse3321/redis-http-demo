package Controllers

import (
	"fmt"
	"github.com/kataras/iris/context"
)

const (
	REDIS_KEY_USER = "AIMY_BIGDATA_USER"
	REDIS_KEY_TOKEN = "AIMY_BIGDATA_TOKEN"
)

type Json context.Map

func SendJson(code int, data interface{}) (result *Json) {
	result = &Json{"code": code, "data": data}
	return
}

func FormatRedisString(name string, cutName interface{}) (result string) {
	switch cutName.(type) {
	case int:
		result = fmt.Sprintf("%s:%d", name, cutName)
	case string:
		result = fmt.Sprintf("%s:%s", name, cutName)
	}
	return
}
