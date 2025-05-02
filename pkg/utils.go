package utils

import (
	"Devenir_dev/internal/api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	"fmt"
)

type Pagedata struct {
	Currentuser models.User
	Users       []models.User
}

// Rendertemplates charge et affiche les templates
func Rendertemplates(res http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("C:\\Users\\PC\\OneDrive\\Documents\\futur\\Devenir_dev\\templates\\" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(res, data)
	if err != nil {
		http.Error(res, "Error executing template", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
	}
}

// VerifyUser vérifie l'authentification de l'utilisateur avec GORM
func VerifyUser(db *gorm.DB, identifier, password string) (bool, models.Role, string) {
	var user models.User

	// Vérifie si l'identifiant est un email ou un nom d'utilisateur
	if strings.Contains(identifier, "@") {
		// Si c'est un email
		if err := db.Where("email = ?", identifier).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return false, "", "User not found."
			}
			log.Println("GORM Error:", err)
			return false, "", "Database error."
		}
	} else {
		// Si c'est un nom d'utilisateur
		if err := db.Where("nom = ?", identifier).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return false, "", "User not found."
			}
			log.Println("GORM Error:", err)
			return false, "", "Database error."
		}
	}

	// Vérifie le mot de passe
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, "", "Incorrect password."
	}

	return true, user.Role, "User verified."
}

// ValidateInput vérifie la validité des champs utilisateur
func ValidateInput(user models.User) (bool, string) {
	// Vérification des champs vides
	if user.Nom == "" || user.Prenom == "" || user.Password == "" || user.Role == "" || user.Email == "" {
		return false, "All fields (nom, prenom, email, password, role) are required."
	}

	// Vérification de l'email avec une expression régulière
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return false, "Invalid email format."
	}

	// Vérification de la longueur du mot de passe (ex: minimum 6 caractères)
	if len(user.Password) < 6 {
		return false, "Password must be at least 6 characters long."
	}

	return true, ""
}
func sanitizeRole(role models.Role) models.Role {
	switch role {
	case models.Admin,models.Professeur, models.Responsable:
		return models.Role(role) // Rôle valide
	default:
		// Retourne un rôle par défaut si le rôle est invalide
		return models.Professeur
	}
}

func SanitizeInput(user *models.User) {
	re := regexp.MustCompile("<.*?>")

	user.Nom = clean(user.Nom, re)
	user.Prenom = clean(user.Prenom, re)
	user.Password = clean(user.Password, re)
	user.Email = clean(user.Email, re)
	user.Role = sanitizeRole(user.Role)
}

// clean supprime les balises HTML et les espaces inutiles
func clean(s string, re *regexp.Regexp) string {
	return re.ReplaceAllString(strings.TrimSpace(s), "")
}

// FormBool vérifie si une case à cocher est activée dans un formulaire
func FormBool(r *http.Request, key string) bool {
	return r.FormValue(key) == "on"
}
