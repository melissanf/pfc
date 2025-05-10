package models

import (
	"time"
)
type Notif struct {
	ID             uint      `gorm:"primaryKey"`
	CommentaireID  uint      `gorm:"not null"`
	Message       string    `gorm:"not null"` // probablement "objet_type" ?
	IsRead         bool      `gorm:"default:false"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}