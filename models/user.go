package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `form:"name" valid:"required"`
	Username string  `form:"username" valid:"required" gorm:"unique"`
	Email    string  `form:"email" valid:"required" gorm:"unique"`
	Password string  `form:"password" valid:"required"`
	Photos   []Photo `gorm:"constrain:onUpdate:CASCADE,onDelete:CASCADE"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
