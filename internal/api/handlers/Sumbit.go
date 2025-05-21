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
	type input struct{
		Nom      string `json:"nom"`
		Prenom   string `json:"prenom"`
		Email    string `null" json:"email"`
		Password string `json:"password"`
		Numero   string `json:"numero"`
		Role     models.Role  `"json:"role"`
		Year_entrance int         `"json:"year_entrance"`
        Grade         string      `"json:"grade"`
	}
    var inputData input

	err := json.NewDecoder(req.Body).Decode(&inputData)
	if err != nil {
		http.Error(res, "Données invalides", http.StatusBadRequest)
		return
	}
	user.Nom = inputData.Nom
	user.Prenom = inputData.Prenom
	user.Email = inputData.Email
	user.Password = inputData.Password
	user.Numero = inputData.Numero
	user.Role = inputData.Role
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
			Year_entrance: inputData.Year_entrance,
			Grade: inputData.Grade,
			ChargeHoraire: 0,
			Heursupp: 0,
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