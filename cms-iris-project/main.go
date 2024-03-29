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
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	// 错误配置
	app.OnErrorCode(iris.StatusNotFound, func(c context.Context) {
		c.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    "page not found",
			"data":   iris.Map{},
		})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    "server interal error ",
			"data":   iris.Map{},
		})
	})

}
