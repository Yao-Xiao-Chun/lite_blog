package models

import (
	"github.com/jinzhu/gorm"
	"mywork/pkg/model"
)

/**
友情链接表
*/
type LiteLink struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);not null;"`  //名称
	Url    string `gorm:"type:varchar(255);not nul;"`   //地址
	Sort   int    `gorm:"type:int;not null;default:0;"` //默认排序
	Status int    `gorm:"type:int;not null;default:1"`  //启用状态 1 启用 0禁用
}

/**
新增
*/

func AddLinkData(data LiteLink) (err error) {
	return model.Db.Create(&data).Error
}

/**
获取展示数据
*/
func GetLinkTotal() (count int, err error) {

	var link []LiteLink

	return count, model.Db.Find(&link).Count(&count).Error
}

/**
查询数据
*/
func SelectLink(page int) (link []LiteLink, err error) {

	return link, model.Db.Order("created_at desc,id desc").Offset((page - 1) * 10).Limit(10).Find(&link).Error

}

/**
查询单条友链
*/
func FindLinkInfo(id int) (link LiteLink, err error) {

	return link, model.Db.Where("id = ?", id).Limit(1).Find(&link).Error
}

/**
更新
*/
func SaveLink(link LiteLink) error {

	return model.Db.Omit("created_at").Save(&link).Error
}

/**
删除
*/

func DeleteLink(id int) error {

	var link LiteLink

	return model.Db.Where("id =?", id).Delete(&link).Error
}

func GetHomeLink() (link []LiteLink, err error) {

	return link, model.Db.Where("status = ?", 1).Select("name,url").Order("sort asc,id desc").Find(&link).Error
}
