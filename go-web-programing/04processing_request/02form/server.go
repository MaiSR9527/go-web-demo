package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.MultipartForm)
	// r.Form
	//map[hello:[sau sheong world] post:[456] thread:[123]]

	// r.PostForm
	// map[hello:[sau sheong] post:[456]]

	//
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
