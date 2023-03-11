package common

import (
	"mywork/internal/app/admin"
	"mywork/internal/app/home"
	"mywork/internal/pkg/dto"
	"mywork/models"

	"github.com/astaxie/beego/logs"
	"github.com/dchest/captcha"
	"github.com/jinzhu/gorm"
	"html/template"
	"strconv"
	"strings"
	"sync"
)

type IndexController struct {
	controllers.HomeBaseController
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
func (c *IndexController) Index() {

	var keywords string

	var tag int

	c.Ctx.Input.Bind(&keywords, "keyword")

	c.Ctx.Input.Bind(&tag, "tag")

	if tag != 0 {

		c.getKeyword("", tag) //查询tag

	} else if keywords == "" {

		c.getHomeArticle() //查询数据

	} else {

		c.getKeyword(keywords, 0) //查询关键词
	}

	c.Data["TagList"] = c.getHomeTags() //展示的是热门标签

	c.Data["Placard"], _ = models.GetPlacard()

	c.GetTop() //获取置顶
	//设置模板路径
	c.TplName = "home/index.html"
}

// 关于

// @router /about [get] 首页
func (c *IndexController) IndexAbout() {

	c.Data["Abort"], _ = models.GetAbort()
	//设置模板路径
	c.TplName = "home/about.html"
}

//消息

// @router /message [get] 首页
func (c *IndexController) IndexMessage() {

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML()) //设置模板xss

	d := struct {
		CaptchaId string
	}{
		captcha.NewLen(4),
	}

	c.Data["CaptchaId"] = d.CaptchaId //生成验证码

	c.Data["Uuid"] = c.GetRandomString(24) //token

	c.Data["RePage"], _ = models.GetHomeReviewCount()

	c.Data["viewList"], _ = models.SelectReview()
	//设置模板路径
	c.TplName = "home/message.html"
}

//详情

// @router /details [get] 首页
func (c *IndexController) IndexDetails() {

	//设置模板路径
	c.TplName = "home/details.html"
}

//详情

// @router /time [get] 时间线
func (c *IndexController) IndexTime() {

	c.getTimeLine()
	//设置模板路径
	c.TplName = "home/time.html"
}

/**
时间线分页
*/

// @router  /time/page/?:id [get]  分页代码
func (c *IndexController) GetTimePage() {

	var tid int

	c.Ctx.Input.Bind(&tid, "id")

	if tid == 0 {

		c.Abort("404")
	}

	line, _ := models.GetPageTimeLine(tid, 10)

	arr := make(map[int]map[string]interface{}, 10)

	if len(line) == 0 {

		c.Data["json"] = map[string]interface{}{
			"code":    "2",
			"data":    "",
			"message": "没有数据了！",
		}

	} else {
		//变量赋值
		for key, val := range line {

			data := make(map[string]interface{}, 4) // 每次使用都要初始化一次

			data["code"] = val.Title

			data["content"] = val.Content

			data["tid"] = val.ID

			data["token"] = val.Token

			arr[key] = data

		}

		c.Data["json"] = map[string]interface{}{
			"code":    "0",
			"data":    arr,
			"message": "请求成功",
		}
	}

	c.ServeJSON()
}

/**
时间线负责前端调用 Page
*/
func (c *IndexController) getTimeLine() {

	line, _ := models.GetHomeTimeLine() //多维结构体

	arr := make(map[int]map[string]interface{}, 10)

	//变量赋值
	for key, val := range line {

		data := make(map[string]interface{}, 3) // 每次使用都要初始化一次

		data["code"] = val.Title

		data["content"] = val.Content

		data["tid"] = val.ID

		arr[key] = data

	}

	num, _ := models.GetHomeCountTimeLine()

	c.Data["count"] = num

	c.Data["line"] = arr

}

