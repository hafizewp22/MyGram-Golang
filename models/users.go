package models

import (
	"fmt"
	"project_final/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GORMModel
	Username string        `gorm:"not null" json:"username" validate:"required-Username is required"`
	Email    string        `gorm:"not null;unique;type:varchar(100)" json:"email" validate:"required-Email is required,email-Invalid email format"`
	Password string        `gorm:"not null" json:"password" validate:"required-Password is required,MinStringLength(6)-Password has to have a minimum length of 6 characters"`
	Age      int           `gorm:"not null" json:"age" validate:"required-Age is required, MinLength(8)-Age has to have a minimum of 8 age"`
	Products []SocialMedia `json:"socialmedias"`
	Photo    []Photo       `json:"photos"`
	Comment  []Comment     `json:"comments"`
}

type APIUser struct {
	GORMModel
	Username string `gorm:"not null" json:"username" validate:"required-Username is required"`
	Email    string `gorm:"not null;unique;type:varchar(100)" json:"email" validate:"required-Email is required,email-Invalid email format"`
	Age      int    `gorm:"not null" json:"age" validate:"required-Age is required, MinLength(8)-Age has to have a minimum of 8 age"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	hashedPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass

	return
}
