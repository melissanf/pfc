package models 
import()
type Affectation struct {
    ID         uint   `gorm:"primaryKey"`
    TeacherID  uint   `gorm:"not null"`
    ModuleID   uint   `gorm:"not null"`
    NiveauID   uint   `gorm:"not null"`
    TypeSeance string `gorm:"not null"`
    Groupe     int    `gorm:"not null"`
    OrganigrammeID  uint   `gorm:"not null"`
    Teacher    Teacher `gorm:"foreignKey:TeacherID"`
    Module     Module  `gorm:"foreignKey:ModuleID"`
    Niveau     Niveau  `gorm:"foreignKey:NiveauID"`
}

