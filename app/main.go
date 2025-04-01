package main

import (
	"fmt"
	"net/http"
	"Devenir_dev/cmd/handlers"
	"Devenir_dev/cmd/database"
	"log"
	"github.com/gorilla/mux"
)

const port = ":3000"
func main (){
	database.InitDB()
	app := mux.NewRouter()
	app.HandleFunc("/login", handlers.Login)
	app.HandleFunc("/Submit", handlers.Submit)
	app.HandleFunc("/Home", handlers.Home)
	app.HandleFunc("/Home/profs", handlers.List)
	app.HandleFunc("/deleteUser", handlers.DeleteUserHandler)
	app.HandleFunc("/Home/Fiche de voeux",handlers.Fiche_de_voeux)
	fmt.Println("(http://localhost:3000/login) le serveur est lancer sur ce lien ")
	
	err := http.ListenAndServe(port, app)
	if err != nil {
		log.Fatal(err)
	}
}
  



 