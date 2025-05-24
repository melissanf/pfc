package services
import(
	"gorm.io/gorm"
	"github.com/melissanf/pfc/backend/internal/api/models"
	"fmt"
	"strings"
)



func GetNiveauBySpecAnnee(db *gorm.DB, nom string) (*models.Niveau, error) {
	parts := strings.SplitN(nom, "-", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("format invalide pour le nom : %s", nom)
	}

	annee := parts[0]
	spec := parts[1]

	var niveau models.Niveau
	err := db.Where("annee = ? AND spec = ?", annee, spec).First(&niveau).Error
	if err != nil {
		return nil, err
	}

	return &niveau, nil
}