package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {

}

// 构建iris Application
func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.StaticWeb("/static", "./static")
	app.StaticWeb("/manage/static", "./static")
	//注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})

	return app
}

// 初始化项目配置
func InitConfig(app *iris.Application) {

}
