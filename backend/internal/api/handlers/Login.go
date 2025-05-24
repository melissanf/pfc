package handlers

import (
	"github.com/melissanf/pfc/backend/internal/database"
	"github.com/melissanf/pfc/backend/pkg"
	"net/http"
	"encoding/json"
)



func Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur dans le formulaire", http.StatusBadRequest)
		return
	}
	type LoginRequest struct {
		Identifier string `json:"email"`
		Password   string `json:"password"`
		code       string `json:"code"`
	}

	var reqData LoginRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	db := database.GetDB()

	// Authentifie l'utilisateur
	verified, user, message := utils.VerifyUser(db,reqData.Identifier, reqData.Password)
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