/**
其他菜单处理逻辑
*/
// @router /category/?:key [get] 菜单处理
func (c *IndexController) TypeArticle() {

	var id int
	var page int
	c.Ctx.Input.Bind(&id, "id")

	c.Ctx.Input.Bind(&page, "page")

	result, num, _ := models.GetMenuAndArticle(id, page)

	if len(result) == 0 {

		c.Abort("500")

	} else {

		var data map[int]map[string]interface{}

		var arrData map[string]interface{}

		data = make(map[int]map[string]interface{})

		for key, val := range result {

			arrData = make(map[string]interface{}, 12)
			arrData["created_time"] = val.CreatedAt.Format("2006-01-02 15:04:05") //创建时间
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
			if val.Title_img != "undefind" {
				arrData["img"] = val.Title_img
			}

			data[key] = arrData
		}

		c.Data["article"] = data

		c.Data["articleNum"] = num
	}

	c.Data["Placard"], _ = models.GetPlacard()

	c.Data["TagList"] = c.getHomeTags() //展示的是热门标签

	c.Data["is_category"] = id

	c.GetTop() //获取标签
	//设置模板路径
	c.TplName = "home/index.html"
}

/**
获取前10个数据
@param ""
@return
*/
func (c *IndexController) getHomeArticle() {

	result, num, _ := models.GetHomeArticle()

	var data map[int]map[string]interface{}

	var arrData map[string]interface{}

	data = make(map[int]map[string]interface{})

	for key, val := range result {

		arrData = make(map[string]interface{}, 12)
		arrData["created_time"] = val.CreatedAt.Format("2006-01-02 15:04:05") //创建时间
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
		if val.Title_img != "undefind" {
			arrData["img"] = val.Title_img
		}

		data[key] = arrData
	}

	c.Data["articleNum"] = num
	c.Data["article"] = data
}

/**
文章分页前台
@param id int
@return []
*/
// @router /article/?:key [get] 文章分页
func (c *IndexController) GetHomePageArticle() {

	var id, category, page, tag int

	var keywords string //关键词查询

	c.Ctx.Input.Bind(&id, "id")

	c.Ctx.Input.Bind(&category, "category")

	c.Ctx.Input.Bind(&page, "page")

	c.Ctx.Input.Bind(&keywords, "keywords")

	c.Ctx.Input.Bind(&tag, "tag")

	if id == 0 {

		c.Data["json"] = map[string]interface{}{
			"code":   "1003",
			"errmsg": "获取参数错误，error",
		}
	} else {
		var result []models.LiteArticle

		if tag > 0 {

			result, _ = c.TagsPage(tag, page)

		} else {

			result, _, _ = models.GetHomeAndPageArticle(id, category, page, keywords)

		}

		if len(result) == 0 {
			c.Data["json"] = map[string]interface{}{
				"code":   "1002",
				"errmsg": "已经最后一页了，别点了",
			}
		} else {

			var data map[int]map[string]interface{}

			var arrData map[string]interface{}

			data = make(map[int]map[string]interface{})

			for key, val := range result {

				arrData = make(map[string]interface{}, 12)
				arrData["created_time"] = val.CreatedAt.Format("2006-01-02 15:04:05") //创建时间
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
				if val.Title_img != "undefind" {
					arrData["img"] = val.Title_img
				}

				data[key] = arrData
			}

			c.Data["json"] = map[string]interface{}{
				"code": "0",
				"data": data,
			}
		}

	}

	c.ServeJSON()
}

/**
获取文章详情
@param id int
@return
*/
// @router /article/info/?:key [get] 文章详情
func (c *IndexController) GetArticleInfo() {

	var id int

	c.Ctx.Input.Bind(&id, "id")

	if id == 0 {

		c.Abort("404")
	}

	res, _ := models.GetHomeArticleInfo(id)

	//更新阅读数
	models.SetArticleAndRead(id)
	c.Data["articleData"] = res
	c.Data["TagList"] = c.getHomeTags() //展示的是热门标签
	c.GetTop()                          //获取置顶
	c.NexArticle(id)                    //文章条数详情
	c.TplName = "home/details.html"
}

/**
留言处理
*/
// @router /message/review [post] 文章详情
func (c *IndexController) SteReview() {

	text := c.GetString("text")

	captchaId := c.GetString("captchaId")

	ver_code := c.GetString("ver_code")

	token := c.GetString("token")

	if !VerifyCaptcha(captchaId, ver_code) {

		c.Data["json"] = map[string]interface{}{
			"code":   "1003",
			"errmsg": "验证码错误",
		}

	} else {

		res, err := models.SelectReviewToken(token)

		if !gorm.IsRecordNotFoundError(err) {

			c.Data["json"] = map[string]interface{}{
				"code":   "1004",
				"errmsg": "请勿重复提交",
			}
		} else {

			var ip string

			ips := c.Ctx.Request.Header.Get("X-Real-IP") //nginx设置反向代理以后获取的值

			if ips == "" {

				ip = c.Ctx.Request.RemoteAddr

				ip = ip[0:strings.LastIndex(ip, ":")]

			} else {

				ip = ips
			}

			res.Token = token

			res.Message, _ = admin.GetSummary(text)

			res.Ip = ip

			res.Address = c.GetAddress(ip)

			models.CreateReview(res)

			c.Data["json"] = map[string]interface{}{
				"code":   "0",
				"errmsg": "留言成功",
			}
		}

	}
	//判断是否

	c.ServeJSON()
}

