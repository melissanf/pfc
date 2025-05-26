package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/melissanf/pfc/backend/internal/database"
	"github.com/melissanf/pfc/backend/internal/api/models"
    "github.com/melissanf/pfc/backend/internal/api/services"
    "github.com/gorilla/mux"
    "strconv"
    "time"
    "fmt"
)
func GetNotifications(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	
	// Configuration des headers CORS et Content-Type
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Vérifier si c'est une route avec paramètre userID
	vars := mux.Vars(r)
	if userIDStr, exists := vars["userID"]; exists {
		// Route : /notifications/user/{userID}
		userID, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			http.Error(w, `{"error": "ID utilisateur invalide"}`, http.StatusBadRequest)
			return
		}

		notifications, err := services.GetNotificationsByUser(db, uint(userID))
		if err != nil {
			http.Error(w, `{"error": "Erreur lors de la récupération des notifications"}`, http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"notifications": notifications,
			"count":         len(notifications),
		}); err != nil {
			http.Error(w, `{"error": "Erreur lors de l'envoi des données"}`, http.StatusInternalServerError)
			return
		}
		return
	}

	// Vérifier les paramètres de requête
	queryUserID := r.URL.Query().Get("user_id")
	queryUsername := r.URL.Query().Get("username")

	if queryUserID != "" {
		// Récupération par user_id
		userID, err := strconv.ParseUint(queryUserID, 10, 32)
		if err != nil {
			http.Error(w, `{"error": "ID utilisateur invalide"}`, http.StatusBadRequest)
			return
		}

		notifications, err := services.GetNotificationsByUser(db, uint(userID))
		if err != nil {
			http.Error(w, `{"error": "Erreur lors de la récupération des notifications"}`, http.StatusInternalServerError)
			return
		}

		// Compter les notifications non lues
		unreadCount, _ := services.GetUnreadNotifsCountByUser(db, uint(userID))

		response := map[string]interface{}{
			"notifications": notifications,
			"total_count":   len(notifications),
			"unread_count":  unreadCount,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, `{"error": "Erreur lors de l'envoi des données"}`, http.StatusInternalServerError)
			return
		}
		return

	} else if queryUsername != "" {
		// Récupération par nom d'utilisateur
		notifications, err := services.GetNotificationsByUsername(db, queryUsername)
		if err != nil {
			http.Error(w, `{"error": "Erreur lors de la récupération des notifications"}`, http.StatusInternalServerError)
			return
		}

		// Compter les notifications non lues
		unreadCount, _ := services.GetUnreadNotifsCountByUsername(db, queryUsername)

		response := map[string]interface{}{
			"notifications": notifications,
			"total_count":   len(notifications),
			"unread_count":  unreadCount,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, `{"error": "Erreur lors de l'envoi des données"}`, http.StatusInternalServerError)
			return
		}
		return

	} else {
		// Récupération de toutes les notifications (pour admin)
		var notifications []models.Notif
		err := db.Order("created_at desc").Find(&notifications).Error
		if err != nil {
			http.Error(w, `{"error": "Erreur lors de la récupération des notifications"}`, http.StatusInternalServerError)
			return
		}

		// Compter toutes les notifications non lues
		unreadCount, _ := services.GetUnreadNotifsCount(db)

		response := map[string]interface{}{
			"notifications": notifications,
			"total_count":   len(notifications),
			"unread_count":  unreadCount,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, `{"error": "Erreur lors de l'envoi des données"}`, http.StatusInternalServerError)
			return
		}
		return
	}
}

// GetNotificationsByCurrentUser récupère les notifications de l'utilisateur connecté
// Utilise le token JWT pour identifier l'utilisateur
func GetNotificationsByCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// TODO: Extraire l'ID utilisateur du token JWT
	// userID := extractUserIDFromToken(r)
	// Pour l'instant, on utilise un paramètre ou on retourne une erreur
	
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, `{"error": "Token d'authentification requis"}`, http.StatusUnauthorized)
		return
	}

	// Ici vous devrez implémenter l'extraction de l'ID utilisateur depuis le token
	// userID := jwt.ExtractUserID(authHeader)
	
	http.Error(w, `{"error": "Authentification non implémentée"}`, http.StatusNotImplemented)
}

