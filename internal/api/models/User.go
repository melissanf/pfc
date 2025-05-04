package models

type Role string

const (
    Admin      Role = "admin"
    Professeur Role = "professeur"
    Responsable Role = "responsable"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Nom      string `gorm:"not null"`
    Prenom   string `gorm:"not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Role     Role   `gorm:"type:enum('admin','professeur','responsable');default:'professeur';not null"`
}