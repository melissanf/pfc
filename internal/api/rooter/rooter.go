package rooter

import (
	"github.com/ilyes-rhdi/Projet_s4/internal/api/handlers"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()

	// ---------- ROUTES PUBLIQUES ----------
	router.HandleFunc("/", handlers.Home).Methods("GET")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout).Methods("POST")
    router.HandleFunc("/submit", handlers.Submit).Methods("GET","POST")
    router.HandleFunc("/home", handlers.Home).Methods("GET")
    router.HandleFunc("/home/profile", handlers.HandelProfile).Methods("GET")
    router.HandleFunc("/home/fiche-de-voeux", handlers.Fiche_de_voeux).Methods("GET", "POST")
    // Routes protégées par JWT et rôle admin pour les utilisateurs
    adminRouter := router.PathPrefix("/admin").Subrouter()
    adminRouter.Use(middleware.IsAdmin)
    // Routes pour les utilisateurs, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
    adminRouter.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    adminRouter.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    adminRouter.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    // Routes pour les enseignants, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/teachers", handlers.GetAllTeachers).Methods("GET")
    adminRouter.HandleFunc("/teachers", handlers.CreateTeacher).Methods("POST")
    adminRouter.HandleFunc("/teachers/{id}", handlers.UpdateTeacher).Methods("PUT")
    adminRouter.HandleFunc("/teachers/{id}", handlers.DeleteTeacher).Methods("DELETE")

    // Routes pour les voeux, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/voeux", handlers.GetAllVoeux).Methods("GET")
    adminRouter.HandleFunc("/voeux", handlers.CreateVoeux).Methods("POST")
    adminRouter.HandleFunc("/voeux/{id}", handlers.UpdateVoeux).Methods("PUT")
    adminRouter.HandleFunc("/voeux/{id}", handlers.DeleteVoeux).Methods("DELETE")

    // Route pour les modules (accessible par les utilisateurs connectés)
    router.HandleFunc("/modules", handlers.GetAllModules).Methods("GET")
    return router
}
