package admin

import (
	"mywork/models"
)

/**
留言管理模块
*/
type AdminReviewController struct {
	AdminBaseController
}

/**
留言管理首页
*/
// @router /admin/review [get] 留言 review列表
func (this *AdminReviewController) ReviewIndex() {

	this.getReviewCount()

	this.TplName = "admin/review/list.html"
}

/**
获取留言评论数据
*/
// @router /admin/review/page/?:key [get] 留言 review列表
func (this *AdminReviewController) GetReviewPage() {

	var page int

	this.Ctx.Input.Bind(&page, "page")

	var res []models.LiteReview

	res = make([]models.LiteReview, 0)

	if page == 0 {

		res, _ = models.SelectReview()
	} else {

		res, _ = models.SelectReviewPage(page)
	}

	this.Data["json"] = map[string]interface{}{
		"code": "0",
		"data": res,
	}

	this.ServeJSON()

}

/**
总条数
*/
func (this *AdminReviewController) getReviewCount() {

	page, _ := models.ReviewCount()

	this.Data["num"] = page
}

/**
删除此条评论
*/
// @router /admin/review/delete/?:key [get] 删除 review列表
func (this *AdminReviewController) DeleteReview() {

	var id int

	this.Ctx.Input.Bind(&id, "id")

	if id == 0 {

		this.Data["json"] = map[string]interface{}{
			"code":   "1003",
			"errmsg": "获取参数错误",
		}

	} else {

		models.DeleteReview(id)

		this.Data["json"] = map[string]interface{}{
			"code":   "0",
			"errmsg": "删除成功",
		}

	}
	this.ServeJSON()
}
