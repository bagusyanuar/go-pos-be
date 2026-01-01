package util

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTAppClaims struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}
