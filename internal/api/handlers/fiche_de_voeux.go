package handlers

import (
	"Devenir_dev/internal/database"
	"Devenir_dev/pkg"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func Fiche_de_voeux(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		utils.Rendertemplates(res, "Fiche", nil)
		return
	}

	if req.Method != http.MethodPost {
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Lire et vérifier le JWT depuis l'en-tête Authorization
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(res, "Token manquant", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // Remplace par ta clé secrète
	})
	if err != nil || !token.Valid {
		http.Error(res, "Token invalide", http.StatusUnauthorized)
		return
	}

	profIDFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		http.Error(res, "Identifiant utilisateur non trouvé", http.StatusUnauthorized)
		return
	}
	profID := int(profIDFloat) // car les claims sont en float64

	// Traitement base de données
	db := database.GetDB()

	// Vérifier si le prof a déjà 3 choix enregistrés
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM fiche WHERE prof_id = ?", profID).Scan(&count)
	if err != nil {
		http.Error(res, "Erreur lors de la vérification", http.StatusInternalServerError)
		return
	}
	if count >= 3 {
		http.Error(res, "Vous avez déjà soumis 3 choix.", http.StatusForbidden)
		return
	}

	// Continue ton traitement ici...
}
