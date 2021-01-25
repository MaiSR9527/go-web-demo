package routers

import (
	"github.com/astaxie/beego"
	"go-project-open-source/simple_beego_project/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
