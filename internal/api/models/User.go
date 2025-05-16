package models

type Role string

const (
    Admin      Role = "admin"
    Professeur Role = "professeur"
    Responsable Role = "responsable"
)

type User struct {
    ID       uint   `gorm:"primaryKey"json:"id"`
    Nom      string `gorm:"not null"json:"nom"`
    Prenom   string `gorm:"not null"json:"prenom"`
    Email    string `gorm:"unique;not null"json:"email"`
    Password string `gorm:"not null"json:"password"`
    Numero   string `gorm:"not null"json:"numero"`
    Role     Role   `gorm:"type:text;default:'professeur';not null"json:"role"`
}