package services

import (
	"gorm.io/gorm"
	"github.com/melissanf/pfc/backend/internal/api/models"
)

// CreateUser crée un nouvel utilisateur dans la base de données.
func CreateUser(db *gorm.DB, user *models.User) error {
    return db.Create(user).Error
}

// GetUserByID récupère un utilisateur par son ID.
func GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
    var user models.User
    if err := db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// GetUserByEmail récupère un utilisateur par son email.
func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
    var user models.User
    if err := db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// UpdateUser met à jour un utilisateur dans la base de données.
func UpdateUser(db *gorm.DB, user *models.User) error {
    return db.Save(user).Error
}

// DeleteUser supprime un utilisateur de la base de données.
func DeleteUser(db *gorm.DB, id uint) error {
    return db.Delete(&models.User{}, id).Error
}