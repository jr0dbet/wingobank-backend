package middleware

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