// MarkNotificationAsRead marque une notification comme lue
func MarkNotificationAsRead(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, `{"error": "Méthode non autorisée"}`, http.StatusMethodNotAllowed)
		return
	}

	db := database.GetDB()
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	notifIDStr := vars["id"]
	if notifIDStr == "" {
		http.Error(w, `{"error": "ID de notification requis"}`, http.StatusBadRequest)
		return
	}

	notifID, err := strconv.ParseUint(notifIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error": "ID de notification invalide"}`, http.StatusBadRequest)
		return
	}

	err = services.MarkNotifAsRead(db, uint(notifID))
	if err != nil {
		http.Error(w, `{"error": "Erreur lors de la mise à jour"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Notification marquée comme lue",
	})
}

// MarkAllNotificationsAsRead marque toutes les notifications d'un utilisateur comme lues
func MarkAllNotificationsAsRead(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, `{"error": "Méthode non autorisée"}`, http.StatusMethodNotAllowed)
		return
	}

	db := database.GetDB()
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userIDStr := vars["userID"]
	if userIDStr == "" {
		http.Error(w, `{"error": "ID utilisateur requis"}`, http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error": "ID utilisateur invalide"}`, http.StatusBadRequest)
		return
	}

	err = services.MarkAllNotifsAsRead(db, uint(userID))
	if err != nil {
		http.Error(w, `{"error": "Erreur lors de la mise à jour"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Toutes les notifications marquées comme lues",
	})
}

type NotificationRequest struct {
	Notifications []NotificationData `json:"notifications"`
}

// NotificationData structure pour une notification individuelle
type NotificationData struct {
	Destinataire  string `json:"destinataire"`
	Type          string `json:"type"`
	Titre         string `json:"titre"`
	Message       string `json:"message"`
	Semestre      string `json:"semestre,omitempty"`
	DateCreation  string `json:"date_creation,omitempty"`
	Lu            bool   `json:"lu,omitempty"`
}

// CreateOrganigrammeNotifications crée les notifications d'organigramme pour tous les utilisateurs
func CreateOrganigrammeNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Méthode non autorisée"}`, http.StatusMethodNotAllowed)
		return
	}

	db := database.GetDB()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Gérer les requêtes OPTIONS pour CORS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var request struct {
		NotifyAllUsers bool   `json:"notify_all_users"`
		Semestre       string `json:"semestre"`
		Type           string `json:"type"`
		Titre          string `json:"titre"`
		Message        string `json:"message"`
		Notifications  []NotificationData `json:"notifications,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error": "Données JSON invalides"}`, http.StatusBadRequest)
		return
	}

	var createdNotifications []models.Notif
	var errors []string

	// Si notify_all_users est true, envoyer à tous les utilisateurs
	if request.NotifyAllUsers {
		// Récupérer tous les utilisateurs
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			http.Error(w, `{"error": "Erreur lors de la récupération des utilisateurs"}`, http.StatusInternalServerError)
			return
		}

		// Valeurs par défaut si non spécifiées
		titre := request.Titre
		if titre == "" {
			titre = fmt.Sprintf("Organigramme %s validé", request.Semestre)
		}
		
		message := request.Message
		if message == "" {
			message = fmt.Sprintf("L'organigramme du semestre %s a été validé par le chef de département.", request.Semestre)
		}

		typeNotif := request.Type
		if typeNotif == "" {
			typeNotif = "organigramme_valide"
		}

		// Créer une notification pour chaque utilisateur
		for _, user := range users {
			notif := models.Notif{
				UserID:       user.ID,
				Destinataire: user.Nom, // ou user.Username selon votre modèle
				Type:         typeNotif,
				Titre:        titre,
				Message:      message,
				IsRead:       false,
				CreatedAt:    time.Now(),
			}

			if err := services.CreateNotif(db, &notif); err != nil {
				errors = append(errors, fmt.Sprintf("Erreur pour %s: %s", user.Nom, err.Error()))
				continue
			}

			createdNotifications = append(createdNotifications, notif)
		}

	} else {
		// Mode original : notifications spécifiques
		for _, notifData := range request.Notifications {
			// Validation des données
			if notifData.Destinataire == "" || notifData.Titre == "" || notifData.Message == "" {
				errors = append(errors, "Destinataire, titre et message sont requis")
				continue
			}

			// Créer la notification en utilisant le service
			err := services.CreateOrganigrammeNotification(
				db,
				notifData.Destinataire,
				notifData.Titre,
				notifData.Message,
				notifData.Semestre,
			)

			if err != nil {
				errors = append(errors, "Erreur lors de la création de notification pour "+notifData.Destinataire+": "+err.Error())
				continue
			}

			// Récupérer la notification créée pour la réponse
			var createdNotif models.Notif
			db.Where("destinataire = ? AND titre = ?", notifData.Destinataire, notifData.Titre).
				Order("created_at desc").First(&createdNotif)
			
			createdNotifications = append(createdNotifications, createdNotif)
		}
	}

	// Préparer la réponse
	response := map[string]interface{}{
		"success":               len(createdNotifications) > 0,
		"created_count":         len(createdNotifications),
		"total_requested":       len(request.Notifications),
		"created_notifications": createdNotifications,
	}

	if len(errors) > 0 {
		response["errors"] = errors
	}

	// Déterminer le code de statut
	statusCode := http.StatusCreated
	if len(errors) > 0 {
		if len(createdNotifications) == 0 {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusPartialContent
		}
	}

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "Erreur lors de l'envoi de la réponse"}`, http.StatusInternalServerError)
		return
	}
}

// CreateSingleNotification crée une seule notification
func CreateSingleNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Méthode non autorisée"}`, http.StatusMethodNotAllowed)
		return
	}

	db := database.GetDB()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var notifData NotificationData
	if err := json.NewDecoder(r.Body).Decode(&notifData); err != nil {
		http.Error(w, `{"error": "Données JSON invalides"}`, http.StatusBadRequest)
		return
	}

	// Validation
	if notifData.Destinataire == "" || notifData.Titre == "" || notifData.Message == "" {
		http.Error(w, `{"error": "Destinataire, titre et message sont requis"}`, http.StatusBadRequest)
		return
	}

	// Trouver l'utilisateur par nom d'utilisateur pour obtenir l'ID
	user, err := services.GetUserByName(db, notifData.Destinataire)
	var userID uint = 0
	if err == nil && user != nil {
		userID = user.ID
	}

	// Créer la notification
	notif := models.Notif{
		UserID:       userID,
		Destinataire: notifData.Destinataire,
		Type:         notifData.Type,
		Titre:        notifData.Titre,
		Message:      notifData.Message,
		IsRead:       false,
		CreatedAt:    time.Now(),
	}

	err = services.CreateNotif(db, &notif)
	if err != nil {
		http.Error(w, `{"error": "Erreur lors de la création de la notification"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":      true,
		"message":      "Notification créée avec succès",
		"notification": notif,
	})
}

// DeleteNotification supprime une notification
func DeleteNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, `{"error": "Méthode non autorisée"}`, http.StatusMethodNotAllowed)
		return
	}

	db := database.GetDB()
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	notifIDStr := vars["id"]
	if notifIDStr == "" {
		http.Error(w, `{"error": "ID de notification requis"}`, http.StatusBadRequest)
		return
	}

	notifID, err := strconv.ParseUint(notifIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error": "ID de notification invalide"}`, http.StatusBadRequest)
		return
	}

	err = services.DeleteNotif(db, uint(notifID))
	if err != nil {
		http.Error(w, `{"error": "Erreur lors de la suppression"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Notification supprimée avec succès",
	})
}