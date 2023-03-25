package models

import (
	"github.com/jinzhu/gorm"
)

/**
小说表
*/

type LiteFiction struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);not null;"`      //文件名称
	Tags     string `gorm:"type:varchar(50);not null"`       //所属类别
	FileName string `gorm:"type:varchar(255);not null;"`     //文件关联地址
	Status   int    `gorm:"type:int(3);default:1;not null;"` //启用类型
	Users    string `gorm:"type:varchar(200);null;"`         //上传者
}
