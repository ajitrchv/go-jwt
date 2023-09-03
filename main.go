package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)


func main(){
	var mux = http.NewServeMux()
	var page = `
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>
			Please Login to view the content!
		</h1>

	</body>
	</html>
	`
	mux.HandleFunc("/", func(w http.ResponseWriter, r* http.Request){
		w.Header().Add("content-type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(page))
	})

	mux.HandleFunc("/home",Home)
	mux.HandleFunc("/refresh", Refresh)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", ClearCookie)


	s := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}


	fmt.Println("Starting server on 8080...")
	log.Fatal(s.ListenAndServe())
}

