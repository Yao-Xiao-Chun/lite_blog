package dto

import (
	"lite_blog/models"
	"lite_blog/pkg/model"
)

/**
加入黑名单
*/
func AddBanned(ip string) {

	var ban models.LiteBlackList

	ban.Ip = ip

	model.Db.Create(&ban)

}

/**
查询ip是否存在
*/
func QueryBanned(ip string) (num int, err error) {

	var ban models.LiteBlackList

	return num, model.Db.Where("ip = ?", ip).Find(&ban).Count(&num).Error

}
