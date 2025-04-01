package handlers
import (
	"net/http"
    "Devenir_dev/cmd/database"
)
func List(res http.ResponseWriter, req *http.Request){
	session, _ := store.Get(req, "session-name")
    db := database.GetDB()
	users, err := GetAllUsers(db)
    if err != nil {
        http.Error(res, "Error fetching users", http.StatusInternalServerError)
        return
    }
	username, ok := session.Values["username"].(string)
    isAdmin, _ := session.Values["isAdmin"].(bool)
    if !ok || username == "" {
        http.Redirect(res, req, "/login", http.StatusFound) // Rediriger si l'utilisateur n'est pas connecté
        return
    }
	data := Pagedata{
        Currentuser: User{
            Name: username,
            Isadmin: isAdmin,
        },
        Users: users,
    }
    Rendertemplates(res,"Home/profs", data)
}