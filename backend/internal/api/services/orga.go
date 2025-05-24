package services
import (
	"gorm.io/gorm"
	"github.com/melissanf/pfc/backend/internal/api/models"
)
func CreateOrganigramme(db *gorm.DB,organigramme *models.Organigramme) (*models.Organigramme, error) {
    if err := db.Create(organigramme).Error; err != nil {
        return nil, err
    }
    return organigramme, nil
}
func GetOrganigrammeByID(db *gorm.DB,id uint) (*models.Organigramme, error) {
    var organigramme models.Organigramme
    if err := db.Preload("Affectations").Where("id = ?", id).First(&organigramme).Error; err != nil {
        return nil, err
    }
    return &organigramme, nil
}
func GetOrganigrammeByYearAndSemester(db *gorm.DB,annee, semestre string) ([]models.Organigramme, error) {
    var organigrammes []models.Organigramme
    if err := db.Preload("Affectations").
        Where("annee = ? AND semestre = ?", annee, semestre).
        Find(&organigrammes).Error; err != nil {
        return nil, err
    }
    return organigrammes, nil
}
func UpdateOrganigramme(db *gorm.DB,id uint, updatedOrganigramme *models.Organigramme) (*models.Organigramme, error) {
    var organigramme models.Organigramme
    if err := db.Where("id = ?", id).First(&organigramme).Error; err != nil {
        return nil, err
    }

    // Mettre à jour les champs
    organigramme.Annee = updatedOrganigramme.Annee
    organigramme.Semestre = updatedOrganigramme.Semestre
    organigramme.IsValide = updatedOrganigramme.IsValide

    // Sauvegarde les modifications
    if err := db.Save(&organigramme).Error; err != nil {
        return nil, err
    }

    return &organigramme, nil
}
func DeleteOrganigramme(db *gorm.DB,id uint) error {
    var organigramme models.Organigramme
    if err := db.Where("id = ?", id).First(&organigramme).Error; err != nil {
        return err
    }

    // Supprimer l'organigramme
    if err := db.Delete(&organigramme).Error; err != nil {
        return err
    }

    return nil
}
