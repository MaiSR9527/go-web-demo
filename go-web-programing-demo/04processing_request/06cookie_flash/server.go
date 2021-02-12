package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("哈哈哈")
	c := http.Cookie{
		Name:  "flash-msg",
		Value: base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("flash-msg"); err == nil {
		rc := http.Cookie{
			Name:    "flash-msg",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		value, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			log.Panic(err.Error())
		}else {
			fmt.Fprintln(w, string(value))
		}

	} else {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/setMessage", setMessage)
	http.HandleFunc("/showMessage", showMessage)
	server.ListenAndServe()
}
