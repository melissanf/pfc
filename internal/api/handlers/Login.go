package handlers

import (
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
	"github.com/ilyes-rhdi/Projet_s4/pkg"
	"net/http"
)



func Login(w http.ResponseWriter, r *http.Request) {
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
	token, err := utils.GenerateJWT(&user)
	if err != nil {
		http.Error(w, "Erreur lors de la création du token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
}
