package models

type Organigramme struct {
    ID         uint    `gorm:"primaryKey"`
    Annee      string  `gorm:"not null"`
    Semestre   string  `gorm:"not null"`
    IsValide   bool    `gorm:"default:false"`
    Affectations []Affectation `gorm:"many2many:organigramme_affectations"`
}