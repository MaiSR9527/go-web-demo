package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 你也可以使用自己的 SecureJSON 前缀.只能用于数组
	r.SecureJsonPrefix("json->',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		//names := []string{"lena", "austin", "foo"}
		//
		m := make(map[string]string)
		m["name"] = "msr"
		m["fullname"] = "maishuren"
		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, m)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
