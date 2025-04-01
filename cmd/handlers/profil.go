package handlers

import "net/http"
func HandelProfile(res http.ResponseWriter,req *http.Request){
	session, _ := store.Get(req, "session-name")
	username, ok := session.Values["username"].(string)
	email, ok := session.Values["email"].(string)
	if !ok || username == "" ||email== "" {
        http.Redirect(res, req, "/login", http.StatusFound) // Rediriger si l'utilisateur n'est pas connecté
        return
    }
	Data := User{
		Name: username,
	    Email: email,
  }
  Rendertemplates(res,"Profil",Data)
}