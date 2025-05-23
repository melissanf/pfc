package database

import (
	"fmt"
	"log"
	"os"
	"github.com/ilyes-rhdi/Projet_s4/internal/api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erreur de chargement du fichier .env :", err)
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN non défini dans .env")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données :", err)
	}

	modelsToMigrate := []interface{}{ 
		&models.User{},
		&models.Teacher{},
		&models.Module{},
		&models.Niveau{},
		&models.Voeux{},
		&models.Speciality{},
		&models.Commentaire{},
		&models.Notif{},
		&models.ModuleNiveau{},
	}
	for _, model := range modelsToMigrate {
    if !DB.Migrator().HasTable(model) {
        if err := DB.AutoMigrate(model); err != nil {
            log.Printf("⚠️ Erreur migration pour %T : %v", model, err)
        } else {
            log.Printf("✅ Table migrée : %T", model)
        }
    } else {
        log.Printf("📌 Table déjà existante : %T", model)
    }
}
	if err != nil {
		log.Fatal("Erreur lors de l'exécution de AutoMigrate :", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Erreur lors de l'accès à la connexion SQL :", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Impossible de se connecter à la base de données :", err)
	}

	fmt.Println("Connexion à la base de données PostgreSQL réussie.")
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
