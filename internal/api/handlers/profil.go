package handlers

import (
	"Devenir_dev/internal/api/models"
	"Devenir_dev/internal/api/services"
	"Devenir_dev/internal/database"
	"Devenir_dev/pkg"
	"net/http"
	"os"
	"strings"
	"github.com/golang-jwt/jwt/v5"
)

func HandelProfile(res http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := &jwt.MapClaims{}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		http.Error(res, "Internal server error: secret key missing", http.StatusInternalServerError)
		return
	}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
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
    db:=database.GetDB()
	// Récupérer l'utilisateur à partir de l'email
	user, err := services.GetUserByEmail(db, email) // Assurez-vous que `db` est accessible ici
	if err != nil {
		http.Error(res, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	// Utilisation des données de l'utilisateur
	Data := models.User{
		Nom:    user.Nom,
		Prenom: user.Prenom,
		Email:  user.Email,
	}

	utils.Rendertemplates(res, "Profil", Data)
}
