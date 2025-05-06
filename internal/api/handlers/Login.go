package handlers

import (
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
	"github.com/ilyes-rhdi/Projet_s4/pkg"
	"encoding/json"
	"fmt"
	"net/http"
)



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
	verified, user, message := utils.VerifyUser(db, identifier, password)
	if !verified {
		http.Error(w, message, http.StatusUnauthorized)
		return
	}
	tokenString, err := utils.GenerateJWT(&user)
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
	fmt.Printf("Token JWT généré pour %s\n", user.Nom)
}
