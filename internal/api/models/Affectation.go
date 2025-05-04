package models 
import()
type Affectation struct {
    ID         uint   `gorm:"primaryKey"`
    TeacherID  uint   `gorm:"not null"`
    ModuleID   uint   `gorm:"not null"`
    NiveauID   uint   `gorm:"not null"`
    TypeSeance string `gorm:"not null"` // "Cours", "TD", "TP"
    Groupe     int    // 0 pour cours, 1-4 pour TD/TP

    Teacher    Teacher `gorm:"foreignKey:TeacherID"`
    Module     Module  `gorm:"foreignKey:ModuleID"`
    Niveau     Niveau  `gorm:"foreignKey:NiveauID"`
}

