package models

import (
	"github.com/jinzhu/gorm"
)

/**
黑名单
*/

type LiteBlackList struct {
	gorm.Model
	Ip string `gorm:"type:varchar(100);not null;index:ip"` //下载用户的ip地址
}
