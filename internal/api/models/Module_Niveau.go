package models 
import()
type ModuleNiveau struct {
    ID          uint   `gorm:"primaryKey"`
    ModuleID    uint
    NiveauID    uint
    NbCours     int 
    NbTD        int 
    NbTP        int 
    Module      Module `gorm:"foreignKey:ModuleID"`
    Niveau      Niveau `gorm:"foreignKey:NiveauID"`
}
