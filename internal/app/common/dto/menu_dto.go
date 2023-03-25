package dto

import (
	"mywork/models"
	"mywork/pkg/model"
)

/**
查询菜单
*/

func GetMenuAndRoule(id int) (menu models.LiteAdminMenu, err error) {

	return menu, model.Db.Where("id = ?", id).Take(&menu).Error

}

// SetMenu /**
func SetMenu(menu models.LiteAdminMenu) {

	model.Db.Create(&menu)
}

/**
查询所有启用的菜单
*/
func GetMenuInfo() (menu []models.LiteAdminMenu, err error) {

	return menu, model.Db.Select([]string{"menu_name", "menu_status", "menu_level", "menu_parent", "id", "menu_roule"}).Where("menu_status = ?", "1").Order("menu_parent asc,menu_roule asc,id desc,created_at desc").Find(&menu).Error

}

/**
查询当前启用菜单的数量
*/
func GetMenuNumber() (count int, err error) {

	var menu []models.LiteAdminMenu

	return count, model.Db.Model(&menu).Where("menu_status = ?", "1").Count(&count).Error
}

/**
列表页
*/
func GetAll() (menu []models.LiteAdminMenu, err error) {

	return menu, model.Db.Select([]string{"menu_name", "menu_status", "menu_level", "menu_parent", "id", "menu_roule"}).Order("menu_parent asc,menu_roule asc,id desc,created_at desc").Find(&menu).Error
}

/**
删除
*/
func DeleteMenu(id uint) (menu models.LiteAdminMenu, err error) {

	return menu, model.Db.Where("id = ?", id).Limit(1).Delete(&menu).Error
}

/**
获取数据的详情
*/
func GetMenuAndFindInfo(id uint) (menu models.LiteAdminMenu, err error) {

	return menu, model.Db.Where("id = ?", id).First(&menu).Limit(1).Error
}

/**
更新数据
*/
func EditMenu(menu models.LiteAdminMenu) bool {

	model.Db.Save(&menu)

	return true
}
