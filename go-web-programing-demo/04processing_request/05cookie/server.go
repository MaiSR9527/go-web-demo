package main

import (
	"fmt"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := &http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := &http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, c1)
	http.SetCookie(w, c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("first_cookie"); err == nil {
		fmt.Fprintln(w, cookie)
	} else {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	fmt.Fprintln(w, r.Cookies())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	server.ListenAndServe()
}
