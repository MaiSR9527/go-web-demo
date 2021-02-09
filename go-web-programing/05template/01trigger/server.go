package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
 *
 *@author MaiShuRen
 *@date 2021/1/20 22:46
 *@version v1.0
 */
func hello(w http.ResponseWriter, r *http.Request) {
	//当用户向ParseFiles 函数或ParseFiles 方法传入多个文件时，
	//ParseFiles 只会返回用户传入的第一个文件的已分析模板，并且这个模板
	//也会根据用户传入的第一个文件的名字进行命名；
	if files, err := template.ParseFiles("05template/01trigger/index.html"); err != nil {
		// 传入单个文件返回的是一个模板，传入多个返回模板集合
		fmt.Println(err.Error())
	} else {
		err := files.Execute(w, "hello template")
		if err != nil {
			panic(err.Error())
		}
	}

}
func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/", hello)
	server.ListenAndServe()
}
