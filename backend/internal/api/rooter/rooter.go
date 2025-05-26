package rooter

import (
	"github.com/melissanf/pfc/backend/internal/api/handlers"
	"github.com/melissanf/pfc/backend/internal/api/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()

	// ---------- ROUTES PUBLIQUES ----------
	router.HandleFunc("/", handlers.Home).Methods("GET")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout).Methods("POST")
    router.HandleFunc("/submit", handlers.Submit).Methods("GET","POST")
    //Routes pour lmes enseignants
    ProfRouter := router.PathPrefix("/Enseignant").Subrouter()
    ProfRouter.HandleFunc("/profile", handlers.HandelProfile).Methods("GET")
    ProfRouter.HandleFunc("/fiche-de-voeux", handlers.Fiche_de_voeux).Methods("POST")
    ProfRouter.HandleFunc("/fiche-de-voeux", handlers.GetVoeuxByTeacherID).Methods("GET")
    ProfRouter.HandleFunc("/modules", handlers.GetAllModules).Methods("GET")
    ProfRouter.HandleFunc("/notifications", handlers.GetNotifications).Methods("GET")
    ProfRouter.HandleFunc("/notifications/user/{userID}", handlers.GetNotifications).Methods("GET")
    ProfRouter.HandleFunc("/notifications/{id}/read", handlers.MarkNotificationAsRead).Methods("PUT", "PATCH")
    ProfRouter.HandleFunc("/notifications/user/{userID}/read-all", handlers.MarkAllNotificationsAsRead).Methods("PUT", "PATCH")
    ProfRouter.HandleFunc("/notifications/{id}", handlers.DeleteNotification).Methods("DELETE")
    ProfRouter.HandleFunc("/notifications", handlers.CreateSingleNotification).Methods("POST")
    ProfRouter.HandleFunc("/Notif", handlers.CreateOrganigrammeNotifications).Methods("POST") 

    staffRouter := router.PathPrefix("/staff").Subrouter()
    staffRouter.HandleFunc("/commentaire", handlers.CreateCommentaire).Methods("POST")

    // Routes protégées par JWT et rôle admin pour les utilisateurs
    adminRouter := router.PathPrefix("/admin").Subrouter()
    adminRouter.Use(middleware.IsAdmin) 

    // Routes pour les utilisateurs, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
    adminRouter.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    adminRouter.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    adminRouter.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    // Routes pour les enseignants, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/teachers/list", handlers.GetAllTeachers).Methods("GET")
    adminRouter.HandleFunc("/teachers", handlers.CreateTeacher).Methods("POST")
    adminRouter.HandleFunc("/teachers/{id}", handlers.UpdateTeacher).Methods("PUT")
    adminRouter.HandleFunc("/teachers/{id}", handlers.DeleteTeacher).Methods("DELETE")

    // Routes pour les voeux, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/voeux/list", handlers.GetAllVoeux).Methods("GET")
    adminRouter.HandleFunc("/voeux", handlers.CreateVoeux).Methods("POST")
    adminRouter.HandleFunc("/voeux/{id}", handlers.UpdateVoeux).Methods("PUT")
    adminRouter.HandleFunc("/voeux/{id}", handlers.DeleteVoeux).Methods("DELETE")

    // Routes pour les commentaire, accessibles uniquement par l'admin
    adminRouter.HandleFunc("/commentaire", handlers.CreateCommentaire).Methods("POST")
    adminRouter.HandleFunc("/commentaire/{id}", handlers.DeleteCommentaire).Methods("DELETE")
    adminRouter.HandleFunc("/commentaire/{id}", handlers.UpdateCommentaire).Methods("PUT")
    adminRouter.HandleFunc("/Notify", handlers.GetNotifications).Methods("GET")
   
    return router
}
