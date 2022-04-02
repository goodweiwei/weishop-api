package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	Id          uint
	Nickname    string
	AuthorityId uint
	jwt.StandardClaims
}
