package models

type Module struct {
    ID       uint    `gorm:"primaryKey"`
    Nom      string  `gorm:"not null"`
    ModuleNiveaux []ModuleNiveau
    Voeux    []Voeux `gorm:"foreignKey:ModuleID"`
}