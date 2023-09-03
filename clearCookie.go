package main

import (
	"net/http"
	"time"
)

func ClearCookie(w http.ResponseWriter, r *http.Request){
	expirationTime := time.Now().Add(time.Minute * -1)
	http.SetCookie(w, &http.Cookie{
		Name: "cookie_cleared",
		Value: "",
		Expires: expirationTime,
	})
}