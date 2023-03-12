package models

import (
	"github.com/jinzhu/gorm"
	"mywork/pkg/model"
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

/**
查询
*/
func SelectReviewToken(token string) (review LiteReview, err error) {

	return review, model.Db.Where("token = ?", token).Limit(1).Find(&review).Error

}

/**
新增
*/
func CreateReview(review LiteReview) error {

	return model.Db.Create(&review).Error

}

/**
查询留言条数
*/
func ReviewCount() (num int, err error) {
	var review []LiteReview

	return num, model.Db.Find(&review).Count(&num).Error
}

/**
查询当前10条数据
*/
func SelectReview() (review []LiteReview, err error) {

	return review, model.Db.Order("created_at desc,id desc").Limit(10).Find(&review).Error
}

/**
分页数据
*/
func SelectReviewPage(page int) (review []LiteReview, err error) {

	return review, model.Db.Order("is_top desc,click desc,created_at desc,id desc").Offset((page - 1) * 10).Limit(10).Find(&review).Error

}

/**
删除数据
*/
func DeleteReview(id int) {

	var reView LiteReview

	model.Db.Where("id = ?", id).Limit(1).Delete(&reView)
}

/**
获取
*/
func GetHomeReviewCount() (count int, err error) {

	var review []LiteReview

	return count, model.Db.Where("status = 1").Find(&review).Count(&count).Error

}

/**
统计今天新增的条数
*/
func GetWhereReviewCount(str string) (count int, err error) {

	var review []LiteReview

	return count, model.Db.Where("created_at > ?", str).Find(&review).Count(&count).Error
}
