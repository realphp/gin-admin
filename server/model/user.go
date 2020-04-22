package model

import "gin-admin/db"

type User struct {
	id       int
	Username string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
}

var table = "ga_user"

func (User) TableName() string {
	return "ga_user"
}

func (u *User) GetUserByName() (user User, err error) {
	err = db.Orm.Where("username=?", u.Username).First(&user).Error
	return
}
