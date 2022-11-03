package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string
}

const (
	PassWordCost = 12
)

// SetPassWord code your password
func (user *User) SetPassWord(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PassWordDigest = string(bytes)
	return nil
}

// CheckPassWord valid your password
func (user *User) CheckPassWord(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return err == nil
}
