package model

import "gin-admin/db"

const DEFAULT_PASSWORD = "123456"

type UserList struct {
	Id       int
	UserName string `json:"user_name" validate:"required"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	RoleId   int    `json:"role_id"`
}

type User struct {
	UserList
	Password string `json:"password"`
}

func (User) TableName() string {
	return "ga_users"
}

func (u *User) GetUserByName() (user User, err error) {
	err = db.Orm.Where("user_name=?", u.UserName).First(&user).Error
	return
}
