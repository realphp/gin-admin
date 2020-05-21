package model

import (
	"time"
)

type Role struct {
	Id        int        `json:"roleId" gorm:"not null;unique;primary_key"`
	Name      string     `json:"roleName"`
	ParentId  string     `json:"parentId"`
	Children  []RoleMenu `json:"children"`
	Menu      []Menu     `json:"menus" gorm:"many2many:ga_menus;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Role) TableName() string {
	return "ga_roles"
}
