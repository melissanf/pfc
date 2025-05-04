package utils

import (
	"Devenir_dev/internal/api/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Pagedata struct {
	Currentuser models.User
	Users       []models.User
}

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
func VerifyUser(db *gorm.DB, identifier, password string) (bool, models.User, string) {
	var user models.User
	if err := db.Where("email = ?", identifier).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, null, "User not found."
		}
		log.Println("GORM Error:", err)
		return false, null, "Database error."
	}
	// Vérifie le mot de passe
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, null, "Incorrect password."
	}

	return true, user, "User verified."
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

func FindTeacher(teacherID int, teachers []models.Teacher) *models.Teacher {
    var t uint 
	for _, t = range teachers {
		if t.ID == teacherID {
			return &t
		}
	}
	return nil
}


func FindModuleForTeacher(teacherID int, slotType string, wishes []models.Voeux, available []models.Module,Niveau[]models.Niveau, currentHours int) *models.Module {
	// Try priorities 1 to 3
	for prio := 1; prio <= 3; prio++ {
		for _, wish := range wishes {
			if wish.TeacherID == teacherID && wish.Priority == prio {
				// Check if teacher wants this type of class
				if (slotType == "cours" && wish.WantsCours) ||
					(slotType == "td" && wish.WantsTD) ||
					(slotType == "tp" && wish.WantsTP) {
					// Find the module in available modules
					for _, module := range available {
						if module.ID == wish.ModuleID {
							hours := GetHoursForType(&module,wish.niveauID,slotType)
							if hours > 0 && currentHours+hours <= 24 {
								return &module
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func GetHoursForType(module *models.Module, niveauID uint, slotType string) float64 {
	for _, mn := range module.ModuleNiveaux {
		if mn.NiveauID == niveauID {
			switch slotType {
			case "cours":
				return mn.ChargeCours
			case "td":
				return mn.ChargeTD
			case "tp":
				return mn.ChargeTP
			}
		}
	}
	return 0
}



// FormBool vérifie si une case à cocher est activée dans un formulaire
func FormBool(r *http.Request, key string) bool {
	return r.FormValue(key) == "on"
}
func FindModuleForTeacher(
    teacherID int,
    niveauID uint,
    slotType string,
    voeux []models.Voeux,
    available []models.Module,
    currentHours float64,
) *models.Module {
    for prio := 1; prio <= 3; prio++ {
        for _, v := range voeux {
            if int(v.TeacherID) == teacherID && v.Priority == prio && v.NiveauID == niveauID {
                if (slotType == "cours" && v.Cours) ||
                   (slotType == "td" && v.Td) ||
                   (slotType == "tp" && v.Tp) {
                    
                    for _, module := range available {
                        if module.ID == v.ModuleID {
                            hours := GetHoursForType(&module, niveauID, slotType)
                            if hours > 0 && currentHours+hours <= 24 {
                                return &module
                            }
                        }
                    }
                }
            }
        }
    }
    return nil
}

func GenerateJWT(user *models.User) (string, error) {
	// Crée les claims avec la date d’expiration
	claims := models.Claims{
		UserID:   user.ID,
		Username: user.Nom + " " + user.Prenom,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24h de validité
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Crée le token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signe le token avec la clé secrète
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}