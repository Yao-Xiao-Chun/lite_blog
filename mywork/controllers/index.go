package controllers

import (
	"mywork/models"

	"html/template"
	"github.com/dchest/captcha"
	"github.com/jinzhu/gorm"
	"strings"

)

type IndexController struct {

	HomeBaseController
}

//使用注解路由

// @router /?:key [get] 首页
func (this *IndexController) Index() {

	var keywords string

	this.Ctx.Input.Bind(&keywords,"keyword")

	if keywords == ""{
		this.getHomeArticle() //查询数据
	}else{
		this.getKeyword(keywords)
	}

	//设置模板路径
	this.TplName = "home/index.html"
}

// 关于

// @router /about [get] 首页
func (this *IndexController) IndexAbout() {

	//设置模板路径
	this.TplName = "home/about.html"
}


//消息

// @router /message [get] 首页
func (this *IndexController) IndexMessage() {

	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())//设置模板xss

	d := struct {
		CaptchaId string
	}{
		captcha.NewLen(4),
	}

	this.Data["CaptchaId"] = d.CaptchaId //生成验证码

	this.Data["Uuid"] = this.GetRandomString(24)//token

	this.Data["RePage"],_ = models.GetHomeReviewCount()

	this.Data["viewList"],_ = models.SelectReview()
	//设置模板路径
	this.TplName = "home/message.html"
}

//详情

// @router /details [get] 首页
func (this *IndexController) IndexDetails() {

	//设置模板路径
	this.TplName = "home/details.html"
}


//详情

// @router /time [get] 时间线
func (this *IndexController) IndexTime() {

	this.getTimeLine()
	//设置模板路径
	this.TplName = "home/time.html"
}


/**
	时间线分页
 */

// @router  /time/page/?:id [get]  分页代码
 func (this *IndexController) GetTimePage(){

	 var tid int

	 this.Ctx.Input.Bind(&tid,"id")

	 if tid == 0{

	 	this.Abort("404")
	 }

	 line,_ := models.GetPageTimeLine(tid,10)

	 arr := make(map[int]map[string]interface{},10)

	 if len(line) == 0{

	 	this.Data["json"] = map[string]interface{}{
			 "code":"2",
			 "data":"",
			 "message":"没有数据了！",
		 }

	 }else{
		 //变量赋值
		 for key,val := range line{

			 data := make(map[string]interface{},3) // 每次使用都要初始化一次

			 data["code"] = val.Title

			 data["content"] = val.Content

			 data["tid"] = val.ID

			 arr[key] = data

		 }

		 this.Data["json"] = map[string]interface{}{
			 "code":"0",
			 "data":arr,
			 "message":"请求成功",
		 }
	 }


	 this.ServeJSON()
 }


/**
	时间线负责前端调用 Page
 */
func (this *IndexController) getTimeLine()  {

	line, _ := models.GetHomeTimeLine() //多维结构体

	arr := make(map[int]map[string]interface{},10)

	//变量赋值
	for key,val := range line{

		data := make(map[string]interface{},3) // 每次使用都要初始化一次

		data["code"] = val.Title

		data["content"] = val.Content

		data["tid"] = val.ID

		arr[key] = data

	}

	num,_:= models.GetHomeCountTimeLine()

	this.Data["count"] = num

	this.Data["line"] = arr

}


/**
	其他菜单处理逻辑
 */
// @router /category/?:key [get] 菜单处理
func (this *IndexController)TypeArticle(){

	var id int

	this.Ctx.Input.Bind(&id,"id")

	result,_ := models.GetMenuAndArticle(id)

	if len(result) == 0{

		this.Abort("500")

	}else{

		var data map[int]map[string]interface{}

		var arrData map[string]interface{}

		data = make(map[int]map[string]interface{})

		for key,val := range result{

			arrData = make(map[string]interface{},12)
			arrData["created_time"] = val.CreatedAt.Format("2006-01-02 15:04:05")//创建时间
			arrData["id"] = val.ID
			arrData["tags"] = models.GetAidAndTagName(val.ID)
			arrData["is_top"] = val.Is_top
			arrData["is_copy"] = val.Priority
			arrData["status"] = val.Status
			arrData["author"] = val.Author
			arrData["click"] = val.Click
			arrData["read"] = val.Read_num
			arrData["title"] = val.Title
			arrData["descript"] = val.Descript
			if val.Title_img != "undefind"{
				arrData["img"] = val.Title_img
			}

			data[key] = arrData
		}


		this.Data["article"] = data
	}

	this.Data["is_category"] = id
	//设置模板路径
	this.TplName = "home/index.html"
}


/**
	获取前10个数据
	@param ""
	@return
 */
 func (this *IndexController) getHomeArticle(){

	 result,_:= models.GetHomeArticle()

	 var data map[int]map[string]interface{}

	 var arrData map[string]interface{}

	 data = make(map[int]map[string]interface{})

	 for key,val := range result{

	 	 arrData = make(map[string]interface{},12)
	 	 arrData["created_time"] = val.CreatedAt.Format("2006-01-02 15:04:05")//创建时间
		 arrData["id"] = val.ID
		 arrData["tags"] = models.GetAidAndTagName(val.ID)
		 arrData["is_top"] = val.Is_top
		 arrData["is_copy"] = val.Priority
		 arrData["status"] = val.Status
		 arrData["author"] = val.Author
		 arrData["click"] = val.Click
		 arrData["read"] = val.Read_num
		 arrData["title"] = val.Title
		 arrData["descript"] = val.Descript
		 if val.Title_img != "undefind"{
			 arrData["img"] = val.Title_img
		 }

	 	data[key] = arrData
	 }


	 this.Data["article"] = data
 }


 /**
 	文章分页前台
 	@param id int
 	@return []
  */
