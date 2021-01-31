package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type RegisterController struct {
	beego.Controller
}

// 登录页
func (this *RegisterController) ShowRegisterPage() {
	this.TplName = "register.html"
}

// 处理注册业务
func (this *RegisterController) HandleRegister() {
	// 1.拿到浏览器传来的数据
	username := this.GetString("username")
	password := this.GetString("password")

	// 2.数据处理:只做简单校验
	if username == "" || password == "" {
		logs.Info("用户名或密码不能为空")
		this.TplName = "register.html"
		return
	}
	// 3.插入数据

	// 4.返回视图
}
