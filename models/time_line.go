package models

import (
	_ "github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

/**
时间线model
*/
type LiteTimeLine struct {
	gorm.Model

	Content string `gorm:"not null;type:text;"`

	Status string `gorm:"type:int;default:1"`

	Uid int `gorm:"type:int(6);not null"` //创建用户

	Token string `gorm:"type:varchar(250);not null;unique_index"` //token,临牌，禁止重复，唯一索引

	Title string `gorm:"null;type:text;"` //标题

}