// @router /article/?:key [get] 文章分页
func (this *IndexController) GetHomePageArticle(){

	var id,category int

	this.Ctx.Input.Bind(&id,"id")

	this.Ctx.Input.Bind(&category,"category")

	if id == 0{

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"errmsg":"获取参数错误，error",
		}
	}else{

		result,_ := models.GetHomeAndPageArticle(id,category)

		if len(result) == 0{
			this.Data["json"] = map[string]interface{}{
				"code":"1002",
				"errmsg":"已经最后一页了，别点了",
			}
		}else{

			var data map[int]map[string]interface{}

			var arrData map[string]interface{}

			data = make(map[int]map[string]interface{})

			for key,val := range result{

				arrData = make(map[string]interface{},12)
				arrData["created_time"] = val.CreatedAt.Format("2006-01-02 15:04:05")//创建时间
				arrData["id"] = val.ID
				arrData["tags"] = models.GetAidAndTagName(val.ID)
				arrData["is_top"] = val.Is_top
				arrData["is_copy"] = val.Priority
				arrData["status"] = val.Status
				arrData["author"] = val.Author
				arrData["click"] = val.Click
				arrData["read"] = val.Read_num
				arrData["title"] = val.Title
				arrData["descript"] = val.Descript
				if val.Title_img != "undefind"{
					arrData["img"] = val.Title_img
				}

				data[key] = arrData
			}

			this.Data["json"] = map[string]interface{}{
				"code":"0",
				"data":data,
			}
		}

	}

	this.ServeJSON()
}


/**
	获取文章详情
	@param id int
	@return
 */
// @router /article/info/?:key [get] 文章详情
func (this *IndexController) GetArticleInfo(){

	var id int

	this.Ctx.Input.Bind(&id,"id")

	if id == 0{

		this.Abort("404")
	}

	res,_ :=models.GetHomeArticleInfo(id)

	//更新阅读数
	models.SetArticleAndRead(id)
	this.Data["articleData"] = res

	this.TplName = "home/details.html"
}

/**
	留言处理
 */
// @router /message/review [post] 文章详情
func (this *IndexController) SteReview()  {

	text := this.GetString("text")

	captchaId := this.GetString("captchaId")

	ver_code := this.GetString("ver_code")

	token := this.GetString("token")

	if !VerifyCaptcha(captchaId,ver_code){

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"errmsg":"验证码错误",
		}

	}else{

		res,err := models.SelectReviewToken(token)

		if !gorm.IsRecordNotFoundError(err){

			this.Data["json"] = map[string]interface{}{
				"code":"1004",
				"errmsg":"请勿重复提交",
			}
		}else{

			ip := this.Ctx.Request.RemoteAddr

			ip = ip[0:strings.LastIndex(ip, ":")]

			res.Token = token

			res.Message,_ = getSummary(text)

			res.Ip = ip

			res.Address = this.GetAddress(ip)

			models.CreateReview(res)


			this.Data["json"] = map[string]interface{}{
				"code":"0",
				"errmsg":"留言成功",
			}
		}

	}
	//判断是否


	this.ServeJSON()
}



/**
	留言分页处理
 */
// @router /message/review/page/?:key [get] 留言分页
func (this *IndexController) HomePageReview(){

	var page int

	this.Ctx.Input.Bind(&page,"page")

	if page == 0{

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"errmsg":"参数错误，获取信息失败",
		}
	}else{

		data,_ := models. SelectReviewPage(page)

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"data":data,
			"errmsg":"请求成功",
		}

	}

	this.ServeJSON()
}


/**
	文章点赞
 */
// @router /article/click/?:key [get] 文章点赞
func (this *IndexController) ArticleClick(){

	var aid int

	this.Ctx.Input.Bind(&aid,"aid")

	if aid == 0{
		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"errmsg":"参数非法,error",
		}
	}else{
		models.SetArticleAndClick(aid)
		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"errmsg":"点赞成功",
		}
	}
	this.ServeJSON()
}



/**
	评论登录
 */
// @router /blog/pages/?:key [get] 评论
func (this *IndexController) CommitArticle(){

	this.Redirect("/",302)
}


/**
	前台首页搜索框查询内容
 */
func (this *IndexController) getKeyword(key string){

	result,_ := models.GetArticleKeywords(key)

	var data map[int]map[string]interface{}

	var arrData map[string]interface{}

	data = make(map[int]map[string]interface{})

	for key,val := range result{

		arrData = make(map[string]interface{},12)
		arrData["created_time"] = val.CreatedAt.Format("2006-01-02 15:04:05")//创建时间
		arrData["id"] = val.ID
		arrData["tags"] = models.GetAidAndTagName(val.ID)
		arrData["is_top"] = val.Is_top
		arrData["is_copy"] = val.Priority
		arrData["status"] = val.Status
		arrData["author"] = val.Author
		arrData["click"] = val.Click
		arrData["read"] = val.Read_num
		arrData["title"] = val.Title
		arrData["descript"] = val.Descript
		if val.Title_img != "undefind"{
			arrData["img"] = val.Title_img
		}

		data[key] = arrData
	}


	this.Data["article"] = data
	this.Data["Keywords"] = key

}