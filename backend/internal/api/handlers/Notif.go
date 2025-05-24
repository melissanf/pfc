package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/melissanf/pfc/backend/internal/database"
	"github.com/melissanf/pfc/backend/internal/api/models"
)
func GetNotifications(w http.ResponseWriter, r *http.Request) {
    db := database.GetDB() // ou comme tu accèdes à ton DB

    var notifs []models.Notif
    err := db.Order("created_at desc").Find(&notifs).Error
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des notifications", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(notifs); err != nil {
		http.Error(w, "Erreur lors de l'envoi des données", http.StatusInternalServerError)
		return
	}
}