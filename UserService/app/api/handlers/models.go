package handlers

import "github.com/dgrijalva/jwt-go"

type UserCredentials struct {
	Name           string `db:"name" json:"name"`
	Password       string `json:"password"`
	HashedPassword []byte `db:"password"`
}

type Claims struct {
	jwt.StandardClaims
}
