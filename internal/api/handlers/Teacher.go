package handlers

import (
	"Devenir_dev/internal/api/models"
	"Devenir_dev/internal/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 🔍 GET /teachers
func GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	var teachers []models.Teacher
	db := database.GetDB() // Obtenir la connexion DB

	if err := db.Preload("User").Preload("Specialities").Find(&teachers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(teachers)
}

// 👤 GET /teachers/{id}
func GetTeacherByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var teacher models.Teacher
	db := database.GetDB() // Obtenir la connexion DB

	if err := db.Preload("User").Preload("Specialities").First(&teacher, id).Error; err != nil {
		http.Error(w, "Professeur non trouvé", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

// ➕ POST /teachers
func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	db := database.GetDB() // Obtenir la connexion DB

	if err := db.Create(&teacher).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

// ✏️ PUT /teachers/{id}
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var teacher models.Teacher
	db := database.GetDB() // Obtenir la connexion DB

	if err := db.First(&teacher, id).Error; err != nil {
		http.Error(w, "Professeur non trouvé", http.StatusNotFound)
		return
	}

	var updated models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	db.Model(&teacher).Updates(updated)
	json.NewEncoder(w).Encode(teacher)
}

// 🗑️ DELETE /teachers/{id}
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	db := database.GetDB() // Obtenir la connexion DB

	if err := db.Delete(&models.Teacher{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
