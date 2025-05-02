package models    



type Niveau struct {
    ID          uint     `gorm:"primaryKey"`
    Spec        string   `gorm:"not null"`
    Modules     []Module `gorm:"many2many:module_niveaux;"`
    Voeux       []Voeux  `gorm:"foreignKey:NiveauID"`
}
