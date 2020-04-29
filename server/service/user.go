package service

import (
	"errors"
	"gin-admin/db"
	"gin-admin/model"
	"gin-admin/utils"
)

func GetUserByName(name string) (user model.User, err error) {
	err = db.Orm.Where("username=?", name).First(&user).Error
	return
}

func AddUser(user model.User) error {
	notRegister := db.Orm.Where("username = ?", user.Username).First(&user).RecordNotFound()
	if !notRegister {
		return errors.New("用户名已注册")
	}
	user.Password = utils.MD5([]byte(user.Password))
	err := db.Orm.Create(&user).Error
	return err
}
