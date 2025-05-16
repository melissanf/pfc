package models

type Teacher struct {
    ID            uint        `gorm:"primaryKey"`
    UserID        uint        `gorm:"not null;unique"`
    YearEntrance  string      `gorm:"unique;not null"`  // Année d'entrée
    Grade         string      `gorm:"not null"`
    ChargeHoraire float64     `gorm:"default:0"` 
    User          User        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    Specialities  []Speciality`gorm:"many2many:teacher_specialities;"`
    Voeux         []Voeux     `gorm:"foreignKey:TeacherID"`
}