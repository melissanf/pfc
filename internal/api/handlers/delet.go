package handlers
import ( 
	"net/http"


)
func DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
    http.Redirect(res, req, "/Home", http.StatusFound)
}
