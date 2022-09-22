package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model

	Name     string `json:"name" gorm:"size:200;not null"`
	Email    string `json:"email" gorm:"size:200;not null;unique"`
	Password string `json:"password,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	u.CreatedAt = now
	u.UpdatedAt = now

	u.HashPassword()

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()

	return
}

func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)
}
