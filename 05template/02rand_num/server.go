package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

/**
 *
 *@author MaiShuRen
 *@date 2021/1/20 23:11
 *@version v1.0
 */
func handleIf(w http.ResponseWriter, r *http.Request) {
	files, _ := template.ParseFiles("05template/02rand_num/index.html")
	rand.Seed(time.Now().Unix())
	err := files.Execute(w, rand.Intn(10) > 5)
	if err != nil {
		log.Println(err)
	}
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", handleIf)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
