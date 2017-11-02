package app

import (
	"github.com/kataras/iris"
	"github.com/xuebing1110/promext/server/router/v1"
)

var (
	iris_app *iris.Application
)

func init() {
	iris_app = iris.New()
}

func GetIrisAPP() *iris.Application {
	//router
	v1.Init(iris_app)
	return iris_app
}