/**
留言分页处理
*/
// @router /message/review/page/?:key [get] 留言分页
func (c *IndexController) HomePageReview() {

	var page int

	c.Ctx.Input.Bind(&page, "page")

	if page == 0 {

		c.Data["json"] = map[string]interface{}{
			"code":   "1003",
			"errmsg": "参数错误，获取信息失败",
		}
	} else {

		data, _ := models.SelectReviewPage(page)

		c.Data["json"] = map[string]interface{}{
			"code":   "0",
			"data":   data,
			"errmsg": "请求成功",
		}

	}

	c.ServeJSON()
}

/**
文章点赞
*/
// @router /article/click/?:key [get] 文章点赞
func (c *IndexController) ArticleClick() {

	var aid int

	c.Ctx.Input.Bind(&aid, "aid")

	if aid == 0 {
		c.Data["json"] = map[string]interface{}{
			"code":   "1003",
			"errmsg": "参数非法,error",
		}
	} else {
		models.SetArticleAndClick(aid)
		c.Data["json"] = map[string]interface{}{
			"code":   "0",
			"errmsg": "点赞成功",
		}
	}
	c.ServeJSON()
}

/**
评论登录
*/
// @router /blog/pages/?:key [get] 评论
func (c *IndexController) CommitArticle() {

	c.Redirect("/", 302)
}

/**
前台首页搜索框查询内容
*/
func (c *IndexController) getKeyword(key string, tag int) {

	var result []models.LiteArticle
	var num int

	if tag != 0 {

		result, num = c.TagsPage(tag, 1) //查询tag表中的数据

	} else {

		result, num, _ = models.GetArticleKeywords(key) //关键词查询的代码
	}

	var data map[int]map[string]interface{}

	var arrData map[string]interface{}

	data = make(map[int]map[string]interface{})

	for key, val := range result {

		arrData = make(map[string]interface{}, 12)
		arrData["created_time"] = val.CreatedAt.Format("2006-01-02 15:04:05") //创建时间
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
		if val.Title_img != "undefind" {
			arrData["img"] = val.Title_img
		}

		data[key] = arrData
	}

	c.Data["article"] = data
	c.Data["Keywords"] = key
	c.Data["articleNum"] = num
	c.Data["PageTag"] = tag

}

/**
设置下载文件
*/
// @router /download/file/?:key [get] 下载文件
func (c *IndexController) DownFile() {

	//获取当前用户的ip地址

	var fileFullPath string

	c.Ctx.Input.Bind(&fileFullPath, "file")

	if fileFullPath == "" {

		c.Abort("404")
	}

	if fileFullPath != "log.txt" {

		c.Abort("404")
	}

	c.Ctx.Output.Download("download/" + fileFullPath)

}

/**
获取前台展示的标签
*/

func (c *IndexController) getHomeTags() (res []TagList) {

	var data []TagList

	data = make([]TagList, 0)

	list, _ := models.FindTagChecked()
	i := 0
	for _, val := range list {
		//解决前台样式
		if i > 3 {
			i = 0
		}
		num := models.CountArticleAndTag(int(val.ID))

		res := TagList{val.Tag_name, num.Total, int(val.ID), i + 1}

		data = append(data, res)
		i++
	}

	return data
}

/**
	处理前台 获取展示
	@param tid int page int 标签id 分页id
   @return
*/
func (c *IndexController) TagsPage(tid int, page int) (list []models.LiteArticle, num int) {

	res, num := models.GetHomeTagsArticle(tid, page)

	var arr []models.LiteArticle

	for _, val := range res {

		//获取文章id进行获取文章
		data, _ := models.GetTagArticleInfo(val.Aid)

		arr = append(arr, data)

	}

	return arr, num
}

