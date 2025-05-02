package models



type Voeux struct {
    ID        uint   `gorm:"primaryKey"`
    TeacherID uint   `gorm:"not null"`  // Un seul enregistrement de voeux par professeur
    Tp        bool   `gorm:"not null"`          // Indique si le professeur souhaite enseigner en TP
    Td        bool   `gorm:"not null"`          // Indique si le professeur souhaite enseigner en TD
    Cours     bool   `gorm:"not null"`          // Indique si le professeur souhaite enseigner en 
	Priority  int    `gorm:"not null"`          // Indique la priorité du voeux (1 étant le plus prioritaire)
    ModuleID  uint   `gorm:"not null"`          // Le module associé au voeux
    Module    Module `gorm:"foreignKey:ModuleID;constraint:OnDelete:CASCADE"`
    
    NiveauID  uint   `gorm:"not null"`          // Le niveau associé au voeux
    Niveau    Niveau `gorm:"foreignKey:NiveauID;constraint:OnDelete:CASCADE"`

    Teacher   Teacher `gorm:"foreignKey:TeacherID;constraint:OnDelete:CASCADE"`  // Lien avec le professeur
}