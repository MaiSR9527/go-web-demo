package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
*/

func main() {
	router := gin.Default()

	router.GET("/post", func(c *gin.Context) {
		// 获取请求路径的参数  ?id=123&page=1
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		//获取请求体中的参数
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	// 匹配请求路径
	//router.GET("/user/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		fmt.Sprintln(b)
	})
	router.Run(":8089")
}
