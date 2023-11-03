package models

import (
	"jwtgo/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) ComparePassowrd(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(user.Password))
	if err != nil {
		return err
	}
	return nil
}
