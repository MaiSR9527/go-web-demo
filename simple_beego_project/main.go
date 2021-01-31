package main

import (
	"github.com/astaxie/beego"
	_ "simple_beego_project/models"
	_ "simple_beego_project/routers"
)

func main() {
	beego.Run()
}
