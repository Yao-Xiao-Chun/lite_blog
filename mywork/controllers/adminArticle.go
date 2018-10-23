package controllers

import (
	"mywork/models"
	"strconv"
	"strings"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
)

type AdminArticleController struct {

	AdminMenuController
}

// @router /admin/article/add [get] 后台首页
func (this *AdminArticleController) Artilce(){

	menuData := this.GetMenuList(false) //获取列表

	this.Data["menu"] = menuData

	this.getTag()//获取标签列表

	this.TplName = "admin/article/add.html"
}

// @router /admin/upload/article [post] 后台首页
func (this *AdminArticleController) UploadArticles(){

	res := this.Upload("file")

	if res["code"] == "0"{

		this.Data["json"] = map[string]interface{}{
			"code":0,
			"name":res["name"],
			"path":res["path"],
		}

	}else{

		this.Data["json"] = map[string]interface{}{
			"code":1002,
			"name":res["msg"],

		}
	}

	this.ServeJSON()
}


/**
	图片编辑
 */
// @router /admin/upload [post] 后台首页
func (this *AdminArticleController) Uploads(){

	res := this.Upload("articleName")

	if res["code"] == "0"{

		var info []interface{}

		info = make([]interface{},1)

		info[0] = res["names"]

		this.Data["json"] = map[string]interface{}{
			"errno":0,
			"name":res["name"],
			"path":res["path"],
			"data":info,
		}

	}else{

		this.Data["json"] = map[string]interface{}{
			"code":1002,
			"name":res["msg"],

		}
	}

	this.ServeJSON()
}

/**
	文章列表页面
 */
// @router /admin/article/list [get] 文章列表页
 func (this *AdminArticleController) GetArticleList(){

 	size,_ := models.GetCountArticle()

 	this.Data["page"] = size

 	this.TplName = "admin/article/list.html"
 }


 /**
 	列表展示页面数据处理
  */
// @router /admin/article/listinfo/?:key [get] 文章列表展示
func (this *AdminArticleController) GetArticleInfo(){

	//获取页数
	var page string //页数
	var limit string //条数

	this.Ctx.Input.Bind(&page,"page")

	this.Ctx.Input.Bind(&limit,"limit")

	if page == "" {

		page = "0"
	}

	list,_ := models.SelectArticle(page)//查询接口数据

	var arr map[int]map[string]interface{} //存放数据

	arr = make(map[int]map[string]interface{})

	var arrData map[string] interface{}

	for key,val := range list{

		arrData = make(map[string] interface{},11)

		arrData["id"] = val.ID
		arrData["tages"] = models.GetAidAndTagName(val.ID)
		arrData["is_top"] = val.Is_top
		arrData["is_copy"] = val.Priority
		arrData["status"] = val.Status
		arrData["author"] = val.Author
		arrData["click"] = val.Click
		arrData["read"] = val.Read_num
		arrData["time"] = val.CreatedAt.Format("2006-01-02 15:04:05")//创建时间
		arrData["title"] = val.Title

		fid,_ := strconv.Atoi(val.Fid_Level)

		arrData["level"],_ = models.GetMenuAndFindInfo(uint(fid))//获取所属父类

		arr[key] = arrData
	}

	this.Data["json"] = map[string]interface{}{
		"code":"0",
		"data":arr,
	}

	this.ServeJSON()

}



 /**
 	新增数据处理页面
  */
// @router /admin/article/add [post] 后台首页
 func (this *AdminArticleController) AddArticle(){

 	//获取传输的值
 	title := this.GetString("title") //文章标题

	contents := this.GetString("content") //文章内容

	status := this.GetString("status") //启用状态

	author := this.GetString("author") //作者

	descript := this.GetString("descript")//描述

	keyword := this.GetString("keyword")//关键词 多个已逗号分隔

	tags := this.GetStrings("tag[]") //标签id

	file := this.GetString("title_img") //上传的图片地址

	is_copy := this.GetString("is_copy") //是否原创

	is_top := this.GetString("is_top") //是否置顶

	menu := this.GetString("menu") //所属分类

	tag := strings.Join(tags,",")

	if title == "" || contents == "" || status == "" || author == "" || keyword == "" ||tag == "" || is_copy == "" || is_top == "" || menu == ""{

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"msg" :"参数丢失，请检查参数完整性",
		}

	}else{

		//是否存在描述 不存在就去截取前300字
		if descript == ""{

			descript,_ = getSummary(contents)
		}

		var article models.LiteArticle

		article.Title = title
		article.Content = contents
		article.Status,_ = strconv.Atoi(status)
		article.Descript = descript
		article.Is_top,_ = strconv.Atoi(is_top)
		article.Title_img = file
		article.Keywords = keyword
		article.Author = author
		article.Fid_Level = menu
		article.Priority,_ = strconv.Atoi(is_copy)

		result,_ := models.AddArticle(article)

		ids := result.ID //获取新创建的id

		models.CreateAidAndTag(tag,ids,this.User.ID)

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"msg" :"文章创建成功",
		}

	}

 	this.ServeJSON()
 }


 /**
 	编辑数据处理
  */