/**
获取置顶推荐
*/
func (c *IndexController) GetTop() {
	//查询前10条置顶推荐

	c.Data["TopArticle"], _ = models.ArticleTopList()

	//查询友情链接
	c.Data["Link"], _ = models.GetHomeLink()

	var data []TimeString

	result, _ := models.SelectReview()

	data = make([]TimeString, 0)

	for _, val := range result {

		arr := TimeString{val, c.TimeRangeComparison(val.CreatedAt)}

		data = append(data, arr)
	}

	//最新留言
	c.Data["Commit"] = data
}

/**
获取详情页中的获取下一条and 上一条
*/
func (c *IndexController) NexArticle(id int) {
	//获取当前id的上一篇和下一篇

	resNext, _ := models.ArticleNext(id)

	resPrev, _ := models.ArticlePrev(id)

	if resPrev.ID == 0 {

		c.Data["articlePrev"] = map[string]string{
			"id":    "0",
			"title": "没有了",
		}
	} else {
		c.Data["articlePrev"] = map[string]string{
			"id":    strconv.Itoa(int(resPrev.ID)),
			"title": resPrev.Title,
		}
	}

	if resNext.ID == 0 {

		c.Data["articleNext"] = map[string]string{
			"id":    "0",
			"title": "没有了",
		}

	} else {

		c.Data["articleNext"] = map[string]string{
			"id":    strconv.Itoa(int(resNext.ID)),
			"title": resNext.Title,
		}
	}

}

/**
小说列表页
*/
func (c *IndexController) GetHomeFiction() {

	c.Data["fictionNum"], _ = models.CountHomeFictionNum()

	c.Data["fictionList"], _ = c.queryFictionData(1, 10)

	c.TplName = "home/fiction.html"
}

/**
小说展示页面
*/
func (c *IndexController) HomeFictionPage() {

	var page, size int

	c.Ctx.Input.Bind(&page, "page")

	c.Ctx.Input.Bind(&size, "size")

	data, err := c.queryFictionData(page, 10)

	if err != nil {

		c.Data["json"] = map[string]interface{}{
			"code": "1002",
			"msg":  "请求错误，请稍后...",
		}
	} else {

		c.Data["json"] = map[string]interface{}{
			"code": "0",
			"data": data,
			"msg":  "获取成功",
		}
	}

	c.ServeJSON()
}

/**
小说页面处理展示
*/
func (c *IndexController) queryFictionData(page, size int) (list []dto.FictionList, err error) {

	data, err := models.FindHomeFictionData(page, 10)

	//var list []FictionList

	if err != nil {

		return list, err

	} else {

		list = make([]dto.FictionList, 0)

		for _, val := range data {

			ids, _ := models.FictionOperation(int(val.ID))

			dataRes := dto.FictionList{LiteFiction: val, Times: val.CreatedAt.Format("2006-01-02 13:04:05"), DownloadNum: ids.DownloadNum}

			if val.Tags != "" {

				res := models.FictionAndTag(val.Tags)

				dataRes.Tags = res.Fname

			}

			list = append(list, dataRes)

		}

		return list, nil

	}
}

/**
小说下载
*/
func (c *IndexController) HomeFictionDownload() {

	var fictionId int

	c.Ctx.Input.Bind(&fictionId, "txt_id")

	if fictionId == 0 {

		c.Abort500(nil)
	}

	//查询是否存在此小说

	data, err := models.FirstFictionDownload(fictionId)

	if err != nil {

		//写入日志
		logs.Error(err)
		c.Abort("500")

	} else {

		//查询是否是封禁的ip
		num, _ := models.QueryBanned(c.GetIP())

		if num > 0 {

			logs.Info(c.GetIP() + "又开始请求下载了")

			c.Abort("404")
		} else {
			//获取下载的详细内容
			c.updateOperation(data)

			c.Ctx.Output.Download(data.FileName, data.Name)
		}

	}
}

//更新日志表和批次表

func (c *IndexController) updateOperation(data models.LiteFiction) bool {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		models.CreateFictionLog(data, c.GetIP())

		wg.Done()
	}()

	go func() {

		models.UpdateOperation(data)

	}()

	wg.Wait() //执行完成，解除阻塞

	return true
}
