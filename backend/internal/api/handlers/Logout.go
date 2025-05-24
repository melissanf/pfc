package handlers
import (
	"net/http")
func Logout(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
	w.Write([]byte("Déconnexion réussie"))
}