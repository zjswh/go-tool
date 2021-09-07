package request

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	Token string
	jwt.StandardClaims
}


