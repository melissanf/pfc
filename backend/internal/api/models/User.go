package models

type Role string

const (
	ChefDepartement Role = "chefDepartement"
	Personnel       Role = "staffAdministrateur"
	Enseignant      Role = "enseignant"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Nom      string `gorm:"not null" json:"nom"`
	Prenom   string `gorm:"not null" json:"prenom"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Numero   string `gorm:"not null" json:"numero"`
	Code     string `gorm:"not null" json:"code"`
	Role     Role   `gorm:"type:text;default:'Enseignant';not null" json:"role"`
}
