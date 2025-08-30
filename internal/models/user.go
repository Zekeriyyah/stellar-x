package models

import (
	"time"

	"github.com/Zekeriyyah/stellar-x/pkg"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Phone     string    `gorm:"unique" json:"phone"`
	Password  string	`json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) SetPassword(password string) error {
	hashed, err := pkg.HashPassword(password)
	if err != nil {
		pkg.Error(err, "error hashing password")
		return err
	}

	u.Password = string(hashed)
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	return pkg.CheckPassword(u.Password, password)
}
