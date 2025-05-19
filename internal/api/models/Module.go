package models

type Module struct {
    ID       uint    `gorm:"primaryKey"json:"id"`
    Nom      string  `gorm:"not null"json:"nom"`
    ModuleNiveaux []ModuleNiveau
    Voeux    []Voeux `gorm:"foreignKey:ModuleID"`
}