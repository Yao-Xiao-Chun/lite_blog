package routers

import (
	"mywork/controllers"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {

	//注册错误处理
	beego.ErrorController(&controllers.ErrorController{})
	
	beego.Include(&controllers.IndexController{})//前台首页

	beego.Include(&controllers.AdminIndexController{})//后台首页

	beego.Include(&controllers.AdminUserController{})//后台登录用户

	beego.Include(&controllers.AdminTimeController{})//时间线控制器

	beego.Include(&controllers.AdminArticleController{})//文章控制器

	beego.Include(&controllers.AdminTagController{})//Tag控制器

	beego.Include(&controllers.AdminMenuController{})//menu控制器

	beego.Include(&controllers.AdminReviewController{})//留言控制器

	beego.Include(&controllers.AdminLinkController{})//友情链接

	beego.Include(&controllers.UploadController{})//大文件上传demo

	beego.Include(&controllers.CronTabController{})//定时任务控制器

	beego.Handler("/captcha/*.png", captcha.Server(120, 38)) //设置验证码

	beego.Router("/ws/index", &controllers.PushSocketController{}) //websocket控制器

	beego.Router("/admin/webpush/index",&controllers.AdminTaskController{},"get:TaskIndex") //推送任务列表首页

	beego.Router("admin/fiction/index",&controllers.AdminFileController{},"get:FileIndex") //小说列表页

	beego.Router("admin/fiction/delete/?:key",&controllers.AdminFileController{},"get:SetFictionStatus") //禁止小说下载

	beego.Router("admin/fiction/index/page/?:key",&controllers.AdminFileController{},"get:FilePage") //小说列表页

	beego.Router("/fiction",&controllers.IndexController{},"get:GetHomeFiction") //前台小说列表页

	beego.Router("/fiction/page/?:key",&controllers.IndexController{},"get:HomeFictionPage") //前台小说列表页

	beego.Router("/fiction/download/?:key",&controllers.IndexController{},"get:HomeFictionDownload") //前台小说下载

	//小说日志路由模块 admin
	beego.Router("admin/fiction/log",&controllers.AdminFileController{},"get:FictionLog") //小说日志访问路由

	beego.Router("/admin/fiction/log/page/?:key",&controllers.AdminFileController{},"get:FictionLogPage") //小说日志列表

	beego.Router("/admin/fiction/banned/?:key",&controllers.AdminFileController{},"get:FictionBanned") //ip黑名单




}


