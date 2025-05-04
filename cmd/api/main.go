package main

import (
	"fmt"
	"net/http"
	"Devenir_dev/internal/api/rooter"
	"Devenir_dev/internal/api/middleware"
	"Devenir_dev/internal/database"
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
  



 