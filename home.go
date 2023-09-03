package main

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)


func Home(w http.ResponseWriter, r *http.Request){
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

	w.Write([]byte(fmt.Sprintf("Hello, %s!", claims.Username)))
	w.Write([]byte("Welcome home"))
	
}