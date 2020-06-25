package service

import (
	"errors"
	"gin-admin/db"
	"gin-admin/model"
	"gin-admin/utils"
)

func ListRole(pageInfo utils.PageInfo) (err error, list []model.Role, total int) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	err = db.Orm.Model(&model.Role{}).Count(&total).Error
	if err != nil {
		return
	} else {
		err = db.Orm.Offset(offset).Limit(limit).Find(&list).Error
	}
	return
}

func AddRole(r model.Role) (err error, role model.Role) {
	notHas := db.Orm.Where("name = ?", r.Name).First(&role).RecordNotFound()
	if !notHas {
		return errors.New("存在相同角色id"), role
	}
	err = db.Orm.Create(&r).Error
	return
}

func DeleteRole(role *model.Role) (err error) {
	err = db.Orm.Delete(role).Error
	return err
}

func UpdateRole(role *model.Role) (err error) {
	err = db.Orm.Model(&role).Update(&role).Error
	return err
}
