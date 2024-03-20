package user

import (
	"OSS/global"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Email    string `json:"email"`
}

func (u *User) BeforeCreate(tx *gorm.Tx) error {
	u.ID = global.IDGenerator[global.UserTable].GenerateID()
	return nil
}
