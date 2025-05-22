package handlers

import (
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/services"
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
	"net/http"
	"encoding/json"
	"strings"
)

func HandelProfile(res http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(res, "Token manquant", http.StatusUnauthorized)
		return
	}


	claims, ok := req.Context().Value("user").(*models.Claims)
	if !ok {
		http.Error(res, "Erreur de récupération des claims", http.StatusInternalServerError)
		return
	}
    db:=database.GetDB()
	// Récupérer l'utilisateur à partir de l'id
	user, err := services.GetUserByID(db, claims.UserID) 
	if err != nil {
		http.Error(res, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}
    type data struct{
		Nom string 
		Prenom string 
		Email string 
		Numero string 

	}
	// Utilisation des données de l'utilisateur
	Data := data{
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
