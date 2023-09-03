package main

import "github.com/golang-jwt/jwt"

type Cred struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user3": "password3",
	"user4": "password4",
	"user5": "password5",
	"user6": "password6",
	"user7": "password7",
	"user8": "password8",
}