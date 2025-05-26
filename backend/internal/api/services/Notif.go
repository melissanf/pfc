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
	err := db.Model(&models.Notif{}).Where("is_read = ?", false).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// NOUVELLES MÉTHODES À AJOUTER :

// GetNotificationsByUser récupère toutes les notifications d'un utilisateur
func GetNotificationsByUser(db *gorm.DB, userID uint) ([]models.Notif, error) {
	var notifications []models.Notif
	err := db.Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

// GetNotificationsByUsername récupère les notifications par nom d'utilisateur
func GetNotificationsByUsername(db *gorm.DB, username string) ([]models.Notif, error) {
	var notifications []models.Notif
	err := db.Where("destinataire = ?", username).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

// CreateOrganigrammeNotification crée une notification spécifique pour l'organigramme
func CreateOrganigrammeNotification(db *gorm.DB, destinataire, titre, message, semestre string) error {
	// Trouver l'utilisateur par nom (si vous stockez par user_id)
	user, err := GetUserByName(db, destinataire)
	var userID uint = 0
	if err == nil && user != nil {
		userID = user.ID
	}
	
	notif := models.Notif{
		UserID:      userID,
		Destinataire: destinataire,
		Type:        "organigramme_valide",
		Titre:       titre,
		Message:     message,
		IsRead:      false,
	}
	
	return db.Create(&notif).Error
}



// GetUnreadNotifsCountByUser compte les notifications non lues pour un utilisateur spécifique
func GetUnreadNotifsCountByUser(db *gorm.DB, userID uint) (int64, error) {
	var count int64
	err := db.Model(&models.Notif{}).Where("user_id = ? AND is_read = ?", userID, false).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetUnreadNotifsCountByUsername compte les notifications non lues par nom d'utilisateur
func GetUnreadNotifsCountByUsername(db *gorm.DB, username string) (int64, error) {
	var count int64
	err := db.Model(&models.Notif{}).Where("destinataire = ? AND is_read = ?", username, false).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func NotifyAdminOnComment(db *gorm.DB, comment *models.Commentaire) error {
	// Créer la notification pour l'admin
	user, _ := GetUserByID(db, comment.AuteurID)
	if user == nil {
		return fmt.Errorf("Utilisateur introuvable")	
	}
	notif := models.Notif{
		UserID:       comment.AuteurID,
		Destinataire: "admin", // Assurez-vous que "admin" est le nom d'utilisateur de l'admin
		Type:         "commentaire",
		Titre:        fmt.Sprintf("Vous avez reçu un commentaire de %s.", user.Nom),
		Message:      comment.Contenu,			
	}
	return db.Create(&notif).Error
}