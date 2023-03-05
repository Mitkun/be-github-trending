package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClains struct {
	UserId string
	Role   string
	jwt.StandardClaims
}
