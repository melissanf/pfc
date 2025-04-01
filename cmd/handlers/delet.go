package handlers
import ( 
	"net/http"
    "Devenir_dev/cmd/database"

)
func DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodPost {
        http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    db := database.GetDB()
    // Récupérer le nom de l'utilisateur à supprimer à partir des paramètres de la requête
    username := req.FormValue("username")
    if username == "" {
        http.Error(res, "Username not provided", http.StatusBadRequest)
        return
    }

    // Supprimer l'utilisateur
    err := DeleteUser(db, username)
    if err != nil {
        http.Error(res, "Error deleting user", http.StatusInternalServerError)
        return
    }

    http.Redirect(res, req, "/Home", http.StatusFound)
}
