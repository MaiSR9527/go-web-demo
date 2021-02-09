package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	static       string
}

var config Config
var logger *log.Logger

// 打印
func p(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	loadConfig()

}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Can not open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Can not get Configuration from file", err)
	}
}

// 抛出错误信息，并重定向到错误页面
func errorMsg(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func getSession(writer http.ResponseWriter, request *http.Request) sess {

}
