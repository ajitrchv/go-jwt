package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)


func Login(w http.ResponseWriter, r *http.Request){
	var credentials Cred
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if (err != nil){
		w.Write([]byte("<h1>Bad Request</h1>"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expected,ok := users[credentials.Username]
	if !ok || expected != credentials.Password{
		w.Write([]byte("Unautherized Access!"))
		w.WriteHeader(http.StatusUnauthorized)
		
		return
	}

	expiryTime := time.Now().Add(time.Minute*5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil{
		w.Write([]byte("Internal Server Error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expiryTime,
	})
}