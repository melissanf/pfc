package handlers

import (
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllUsers ➡️ GET /users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db := database.GetDB() // Obtenir la connexion DB

	if err := db.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// GetUserByID ➡️ GET /users/{id}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	var user models.User
	db := database.GetDB() // Obtenir la connexion DB

	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// CreateUser ➕ POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	db := database.GetDB() // Obtenir la connexion DB

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// UpdateUser ✏️ PUT /users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	var user models.User
	db := database.GetDB() // Obtenir la connexion DB

	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	var updatedData models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	db.Model(&user).Updates(updatedData)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser 🗑️ DELETE /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	db := database.GetDB() // Obtenir la connexion DB

	if err := db.Delete(&models.User{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
