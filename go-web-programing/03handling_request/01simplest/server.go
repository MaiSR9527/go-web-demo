package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe("", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
