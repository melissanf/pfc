package services
import (
	"gorm.io/gorm"
	"Devenir_dev/internal/api/models"
	)	
func GetModuleByName(db *gorm.DB, name string) (*models.Module, error) {
    var module models.Module
    err := db.Where("name = ?", name).First(&module).Error
    if err != nil {
        return nil, err
    }
    return &module, nil
}