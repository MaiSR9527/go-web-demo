package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {
	// 1. 创建数据库引擎对象
	engine, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	engine.SetMapper(core.GonicMapper{})
	engine.Sync2(new(PersonTable))

	personEmpty, err := engine.IsTableEmpty(new(PersonTable))
	if err != nil {
		panic(err.Error())
	}
	if personEmpty {
		fmt.Println(" 人员表是空的 ")
	} else {
		fmt.Println(" 人员表不为空 ")
	}

	//判断表结构是否存在
	studentExist, err := engine.IsTableExist(new(StudentTable))
	if err != nil {
		panic(err.Error())
	}
	if studentExist {
		fmt.Println("学生表存在")
	} else {
		fmt.Println("学生表不存在")
	}

	//二、条件查询
	//1.ID查询
	var person PersonTable
	// select * from person_table where id = 1
	engine.Id(1).Get(&person)
	fmt.Println(person.PersonName)
	fmt.Println()

	//2.where多条件查询
	var person1 PersonTable
	// select * from person_table where person_age = 26 and person_sex = 2
	engine.Where(" person_age = ? and person_sex = ?", 26, 2).Get(&person1)
	fmt.Println(person1.PersonName)
	fmt.Println()

	//3.And条件查询
	var persons []PersonTable
	//select * from person_table where person_age = 26 and person_sex = 2
	err = engine.Where(" person_age = ? ", 26).And("person_sex = ? ", 2).Find(&persons)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(persons)
	fmt.Println()

	//4、Or条件查询
	var personArr []PersonTable
	//select * from person_table where person_age = 26 or person_sex = 1
	err = engine.Where(" person_age = ? ", 26).Or("person_sex = ? ", 1).Find(&personArr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(personArr)
	fmt.Println()

	//5、原生SQL语句查询支持 like语法
	var personsNative []PersonTable
	err = engine.SQL(" select * from person_table where person_name like 't%' ").Find(&personsNative)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(personsNative)
	fmt.Println()

	//6、排序条件查询
	var personsOrderBy []PersonTable
	//select * from person_table orderby person_age  升序排列
	//engine.OrderBy(" person_age ").Find(&personsOrderBy)
	engine.OrderBy(" person_age desc ").Find(&personsOrderBy)
	fmt.Println(personsOrderBy)
	fmt.Println()

	//7、查询特定字段
	var personsCols []PersonTable
	engine.Cols("person_name", "person_age").Find(&personsCols)
	for _, col := range personsCols {
		fmt.Println(col)
	}

	//三、增加记录操作
	personInsert := PersonTable{
		PersonName: "Hello",
		PersonAge:  18,
		PersonSex:  1,
	}
	//rowNum, err := engine.Insert(&personInsert)
	//fmt.Println(rowNum) //rowNum 受影响的记录条数
	//fmt.Println()

	//四、删除操作
	rowNum, err := engine.Delete(&personInsert)
	fmt.Println(rowNum) //rowNum 受影响的记录条数
	fmt.Println()

	//五、更新操作
	rowNum, err = engine.Id(7).Update(&personInsert)
	fmt.Println(rowNum) //rowNum 受影响的记录条数
	fmt.Println()

	//六、统计功能count
	count, err := engine.Count(new(PersonTable))
	fmt.Println("PersonTable表总记录条数：", count)

	//七、事务操作
	personsArray := []PersonTable{
		PersonTable{
			PersonName: "Jack",
			PersonAge:  28,
			PersonSex:  1,
		},
		PersonTable{
			PersonName: "Mali",
			PersonAge:  28,
			PersonSex:  1,
		},
		PersonTable{
			PersonName: "Ruby",
			PersonAge:  28,
			PersonSex:  1,
		},
	}

	session := engine.NewSession()
	session.Begin()
	for i := 0; i < len(personsArray); i++ {
		_, err = session.Insert(personsArray[i])
		if err != nil {
			session.Rollback()
			session.Close()
		}
	}
	err = session.Commit()
	session.Close()
	if err != nil {
		panic(err.Error())
	}
}

type UserTable struct {
	UserId   int64  `xorm:"pk autoincr"`
	UserName string `xorm:"varchar(32)"` //用户名
	UserAge  int64  `xorm:"default 1"`   //用户年龄
	UserSex  int64  `xorm:"default 0"`   //用户性别
}

/**
 * 学生表
 */
type StudentTable struct {
	Id          int64  `xorm:"pk autoincr"` //主键 自增
	StudentName string `xorm:"varchar(24)"` //
	StudentAge  int    `xorm:"int default 0"`
	StudentSex  int    `xorm:"index"` //sex为索引
}

/**
 * 人员结构表
 */
type PersonTable struct {
	Id         int64     `xorm:"pk autoincr"`   //主键自增
	PersonName string    `xorm:"varchar(24)"`   //可变字符
	PersonAge  int       `xorm:"int default 0"` //默认值
	PersonSex  int       `xorm:"notnull"`       //不能为空
	City       CityTable `xorm:"-"`             //不映射该字段
}

type CityTable struct {
	CityName      string
	CityLongitude float32
	CityLatitude  float32
}

//func test1() {
//	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1)/test?loc=Local&charset=utf8")
//
//	if err != nil {
//		panic(err.Error())
//	}
//	defer engine.Close()
//
//	engine.ShowSQL(true)
//	engine.Logger().SetLevel(core.LOG_DEBUG)
//	engine.SetMaxOpenConns(10)
//
//	session := engine.Table("payment")
//	if count, err := session.Count(); err == nil {
//		fmt.Println("总数：", count)
//	} else {
//		panic(err.Error())
//	}
//}
