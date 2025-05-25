package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/melissanf/pfc/backend/internal/database"
	utils "github.com/melissanf/pfc/backend/pkg"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur dans le formulaire", http.StatusBadRequest)
		return
	}
	type LoginRequest struct {
		Identifier string `json:"email"`
		Password   string `json:"password"`
		Code       string `json:"code"`
	}

	var reqData LoginRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	db := database.GetDB()

	// Authentifie l'utilisateur
	verified, user, message := utils.VerifyUser(db, reqData.Identifier, reqData.Password)
	if !verified {
		http.Error(w, message, http.StatusUnauthorized)
		return
	}
	token, err := utils.GenerateJWT(&user)
	if err != nil {
		http.Error(w, "Erreur lors de la création du token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)

	response := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	json.NewEncoder(w).Encode(response)
}
