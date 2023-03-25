package dto

import (
	"mywork/models"
	"mywork/pkg/model"
)

/**
小说表
*/

/**
新增小说列表
*/

func CreateFiction(data models.LiteFiction) {

	model.Db.Create(&data)
}

/**
小说列表
*/
func FindFictionData(page, size int) (data []models.LiteFiction, err error) {

	return data, model.Db.Order("created_at desc,id desc").Offset((page - 1) * size).Limit(size).Find(&data).Error
}

/**
小说数量
*/
func CountFictionNum() (num int, err error) {

	var fiction []models.LiteFiction

	return num, model.Db.Find(&fiction).Count(&num).Error
}

/**
更改状态
*/
func UpdateFictionStatus(id int) {

	var fiction models.LiteFiction

	model.Db.Model(&fiction).Omit("created_at").Where("id = ?", id).Update("status", "0")
}

/**
前台可供下载的数量
*/
func CountHomeFictionNum() (num int, err error) {

	var fiction []models.LiteFiction

	return num, model.Db.Where("status = ?", 1).Find(&fiction).Count(&num).Error
}

/**
小说列表
*/
func FindHomeFictionData(page, size int) (data []models.LiteFiction, err error) {

	return data, model.Db.Where("status = ?", 1).Order("created_at desc,id desc").Offset((page - 1) * size).Limit(size).Find(&data).Error
}

/**
查询小说是否存在
*/
func FirstFictionDownload(id int) (data models.LiteFiction, err error) {

	return data, model.Db.Where("status = 1 and id = ?", id).Find(&data).Error
}
