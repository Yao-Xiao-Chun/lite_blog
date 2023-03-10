package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "AddArticle",
			Router:           "/admin/article/add",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "Article",
			Router:           "/admin/article/add",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "DelArticle",
			Router:           "/admin/article/delete/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "ArticleEdit",
			Router:           "/admin/article/edit/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "EditArticle",
			Router:           "/admin/article/edit/?:key",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "GetArticleList",
			Router:           "/admin/article/list",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "GetArticleInfo",
			Router:           "/admin/article/listinfo/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "Uploads",
			Router:           "/admin/upload",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "UploadArticles",
			Router:           "/admin/upload/article",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminArticleController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminArticleController"],
		beego.ControllerComments{
			Method:           "GetUser",
			Router:           "/admin/user/add",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "AdminIndex",
			Router:           "/admin",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "AdminGettime",
			Router:           "/admin/addtime",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "SetPlacard",
			Router:           "/admin/baseplacard",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "PlacardFormData",
			Router:           "/admin/baseplacard",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "SetAbort",
			Router:           "/admin/baseseting",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "AbortFormData",
			Router:           "/admin/baseseting",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "SetClear",
			Router:           "/admin/clear",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "DeleteFile",
			Router:           "/admin/clear/delete/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "DownCacheLog",
			Router:           "/admin/clear/download/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "DownFile",
			Router:           "/admin/clear/download/logs/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "GetFileName",
			Router:           "/admin/clear/log/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "GetLogList",
			Router:           "/admin/log/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "GetLogPage",
			Router:           "/admin/log/page/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "AdminLogin",
			Router:           "/admin/login",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "GetUserMessage",
			Router:           "/admin/user/message",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "DownLog",
			Router:           "/download/log",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "AdminMain",
			Router:           "/main",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "ReadExcel",
			Router:           "/upload/done",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminIndexController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminIndexController"],
		beego.ControllerComments{
			Method:           "ToExcel",
			Router:           "/upload/excel",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminLinkController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminLinkController"],
		beego.ControllerComments{
			Method:           "LinkAdd",
			Router:           "/admin/link/add",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminLinkController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminLinkController"],
		beego.ControllerComments{
			Method:           "LinkAddForm",
			Router:           "/admin/link/add",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminLinkController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminLinkController"],
		beego.ControllerComments{
			Method:           "LinkDel",
			Router:           "/admin/link/delete/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminLinkController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminLinkController"],
		beego.ControllerComments{
			Method:           "LinkFormData",
			Router:           "/admin/link/edit",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminLinkController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminLinkController"],
		beego.ControllerComments{
			Method:           "LinkInfo",
			Router:           "/admin/link/edit/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminLinkController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminLinkController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           "/admin/link/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminLinkController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminLinkController"],
		beego.ControllerComments{
			Method:           "GetLinkIndex",
			Router:           "/admin/link/page/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminMenuController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminMenuController"],
		beego.ControllerComments{
			Method:           "GetList",
			Router:           "/admin/menu",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminMenuController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminMenuController"],
		beego.ControllerComments{
			Method:           "GetAdd",
			Router:           "/admin/menu/add",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminMenuController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminMenuController"],
		beego.ControllerComments{
			Method:           "MenuAddForm",
			Router:           "/admin/menu/add",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminMenuController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminMenuController"],
		beego.ControllerComments{
			Method:           "DeleteMenu",
			Router:           "/admin/menu/del/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminMenuController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminMenuController"],
		beego.ControllerComments{
			Method:           "EditPost",
			Router:           "/admin/menu/edit",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminMenuController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminMenuController"],
		beego.ControllerComments{
			Method:           "EditMenu",
			Router:           "/admin/menu/edit/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminReviewController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminReviewController"],
		beego.ControllerComments{
			Method:           "ReviewIndex",
			Router:           "/admin/review",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminReviewController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminReviewController"],
		beego.ControllerComments{
			Method:           "DeleteReview",
			Router:           "/admin/review/delete/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminReviewController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminReviewController"],
		beego.ControllerComments{
			Method:           "GetReviewPage",
			Router:           "/admin/review/page/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTagController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTagController"],
		beego.ControllerComments{
			Method:           "TagList",
			Router:           "/admin/tag",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTagController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTagController"],
		beego.ControllerComments{
			Method:           "TagAdd",
			Router:           "/admin/tag/add",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTagController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTagController"],
		beego.ControllerComments{
			Method:           "TagAddPost",
			Router:           "/admin/tag/add",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTagController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTagController"],
		beego.ControllerComments{
			Method:           "DelTag",
			Router:           "/admin/tag/del/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTagController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTagController"],
		beego.ControllerComments{
			Method:           "GatTagInfo",
			Router:           "/admin/tag/edit/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTagController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTagController"],
		beego.ControllerComments{
			Method:           "GetTagList",
			Router:           "/admin/tag/list/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTagController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTagController"],
		beego.ControllerComments{
			Method:           "SetTagInfo",
			Router:           "/admin/tag/update",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTimeController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTimeController"],
		beego.ControllerComments{
			Method:           "DelTime",
			Router:           "/admin/deltime/:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTimeController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTimeController"],
		beego.ControllerComments{
			Method:           "SetTimeInfo",
			Router:           "/admin/formtime/:key",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTimeController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTimeController"],
		beego.ControllerComments{
			Method:           "GetTimeInfo",
			Router:           "/admin/timeinfo/?:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminTimeController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminTimeController"],
		beego.ControllerComments{
			Method:           "AddTime",
			Router:           "/admin/timepost/:key",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminUserController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminUserController"],
		beego.ControllerComments{
			Method:           "DoLogin",
			Router:           "/admin/dologin",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminUserController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminUserController"],
		beego.ControllerComments{
			Method:           "LoginOut",
			Router:           "/admin/loginout",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminUserController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminUserController"],
		beego.ControllerComments{
			Method:           "UserList",
			Router:           "/admin/user",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminUserController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminUserController"],
		beego.ControllerComments{
			Method:           "CreateUser",
			Router:           "/admin/user/create",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminUserController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminUserController"],
		beego.ControllerComments{
			Method:           "DelUser",
			Router:           "/admin/user/del/?:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminUserController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminUserController"],
		beego.ControllerComments{
			Method:           "EditUser",
			Router:           "/admin/user/edit/?:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminUserController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminUserController"],
		beego.ControllerComments{
			Method:           "UserPage",
			Router:           "/admin/user/page/?:page",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:AdminUserController"] = append(beego.GlobalControllerRouter["mywork/internal:AdminUserController"],
		beego.ControllerComments{
			Method:           "EditUserData",
			Router:           "/admin/user/update",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:CronTabController"] = append(beego.GlobalControllerRouter["mywork/internal:CronTabController"],
		beego.ControllerComments{
			Method:           "TaskAdd",
			Router:           "/admin/crontab/add",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:CronTabController"] = append(beego.GlobalControllerRouter["mywork/internal:CronTabController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/admin/crontab/add",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:CronTabController"] = append(beego.GlobalControllerRouter["mywork/internal:CronTabController"],
		beego.ControllerComments{
			Method:           "StopAllTask",
			Router:           "/admin/crontab/allstop",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:CronTabController"] = append(beego.GlobalControllerRouter["mywork/internal:CronTabController"],
		beego.ControllerComments{
			Method:           "CrontabDelete",
			Router:           "/admin/crontab/delete/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:CronTabController"] = append(beego.GlobalControllerRouter["mywork/internal:CronTabController"],
		beego.ControllerComments{
			Method:           "Task",
			Router:           "/admin/crontab/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:CronTabController"] = append(beego.GlobalControllerRouter["mywork/internal:CronTabController"],
		beego.ControllerComments{
			Method:           "GetTaskPage",
			Router:           "/admin/crontab/page/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:CronTabController"] = append(beego.GlobalControllerRouter["mywork/internal:CronTabController"],
		beego.ControllerComments{
			Method:           "StartTask",
			Router:           "/admin/crontab/startTask/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:CronTabController"] = append(beego.GlobalControllerRouter["mywork/internal:CronTabController"],
		beego.ControllerComments{
			Method:           "StopTask",
			Router:           "/admin/crontab/stopTask/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           "/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "IndexAbout",
			Router:           "/about",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "GetHomePageArticle",
			Router:           "/article/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "ArticleClick",
			Router:           "/article/click/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "GetArticleInfo",
			Router:           "/article/info/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "CommitArticle",
			Router:           "/blog/pages/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "TypeArticle",
			Router:           "/category/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "IndexDetails",
			Router:           "/details",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "DownFile",
			Router:           "/download/file/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "IndexMessage",
			Router:           "/message",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "SteReview",
			Router:           "/message/review",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "HomePageReview",
			Router:           "/message/review/page/?:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "IndexTime",
			Router:           "/time",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:IndexController"] = append(beego.GlobalControllerRouter["mywork/internal:IndexController"],
		beego.ControllerComments{
			Method:           "GetTimePage",
			Router:           "/time/page/?:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:UploadController"] = append(beego.GlobalControllerRouter["mywork/internal:UploadController"],
		beego.ControllerComments{
			Method:           "UploadFileDone",
			Router:           "/admin/upload/file",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:UploadController"] = append(beego.GlobalControllerRouter["mywork/internal:UploadController"],
		beego.ControllerComments{
			Method:           "UploadIndex",
			Router:           "/admin/upload/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mywork/internal:UploadController"] = append(beego.GlobalControllerRouter["mywork/internal:UploadController"],
		beego.ControllerComments{
			Method:           "SetUpload",
			Router:           "/admin/upload/success",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
