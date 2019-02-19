package Routes

import (
	"github.com/kataras/iris"
	"login-demo/Controllers"
)

func Configure(app *iris.Application) {
	app.Post("/login", Controllers.Login)
}