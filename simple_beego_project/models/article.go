package models

import "time"

type Article struct {
	Id          int       `orm:"pk;auto"`
	ArticleName string    `orm:"size(20)"`
	CreateAt    time.Time `orm:"auto_now"`
	Count       int       `orm:"default(0);null"`
	Content     string    `orm:"size(500)"`
	Img         string    `orm:"size(100)"`
}
