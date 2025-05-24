package services
import (
	"gorm.io/gorm"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
    "errors"
)


func CreateTeacher(db *gorm.DB, teacher *models.Teacher) error {
    return db.Create(teacher).Error
}

// Get Teacher by ID
func GetTeacherByID(db *gorm.DB, id uint) (*models.Teacher, error) {
    var teacher models.Teacher
    if err := db.First(&teacher, id).Error; err != nil {
        return nil, err
    }
    return &teacher, nil
}
func GetTeacherByUserID(db *gorm.DB, userID uint) (*models.Teacher, error) {
    var teacher models.Teacher
    err := db.Where("user_id = ?", userID).First(&teacher).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // Pas trouvé
        }
        return nil, err // Erreur DB
    }
    return &teacher, nil
}

// Update Teacher
func UpdateTeacher(db *gorm.DB, teacher *models.Teacher) error {
    return db.Save(teacher).Error
}

// Delete Teacher
func DeleteTeacher(db *gorm.DB, id uint) error {
    return db.Delete(&models.Teacher{}, id).Error
}