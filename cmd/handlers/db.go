package handlers
import (
	"database/sql"
	"log"
	"fmt"
)
func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:ilyesgamer2005@@tcp(localhost:3306)/gestion_universite")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
}