package admin

import (
	"github.com/astaxie/beego/logs"
	"mywork/internal/app/common/dto"
	"mywork/internal/pkg"
	"mywork/models"
	"strconv"
	"strings"
)

type ArticleController struct {
	AdminMenuController
}

// @router /admin/article/add [get] 后台首页
func (c *ArticleController) Article() {

	menuData := c.GetMenuList(false) //获取列表

	c.Data["menu"] = menuData

	c.getTag() //获取标签列表

	c.TplName = "admin/article/add.html"
}

// @router /admin/upload/article [post] 后台首页
func (c *ArticleController) UploadArticles() {

	res := c.Upload("file")

	if res["code"] == "0" {

		c.Data["json"] = map[string]interface{}{
			"code": 0,
			"name": res["name"],
			"path": res["path"],
		}

	} else {

		c.Data["json"] = map[string]interface{}{
			"code": 1002,
			"name": res["msg"],
		}
	}

	c.ServeJSON()
}

/**
图片编辑
*/
// @router /admin/upload [post] 后台首页
func (c *ArticleController) Uploads() {

	res := c.Upload("articleName")

	if res["code"] == "0" {

		var info []interface{}

		info = make([]interface{}, 1)

		info[0] = res["names"]

		c.Data["json"] = map[string]interface{}{
			"errno": 0,
			"name":  res["name"],
			"path":  res["path"],
			"data":  info,
		}

	} else {

		c.Data["json"] = map[string]interface{}{
			"code": 1002,
			"name": res["msg"],
		}
	}

	c.ServeJSON()
}

/**
文章列表页面
*/
// @router /admin/article/list [get] 文章列表页
func (c *ArticleController) GetArticleList() {

	size, _ := dto.GetCountArticle()

	c.Data["page"] = size

	c.TplName = "admin/article/list.html"
}

/**
列表展示页面数据处理
*/
// @router /admin/article/listinfo/?:key [get] 文章列表展示
func (c *ArticleController) GetArticleInfo() {

	//获取页数
	var page string  //页数
	var limit string //条数

	c.Ctx.Input.Bind(&page, "page")

	c.Ctx.Input.Bind(&limit, "limit")

	if page == "" {

		page = "0"
	}

	list, _ := dto.SelectArticle(page) //查询接口数据

	var arr map[int]map[string]interface{} //存放数据

	arr = make(map[int]map[string]interface{})

	var arrData map[string]interface{}

	for key, val := range list {

		arrData = make(map[string]interface{}, 11)

		arrData["id"] = val.ID
		arrData["tages"] = dto.GetAidAndTagName(val.ID)
		arrData["is_top"] = val.Is_top
		arrData["is_copy"] = val.Priority
		arrData["status"] = val.Status
		arrData["author"] = val.Author
		arrData["click"] = val.Click
		arrData["read"] = val.Read_num
		arrData["time"] = val.CreatedAt.Format("2006-01-02 15:04:05") //创建时间
		arrData["title"] = val.Title

		fid, _ := strconv.Atoi(val.Fid_Level)

		arrData["level"], _ = dto.GetMenuAndFindInfo(uint(fid)) //获取所属父类

		arr[key] = arrData
	}

	c.Data["json"] = map[string]interface{}{
		"code": "0",
		"data": arr,
	}

	c.ServeJSON()

}

/**
新增数据处理页面
*/
// @router /admin/article/add [post] 后台首页
func (c *ArticleController) AddArticle() {

	//获取传输的值
	title := c.GetString("title") //文章标题

	contents := c.GetString("content") //文章内容

	status := c.GetString("status") //启用状态

	author := c.GetString("author") //作者

	descript := c.GetString("descript") //描述

	keyword := c.GetString("keyword") //关键词 多个已逗号分隔

	tags := c.GetStrings("tag[]") //标签id

	file := c.GetString("title_img") //上传的图片地址

	isCopy := c.GetString("is_copy") //是否原创

	isTop := c.GetString("is_top") //是否置顶

	menu := c.GetString("menu") //所属分类

	tag := strings.Join(tags, ",")

	if title == "" || contents == "" || status == "" || author == "" || keyword == "" || tag == "" || isCopy == "" || isTop == "" || menu == "" {

		c.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  "参数丢失，请检查参数完整性",
		}

	} else {

		//是否存在描述 不存在就去截取前300字
		if descript == "" {

			descript, _ = pkg.GetSummary(contents)
		}

		var article models.LiteArticle

		article.Title = title
		article.Content = contents
		article.Status, _ = strconv.Atoi(status)
		article.Descript = descript
		article.Is_top, _ = strconv.Atoi(isTop)
		article.Title_img = file
		article.Keywords = keyword
		article.Author = author
		article.Fid_Level = menu
		article.Priority, _ = strconv.Atoi(isCopy)

		result, _ := dto.AddArticle(article)

		ids := result.ID //获取新创建的id

		dto.CreateAidAndTag(tag, ids, c.User.ID)

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "文章创建成功",
		}

	}

	c.ServeJSON()
}

