package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"login-demo/Cache"
	"login-demo/Routes"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(logger.New())
	app.Configure(Routes.Configure)

	issue := Cache.Instance().Init()
	if issue != nil {
		app.Logger().Error(issue)
	}

	app.Run(
		iris.Addr(":3000"),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}