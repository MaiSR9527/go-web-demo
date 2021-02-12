package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func file(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	// 获取<input type="file" name="uploaded"> 上传的文件 [0]是多文件上传的时候读取第一个文件

	// 单个文件获取
	//formFile, _, err := r.FormFile("uploaded")

	fileHeader := r.MultipartForm.File["uploaded"][0]
	if file, err := fileHeader.Open(); err == nil {
		if data, err := ioutil.ReadAll(file); err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", file)
	server.ListenAndServe()
}
