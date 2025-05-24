package models 
import()
type Affectation struct {
    ID         uint   `gorm:"primaryKey" json:"id"`
    TeacherID  uint   `gorm:"not null" json:"Teacherid"`
    ModuleID   uint   `gorm:"not null" json:"Moduleid"`
    NiveauID   uint   `gorm:"not null" json:"Niveauid"`
    TypeSeance string `gorm:"not null" json:"typeSeance"`
    Groupe     int    `gorm:"not null" json:"Groupe"`
    OrganigrammeID  uint   `gorm:"not null" json:"Orgaid"`
    Teacher    Teacher `gorm:"foreignKey:TeacherID"`
    Module     Module  `gorm:"foreignKey:ModuleID"`
    Niveau     Niveau  `gorm:"foreignKey:NiveauID"`
}