// @router /admin/article/edit/?:key [post] 后台首页
func (this *AdminArticleController) EditArticle(){

	//获取传输的值
	title := this.GetString("title") //文章标题

	aid := this.GetString("aid") //主键id

	contents := this.GetString("content") //文章内容
	//logs.Info(this.Ctx.Request.Form.Get("content"))
	status := this.GetString("status") //启用状态

	author := this.GetString("author") //作者

	descript := this.GetString("descript")//描述

	keyword := this.GetString("keyword")//关键词 多个已逗号分隔

	tags := this.GetStrings("tag[]") //标签id 数组获取

	file := this.GetString("title_img") //上传的图片地址

	is_copy := this.GetString("is_copy") //是否原创

	is_top := this.GetString("is_top") //是否置顶

	menu := this.GetString("menu") //所属分类

	tag := strings.Join(tags,",")

	if title == "" || contents == "" || status == "" || author == "" || keyword == "" ||tag == "" || is_copy == "" || is_top == "" || menu == ""{

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"msg" :"参数丢失，请检查参数完整性",
		}

	}else{

		//是否存在描述
		if descript == ""{

			descript,_ = getSummary(contents)
		}

		var article models.LiteArticle
		id,_ := strconv.Atoi(aid)
		article.ID =uint(id)
		article.Title = title
		article.Content = contents
		article.Status,_ = strconv.Atoi(status)
		article.Descript = descript
		article.Is_top,_ = strconv.Atoi(is_top)
		article.Title_img = file
		article.Keywords = keyword
		article.Author = author
		article.Fid_Level = menu
		article.Priority,_ = strconv.Atoi(is_copy)

		models.EditArticle(article)//更新数据


		models.UpdateAidAndTag(tag,uint(id),this.User.ID)//更新关系表

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"msg" :"文章编辑成功",
		}

	}

	this.ServeJSON()

}

 /**
 	编辑文章列表
  */
// @router /admin/article/edit/?:key [get] 文章详情
func (this *AdminArticleController) ArticleEdit(){

	var id int

	this.Ctx.Input.Bind(&id,"id")

	//获取文章列表数据

	artist,_ := models.FindArticleInfo(id)

	tagsIds := models.GetAidAndTagName(uint(id))//当前文章启用的标签

	menuData := this.GetMenuList(false) //获取列表

	this.Data["tag_ids"] = this.checkTag(tagsIds.Mtid)//查询到值

	this.Data["menu"] = menuData

	this.Data["list"] = artist //列表详情页

	lid,_ := strconv.Atoi(artist.Fid_Level)

	this.Data["lid"] = lid //列表详情页

	this.getTag()//获取标签列表

	this.TplName = "admin/article/edit.html"
}



/**
	文章删除
 */
// @router /admin/article/delete/?:key [get] 文章删除
func (this *AdminArticleController) DelArticle(){

	var id int

	this.Ctx.Input.Bind(&id,"id")

	logs.Info(id)

	if id == 0{

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"errmsg":"获取参数错误",
		}

	}else{

		models.DeleArticle(id) //删除文章，

		models.DeleteArticleAndTag(id)//删除

		this.Data["json"] = map[string]interface{}{
			"code":0,
			"errmsg":"获取参数错误",
		}

	}

	this.ServeJSON()
}


 /**
 	获取当前启用的标签
  */
  func (this *AdminArticleController) getTag(){

  		list,_ := models.FindTagChecked()

  		this.Data["Tag"] = list
  }


  /**
  	 对比标签数据是否存在
  	 @param
     @return
   */
   func (this *AdminArticleController) checkTag(str string) ([]string){

   		var ids []string

   		ids = strings.Split(str,",")

   		return  ids
   }

   /**
     文章截取前300字
    */
	func getSummary(content string) (string, error) {

		// bytes.Buffer，非常常用。
		var buf bytes.Buffer

		buf.Write([]byte(content))

		// 用goquery来解析content
		doc, err := goquery.NewDocumentFromReader(&buf)

		if err != nil {

			return "", err

		}

		// Text() 得到body元素下的文本内容（去掉html元素）
		str := doc.Find("body").Text()

		// 截取字符串
		if len(str) > 300 {

			str = str[0:300] + "..."
		}

		return str, nil

	}

