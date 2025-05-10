package database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dsn := "root:ilyesgamer2005@@tcp(127.0.0.1:3306)/gestion_universite?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Teacher{},
		&models.Module{},
		&models.Niveau{},
		&models.Voeux{},
		&models.Speciality{},
		&models.Commentaire{},
		&models.Notif{},
	)

	if err != nil {
		log.Fatal("Erreur de migration:", err)
	}
	
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Erreur lors de l'accès à la connexion SQL:", err)
	}


	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Impossible de se connecter à la base de données:", err)
	}

	fmt.Println("Connexion à la base de données MySQL réussie.")
}

// GetDB retourne la connexion à la base de données
func GetDB() *gorm.DB {
	return DB
}

// CloseDB ferme la connexion à la base de données
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Erreur lors de la fermeture de la base de données:", err)
	}
	sqlDB.Close()
}
