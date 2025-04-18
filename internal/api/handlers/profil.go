package handlers

import (
	"net/http"
	"Devenir_dev/pkg"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"Devenir_dev/internal/api/models"
)
func HandelProfile(res http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" || ! strings.HasPrefix(authHeader, "Bearer ") {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // remplace "your-secret-key" par ta vraie clé
	})

	if err != nil || !token.Valid {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	username, _ := (*claims)["username"].(string)
	email, _ := (*claims)["email"].(string)

	if username == "" || email == "" {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	Data := models.User{
		Username: username,
		Email:    email,
	}
	utils.Rendertemplates(res, "Profil", Data)
}
