package models 
import()
type ModuleNiveau struct {
    ID          uint   `gorm:"primaryKey"`
    ModuleID    uint
    NiveauID    uint
    ChargeCours uint 
    ChargeTD    uint 
    ChargeTP    uint 

    Module      Module `gorm:"foreignKey:ModuleID"`
    Niveau      Niveau `gorm:"foreignKey:NiveauID"`
}
