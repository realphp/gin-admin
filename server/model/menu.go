package model

import "time"

type Menu struct {
	ID        uint   `gorm:"primary_key"`
	MenuLevel uint   `json:"-"`
	ParentId  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Meta      `json:"meta"`
	Children  []Menu `json:"children"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive"`
	DefaultMenu bool   `json:"defaultMenu"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
}

func (Menu) TableName() string {
	return "ga_menus"
}
