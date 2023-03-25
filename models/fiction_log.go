package models

import (
	"github.com/jinzhu/gorm"
)

/**
下载小说记录
*/
type LiteFictionLog struct {
	gorm.Model
	FictionId int    `gorm:"type:int;not null;index:fiction_id"`  //关联小说表id
	Ip        string `gorm:"type:varchar(100);not null;index:ip"` //下载用户的ip地址
	SubNum    int    `gorm:"type:int;default:1;not null"`         //下载本次书籍次数
}

/**
下载小说批次表
*/
type LiteFictionOperation struct {
	gorm.Model
	FictionId   int `gorm:"type:int;not null;index:fiction_id"` //关联小说表id
	DownloadNum int `gorm:"type:int;default:0;not null"`        //改小说下载的量

}
