package model

type RoleMenu struct {
	Menu
	MenuId      string     `json:"menuId"`
	AuthorityId string     `json:"-"`
	Children    []RoleMenu `json:"children"`
}

func (RoleMenu) TableName() string {
	return "ga_role_menus"
}
