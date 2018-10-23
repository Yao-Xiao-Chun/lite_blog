package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"strconv"
)

type LiteArticleTag struct {
	gorm.Model
	Aid int `gorm:"not null;type:int;index:aid"` //文章id
	Tid string `gorm:"not null;type:varchar(255)"`//标签id
	Create_name string `gorm:"type:varchar(255);not null;"`//创建时间
	Uid int `gorm:"type:int;not null"`//用户

}

type Result struct {
	Mid int
	Maid string
	Mtid string
	Mnames string
}

func CreateAidAndTag(str string,aid uint,uid uint){

	var artTags LiteArticleTag
	artTags.Aid = int(aid)
	artTags.Tid = str
	artTags.Uid = int(uid)
	artTags.Create_name = time.Now().Format("2006-01-02 15:04:05")

	db.Create(&artTags)

}

func UpdateAidAndTag(str string,aid uint,uid uint){
	var artTags LiteArticleTag
	artTags.Aid = int(aid)
	artTags.Tid = str
	artTags.Uid = int(uid)
	db.Where("aid = ?",aid).Save(&artTags)
}

/**
	获取文章对应标签
 */
func GetAidAndTagName(aid uint) (res Result) {

	id := strconv.Itoa(int(aid)) //文章id

	var str string

	str = `SELECT
	m.id as mid,
		m.aid as maid,
		m.tid as mtid,
		GROUP_CONCAT(tag.tag_name) AS mnames
	FROM
	lite_article_tags AS  m
	JOIN lite_tags AS tag ON Find_IN_SET(tag.id, m.tid) WHERE m.aid =`+id+`
	GROUP BY
	m.id `


	db.Raw(str).Scan(&res) //手动执行sql语句

	return res
}


/**
	删除 管理标签
 */
 func DeleteArticleAndTag(id int)(err error){

 	var artTag LiteArticleTag

 	return db.Where("aid = ?",id).Delete(&artTag).Limit(1).Error
 }