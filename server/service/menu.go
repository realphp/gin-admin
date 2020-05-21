package service

import (
	"gin-admin/db"
	"gin-admin/model"
)

// @title   getMenuTreeMap
// @description    获取路由总树map
// @auth       qm      (2020/05/06 10:26)
// @return     err             error
// @return    menusMsp            map{string}[]SysBaseMenu
func getMenuTreeMap(roleId string) (err error, treeMap map[string][]model.RoleMenu) {
	var allMenus []model.RoleMenu
	treeMap = make(map[string][]model.RoleMenu)
	db := db.Orm.Select("*").Joins("left join  ga_menus on ga_menus.id=ga_role_menus.menu_id")
	if roleId == "2" {
		db.Find(&allMenus)
	} else {
		db.Where("role_id=?", roleId).Find(&allMenus)
	}
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// @title    GetMenuTree
// @description   获取动态菜单树
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return    err             error
// @return    menus           []model.SysMenu
func GetMenuTree(roleId string) (err error, menus []model.RoleMenu) {
	err, menuTree := getMenuTreeMap(roleId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

// @title    getChildrenList
// @description   获取子菜单
// @auth                     （2020/04/05  20:22）
// @param     menu            *model.SysMenu
// @param     sql             string
// @return    err             error

func getChildrenList(menu *model.RoleMenu, treeMap map[string][]model.RoleMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
