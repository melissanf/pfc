package models

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims représente la structure des données stockées dans le token JWT.
type Claims struct {
    UserID uint   `json:"user_id"`
    Username string`json:"username"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}