package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/melissanf/pfc/backend/internal/api/models"
	"github.com/melissanf/pfc/backend/internal/api/services"
	"github.com/melissanf/pfc/backend/internal/database"
)


func Fiche_de_voeux(res http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(res, "Token manquant", http.StatusUnauthorized)
		return
	}

	claims, ok := req.Context().Value("user").(*models.Claims)
	if !ok {
		http.Error(res, "Erreur de récupération des claims", http.StatusInternalServerError)
		return
	}
	db := database.GetDB()
	teacher , err :=services.GetTeacherByUserID(db, claims.UserID)
	if err != nil {	
		http.Error(res, "Erreur lors de la récupération du professeur", http.StatusInternalServerError)
		return	
	}
	profID := teacher.ID
	count, err := services.CountVoeuxByTeacherID(db, profID)
	if err != nil {
		http.Error(res, "Erreur lors de la vérification", http.StatusInternalServerError)
		return
	}
	if count > 3 {
		http.Error(res, "Vous avez déjà soumis 3 vœux.", http.StatusForbidden)
		return
	}
    type ModuleData struct {
		ModuleName string `json:"module_name"`
		NiveauName   string   `json:"niveau_name"`
		TP         bool   `json:"tp"`
		TD         bool   `json:"td"`
		Cour       bool   `json:"cour"`
        Heursupp   bool   `json:"hr"`
	}
	var modules []ModuleData
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&modules)
	if err != nil {
		http.Error(res, "JSON invalide", http.StatusBadRequest)
		return
	}
	for i, Module := range modules {
        if Module.ModuleName == "" || Module.NiveauName ==""{
			http.Error(res, fmt.Sprintf("donner recus vide  "), http.StatusNotFound)
			return
		}
		priority := i+1
		module, err := services.GetModuleByName(db, Module.ModuleName)
		if err != nil {
			http.Error(res, fmt.Sprintf("Module introuvable : %s", Module.ModuleName), http.StatusNotFound)
			return
		}
		niveau, err := services.GetNiveauBySpecAnnee(db,Module.NiveauName)
		if err != nil {
			http.Error(res, fmt.Sprintf("Niveau introuvable : %s", Module.NiveauName), http.StatusNotFound)
			return
		}
		exists, err := services.VoeuxExactExists(db, profID, module.ID, niveau.ID, Module.TP, Module.TD, Module.Cour)
		if err != nil {
			http.Error(res, "Erreur lors de la vérification des doublons", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(res, fmt.Sprintf("Un vœu identique existe déjà pour le module '%s' et niveau spécifié.", Module.ModuleName), http.StatusBadRequest)
			return
		}
        if Module.Heursupp{
			priority = -1
		}
		voeux := &models.Voeux{
			TeacherID: profID,
			ModuleID:  module.ID,
			NiveauID:  niveau.ID,
			Tp:        Module.TP,
			Td:        Module.TD,
			Cours:     Module.Cour,
			Priority:  priority,
		}
        
		if err := services.CreateVoeux(db, voeux); err != nil {
			fmt.Println(err)
			http.Error(res, "Erreur lors de l'insertion", http.StatusInternalServerError)
			return
		}
	}
    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	if err :=json.NewEncoder(res).Encode(map[string]string{"message": "Les vœux ont bien été enregistrés"}); err!=nil {	
		http.Error(res, "Erreur lors de l'envoi des données", http.StatusInternalServerError)
		return
	}
}
func GetAllVoeux(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var voeux []models.Voeux
	if err := db.Preload("Teacher.User").Preload("Module").Preload("Niveau").Find(&voeux).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err:=json.NewEncoder(w).Encode(voeux); err!=nil {	
		http.Error(w, "Erreur lors de l'envoi des données", http.StatusInternalServerError)
		return
	}
}

func GetVoeuxByID(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var voeux models.Voeux
	if err := db.Preload("Teacher.User").Preload("Module").Preload("Niveau").First(&voeux, id).Error; err != nil {
		http.Error(w, "Vœu non trouvé", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err:=json.NewEncoder(w).Encode(voeux); err!=nil {	
		http.Error(w, "Erreur lors de l'envoi des données", http.StatusInternalServerError)
		return
	}
}

func GetVoeuxByTeacherID(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	// Récupération des claims (authentification)
	claims, ok := r.Context().Value("user").(*models.Claims)
	if !ok || claims == nil {
		http.Error(w, "Utilisateur non authentifié", http.StatusUnauthorized)
		return
	}

	// Récupération de l'objet Teacher lié à l'utilisateur connecté
	teacher, err := services.GetTeacherByUserID(db, claims.UserID)
	if err != nil {
		http.Error(w, "Professeur introuvable", http.StatusNotFound)
		return
	}

	// Récupération des vœux liés à ce professeur
	var voeux []models.Voeux
	if err := db.Preload("Teacher.User").
		Preload("Module").
		Preload("Niveau").
		Where("teacher_id = ?", teacher.ID).
		Find(&voeux).Error; err != nil || len(voeux) == 0 {
		http.Error(w, "Aucun vœu trouvé pour ce professeur", http.StatusNotFound)
		return
	}

	// Envoi de la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(voeux); err != nil {
		http.Error(w, "Erreur lors de l'envoi des données", http.StatusInternalServerError)
	}
}

func CreateVoeux(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var voeux models.Voeux
	if err := json.NewDecoder(r.Body).Decode(&voeux); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	if err := db.Create(&voeux).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err:=json.NewEncoder(w).Encode(voeux); err!=nil {	
		http.Error(w, "Erreur lors de l'envoi des données", http.StatusInternalServerError)
		return
	}
}

func UpdateVoeux(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var voeux models.Voeux
	if err := db.First(&voeux, id).Error; err != nil {
		http.Error(w, "Vœu non trouvé", http.StatusNotFound)
		return
	}

	var updated models.Voeux
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	db.Model(&voeux).Updates(updated)
	w.Header().Set("Content-Type", "application/json")
	if err:=json.NewEncoder(w).Encode(voeux); err!=nil {	
		http.Error(w, "Erreur lors de l'envoi des données", http.StatusInternalServerError)
		return
	}
}


func DeleteVoeux(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := db.Delete(&models.Voeux{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}