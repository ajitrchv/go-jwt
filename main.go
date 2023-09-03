package main

import "net/http"

func main(){
	http.HandleFunc("/",Home)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/login", Login)
	http.ListenAndServe(":8080",nil)
}

