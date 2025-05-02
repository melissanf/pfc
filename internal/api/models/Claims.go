package models

import (
	"github.com/dgrijalva/jwt-go"
)
// Claims représente la structure des données stockées dans le token JWT.
type Claims struct {
    UserID uint   `json:"user_id"`
    Role   string `json:"role"`
    jwt.StandardClaims
}