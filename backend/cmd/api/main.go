package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/melissanf/pfc/backend/internal/api/middleware"
	"github.com/melissanf/pfc/backend/internal/api/rooter"
	"github.com/melissanf/pfc/backend/internal/database"
	"github.com/rs/cors"
)

func main() {
	database.InitDB()
	app := rooter.NewRouter()
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.JwtMiddleware)

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("http://localhost:8000/")
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
