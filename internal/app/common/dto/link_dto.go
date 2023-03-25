package dto

import (
	"lite_blog/models"
	"lite_blog/pkg/model"
)

/**
新增
*/

func AddLinkData(data models.LiteLink) (err error) {
	return model.Db.Create(&data).Error
}

/**
获取展示数据
*/
func GetLinkTotal() (count int, err error) {

	var link []models.LiteLink

	return count, model.Db.Find(&link).Count(&count).Error
}

/**
查询数据
*/
func SelectLink(page int) (link []models.LiteLink, err error) {

	return link, model.Db.Order("created_at desc,id desc").Offset((page - 1) * 10).Limit(10).Find(&link).Error

}

/**
查询单条友链
*/
func FindLinkInfo(id int) (link models.LiteLink, err error) {

	return link, model.Db.Where("id = ?", id).Limit(1).Find(&link).Error
}

/**
更新
*/
func SaveLink(link models.LiteLink) error {

	return model.Db.Omit("created_at").Save(&link).Error
}

/**
删除
*/

func DeleteLink(id int) error {

	var link models.LiteLink

	return model.Db.Where("id =?", id).Delete(&link).Error
}

func GetHomeLink() (link []models.LiteLink, err error) {

	return link, model.Db.Where("status = ?", 1).Select("name,url").Order("sort asc,id desc").Find(&link).Error
}
