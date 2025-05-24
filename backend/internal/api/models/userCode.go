package models

type UserCode struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Code   string `gorm:"not null;unique" json:"code"`
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
}