/**
编辑数据处理
*/
// @router /admin/article/edit/?:key [post] 后台首页
func (c *ArticleController) EditArticle() {

	//获取传输的值
	title := c.GetString("title") //文章标题

	aid := c.GetString("aid") //主键id

	contents := c.GetString("content") //文章内容
	//logs.Info(c.Ctx.Request.Form.Get("content"))
	status := c.GetString("status") //启用状态

	author := c.GetString("author") //作者

	descript := c.GetString("descript") //描述

	keyword := c.GetString("keyword") //关键词 多个已逗号分隔

	tags := c.GetStrings("tag[]") //标签id 数组获取

	file := c.GetString("title_img") //上传的图片地址

	is_copy := c.GetString("is_copy") //是否原创

	is_top := c.GetString("is_top") //是否置顶

	menu := c.GetString("menu") //所属分类

	tag := strings.Join(tags, ",")

	if title == "" || contents == "" || status == "" || author == "" || keyword == "" || tag == "" || is_copy == "" || is_top == "" || menu == "" {

		c.Data["json"] = map[string]interface{}{
			"code": "1003",
			"msg":  "参数丢失，请检查参数完整性",
		}

	} else {

		//是否存在描述
		if descript == "" {

			descript, _ = pkg.GetSummary(contents)
		}

		var article models.LiteArticle
		id, _ := strconv.Atoi(aid)
		article.ID = uint(id)
		article.Title = title
		article.Content = contents
		article.Status, _ = strconv.Atoi(status)
		article.Descript = descript
		article.Is_top, _ = strconv.Atoi(is_top)
		article.Title_img = file
		article.Keywords = keyword
		article.Author = author
		article.Fid_Level = menu
		article.Priority, _ = strconv.Atoi(is_copy)

		dto.EditArticle(article) //更新数据

		dto.UpdateAidAndTag(tag, uint(id), c.User.ID) //更新关系表

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"msg":  "文章编辑成功",
		}

	}

	c.ServeJSON()

}

/**
编辑文章列表
*/
// @router /admin/article/edit/?:key [get] 文章详情
func (c *ArticleController) ArticleEdit() {

	var id int

	c.Ctx.Input.Bind(&id, "id")

	//获取文章列表数据

	artist, _ := dto.FindArticleInfo(id)

	tagsIds := dto.GetAidAndTagName(uint(id)) //当前文章启用的标签

	menuData := c.GetMenuList(false) //获取列表

	c.Data["tag_ids"] = c.checkTag(tagsIds.Mtid) //查询到值

	c.Data["menu"] = menuData

	c.Data["list"] = artist //列表详情页

	lid, _ := strconv.Atoi(artist.Fid_Level)

	c.Data["lid"] = lid //列表详情页

	c.getTag() //获取标签列表

	c.TplName = "admin/article/edit.html"
}

/**
文章删除
*/
// @router /admin/article/delete/?:key [get] 文章删除
func (c *ArticleController) DelArticle() {

	var id int

	c.Ctx.Input.Bind(&id, "id")

	logs.Info(id)

	if id == 0 {

		c.Data["json"] = map[string]interface{}{
			"code":   "1003",
			"errmsg": "获取参数错误",
		}

	} else {

		dto.DeleArticle(id) //删除文章，

		dto.DeleteArticleAndTag(id) //删除

		c.Data["json"] = map[string]interface{}{
			"code":   0,
			"errmsg": "获取参数错误",
		}

	}

	c.ServeJSON()
}

/**
获取当前启用的标签
*/
func (c *ArticleController) getTag() {

	list, _ := dto.FindTagChecked()

	c.Data["Tag"] = list
}

/**
  	 对比标签数据是否存在
  	 @param
     @return
*/
func (c *ArticleController) checkTag(str string) []string {

	var ids []string

	ids = strings.Split(str, ",")

	return ids
}
