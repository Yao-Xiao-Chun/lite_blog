package models

import "github.com/jinzhu/gorm"

/**
	黑名单
 */

 type LiteBlackList struct {

 	gorm.Model
 	Ip string `gorm:"type:varchar(100);not null;index:ip"`//下载用户的ip地址
 }


 /**
 	加入黑名单
  */
func AddBanned(ip string){

	var ban LiteBlackList

	ban.Ip = ip

	db.Create(&ban)

}

/**
	查询ip是否存在
 */
func QueryBanned(ip string) (num int,err error) {

	var ban LiteBlackList

	return num,db.Where("ip = ?",ip).Find(&ban).Count(&num).Error

}