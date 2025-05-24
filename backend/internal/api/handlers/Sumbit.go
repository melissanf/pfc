package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/melissanf/pfc/backend/internal/api/models"
	"github.com/melissanf/pfc/backend/internal/api/services"
	"github.com/melissanf/pfc/backend/internal/database"
	utils "github.com/melissanf/pfc/backend/pkg"
	"golang.org/x/crypto/bcrypt"
)

func Submit(res http.ResponseWriter, req *http.Request) {
	db := database.GetDB()

	type input struct {
		Nom      string      `json:"nom"`
		Prenom   string      `json:"prenom"`
		Email    string      `json:"email"`
		Password string      `json:"password"`
		Numero   string      `json:"numero"`
		Role     models.Role `json:"role"`
		Code     string      `json:"code"`
	}
	var inputData input

	err := json.NewDecoder(req.Body).Decode(&inputData)
	if err != nil {
		http.Error(res, "Données invalides", http.StatusBadRequest)
		return
	}
	ok, err := utils.Ismatch(db, inputData.Role, inputData.Code)
	if err != nil {
		http.Error(res, "Erreur lors de la vérification du code", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(res, "Le code ne correspond pas au rôle sélectionné", http.StatusBadRequest)
		return
	}

	var user models.User
	user.Nom = inputData.Nom
	user.Prenom = inputData.Prenom
	user.Email = inputData.Email
	user.Password = inputData.Password
	user.Numero = inputData.Numero
	user.Role = inputData.Role
	user.Code = inputData.Code

	utils.ValidateInput(user)
	utils.SanitizeInput(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	if err := services.CreateUser(db, &user); err != nil {
		http.Error(res, "Échec de l'insertion de l'utilisateur dans la base de données", http.StatusInternalServerError)
		return
	}

	if user.Role == "enseignant" {
		teacher := models.Teacher{
			UserID:        user.ID,
			ChargeHoraire: 0,
			Heursupp:      0,
		}
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

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"nom":    user.Nom,
			"prenom": user.Prenom,
			"email":  user.Email,
			"numero": user.Numero,
			"role":   user.Role,
			"code":   user.Code,
		},
	})
}
