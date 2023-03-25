package dto

import (
	"mywork/models"
	"mywork/pkg/model"
)

/**
查询
*/
func SelectReviewToken(token string) (review models.LiteReview, err error) {

	return review, model.Db.Where("token = ?", token).Limit(1).Find(&review).Error

}

/**
新增
*/
func CreateReview(review models.LiteReview) error {

	return model.Db.Create(&review).Error

}

/**
查询留言条数
*/
func ReviewCount() (num int, err error) {
	var review []models.LiteReview

	return num, model.Db.Find(&review).Count(&num).Error
}

/**
查询当前10条数据
*/
func SelectReview() (review []models.LiteReview, err error) {

	return review, model.Db.Order("created_at desc,id desc").Limit(10).Find(&review).Error
}

/**
分页数据
*/
func SelectReviewPage(page int) (review []models.LiteReview, err error) {

	return review, model.Db.Order("is_top desc,click desc,created_at desc,id desc").Offset((page - 1) * 10).Limit(10).Find(&review).Error

}

/**
删除数据
*/
func DeleteReview(id int) {

	var reView models.LiteReview

	model.Db.Where("id = ?", id).Limit(1).Delete(&reView)
}

/**
获取
*/
func GetHomeReviewCount() (count int, err error) {

	var review []models.LiteReview

	return count, model.Db.Where("status = 1").Find(&review).Count(&count).Error

}

/**
统计今天新增的条数
*/
func GetWhereReviewCount(str string) (count int, err error) {

	var review []models.LiteReview

	return count, model.Db.Where("created_at > ?", str).Find(&review).Count(&count).Error
}
