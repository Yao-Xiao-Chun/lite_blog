package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"mywork/internal/app/admin"
	"mywork/internal/app/common"
	"mywork/internal/app/home"
)

func initRoute() {

	//注册错误处理
	beego.ErrorController(&common.ErrorController{})

	beego.Include(&admin.AdminUserController{}) //后台登录用户

	beego.Include(&admin.AdminTimeController{}) //时间线控制器

	beego.Include(&admin.ArticleController{}) //文章控制器

	beego.Include(&admin.AdminTagController{}) //Tag控制器

	beego.Include(&admin.AdminMenuController{}) //menu控制器

	beego.Include(&admin.ReviewController{}) //留言控制器

	beego.Include(&admin.AdminLinkController{}) //友情链接

	beego.Include(&admin.UploadController{}) //大文件上传demo

	beego.Include(&admin.CronTabController{}) //定时任务控制器

}

// HomeApi TODO
func HomeApi() {

	beego.Router("/", &home.IndexController{}, "*:Index")

	beego.Handler("/captcha/*.png", captcha.Server(120, 38)) //设置验证码

	beego.Router("/fiction", &home.IndexController{}, "get:GetHomeFiction") //前台小说列表页

	beego.Router("/fiction/page/?:key", &home.IndexController{}, "get:HomeFictionPage") //前台小说列表页

	beego.Router("/fiction/download/?:key", &home.IndexController{}, "get:HomeFictionDownload") //前台小说下载
	beego.Router("/about", &home.IndexController{}, "get:IndexAbout")                           //关于
	beego.Router("/message", &home.IndexController{}, "get:IndexMessage")
	beego.Router("/details", &home.IndexController{}, "get:IndexDetails")
	beego.Router("/time", &home.IndexController{}, "get:IndexTime")
	beego.Router("/time/page/?:id", &home.IndexController{}, "get:GetTimePage")
	beego.Router("/category/?:key", &home.IndexController{}, "get:TypeArticle")
	beego.Router("/article/?:key", &home.IndexController{}, "get:GetHomePageArticle")
	beego.Router("/article/info/?:key", &home.IndexController{}, "get:GetArticleInfo")
	beego.Router("/message/review", &home.IndexController{}, "post:SteReview")
	beego.Router("/message/review/page/?:key", &home.IndexController{}, "get:HomePageReview")
	beego.Router("/article/click/?:key", &home.IndexController{}, "get:ArticleClick") //文章点赞
	beego.Router("/blog/pages/?:key", &home.IndexController{}, "get:CommitArticle")   //文章点赞
	beego.Router("/download/file/?:key", &home.IndexController{}, "get:DownFile")     //文章点赞
	beego.Router("/download/file/?:key", &home.IndexController{}, "get:DownFile")     //文章点赞

	includeApi()
	initRoute()
}
 
func includeApi() {

	beego.Include(&admin.AdminIndexController{}) //后台首页

	beego.Router("/admin/webpush/index", &admin.AdminTaskController{}, "get:TaskIndex") //推送任务列表首页

	beego.Router("admin/fiction/index", &admin.FileController{}, "get:FileIndex") //小说列表页

	beego.Router("admin/fiction/delete/?:key", &admin.FileController{}, "get:SetFictionStatus") //禁止小说下载

	beego.Router("admin/fiction/index/page/?:key", &admin.FileController{}, "get:FilePage") //小说列表页

	//小说日志路由模块 admin
	beego.Router("admin/fiction/log", &admin.FileController{}, "get:FictionLog") //小说日志访问路由

	beego.Router("/admin/fiction/log/page/?:key", &admin.FileController{}, "get:FictionLogPage") //小说日志列表

	beego.Router("/admin/fiction/banned/?:key", &admin.FileController{}, "get:FictionBanned") //ip黑名单
	beego.Router("/ws/index", &admin.PushSocketController{})                                  //websocket控制器
}
