package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/middleware"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/rooter"
	"github.com/ilyes-rhdi/Projet_s4/internal/database"
)

func main() {
	database.InitDB()
	app := rooter.NewRouter()
	app.Use(middleware.JwtMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Port par défaut localement
	}

    fmt.Println("http://localhost:8000/")
	err := http.ListenAndServe(":"+port, app)
	if err != nil {
		log.Fatal(err)
	}
}
  



 