package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int    `orm:"pk;auto"`
	UserName string `orm:"unique;size(20)"`
	Password string `orm:size(200)`
}

func init() {
	//ORM操作数据库
	//获取连接对象
	orm.RegisterDataBase("default", "mysql", "root:mai1208142397@tcp(127.0.0.1:3306)/go_test?charset=utf8")

	//创建表
	orm.RegisterModel(new(User))

	//生成表
	//第一个参数是数据库别名，第二个参数是是否强制更新
	orm.RunSyncdb("default", false, true)
	//操作表

}
