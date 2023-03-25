package dto

import (
	"github.com/jinzhu/gorm"
	"lite_blog/models"
	"lite_blog/pkg/model"
	"strconv"
)

/**
新增文章
*/
func AddArticle(article models.LiteArticle) (articles models.LiteArticle, err error) {

	return article, model.Db.Create(&article).Error

}

/**
编辑文章
*/
func EditArticle(article models.LiteArticle) bool {

	model.Db.Omit("created_at", "click", "read_num").Save(&article) //更新的时候忽略某个字段
	//db.Model(&article).Update("title", "content","priority","is_top","status","click","read_num","title_img","keywords","descript","author","fid_level","updated_at")
	return true
}

// GetCountArticle /**
func GetCountArticle() (count int, err error) {

	var article []models.LiteArticle

	return count, model.Db.Where("status = ?", 1).Find(&article).Count(&count).Error
}

/**
数据展示
*/
func SelectArticle(page string) (articles []models.LiteArticle, err error) {

	pages, _ := strconv.Atoi(page)

	limit := 10

	return articles, model.Db.Select("created_at,id,title,priority,is_top,status,click,read_num,author,fid_level").Order("created_at desc,id desc").Offset((pages - 1) * limit).Limit(limit).Find(&articles).Error
}

/**
获取单条展示数据
*/
func FindArticleInfo(id int) (article models.LiteArticle, err error) {

	return article, model.Db.Where("id = ?", id).Limit(1).Find(&article).Error
}

/**
删除数据
*/
func DeleArticle(id int) (err error) {

	var article models.LiteArticle
	return model.Db.Where("id = ?", id).Delete(&article).Error
}

/**
前台获取
*/
func GetHomeArticle() (article []models.LiteArticle, num int, err error) {

	return article, num, model.Db.Where("status = ?", 1).Limit(10).Order("is_top asc,read_num desc, click desc,created_at desc,id desc").Select("created_at,title,priority,is_top,click,read_num,title_img,keywords,descript,author,id").Find(&article).Count(&num).Error

}

/**
前台获取分页 10
@param id 文章id fid 所属分类 page 分页条数
@return count 这个值没有任何用处 丢弃
*/
func GetHomeAndPageArticle(id int, category int, page int, keyword string) (article []models.LiteArticle, count int, err error) {

	if category != 0 {

		return article, 0, model.Db.Where("status = ? and fid_level = ?", 1, category).Offset((page - 1) * 10).Limit(10).Order("is_top asc,read_num desc, click desc,created_at desc,id desc").Select("created_at,title,priority,is_top,click,read_num,title_img,keywords,descript,author,id").Find(&article).Error

	} else if keyword != "" {

		return article, 0, model.Db.Where("status = ? and title LIKE ?", 1, `%`+keyword+`%`).Or("status = ? and keywords LIKE ?", 1, `%`+keyword+`%`).Or("status = ? and descript LIKE ?", 1, `%`+keyword+`%`).Or("status = ? and author LIKE ?", 1, `%`+keyword+`%`).Offset((page - 1) * 10).Limit(10).Order("is_top asc,read_num desc, click desc,created_at desc,id desc").Select("created_at,title,priority,is_top,click,read_num,title_img,keywords,descript,author,id").Find(&article).Error

	} else {

		return article, 0, model.Db.Where("status = ?", 1).Offset((page - 1) * 10).Limit(10).Order("is_top asc,read_num desc, click desc,created_at desc,id desc").Select("created_at,title,priority,is_top,click,read_num,title_img,keywords,descript,author,id").Find(&article).Error

	}

}

/**
前台获取文章详情
*/
func GetHomeArticleInfo(id int) (article models.LiteArticle, err error) {

	return article, model.Db.Where("status = ? and id = ?", 1, id).Limit(1).Select("created_at,title,priority,click,read_num,keywords,author,id,content").Find(&article).Error

}

/**
根据菜单查询文章
*/
func GetMenuAndArticle(mid int, page int) (article []models.LiteArticle, count int, err error) {

	return article, count, model.Db.Where("status = ? and fid_level = ?", 1, mid).Limit(10).Order("is_top asc,read_num desc, click desc,created_at desc,id desc").Select("created_at,title,priority,is_top,click,read_num,title_img,keywords,descript,author,id").Find(&article).Count(&count).Error

}

/**
文章阅读数
*/
func SetArticleAndRead(id int) {
	var article = models.LiteArticle{}

	model.Db.Model(&article).Where("id = ? and status = ?", id, 1).Update("read_num", gorm.Expr("read_num + 1"))

}

/**
点击次数加+1
*/
func SetArticleAndClick(id int) {

	var article = models.LiteArticle{}

	model.Db.Model(&article).Where("id = ? and status = ?", id, 1).Update("click", gorm.Expr("click + 1"))
}

/**
关键词查询出现的数据
*/
func GetArticleKeywords(keyword string) (article []models.LiteArticle, num int, err error) {

	return article, num, model.Db.Where("status = ? and title LIKE ?", 1, `%`+keyword+`%`).Or("status = ? and keywords LIKE ?", 1, `%`+keyword+`%`).Or("status = ? and descript LIKE ?", 1, `%`+keyword+`%`).Or("status = ? and author LIKE ?", 1, `%`+keyword+`%`).Limit(10).Order("is_top asc,read_num desc, click desc,created_at desc,id desc").Select("created_at,title,priority,is_top,click,read_num,title_img,keywords,descript,author,id").Find(&article).Count(&num).Error

}

/**
	用于tag查询文章详情
    @param aid int 文章id
    @return
*/
func GetTagArticleInfo(aid int) (article models.LiteArticle, err error) {

	return article, model.Db.Where("status = ? and id = ?", 1, aid).Limit(1).Select("created_at,title,priority,is_top,click,read_num,title_img,keywords,descript,author,id").Find(&article).Error
}

/**
文章置顶
*/
func ArticleTopList() (article []models.LiteArticle, err error) {

	return article, model.Db.Where("status = ? and is_top = 1", 1).Limit(20).Order("created_at desc,id desc").Select("title,id").Find(&article).Error

}

func ArticleNext(id int) (article models.LiteArticle, err error) {

	return article, model.Db.Where("id > ?", id).Select("id,title").Limit(1).Find(&article).Error
}

func ArticlePrev(id int) (article models.LiteArticle, err error) {

	return article, model.Db.Where("id < ?", id).Select("id,title").Limit(1).Find(&article).Error
}
