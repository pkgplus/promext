package app

import (
	"github.com/kataras/iris"

	"github.com/xuebing1110/promext/pkg/log"
	"github.com/xuebing1110/promext/pkg/prometheus"
	"github.com/xuebing1110/promext/server/router/v1"
)

var (
	iris_app *iris.Application
)

func init() {
	iris_app = iris.New()

	// logger
	var logger log.Logger = iris_app.Logger()
	prometheus.SetLogger(logger)
}

func GetIrisAPP() *iris.Application {
	//router
	v1.Init(iris_app)
	return iris_app
}
