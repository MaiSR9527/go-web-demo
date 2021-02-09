package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("05template/04set/index.html")
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
