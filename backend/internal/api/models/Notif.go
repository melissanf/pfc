package models

import (
	"time"
)
type Notif struct {
    ID           uint   `json:"id" gorm:"primaryKey"`
    UserID       uint   `gorm:"not null" json:"user_id"`
    Destinataire string `gorm:"not null" json:"destinataire"`
    Type         string `gorm:"not null"  json:"type"`
    Titre        string `gorm:"not null"  json:"titre"`
    Message      string `gorm:"not null" json:"message"`
    IsRead       bool   `json:"is_read" gorm:"default:false"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
	User 		User   `gorm:"foreignKey:UserID"` // Association avec le modèle User
}