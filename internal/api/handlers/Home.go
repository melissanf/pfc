package handlers
import( 
    "Devenir_dev/pkg"
	"net/http"
 
)


func Home(res http.ResponseWriter, req *http.Request){


	utils.Rendertemplates(res,"Home", nil)
		
}