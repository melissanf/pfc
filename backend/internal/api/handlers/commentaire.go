package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/melissanf/pfc/backend/internal/database"
	"github.com/melissanf/pfc/backend/internal/api/models"
	"github.com/melissanf/pfc/backend/internal/api/services"
)
func CreateCommentaire(w http.ResponseWriter, r *http.Request) {
	var commentaire models.Commentaire
	db := database.GetDB()
	if err := json.NewDecoder(r.Body).Decode(&commentaire); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _,err := services.CreateCommentaire(db, &commentaire); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(commentaire)
}
func DeleteCommentaire(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
    db := database.GetDB()
	if err := services.DeleteCommentaire(db,uint(id)) ;err != nil {
		http.Error(w, "Erreur lors de la suppression", http.StatusInternalServerError)
		return
	}
    
	w.WriteHeader(http.StatusNoContent)
}
func UpdateCommentaire(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
    db := database.GetDB()
	var commentaire models.Commentaire
	if err := db.First(&commentaire, id).Error; err != nil {
		http.Error(w, "Commentaire introuvable", http.StatusNotFound)
		return
	}

	var input struct {
		Contenu string `json:"contenu"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Corps invalide", http.StatusBadRequest)
		return
	}

	commentaire.Contenu = input.Contenu
	db.Save(&commentaire)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err :=json.NewEncoder(w).Encode(commentaire); err!=nil {	
		http.Error(w, "Erreur lors de l'envoi de la réponse", http.StatusInternalServerError)
		return
	}
}
