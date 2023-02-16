package models

import "github.com/jinzhu/gorm"

/**
	小说表
 */

type LiteFiction struct {

	gorm.Model
	Name string `gorm:"type:varchar(50);not null;"` //文件名称
	Tags string `gorm:"type:varchar(50);not null"`//所属类别
	FileName string `gorm:"type:varchar(255);not null;"`//文件关联地址
	Status int  `gorm:"type:int(3);default:1;not null;"`//启用类型
	Users string `gorm:"type:varchar(200);null;"` //上传者
}


/**
	新增小说列表
 */

func CreateFiction(data LiteFiction){

	db.Create(&data)
}

/**
	小说列表
 */
func FindFictionData(page ,size int)(data []LiteFiction,err error){

	return data,db.Order("created_at desc,id desc").Offset((page - 1) * size).Limit(size).Find(&data).Error
}



/**
	小说数量
 */
func CountFictionNum()(num int,err error){

	var fiction []LiteFiction

	return num,db.Find(&fiction).Count(&num).Error
}


/**
	更改状态
 */
func UpdateFictionStatus(id int){

	var fiction LiteFiction

	db.Model(&fiction).Omit("created_at").Where("id = ?",id).Update("status","0")
}

/**
	前台可供下载的数量
 */
func CountHomeFictionNum()(num int,err error){

	var fiction []LiteFiction

	return num,db.Where("status = ?",1).Find(&fiction).Count(&num).Error
}

/**
	小说列表
 */
func FindHomeFictionData(page ,size int)(data []LiteFiction,err error){

	return data,db.Where("status = ?",1).Order("created_at desc,id desc").Offset((page - 1) * size).Limit(size).Find(&data).Error
}

/**
	查询小说是否存在
 */
func FirstFictionDownload(id int)(data LiteFiction,err error){

	return data,db.Where("status = 1 and id = ?",id).Find(&data).Error
}