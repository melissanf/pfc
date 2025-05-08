package models    



type Niveau struct {
    ID          uint     `gorm:"primaryKey"`
    Spec        string   `gorm:"not null"`
    Annee       string   `gorm:"not null"`
    Section     string   `gorm:"not null"`
    ModuleNiveaux []ModuleNiveau
    Voeux       []Voeux  `gorm:"foreignKey:NiveauID"`
}
