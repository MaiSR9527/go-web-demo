package model

type OrderStatus struct {
	Id         int    `xorm:"pk antoincr" json:"id"` //主键
	StatusId   int    //订单状态编号
	StatusDesc string `xorm:"varchar(255)"` // 订单状态描述

}
