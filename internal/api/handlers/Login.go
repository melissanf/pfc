package handlers

import (
	"Devenir_dev/cmd/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
    "Devenir_dev/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
)
const SecretKey="iliii"
var store = sessions.NewCookieStore([]byte(SecretKey)) // Clé secrète pour sécurise

func Login(res http.ResponseWriter, req *http.Request) {
    if req.Method == http.MethodGet {
        utils.Rendertemplates(res, "Login",nil)
        return
    }
    db := database.GetDB()

    if req.Method == http.MethodPost {
        // Parse the form data
        if err := req.ParseForm(); err != nil {
            http.Error(res, "Error parsing form data", http.StatusBadRequest)
            return
        }

        // Extract data from the form
        identifier := req.FormValue("identifier")
        password := req.FormValue("password")



        // Verify if the user exists and the credentials are correct
        verified, isAdmin, message := utils.VerifyUser(db, identifier, password)
        if !verified {
            http.Error(res, message, http.StatusUnauthorized)
            return
        }
        var query string
        var name string
        if verified {
        session, _ := store.Get(req, "session-name")
        if strings.Contains(identifier, "@") {
            query = "SELECT name FROM users WHERE email = ?"
            db.QueryRow(query, identifier).Scan(&name)
            session.Values["username"]= name
        } else {
            session.Values["username"]= identifier
        } 
    
        session.Values["isAdmin"] = isAdmin
        session.Save(req, res) 
       }
       claims:=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
        Issuer : strconv.Itoa(int(user.id)),
        ExpirAt: time.Now().Add(time.Hour*24).Unix(),
})
       token,err:= claims.SignedString([]byte(SecretKey))
       if err != nil {

       }
       return app.Json(token)
        // Redirect based on admin status
            http.Redirect(res, req, "/Home", http.StatusFound)
       

        fmt.Fprintf(res, "Login successful. Welcome %s!\n", identifier)
    }
}



