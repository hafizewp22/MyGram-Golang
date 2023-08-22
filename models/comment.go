package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GORMModel
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Name    string `gorm:"not null" json:"name" validate:"required-Name is required"`
	Message string `gorm:"not null" json:"message" validate:"required-Message is required"`
	User    *APIUser
	Photo   *Photo
}

func (p *SocialMedia) BeforeCreateComment(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *SocialMedia) BeforeUpdateComment(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
