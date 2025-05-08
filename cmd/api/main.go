package main

import (
	"fmt"
	"net/http"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/rooter"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/middleware"
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
	"log"
)

const port = ":8000"
func main (){
	database.InitDB()
	app := rooter.NewRouter()
	app.Use(middleware.JwtMiddleware)
	fmt.Println("(http://localhost"+port+"/login) le serveur est lancer sur ce lien ")
	
	err := http.ListenAndServe(port, app)
	if err != nil {
		log.Fatal(err)
	}
}
  



 