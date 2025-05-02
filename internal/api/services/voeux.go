package services
import (
	"gorm.io/gorm"
	"Devenir_dev/internal/api/models"
)
func CreateVoeux(db *gorm.DB, voeux *models.Voeux) error {
    return db.Create(voeux).Error
}

func GetVoeuxByID(db *gorm.DB, id uint) (*models.Voeux, error) {
    var voeux models.Voeux
    if err := db.Preload("Teacher").Preload("Module").Preload("Niveau").First(&voeux, id).Error; err != nil {
        return nil, err
    }
    return &voeux, nil
}

func UpdateVoeux(db *gorm.DB, voeux *models.Voeux) error {
    return db.Save(voeux).Error
}

func DeleteVoeux(db *gorm.DB, id uint) error {
    return db.Delete(&models.Voeux{}, id).Error
}
func CountVoeuxByTeacherID(db *gorm.DB, teacherID uint) (int64, error) {
    var count int64
    err := db.Model(&models.Voeux{}).Where("teacher_id = ?", teacherID).Count(&count).Error
    return count, err
}
func VoeuxExactExists(db *gorm.DB, teacherID, moduleID, niveauID uint, tp, td, cours bool) (bool, error) {
	var count int64
	err := db.Model(&models.Voeux{}).
		Where("teacher_id = ? AND module_id = ? AND niveau_id = ? AND tp = ? AND td = ? AND cours = ?",
			teacherID, moduleID, niveauID, tp, td, cours).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}