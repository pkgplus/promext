package main

import (
	"github.com/kataras/iris"
	"github.com/xuebing1110/promext/server/app"
)

func main() {
	irisApp := app.GetIrisAPP()
	irisApp.Run(iris.Addr(":8080"))
}
