package main

import "github.com/kataras/iris"

func main() {
	//1.创建app结构体对象
	app := iris.New()
	//2.端口监听
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

	////application.Run(iris.Addr(":8080"))//第一种
	//application.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed)) //第二种
}
