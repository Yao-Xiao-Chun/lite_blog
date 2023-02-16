package controllers

import (
	"mywork/models"

	"html/template"
	"github.com/dchest/captcha"
	"github.com/jinzhu/gorm"
	"strings"
	"strconv"
	"github.com/astaxie/beego/logs"
	"sync"
)

type IndexController struct {

	HomeBaseController
}

type TagList struct {
	Name string
	Num  int
	Tid  int
	Sort int
}


/**
	时间结构体 仅供前天调用 xx时间前留言
 */
type TimeString struct {
	models.LiteReview

	TimeStr string
}

//使用注解路由

// @router /?:key [get] 首页
func (this *IndexController) Index() {

	var keywords string

	var tag int

	this.Ctx.Input.Bind(&keywords,"keyword")

	this.Ctx.Input.Bind(&tag,"tag")

	if tag != 0{

		this.getKeyword("",tag)//查询tag

	}else if keywords == ""{

		this.getHomeArticle() //查询数据

	}else{

		this.getKeyword(keywords,0) //查询关键词
	}

	this.Data["TagList"] = this.getHomeTags() //展示的是热门标签

	this.Data["Placard"],_ = models.GetPlacard()

	this.GetTop()//获取置顶
	//设置模板路径
	this.TplName = "home/index.html"
}

// 关于

// @router /about [get] 首页
func (this *IndexController) IndexAbout() {

	this.Data["Abort"],_ = models.GetAbort()
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

			 data := make(map[string]interface{},4) // 每次使用都要初始化一次

			 data["code"] = val.Title

			 data["content"] = val.Content

			 data["tid"] = val.ID

			 data["token"] = val.Token

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
    var page int
	this.Ctx.Input.Bind(&id,"id")

	this.Ctx.Input.Bind(&page,"page")

	result,num,_ := models.GetMenuAndArticle(id,page)

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

		this.Data["articleNum"] = num
	}

	this.Data["Placard"],_ = models.GetPlacard()

	this.Data["TagList"] = this.getHomeTags() //展示的是热门标签

	this.Data["is_category"] = id

	this.GetTop() //获取标签
	//设置模板路径
	this.TplName = "home/index.html"
}


/**
	获取前10个数据
	@param ""
	@return
 */
 func (this *IndexController) getHomeArticle(){

	 result,num,_:= models.GetHomeArticle()

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

	 this.Data["articleNum"] = num
	 this.Data["article"] = data
 }


 /**
 	文章分页前台
 	@param id int
 	@return []
  */
