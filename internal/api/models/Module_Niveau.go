package models 
import()
type ModuleNiveau struct {
    ID          uint   `gorm:"primaryKey"`
    ModuleID    uint   `gorm:"not null;unique"`
    NiveauID    uint   `gorm:"not null;unique"`
    NbCours     int    `gorm:"not null"`
    NbTD        int    `gorm:"not null"`
    NbTP        int    `gorm:"not null"`
    Module      Module `gorm:"foreignKey:ModuleID"`
    Niveau      Niveau `gorm:"foreignKey:NiveauID"`
}
