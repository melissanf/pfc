package models



type Voeux struct {
    ID        uint   `gorm:"primaryKey"`
    TeacherID uint   `gorm:"not null"`  
    Tp        bool   `gorm:"not null"`        
    Td        bool   `gorm:"not null"`         
    Cours     bool   `gorm:"not null"`          
	Priority  int    `gorm:"not null"`          
    ModuleID  uint   `gorm:"not null"`          
    Module    Module `gorm:"foreignKey:ModuleID;constraint:OnDelete:CASCADE"`
    NiveauID  uint   `gorm:"not null"`          
    Niveau    Niveau `gorm:"foreignKey:NiveauID;constraint:OnDelete:CASCADE"`
    Teacher   Teacher `gorm:"foreignKey:TeacherID;constraint:OnDelete:CASCADE"`  
}