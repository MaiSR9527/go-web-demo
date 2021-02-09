package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 响应发送html
func responseWriteHtml(w http.ResponseWriter, r *http.Request) {
	str := `<html>
				<head><title>Go Web Programming</title></head>
				<body><h1>Hello World</h1></body>
			</html>`
	w.Write([]byte(str))
}

// 设置响应状态码
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

// 设置重定向
func setHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.baidu.com")
	w.WriteHeader(302)
}

type User struct {
	User    string
	Threads []string
}

// 响应json
func responseJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	data := &User{
		User:    "msr",
		Threads: []string{"102", "8521", "8854"},
	}
	if marshal, err := json.Marshal(data); err == nil {
		w.Write(marshal)
	} else {
		log.Panic(err.Error())
	}

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", responseWriteHtml)
	http.HandleFunc("/login", writeHeaderExample)
	http.HandleFunc("/set", setHeaderExample)
	http.HandleFunc("/json", responseJson)

	server.ListenAndServe()
}
