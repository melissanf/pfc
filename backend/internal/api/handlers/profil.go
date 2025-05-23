package handlers

import (
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/services"
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
	"net/http"
	"os"
	"encoding/json"
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
	user, err := services.GetUserByEmail(db, email) 
	if err != nil {
		http.Error(res, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	// Utilisation des données de l'utilisateur
	Data := models.User{
		Nom:    user.Nom,
		Prenom: user.Prenom,
		Email:  user.Email,
		Numero: user.Numero,
	}
	if err :=json.NewEncoder(res).Encode(Data); err!=nil {	
		http.Error(res, "Erreur lors de l'envoi de la réponse", http.StatusInternalServerError)
		return
	}
}
