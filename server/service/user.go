package service

import (
	"errors"
	"gin-admin/db"
	"gin-admin/model"
	"gin-admin/utils"
)

func GetUserByName(name string) (user model.User, err error) {
	err = db.Orm.Where("user_name=?", name).First(&user).Error
	return
}

func AddUser(user model.User) error {
	notRegister := db.Orm.Where("user_name = ?", user.UserName).First(&user).RecordNotFound()
	if !notRegister {
		return errors.New("用户名已注册")
	}
	user.Password = utils.MD5([]byte(model.DEFAULT_PASSWORD))
	err := db.Orm.Create(&user).Error
	return err
}

func UpdateUser(user *model.User) (err error) {
	err = db.Orm.Model(&user).Omit("password").Update(&user).Error
	return err
}

func DeleteUser(user *model.User) (err error) {
	err = db.Orm.Delete(user).Error
	return err
}

func ListUser(info utils.PageInfo) (err error, list []model.UserList, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = db.Orm.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return
	} else {
		err = db.Orm.Table("ga_users").Offset(offset).Limit(limit).Find(&list).Error
	}
	return
}
