package utils
import (
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	"log"	
	"regexp"
	"strings"
	"time"
	"os"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Pagedata struct {
	Currentuser models.User
	Users       []models.User
}
func VerifyUser(db *gorm.DB, identifier, password string) (bool, models.User, string) {
	var user models.User
	if err := db.Where("email = ?", identifier).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, user, "User not found."
		}
		log.Println("GORM Error:", err)
		return false, user, "Database error."
	}
	// Vérifie le mot de passe
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, user, "Incorrect password."
	}

	return true, user, "User verified."
}

func ValidateInput(user models.User) (bool, string) {
	if user.Nom == "" || user.Prenom == "" || user.Password == "" || user.Role == "" || user.Email == "" || user.Numero == "" {
		return false, "All fields (nom, prenom, email, password, role) are required."
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return false, "Invalid email format."
	}
	if len(user.Password) < 6 {
		return false, "Password must be at least 6 characters long."
	}
	if len(user.Numero) != 10  || user.Numero[0] != '0' {
		return false, "Phone number must be 10 digits long."
	}

	return true, ""
}
func sanitizeRole(role models.Role) models.Role {
	switch role {
	case models.Chef_de_Departement,models.Personnel, models.Enseignant:
		return models.Role(role) // Rôle valide
	default:
		// Retourne un rôle par défaut si le rôle est invalide
		return models.Enseignant
	}
}

func SanitizeInput(user *models.User) {
	re := regexp.MustCompile("<.*?>")

	user.Nom = clean(user.Nom, re)
	user.Prenom = clean(user.Prenom, re)
	user.Password = clean(user.Password, re)
	user.Email = clean(user.Email, re)
	user.Numero = clean(user.Numero, re)
	user.Role = sanitizeRole(user.Role)
}

// clean supprime les balises HTML et les espaces inutiles
func clean(s string, re *regexp.Regexp) string {
	return re.ReplaceAllString(strings.TrimSpace(s), "")
}

func FindTeacher(teacherID uint, teachers []models.Teacher) *models.Teacher {
	for _, t := range teachers {
		if t.ID == teacherID {
			return &t
		}
	}
	return nil
}



func FindModuleForTeacher(teacherID int,niveauID uint , slotType string, wishes []models.Voeux, available []models.Module,Niveau[]models.Niveau, currentHours int) *models.Module {
    for prio := 1; prio <= 3; prio++ {
        for _, v := range wishes {
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

func GetHoursForType(module *models.Module, niveauID uint, slotType string) int {
	for _, mn := range module.ModuleNiveaux {
		if mn.NiveauID == niveauID {
			switch slotType {
			case "cours":
				return mn.NbCours
			case "td":
				return mn.NbTD
			case "tp":
				return mn.NbTP
			}
		}
	}
	return 0
}


var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY")) 
func GenerateJWT(user *models.User) (string, error) {
	// Crée les claims avec la date d’expiration
	claims := models.Claims{
		UserID:   user.ID,
		Username: user.Nom + " " + user.Prenom,
		Role:     user.Role,
		Code :    user.Code ,
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
func RecalculerChargeHoraire(db *gorm.DB) error {
	coefficients := map[string]float64{
		"Cours": 3.0,
		"TD":    2.0,
		"TP":    1.5,
	}

	var affectations []models.Affectation
	if err := db.Find(&affectations).Error; err != nil {
		return err
	}

	chargeParProf := make(map[uint]float64)

	for _, a := range affectations {
		var mn models.ModuleNiveau
		err := db.Where("module_id = ? AND niveau_id = ?", a.ModuleID, a.NiveauID).First(&mn).Error
		if err != nil {
			continue // ignorer si le lien module-niveau est manquant
		}

		var nbSeances int
		switch a.TypeSeance {
		case "Cours":
			nbSeances = mn.NbCours
		case "TD":
			nbSeances = mn.NbTD
		case "TP":
			nbSeances = mn.NbTP
		default:
			continue
		}

		coef := coefficients[a.TypeSeance]
		charge := float64(nbSeances) * coef * float64(a.Groupe) // charge totale pour ce prof
		chargeParProf[a.TeacherID] += charge
	}

	for teacherID, charge := range chargeParProf {
		if err := db.Model(&models.Teacher{}).Where("id = ?", teacherID).Update("charge_horaire", charge).Error; err != nil {
			return err
		}
	}

	return nil
}
func abbrevRole(role models.Role) string {
    switch role {
	case "Enseignant":
		return "E"
	case "Chef_de_Departement":
		 return "D"
	case "Personnel Administratif":
		return "P"
	default :
		return "/"
	}
}

// Fonction pour générer le code
func GenerateUserCode(user *models.User) string {
    initials := strings.ToUpper(string(user.Prenom[0]) + string(user.Nom[0]))
    roleCode := abbrevRole(user.Role)
    idStr := fmt.Sprintf("%04d", user.ID) // padding avec 0 jusqu'à 4 chiffres
    return fmt.Sprintf("%s-%s-%s", roleCode, initials, idStr) 
}