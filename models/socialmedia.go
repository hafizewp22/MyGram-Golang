package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GORMModel
	UserID         uint   `gorm:"not null" json:"user_id"`
	Name           string `gorm:"not null" json:"name" validate:"required-Name is required"`
	SosialMediaURL string `gorm:"not null" json:"social_media_url" validate:"required-Sosial Media URL is required"`
	User           *User
}

func (p *SocialMedia) BeforeCreateSocialMedia(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *SocialMedia) BeforeUpdateSocialMedia(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
