package model

import "gin-admin/db"

type User struct {
	Id       int
	Username string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	RoleId   string `json:"role_id"`
}

func (User) TableName() string {
	return "ga_users"
}

func (u *User) GetUserByName() (user User, err error) {
	err = db.Orm.Where("user_name=?", u.Username).First(&user).Error
	return
}
