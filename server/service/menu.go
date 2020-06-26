package service

import (
	"gin-admin/db"
	"gin-admin/model"
)

// @title   getMenuTreeMap
// @description    获取路由总树map
// @return     err             error
// @return    menusMsp            map{string}[]SysBaseMenu
func getMenuTreeMap(roleId int) (err error, treeMap map[uint][]model.RoleMenu) {
	var allMenus []model.RoleMenu
	treeMap = make(map[uint][]model.RoleMenu)
	if roleId == 1 {
		err = db.Orm.Table("ga_menus").Find(&allMenus).Error
	} else {
		err = db.Orm.Select("*").Joins("left join  ga_menus on ga_menus.id=ga_role_menus.menu_id").Where("role_id=?", roleId).Find(&allMenus).Error
	}
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// @title    GetMenuTree
// @description   获取动态菜单树
// @param     authorityId     string
// @return    err             error
// @return    menus           []model.SysMenu
func GetMenuTree(roleId int) (err error, menus []model.RoleMenu) {
	err, menuTree := getMenuTreeMap(roleId)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

// @title    getChildrenList
// @description   获取子菜单
// @param     menu            *model.SysMenu
// @param     sql             string
// @return    err             error

func getChildrenList(menu *model.RoleMenu, treeMap map[uint][]model.RoleMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// @title    GetMenuList
// @description   获取路由分页
// @param     info            request.PageInfo
// @return    err             error
// @return    list            interface{}
// @return    total           int

func GetMenuList() (err error, list interface{}, total int) {
	var menuList []model.Menu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap[0]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

// @title   getBaseMenuTreeMap
// @description    获取路由总树map
// @return     err             error
// @return    menusMsp            map{string}[]RoleMenu

func getBaseMenuTreeMap() (err error, treeMap map[uint][]model.Menu) {
	var allMenus []model.Menu
	treeMap = make(map[uint][]model.Menu)
	err = db.Orm.Order("sort", true).Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// @title    getBaseChildrenList
// @description   get children of menu, 获取菜单的子菜单
// @param     menu            *model.SysBaseMenu
// @return    err             error

func getBaseChildrenList(menu *model.Menu, treeMap map[uint][]model.Menu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
