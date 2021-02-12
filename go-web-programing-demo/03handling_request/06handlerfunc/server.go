package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello!!!"))
}

func world(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello!!!"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
