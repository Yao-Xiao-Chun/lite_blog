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
