package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}
	err := server.ListenAndServeTLS("cert.pem", "key.pom")
	if err != nil {
		fmt.Println(err.Error())
	}
}
