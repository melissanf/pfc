package models
import (
	"time"
)
type Commentaire struct {
    ID            uint      `gorm:"primaryKey"`
    Contenu       string    `gorm:"type:text;not null"`
    AuteurID      uint      `gorm:"not null"`  // Utilisateur qui a commenté
    CreatedAt     time.Time `gorm:"autoCreateTime"`
    UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}