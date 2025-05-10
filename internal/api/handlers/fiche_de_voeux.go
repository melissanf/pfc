package handlers

import (
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/services"
	"github.com/ilyes-rhdi/Projet_s4/pkg"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)


func Fiche_de_voeux(res http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(res, "Token manquant", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		http.Error(res, "Token invalide", http.StatusUnauthorized)
		return
	}

	profIDFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		http.Error(res, "ID utilisateur non trouvé", http.StatusUnauthorized)
		return
	}
	profID := uint(profIDFloat)

	db := database.GetDB()

	count, err := services.CountVoeuxByTeacherID(db, profID)
	if err != nil {
		http.Error(res, "Erreur lors de la vérification", http.StatusInternalServerError)
		return
	}
	if count >= 3 {
		http.Error(res, "Vous avez déjà soumis 3 vœux.", http.StatusForbidden)
		return
	}

	for i := 1; i <= 3; i++ {
		moduleName := req.FormValue(fmt.Sprintf("module_name_%d", i))
		if moduleName == "" {
			continue
		}
		niveauIDStr := req.FormValue(fmt.Sprintf("niveau_id_%d", i))
		niveauIDUint, err := strconv.ParseUint(niveauIDStr, 10, 64)
		if err != nil || niveauIDUint == 0 {
			http.Error(res, "Niveau invalide", http.StatusBadRequest)
			return
		}
		niveauID := uint(niveauIDUint)

		isTP := utils.FormBool(req, fmt.Sprintf("tp_%d", i))
		isTD := utils.FormBool(req, fmt.Sprintf("td_%d", i))
		isCour := utils.FormBool(req, fmt.Sprintf("cour_%d", i))
		priority := i

		module, err := services.GetModuleByName(db, moduleName)
		if err != nil {
			http.Error(res, fmt.Sprintf("Module introuvable : %s", moduleName), http.StatusNotFound)
			return
		}
		exists, err := services.VoeuxExactExists(db, profID, module.ID, niveauID, isTP, isTD, isCour)
		if err != nil {
			http.Error(res, "Erreur lors de la vérification des doublons", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(res, fmt.Sprintf("Un vœu identique existe déjà pour le module '%s' et niveau spécifié.", moduleName), http.StatusBadRequest)
			return
		}

		voeux := &models.Voeux{
			TeacherID: profID,
			ModuleID:  module.ID,
			NiveauID:  niveauID,
			Tp:        isTP,
			Td:        isTD,
			Cours:     isCour,
			Priority:  priority,
		}
		if err := services.CreateVoeux(db, voeux); err != nil {
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