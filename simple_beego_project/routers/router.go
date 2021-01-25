package routers

import (
	"github.com/astaxie/beego"
	"simple_beego_project/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegisterPage;post:HandleRegister")
}
