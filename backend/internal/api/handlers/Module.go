package handlers

import (
	"github.com/melissanf/pfc/backend/internal/database"
	"github.com/melissanf/pfc/backend/internal/api/models"
	"github.com/melissanf/pfc/backend/internal/api/services"
	"encoding/json"
	"net/http"
)

func GetAllModules(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var modules []models.Module

	if err := db.
		Preload("ModuleNiveaux.Niveau"). // ✅ Pour charger les niveaux liés
		Find(&modules).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(modules); err != nil {
		http.Error(w, "Erreur lors de l'envoi des données", http.StatusInternalServerError)
	}
}

func Createmodule(w http.ResponseWriter, r *http.Request) {
	var module models.Module
	if err := json.NewDecoder(r.Body).Decode(&module); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	db := database.GetDB() 

	if _,err :=services.Createmodule(db, &module); err != nil {
	   json.NewEncoder(w).Encode(module)
   	   return
    }
}