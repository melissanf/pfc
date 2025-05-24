package models

type Speciality struct {
    ID       uint      `gorm:"primaryKey"`
    Nom      string    `gorm:"not null"`
    Teachers []Teacher `gorm:"many2many:teacher_specialities;"`
}