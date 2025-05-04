package handlers

import (
	"Devenir_dev/internal/api/models"
	"Devenir_dev/internal/api/services"
	"Devenir_dev/internal/database"
	"Devenir_dev/pkg"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
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
	user.Password = string(hashedPassword)

	if err := services.CreateUser(db, &user); err != nil {
		http.Error(res, "Failed to insert user into database", http.StatusInternalServerError)
		return
	}


	claims := models.Claims{
		UserID:   user.ID,
		Username: user.Nom + " " + user.Prenom,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		http.Error(res, "Erreur lors de la génération du token", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Authorization", "Bearer "+tokenString)

	// Redirection ou réponse JSON selon le cas
	http.Redirect(res, req, "/Home", http.StatusFound)
}