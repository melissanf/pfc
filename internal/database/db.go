package database
import (
	"database/sql"
	"log"
	 
)
var db *sql.DB

func InitDB() {
	var err error
	dsn := "root:ilyesgamer2005@@tcp(127.0.0.1:3306)/gestion_universite"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	return  &db
}

func GetDB() *sql.DB {
	return db
}