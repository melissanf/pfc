package rooter

import (
    "Devenir_dev/internal/api/handlers"
    "Devenir_dev/internal/api/middleware" // Assurez-vous que le bon chemin vers le middleware est utilisé
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // Routes accessibles à tous
    router.HandleFunc("/login", handlers.Login).Methods("GET", "POST")
    router.HandleFunc("/submit", handlers.Submit).Methods("POST")
    router.HandleFunc("/home", handlers.Home).Methods("GET")
    router.HandleFunc("/home/profs", handlers.List).Methods("GET")
    router.HandleFunc("/home/fiche-de-voeux", handlers.Fiche_de_voeux).Methods("GET", "POST")

    // Routes protégées par JWT et rôle admin pour les utilisateurs
    adminRouter := router.PathPrefix("/admin").Subrouter()
    adminRouter.Use(middleware.IsAdmin)

    // Routes pour les utilisateurs, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
    adminRouter.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
    adminRouter.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    adminRouter.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    adminRouter.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    // Routes pour les enseignants, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/teachers", handlers.GetAllTeachers).Methods("GET")
    adminRouter.HandleFunc("/teachers/{id}", handlers.GetTeacherByID).Methods("GET")
    adminRouter.HandleFunc("/teachers", handlers.CreateTeacher).Methods("POST")
    adminRouter.HandleFunc("/teachers/{id}", handlers.UpdateTeacher).Methods("PUT")
    adminRouter.HandleFunc("/teachers/{id}", handlers.DeleteTeacher).Methods("DELETE")

    // Routes pour les voeux, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/voeux", handlers.GetAllVoeux).Methods("GET")
    adminRouter.HandleFunc("/voeux/{id}", handlers.GetVoeuxByID).Methods("GET")
    adminRouter.HandleFunc("/voeux", handlers.CreateVoeux).Methods("POST")
    adminRouter.HandleFunc("/voeux/{id}", handlers.UpdateVoeux).Methods("PUT")
    adminRouter.HandleFunc("/voeux/{id}", handlers.DeleteVoeux).Methods("DELETE")

    // Route pour les modules (accessible par les utilisateurs connectés)
    router.HandleFunc("/modules", handlers.GetAllModules).Methods("GET")
    return router
}
