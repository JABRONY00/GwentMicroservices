package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
}

type Token struct {
}

type PlayerInfo struct {
	ID    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

type PlayerInfoPassword struct {
	PlayerInfo
	Password     string `json:"password"`
	PasswordHash []byte `db:"password"`
}
