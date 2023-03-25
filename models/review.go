package models

import (
	"github.com/jinzhu/gorm"
)

// LiteReview /**
type LiteReview struct {
	gorm.Model
	Ip      string `gorm:"type:varchar(200);not null;comment:'ip地址'"`                   //ip用户地址
	Message string `gorm:"type varchar(200);not null;comment:'消息'"`                     //用户评论信息
	Status  int    `gorm:"type int;not null;default:1;comment:'评论状态'"`                  //评论启用状态
	Click   int    `gorm:"type int;not null;default:0;comment:'点赞数'"`                   //点赞数
	Is_top  int    `gorm:"type int;not null;default:0;comment:'置顶留言'"`                  //置顶留言
	Token   string `gorm:"type varchar(200);not null;default:0;unique;comment:'token'"` //token
	Address string `gorm:"type varchar(200); null;default:null;comment:'地址'"`           //token
}
