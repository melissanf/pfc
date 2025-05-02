package handlers

import (
	"Devenir_dev/internal/database"
	"Devenir_dev/internal/api/services"
	"Devenir_dev/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v5"	
	"os"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY")) 

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.Rendertemplates(w, "Login", nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur dans le formulaire", http.StatusBadRequest)
		return
	}

	identifier := r.FormValue("identifier")
	password := r.FormValue("password")
	db := database.GetDB()

	// Authentifie l'utilisateur
	verified, isAdmin, message := utils.VerifyUser(db, identifier, password)
	if !verified {
		http.Error(w, message, http.StatusUnauthorized)
		return
	}

	var name string
	user, err := services.GetUserByEmail(db, identifier)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'utilisateur", http.StatusInternalServerError)
		return
	}
	name = user.Nom + " " + user.Prenom

	// Création du token JWT
	claims := jwt.MapClaims{
		"username": name,
		"role":     isAdmin,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		http.Error(w, "Erreur lors de la création du token", http.StatusInternalServerError)
		return
	}

	// Réponse JSON avec le token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})

	// Optionnel : Stocker le token dans un cookie HTTP-only sécurisé
	/*
		http.SetCookie(w, &http.Cookie{
			Name:     "auth-token",
			Value:    tokenString,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
		})
	*/

	// Log de l'opération
	fmt.Printf("Token JWT généré pour %s\n", name)
}
