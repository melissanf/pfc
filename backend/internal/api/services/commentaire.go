package services

import (
	"github.com/melissanf/pfc/backend/internal/api/models"
	"gorm.io/gorm"
)
	
func DeleteCommentaire(db *gorm.DB,id uint) error {
    var commentaire models.Commentaire
    if err := db.Where("id = ?", id).First(&commentaire).Error; err != nil {
        return err
    }
    
    if err := db.Delete(&commentaire).Error; err != nil {
        return err
    }
    
    return nil
}
func UpdateCommentaire(db *gorm.DB,id uint, contenu string) (*models.Commentaire, error) {
    var commentaire models.Commentaire
    if err := db.Where("id = ?", id).First(&commentaire).Error; err != nil {
        return nil, err
    }
    
    commentaire.Contenu = contenu
    if err := db.Save(&commentaire).Error; err != nil {
        return nil, err
    }

    return &commentaire, nil
}
func CreateCommentaire(db *gorm.DB,commentaire *models.Commentaire) (*models.Commentaire, error) {
    if err := db.Create(commentaire).Error; err != nil {
        return nil, err
    }
	if err := NotifyAdminOnComment(db,commentaire); err != nil {
        return nil, err
    }
    return commentaire, nil
}
