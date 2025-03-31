package handlers
import( 

	"net/http"
 
)


func Home(res http.ResponseWriter, req *http.Request){


	Rendertemplates(res,"Home", nil)
		
}