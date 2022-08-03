package model

import (
	"auth-service/helper"
)

// 由于不涉及更新操作，所以用name作为key
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Role struct {
	Name string `json:"name"`
}

func (u *User) EncryptPassword() {
	u.Password = helper.Md5Encrypt(u.Password)
}
