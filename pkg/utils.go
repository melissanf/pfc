package utils

import (
	"Devenir_dev/internal/api/models"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)


type Pagedata struct {
	Currentuser models.User
	Users []models.User 
}

func Rendertemplates(res http.ResponseWriter,tmpl string ,data interface{}){
	t, err:= template.ParseFiles("C:\\Users\\PC\\OneDrive\\Documents\\futur\\Devenir_dev\\templates\\"+tmpl+".page.tmpl")
	if err !=nil  {
	   http.Error(res,err.Error(),http.StatusInternalServerError)
	   return
	}
	err =t.Execute(res , data)
	if err != nil {
	   http.Error(res, "Error executing template", http.StatusInternalServerError)
	   fmt.Println("Error executing template:", err)
     }
   }
func VerifyUser(db *sql.DB, identifier, password string) (bool, bool, string) {
	var storedPassword []byte
	var isAdmin bool
	var query string
	
	// Check if identifier is an email or username
	if strings.Contains(identifier, "@") {
	    query = "SELECT password, isAdmin FROM users WHERE email = ?"
	        } else {
			    query = "SELECT password, isAdmin FROM users WHERE name = ?"
		    }
	
		// Execute the query
	err := db.QueryRow(query, identifier).Scan(&storedPassword, &isAdmin)
	
		// Handle case where the user is not found or other SQL errors occur
	if err != nil {
		if err == sql.ErrNoRows {
			return false, false, "User not found."
		}
		    log.Println("SQL Error:", err)
			return false, false, "Database error."
		}
	
		// Compare provided password with stored password (ensure passwords are hashed)
	
	if err =bcrypt.CompareHashAndPassword(storedPassword,[]byte(password));err !=nil{
		return true, isAdmin, "User verified."
	} else {
		return false, false, "Incorrect password."
	}
}

func ValidateInput(user models.User) (bool, string) {
	// Vérification des champs vides
	if user.Username == "" || user.Email == "" || user.PasswordHash == "" || user.Role == "" || user.FullName == "" {
		return false, "All fields (name, email, password,Speciality ,Year_entrance, Grade) are required."
	}

	// Vérification de l'email avec une expression régulière
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return false, "Invalid email format."
	}

	// Vérification de la longueur du mot de passe (ex: minimum 6 caractères)
	if len(user.PasswordHash) < 6 {
		return false, "Password must be at least 6 characters long."
	}
    
	return true, ""
}
func SanitizeInput(user *models.User) {
	re := regexp.MustCompile("<.*?>")

	user.Username = clean(user.Username, re)
	user.Email = clean(user.Email, re)
	user.PasswordHash = clean(user.PasswordHash, re)
	user.Role = clean(user.Role, re)
	user.FullName = clean(user.FullName, re)
}

// clean supprime les balises HTML et les espaces inutiles
func clean(s string, re *regexp.Regexp) string {
	return re.ReplaceAllString(strings.TrimSpace(s), "")
}
func formBool(r *http.Request, key string) bool {
    return r.FormValue(key) == "on"
}