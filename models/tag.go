package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type LiteTag struct {
	gorm.Model
	Tag_name string `gorm:"not null;type:varchar(255)"` //标签名称
	Is_status int `gorm:"not null;default:0"`

}


/**
	创建数据
 */
 func CreateTag(tag LiteTag){

 	db.Create(&tag)
 }


 /**
 	查询数据
  */
  func QueryTagList(pages string,size int)(tag []LiteTag,count string,err error){

  	if size == 0{

  		size = 10
	}
	page,_ := strconv.Atoi(pages)

	if page == 0{

		page = 1;
	}

	var list []LiteTag

	return list,count,db.Order("created_at desc,id desc").Offset((page -1 )*size).Limit(size).Find(&list).Count(&count).Error

  }


  /**
  	删除
   */
   func TagDel(id int)(tag LiteTag,err error){

   	return tag,db.Where("id = ?",id).Limit(1).Delete(&tag).Error
   }


  /**
  	 查询个数
   */
   func QueryCountTag() (num int,err error){

   	var count int

   	var tag []LiteTag

   	return count,db.Order("id desc").Find(&tag).Count(&count).Error
   }

   /**
   		查询当前编辑的详情
    */
    func QueryTagFirst(id int) (tag LiteTag,err error){

    	if id == 0{
			return
		}

		return tag,db.Where("id = ?",id).Select([]string{"tag_name","id","is_status"}).First(&tag).Error
	}

	/**
		更新
	 */
	 func UpdateTagFirst(tag LiteTag){

		 db.Model(&tag).Updates(map[string]interface{}{"tag_name": tag.Tag_name,"is_status":tag.Is_status})
	 }

	 /**
	 	获取当前选择的所有标签
	  */
	  func FindTagChecked() (tag []LiteTag,err error){

	  	return tag,db.Where("is_status = ?",1).Select([]string{"tag_name","id"}).Find(&tag).Error

	  }
/**
	获取小说的标签 状态为0的 假设
 */
func FindTagTypeTwo() (tag []LiteTag,err error){

	return tag,db.Where("is_status = ?",0).Select([]string{"tag_name","id"}).Find(&tag).Error

}