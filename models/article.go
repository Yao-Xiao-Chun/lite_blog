package models

import (
	"github.com/jinzhu/gorm"
)

/**
文章编辑代码
*/
type LiteArticle struct {
	gorm.Model //继承gorm类

	Title     string `gorm:"type:varchar(255);not null"`
	Content   string `gorm:"type:text;null"`
	Priority  int    `gorm:"type: int;not null;default:0"` //原创
	Is_top    int    `gorm:"type:int;not null;default:0"`  //置顶
	Status    int    `gorm:"type:int;not null;default:1"`  //状态
	Click     int    `gorm:"type:int;not null;default:0"`  //点击数
	Read_num  int    `gorm:"type:int;not null;default:0"`  //阅读数量
	Title_img string `gorm:"type:varchar(255);null"`       //预览标题图片
	Keywords  string `gorm:"type:varchar(255);null"`       //关键词
	Descript  string `gorm:"type:varchar(255);not null"`   //描述
	Author    string `gorm:"type:varchar(255);not null"`   // 作者
	Fid_Level string `gorm:"type:varchar(10);null"`        //所属分类

}
