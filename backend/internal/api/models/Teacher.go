package models

type Teacher struct {
    ID            uint        `gorm:"primaryKey"`
    UserID        uint        `gorm:"not null;unique"`
    ChargeHoraire float64     `gorm:"default:0"` 
    Heursupp      float64     `gorm:"default:0"`
    Year_entrance int         `gorm:"not null"`
    Grade         string      `gorm:"not null"`
    User          User        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    Specialities  []Speciality `gorm:"many2many:teacher_specialities;"`
    Voeux         []Voeux     `gorm:"foreignKey:TeacherID"`
}