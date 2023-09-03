package main

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)


func Home(w http.ResponseWriter, r *http.Request){
	var page = `
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>
			Hi, welcome to my random learning!
		</h1>
		<p>Please help me to get a good job!</p>
	</body>
	</html>
	`
	cookie, err := r.Cookie("token")
	if err != nil{
		if err == http.ErrNoCookie{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	tokenStr := cookie.Value
	
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
	func(t *jwt.Token) (interface{}, error){
		return jwtKey, nil
	})
	if err != nil{
		if err == jwt.ErrSignatureInvalid{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("<h1>Hello, Mr.%s!</h1>\n", claims.Username)))
	w.Write([]byte(page))
}


