package rooter

import (
    "Devenir_dev/internal/api/handlers"
    "github.com/gorilla/mux"

)

func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // Définition des routes
    router.HandleFunc("/login", handlers.Login).Methods("GET", "POST")
    router.HandleFunc("/submit", handlers.Submit).Methods("POST")
    router.HandleFunc("/home", handlers.Home).Methods("GET")
    router.HandleFunc("/home/profs", handlers.List).Methods("GET")
    router.HandleFunc("/deleteUser", handlers.DeleteUserHandler).Methods("POST")
    router.HandleFunc("/home/fiche-de-voeux", handlers.Fiche_de_voeux).Methods("GET", "POST")

    return router
}