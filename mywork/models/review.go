package models

import "github.com/jinzhu/gorm"

/**
	model
 */
type LiteReview struct {
	gorm.Model
	Ip string `gorm:"type:varchar(255);not null;"` //ip用户地址
	Message string `gorm:"type varchar(255);not null"`//用户评论信息
	Status int `gorm:"type int;not null;default:1"`//评论启用状态
	Click int `gorm:"type int;not null;default:0"`//点赞数
	Is_top int `gorm:"type int;not null;default:0"`//置顶留言
	Token string `gorm:"type varchar(200);not null;default:0;unique"`//token
	Address string `gorm:"type varchar(255); null;default:null"`//token
}

/**
	查询
 */
func SelectReviewToken(token string) (review LiteReview,err error) {

	return review,db.Where("token = ?",token).Limit(1).Find(&review).Error

}

/**
	新增
 */
func CreateReview(review LiteReview) error  {

	return db.Create(&review).Error

}

/**
	查询留言条数
 */
 func ReviewCount() (num int,err error){
 	var review []LiteReview

 	return num,db.Find(&review).Count(&num).Error
 }


 /**
 	查询当前10条数据
  */
  func SelectReview()(review []LiteReview,err error){

  	return review,db.Order("created_at desc,id desc").Limit(10).Find(&review).Error
  }

  /**
  	分页数据
   */
   func SelectReviewPage(page int)(review []LiteReview,err error){

	return review,db.Order("is_top desc,click desc,created_at desc,id desc").Offset((page -1) * 10).Limit(10).Find(&review).Error

   }


	/**
		删除数据
	 */
	 func DeleteReview(id int){

	 	var reView LiteReview

	 	db.Where("id = ?",id).Limit(1).Delete(&reView)
	 }

	 /**
	 	获取
	  */
	 func GetHomeReviewCount() (count int,err error){

	 	var review []LiteReview

	 	return count,db.Where("status = 1").Find(&review).Count(&count).Error

	 }

	 /**

	  */