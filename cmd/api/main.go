package main

import (
	"fmt"
	"net/http"
	"Devenir_dev/internal/api/rooter"
	"Devenir_dev/internal/database"
	"log"
)

const port = ":3000"
func main (){
	database.InitDB()
	app := rooter.NewRouter()
	fmt.Println("(http://localhost:3000/login) le serveur est lancer sur ce lien ")
	
	err := http.ListenAndServe(port, app)
	if err != nil {
		log.Fatal(err)
	}
}
  



 