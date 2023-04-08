package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GORMModel
	UserID   uint   `gorm:"not null" json:"user_id"`
	Title    string `gorm:"not null" json:"title" validate:"required-Title is required"`
	Caption  string `gorm:"not null" json:"caption" validate:"required-Caption is required"`
	PhotoURL string `gorm:"not null" json:"photo_url" validate:"required-Photo URL is required"`
	User     *User
	Comment  []Comment `json:"comments"`
}

func (p *SocialMedia) BeforeCreatePhoto(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *SocialMedia) BeforeUpdatePhoto(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
