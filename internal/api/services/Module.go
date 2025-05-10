package services
import (
	"gorm.io/gorm"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	)	
func GetModuleByName(db *gorm.DB, name string) (*models.Module, error) {
    var module models.Module
    err := db.Where("name = ?", name).First(&module).Error
    if err != nil {
        return nil, err
    }
    return &module, nil
}
func Createmodule(db *gorm.DB, module *models.Module) (*models.Module, error) {
	if err := db.Create(module).Error; err != nil {
		return nil, err
	}
	return module, nil
}
func RemoveModule(moduleID int, modules []models.Module) []models.Module {
	var result []models.Module
	for _, m := range modules {
		if int(m.ID) != moduleID {
			result = append(result, m)
		}
	}
	return result
}