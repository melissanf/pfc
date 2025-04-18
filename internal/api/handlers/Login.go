package handlers

import (
	"Devenir_dev/internal/database"
	"Devenir_dev/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/golang-jwt/jwt/v5"
)
const SecretKey = "ilyes"
var jwtSecretKey = []byte(SecretKey)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.Rendertemplates(w, "Login", nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Parse form
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

	// Récupération du nom si email
	var name string
	if strings.Contains(identifier, "@") {
		query := "SELECT name FROM users WHERE email = ?"
		if err := db.QueryRow(query, identifier).Scan(&name); err != nil {
			http.Error(w, "Erreur lors de la récupération du nom", http.StatusInternalServerError)
			return
		}
	} else {
		name = identifier
	}

	// Création du token JWT
	claims := jwt.MapClaims{
		"username": name,
		"isAdmin":  isAdmin,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		http.Error(w, "Erreur lors de la création du token", http.StatusInternalServerError)
		return
	}

	// Envoi du token (au choix : en JSON ou en cookie)
	// OPTION 1 : Réponse JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})

	// OPTION 2 (alternative) : stocker en cookie HTTP-only sécurisé
	/*
		http.SetCookie(w, &http.Cookie{
			Name:     "auth-token",
			Value:    tokenString,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
		})
	*/

	fmt.Printf("Token JWT généré pour %s\n", name)
}
