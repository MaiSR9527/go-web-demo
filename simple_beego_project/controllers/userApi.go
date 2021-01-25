package controllers

import "github.com/astaxie/beego"

type RegisterController struct {
	beego.Controller
}

// 登录页
func (this *RegisterController) ShowRegisterPage() {
	this.TplName = "index.tpl"
}

// 处理注册业务
func (this *RegisterController) HandleRegister() {
	// 1.拿到浏览器传来的数据

	// 2.数据处理

	// 3.插入数据

	// 4.返回视图
}
