package handlers
import( 
    "Devenir_dev/pkg"
	"Devenir_dev/internal/api/models"
	"net/http"
 
)


func Home(res http.ResponseWriter, req *http.Request){


	utils.Rendertemplates(res,"Home", nil)
		
}