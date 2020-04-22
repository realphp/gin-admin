package service

import (
	"gin-admin/db"
	"gin-admin/model"
)

func GetUserByName(name string) (user model.User, err error) {
	err = db.Orm.Where("username=?", name).First(&user).Error
	return
}
