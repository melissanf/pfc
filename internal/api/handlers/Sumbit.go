package handlers

import (
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/services"
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
	"github.com/ilyes-rhdi/Projet_s4/pkg"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)
func Submit(res http.ResponseWriter, req *http.Request) {
	db := database.GetDB()
    var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, "Données invalides", http.StatusBadRequest)
		return
	}
	utils.ValidateInput(user)
	utils.SanitizeInput(&user)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword[:])

	if err := services.CreateUser(db, &user); err != nil {
		http.Error(res, "Failed to insert user into database", http.StatusInternalServerError)
		return
	}
	if user.Role == "professeur" {
		teacher := models.Teacher{
			UserID: user.ID,
			ChargeHoraire: 0,
		};
		if err := services.CreateTeacher(db, &teacher); err != nil {
			http.Error(res, "Failed to insert user into database", http.StatusInternalServerError)
			return
		}
	}


	token, err := utils.GenerateJWT(&user)
	if err != nil {
		http.Error(res, "Erreur lors de la génération du token", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Authorization", "Bearer "+token)
}