// @router /article/?:key [get] 文章分页
func (this *IndexController) GetHomePageArticle(){

	var id,category,page,tag int

	var keywords string //关键词查询

	this.Ctx.Input.Bind(&id,"id")

	this.Ctx.Input.Bind(&category,"category")

	this.Ctx.Input.Bind(&page,"page")

	this.Ctx.Input.Bind(&keywords,"keywords")

	this.Ctx.Input.Bind(&tag,"tag")

	if id == 0{

		this.Data["json"] = map[string]interface{}{
			"code":"1003",
			"errmsg":"获取参数错误，error",
		}
	}else{
		var result []models.LiteArticle

		if tag > 0{

			result,_ = this.TagsPage(tag,page)

		}else{

			result,_,_ = models.GetHomeAndPageArticle(id,category,page,keywords)

		}

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
	this.Data["TagList"] = this.getHomeTags() //展示的是热门标签
	this.GetTop()//获取置顶
	this.NexArticle(id) //文章条数详情
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

			var ip string

			ips := this.Ctx.Request.Header.Get("X-Real-IP") //nginx设置反向代理以后获取的值

			if ips == ""{

				ip = this.Ctx.Request.RemoteAddr

				ip = ip[0:strings.LastIndex(ip, ":")]

			}else{

				ip = ips
			}

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

		data,_ := models.SelectReviewPage(page)

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
func (this *IndexController) getKeyword(key string,tag int){

	var result []models.LiteArticle
	var num int

	if tag != 0{

		result,num = this.TagsPage(tag,1)//查询tag表中的数据

	}else{

		result,num,_ = models.GetArticleKeywords(key) //关键词查询的代码
	}

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
	this.Data["articleNum"] = num
	this.Data["PageTag"] = tag

}

/**
	设置下载文件
 */
// @router /download/file/?:key [get] 下载文件
func (this *IndexController) DownFile(){

	//获取当前用户的ip地址

	var fileFullPath string

	this.Ctx.Input.Bind(&fileFullPath,"file")

	if fileFullPath == ""{

		this.Abort("404")
	}

	if fileFullPath != "log.txt"{

		this.Abort("404")
	}

	this.Ctx.Output.Download("download/"+fileFullPath)

}


/**
	获取前台展示的标签
 */

 func (this *IndexController) getHomeTags()(res []TagList){

 	var data []TagList

 	data = make([]TagList,0)

 	list,_:= models.FindTagChecked()
	i := 0
	for _,val:=range list{
		//解决前台样式
		if i > 3{
			i = 0
		}
		num := models.CountArticleAndTag(int(val.ID))

		res := TagList{val.Tag_name,num.Total,int(val.ID),i + 1}

		data = append(data,res)
		i++
	}

	return data
 }


 /**
 	处理前台 获取展示
 	@param tid int page int 标签id 分页id
    @return
  */
 func (this *IndexController) TagsPage(tid int,page int)(list []models.LiteArticle,num int){

 	res,num := models.GetHomeTagsArticle(tid,page)

 	var arr []models.LiteArticle

 	for _,val := range res{

		//获取文章id进行获取文章
 		data,_ := models.GetTagArticleInfo(val.Aid)

 		arr = append(arr,data)

	}

 	return arr,num
 }


 /**
 	获取置顶推荐
  */
 func (this *IndexController) GetTop(){
	//查询前10条置顶推荐

	this.Data["TopArticle"],_ = models.ArticleTopList()

	//查询友情链接
	this.Data["Link"],_ = models.GetHomeLink()

	var data []TimeString

	result,_ := models.SelectReview()

	data = make([]TimeString,0)

	for _,val := range result{

		arr := TimeString{val,this.TimeRangeComparison(val.CreatedAt)}

		data = append(data,arr)
	}

	 //最新留言
	 this.Data["Commit"] = data
 }


 /**
 	获取详情页中的获取下一条and 上一条
  */
 func(this *IndexController) NexArticle(id int){
	//获取当前id的上一篇和下一篇

	resNext,_ := models.ArticleNext(id)

	resPrev,_ := models.ArticlePrev(id)

	if resPrev.ID == 0{

		this.Data["articlePrev"] = map[string]string{
			"id":"0",
			"title":"没有了",
		}
	}else{
		this.Data["articlePrev"] = map[string]string{
			"id":strconv.Itoa(int(resPrev.ID)),
			"title":resPrev.Title,
		}
	}

	if resNext.ID == 0{

		this.Data["articleNext"] = map[string]string{
			"id":"0",
			"title":"没有了",
		}

	}else{

		this.Data["articleNext"] = map[string]string{
			"id":strconv.Itoa(int(resNext.ID)),
			"title":resNext.Title,
		}
	}


 }



 /**
 	小说列表页
  */
 func (this *IndexController) GetHomeFiction(){

 	this.Data["fictionNum"],_ = models.CountHomeFictionNum()

 	this.Data["fictionList"],_ = this.queryFictionData(1,10)

 	this.TplName = "home/fiction.html"
 }

 /**
 	小说展示页面
  */
 func (this *IndexController) HomeFictionPage(){

 	var page,size int

 	this.Ctx.Input.Bind(&page,"page")

 	this.Ctx.Input.Bind(&size,"size")

 	data,err := this.queryFictionData(page,10)

 	if err != nil{

		this.Data["json"] = map[string]interface{}{
			"code":"1002",
			"msg":"请求错误，请稍后...",
		}
	}else{

		this.Data["json"] = map[string]interface{}{
			"code":"0",
			"data":data,
			"msg":"获取成功",
		}
	}

 	this.ServeJSON()
 }
 
 /**
 	小说页面处理展示
  */
func (this *IndexController) queryFictionData(page,size int) (list []FictionList,err error)  {

	data,err := models.FindHomeFictionData(page,10)

	//var list []FictionList

	if err != nil{

		return list,err

	}else{

		list = make([]FictionList,0)

		for _,val := range data {

			ids,_ := models.FictionOperation(int(val.ID))

			dataRes := FictionList{val,"",val.CreatedAt.Format("2006-01-02 13:04:05"),ids.DownloadNum}

			if val.Tags != ""{

				res := models.FictionAndTag(val.Tags)

				dataRes.Tags = res.Fname

			}

			list = append(list,dataRes)

		}

		return list,nil

	}
}

/**
	小说下载
 */
func (this *IndexController) HomeFictionDownload()  {

	var fictionId int

	this.Ctx.Input.Bind(&fictionId,"txt_id")

	if fictionId == 0{

		this.Abort500(nil)
	}

	//查询是否存在此小说

	data,err := models.FirstFictionDownload(fictionId)

	if err != nil{

		//写入日志
		logs.Error(err)
		this.Abort("500")

	}else{

		//查询是否是封禁的ip
		num,_:= models.QueryBanned(this.GetIP())

		if num > 0{

			logs.Info(this.GetIP()+"又开始请求下载了")

			this.Abort("404")
		}else{
			//获取下载的详细内容
			this.updateOperation(data)

			this.Ctx.Output.Download(data.FileName,data.Name)
		}

	}
}

//更新日志表和批次表

func (this *IndexController) updateOperation(data models.LiteFiction) bool {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		models.CreateFictionLog(data,this.GetIP())

		wg.Done()
	}()

	go func() {

		models.UpdateOperation(data)

	}()

	wg.Wait() //执行完成，解除阻塞

	return  true
}