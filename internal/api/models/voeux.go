package models



type Voeux struct {
    ID        uint   `gorm:"primaryKey"json:"id"`
    TeacherID uint   `gorm:"not null"json:"Teacher_id"`  
    Tp        bool   `gorm:"not null"json:"tp"`        
    Td        bool   `gorm:"not null"json:"td"`         
    Cours     bool   `gorm:"not null"json:"cours"`          
	Priority  int    `gorm:"not null"`          
    ModuleID  uint   `gorm:"not null"json:"Module_id"`          
    Module    Module `gorm:"foreignKey:ModuleID;constraint:OnDelete:CASCADE"`
    NiveauID  uint   `gorm:"not null"json:"Niveau_id"`          
    Niveau    Niveau `gorm:"foreignKey:NiveauID;constraint:OnDelete:CASCADE"`
    Teacher   Teacher `gorm:"foreignKey:TeacherID;constraint:OnDelete:CASCADE"`  
}