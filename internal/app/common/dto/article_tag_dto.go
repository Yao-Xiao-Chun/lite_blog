package dto

import (
	"mywork/models"
	"mywork/pkg/model"
	"strconv"
	"time"
)

type Result struct {
	Mid    int
	Maid   string
	Mtid   string
	Mnames string
}

type HomeTag struct {
	Total int
}

type FciData struct {
	Fname string
	Id    int
}

/**
前台获取tage文章
*/
type HomeArticleTagList struct {
	Aid int
}

func CreateAidAndTag(str string, aid uint, uid uint) {

	var artTags models.LiteArticleTag
	artTags.Aid = int(aid)
	artTags.Tid = str
	artTags.Uid = int(uid)
	artTags.Create_name = time.Now().Format("2006-01-02 15:04:05")

	model.Db.Create(&artTags)

}

func UpdateAidAndTag(str string, aid uint, uid uint) {
	var artTags models.LiteArticleTag
	artTags.Aid = int(aid)
	artTags.Tid = str
	artTags.Uid = int(uid)
	model.Db.Where("aid = ?", aid).Save(&artTags)
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
	JOIN lite_tags AS tag ON Find_IN_SET(tag.id, m.tid) WHERE m.aid =` + id + `
	GROUP BY
	m.id `

	model.Db.Raw(str).Scan(&res) //手动执行sql语句

	return res
}

/**
删除 管理标签
*/
func DeleteArticleAndTag(id int) (err error) {

	var artTag models.LiteArticleTag

	return model.Db.Where("aid = ?", id).Delete(&artTag).Limit(1).Error
}

/**
统计各标签关联的文章数量
*/
func CountArticleAndTag(tid int) (num HomeTag) {

	id := strconv.Itoa(int(tid)) //文章id

	var str string

	str = `SELECT
	count(m.id) as total
	FROM
	lite_article_tags AS  m left join lite_articles as a on a.id = m.aid
	WHERE FIND_IN_SET(` + id + `,m.tid) and a.status = 1`

	model.Db.Raw(str).Scan(&num)

	return num
}

/**
前台获取标签关联的文章
*/
func GetHomeTagsArticle(tid, page int) (list []HomeArticleTagList, total int) {

	id := strconv.Itoa(int(tid)) //文章id

	var str, sql string

	var pages int

	var num HomeTag

	pages = (page - 1) * 10

	pageStr := strconv.Itoa(pages)

	str = `SELECT
	aid
	FROM
	lite_article_tags AS  m
	WHERE FIND_IN_SET(` + id + `,m.tid) order by id desc limit ` + pageStr + `,10`

	model.Db.Raw(str).Scan(&list)

	sql = `SELECT
	count(m.id) as total
	FROM
	lite_article_tags AS  m
	WHERE FIND_IN_SET(` + id + `,m.tid)`
	model.Db.Raw(sql).Scan(&num)

	return list, num.Total
}

/**
获取小说的标签名称
*/
func FictionAndTag(fid string) (data FciData) {

	var sql string

	sql = `select f.id,group_concat(t.tag_name) as fname from lite_fictions as f left join lite_tags as t ON FIND_IN_SET(t.id,f.tags) where f.tags="` + fid + `"`

	model.Db.Raw(sql).Scan(&data)

	return data
}
