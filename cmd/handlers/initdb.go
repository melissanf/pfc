package handlers
import (
	"net/http"
	"database/sql"
)
func Initdb(res http.ResponseWriter, req *http.Request){
	dsn := "root:ilyesgamer2005@@tcp(127.0.0.1:3306)/gestion_universite"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		http.Error(res, "Database connection error", http.StatusInternalServerError)
		return 
	}
	defer db.Close()
	
}