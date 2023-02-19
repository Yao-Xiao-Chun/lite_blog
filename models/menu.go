package models

import (
	"github.com/jinzhu/gorm"
)

// LiteAdminMenu /**
type LiteAdminMenu struct {
	gorm.Model

	Menu_name string `gorm:"null;type:varchar(255);commit:'菜单名称'"` // 菜单名称

	Menu_roule int `gorm:"default:0;not null;comment:'菜单类型 0代表顶级 1代表一级分类'"` //菜单类型 0代表顶级 1代表一级分类

	Menu_parent int `gorm:"null;comment:'所属父类 默认为空'"` // 所属父类 默认为空

	Menu_status int `gorm:"not null;default:1;comment:'启用状态 1代表启用 0代表禁用'"` // 启用状态 1代表启用 0代表禁用

	Menu_key string `gorm:"not null;default:'-';comment:'菜单key'"` //菜单key

	Menu_level int `gorm:"not null;default:0;comment:'菜单等级 |-'"` //菜单等级 |-

}

/**
查询菜单
*/

func GetMenuAndRoule(id int) (menu LiteAdminMenu, err error) {

	return menu, db.Where("id = ?", id).Take(&menu).Error

}

// SetMenu /**
func SetMenu(menu LiteAdminMenu) {

	db.Create(&menu)
}

/**
查询所有启用的菜单
*/
func GetMenuInfo() (menu []LiteAdminMenu, err error) {

	return menu, db.Select([]string{"menu_name", "menu_status", "menu_level", "menu_parent", "id", "menu_roule"}).Where("menu_status = ?", "1").Order("menu_parent asc,menu_roule asc,id desc,created_at desc").Find(&menu).Error

}

/**
查询当前启用菜单的数量
*/
func GetMenuNumber() (count int, err error) {

	var menu []LiteAdminMenu

	return count, db.Model(&menu).Where("menu_status = ?", "1").Count(&count).Error
}

/**
列表页
*/
func GetAll() (menu []LiteAdminMenu, err error) {

	return menu, db.Select([]string{"menu_name", "menu_status", "menu_level", "menu_parent", "id", "menu_roule"}).Order("menu_parent asc,menu_roule asc,id desc,created_at desc").Find(&menu).Error
}

/**
删除
*/
func DeleteMenu(id uint) (menu LiteAdminMenu, err error) {

	return menu, db.Where("id = ?", id).Limit(1).Delete(&menu).Error
}

/**
获取数据的详情
*/
func GetMenuAndFindInfo(id uint) (menu LiteAdminMenu, err error) {

	return menu, db.Where("id = ?", id).First(&menu).Limit(1).Error
}

/**
更新数据
*/
func EditMenu(menu LiteAdminMenu) bool {

	db.Save(&menu)

	return true
}
