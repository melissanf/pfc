package handlers

import (
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/services"
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
	"github.com/ilyes-rhdi/Projet_s4/pkg"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)
func Submit(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		utils.Rendertemplates(res, "Submit", nil)
		return
	}

	if req.Method != http.MethodPost {
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	db := database.GetDB()

	err := req.ParseForm()
	if err != nil {
		http.Error(res, "Error parsing form data", http.StatusBadRequest)
		return
	}

	user := models.User{
		Nom:      req.FormValue("nom"),
		Prenom:   req.FormValue("prenom"),
		Email:    req.FormValue("email"),
		Password: req.FormValue("password"),
		Role:     models.Role(req.FormValue("role")), // Assure-toi que la valeur correspond bien à un rôle valide
	}

	utils.ValidateInput(user)
	utils.SanitizeInput(&user)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword[:])

	if err := services.CreateUser(db, &user); err != nil {
		http.Error(res, "Failed to insert user into database", http.StatusInternalServerError)
		return
	}


	token, err := utils.GenerateJWT(&user)
	if err != nil {
		http.Error(res, "Erreur lors de la génération du token", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Authorization", "Bearer "+token)
	

	// Redirection ou réponse JSON selon le cas
	http.Redirect(res, req, "/home", http.StatusFound)
}