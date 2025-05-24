package services
import(
	"github.com/melissanf/pfc/backend/internal/api/models"
	"gorm.io/gorm"
	"fmt"
)
func CreateNotif(db *gorm.DB, notif *models.Notif) error {
	return db.Create(notif).Error
}
func MarkNotifAsRead(db *gorm.DB, id uint) error {
	return db.Model(&models.Notif{}).Where("id = ?", id).Update("is_read", true).Error
}
func DeleteNotif(db *gorm.DB, id uint) error {
	return db.Delete(&models.Notif{}, id).Error
}
func MarkAllNotifsAsRead(db *gorm.DB, userID uint) error {
	return db.Model(&models.Notif{}).Where("user_id = ?", userID).Update("is_read", true).Error
}
func GetUnreadNotifsCount(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&models.Notif{}).Where("is_read = ?",false).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func NotifyAdminOnComment(db *gorm.DB,comment *models.Commentaire) error {
	// Créer la notification pour l'admin
	user,_:=GetUserByID(db, comment.AuteurID)
	if user == nil {
		return fmt.Errorf("Utilisateur introuvable")	
	}
	notif := models.Notif{
		CommentaireID: comment.ID,
		Message:       fmt.Sprintf("Vous avez reçu un commentaire de %s.",user.Nom),
	}
	return db.Create(&notif).Error
}