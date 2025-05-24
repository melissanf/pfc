package models

type TeacherSpeciality struct {
	ID         uint       `gorm:"primaryKey"`
	SpecialityID uint     `gorm:"not null" json:"speciality_id"`
	TeacherID  uint       `gorm:"not null" json:"teacher_id"`
	Speciality Speciality `gorm:"foreignKey:SpecialityID;constraint:OnDelete:CASCADE" json:"speciality"`	
	Teacher    Teacher    `gorm:"foreignKey:TeacherID;constraint:OnDelete:CASCADE" json:"teacher"` 
}