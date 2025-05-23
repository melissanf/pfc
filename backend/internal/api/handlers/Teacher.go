package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/services"
	"github.com/ilyes-rhdi/Projet_s4/internal/database"

	"github.com/gorilla/mux"
)

func GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	var teachers []models.Teacher
	db := database.GetDB() 

	if err := db.Preload("User").Preload("Specialities").Find(&teachers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(teachers)
}

func GetTeacherByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var teacher models.Teacher
	db := database.GetDB()

	if err := db.Preload("User").Preload("Specialities").First(&teacher, id).Error; err != nil {
		http.Error(w, "Professeur non trouvé", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	db := database.GetDB() 

	if err :=services.CreateTeacher(db, &teacher); err != nil {
	   json.NewEncoder(w).Encode(teacher)
   	   return
    }
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var teacher models.Teacher
	db := database.GetDB() 

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

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	db := database.GetDB()

	if err := db.Delete(&models.Teacher{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
