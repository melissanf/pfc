package handlers
import( 
    "github.com/ilyes-rhdi/Projet_s4/pkg"
	"net/http"
 
)


func Home(res http.ResponseWriter, req *http.Request){


	utils.Rendertemplates(res,"Home", nil)
		
}